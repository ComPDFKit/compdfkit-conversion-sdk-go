package compdf

/*
#cgo CFLAGS: -I. -Iinclude -D_CRT_SECURE_NO_WARNINGS
#cgo windows LDFLAGS: -L${SRCDIR}/lib/windows/amd64 -lcpdfconversionsdk
#cgo linux   LDFLAGS: -L${SRCDIR}/lib/linux/amd64   -lcpdfconversionsdk -Wl,-rpath,$ORIGIN
#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/lib/darwin/amd64 -lcpdfconversionsdk
#cgo darwin,arm64 LDFLAGS: -L${SRCDIR}/lib/darwin/arm64 -lcpdfconversionsdk

#include <stdlib.h>
#include "bridge.h"
*/
import "C"

import (
	"reflect"
	"runtime/cgo"
	"sync"
	"unsafe"
)

// callbackRegistry stores active ConvertCallback values keyed by a cgo.Handle
// value that gets passed to the C bridge as an opaque integer.
//
// We use cgo.Handle (Go 1.17+) directly: NewHandle returns a uintptr that the
// C side carries verbatim. The trampolines call back into Go and resolve the
// handle back to the original ConvertCallback.
type callbackRegistry struct {
	results sync.Map // uintptr(handle) -> *callbackBuffers
}

// callbackBuffers holds C-strings returned by the get_*_result trampolines.
// The C SDK requires the returned pointer remain valid until the next call to
// the same getter — we own one CString per result type per active callback
// and free it on the next overwrite or when the conversion finishes.
type callbackBuffers struct {
	mu        sync.Mutex
	ocr       *C.char
	layout    *C.char
	tableRes  *C.char
}

func (b *callbackBuffers) replace(target **C.char, val string) *C.char {
	if *target != nil {
		C.free(unsafe.Pointer(*target))
		*target = nil
	}
	if val == "" {
		return nil
	}
	*target = C.CString(val)
	return *target
}

func (b *callbackBuffers) free() {
	b.mu.Lock()
	defer b.mu.Unlock()
	for _, p := range [...]**C.char{&b.ocr, &b.layout, &b.tableRes} {
		if *p != nil {
			C.free(unsafe.Pointer(*p))
			*p = nil
		}
	}
}

var registry callbackRegistry

// baseCallbackMethodPCs caches the function PC of each BaseCallback method so
// that we can tell at runtime whether the user overrode them in their own
// ConvertCallback type. We have to do this lazily because Go's reflect package
// can't be used in package init in some constrained environments, but a single
// sync.Once gets us a global table cheaply.
var (
	baseCallbackPCs     map[string]uintptr
	baseCallbackPCsOnce sync.Once
)

func loadBaseCallbackPCs() {
	baseCallbackPCsOnce.Do(func() {
		base := reflect.ValueOf(BaseCallback{})
		baseCallbackPCs = map[string]uintptr{}
		for _, name := range []string{
			"OnOCR", "OnLayout", "OnTable",
			"GetOCRResult", "GetLayoutResult", "GetTableResult",
		} {
			m := base.MethodByName(name)
			if m.IsValid() {
				baseCallbackPCs[name] = m.Pointer()
			}
		}
	})
}

// overrideFlags returns the GO_CB_FLAG_* bitmask describing which optional
// methods on cb are NOT the BaseCallback default. Promoted methods from an
// embedded BaseCallback share the same function PC as BaseCallback's own
// method values, so a simple pointer comparison is sufficient. If reflection
// can't resolve a method (e.g. user implemented ConvertCallback directly
// without embedding BaseCallback), we treat it as overridden to preserve
// existing behaviour for that path.
func overrideFlags(cb ConvertCallback) C.int {
	if cb == nil {
		return 0
	}
	loadBaseCallbackPCs()
	v := reflect.ValueOf(cb)
	var flags C.int
	check := func(name string, bit C.int) {
		m := v.MethodByName(name)
		if !m.IsValid() {
			return
		}
		basePC, ok := baseCallbackPCs[name]
		if !ok || m.Pointer() != basePC {
			flags |= bit
		}
	}
	check("OnOCR", C.GO_CB_FLAG_OCR)
	check("OnLayout", C.GO_CB_FLAG_LAYOUT)
	check("OnTable", C.GO_CB_FLAG_TABLE)
	check("GetOCRResult", C.GO_CB_FLAG_GET_OCR_RESULT)
	check("GetLayoutResult", C.GO_CB_FLAG_GET_LAYOUT_RESULT)
	check("GetTableResult", C.GO_CB_FLAG_GET_TABLE_RESULT)
	return flags
}

// registerCallback wraps cb in a handle and a result-buffer record. The
// returned cleanup must be called once the C SDK is guaranteed to no longer
// touch the callback.
func registerCallback(cb ConvertCallback) (C.GoCallbackHandle, func()) {
	if cb == nil {
		return 0, func() {}
	}
	h := cgo.NewHandle(cb)
	buf := &callbackBuffers{}
	registry.results.Store(uintptr(h), buf)
	cleanup := func() {
		registry.results.Delete(uintptr(h))
		buf.free()
		h.Delete()
	}
	return C.GoCallbackHandle(uintptr(h)), cleanup
}

func resolveCallback(h C.GoCallbackHandle) (ConvertCallback, *callbackBuffers) {
	if h == 0 {
		return nil, nil
	}
	val := cgo.Handle(uintptr(h)).Value()
	cb, _ := val.(ConvertCallback)
	bufVal, _ := registry.results.Load(uintptr(h))
	buf, _ := bufVal.(*callbackBuffers)
	return cb, buf
}

/* ===== Exported trampolines invoked from the C bridge ===== */

//export goCBProgress
func goCBProgress(h C.GoCallbackHandle, current, total C.int) {
	cb, _ := resolveCallback(h)
	if cb != nil {
		cb.OnProgress(int(current), int(total))
	}
}

//export goCBCancel
func goCBCancel(h C.GoCallbackHandle) C.bool {
	cb, _ := resolveCallback(h)
	if cb != nil && cb.IsCancelled() {
		return C.bool(true)
	}
	return C.bool(false)
}

//export goCBOcr
func goCBOcr(h C.GoCallbackHandle, imagePath *C.char) C.bool {
	cb, _ := resolveCallback(h)
	if cb == nil {
		return C.bool(false)
	}
	return C.bool(cb.OnOCR(C.GoString(imagePath)))
}

//export goCBLayout
func goCBLayout(h C.GoCallbackHandle, imagePath *C.char) C.bool {
	cb, _ := resolveCallback(h)
	if cb == nil {
		return C.bool(false)
	}
	return C.bool(cb.OnLayout(C.GoString(imagePath)))
}

//export goCBTable
func goCBTable(h C.GoCallbackHandle, imagePath *C.char) C.bool {
	cb, _ := resolveCallback(h)
	if cb == nil {
		return C.bool(false)
	}
	return C.bool(cb.OnTable(C.GoString(imagePath)))
}

//export goCBGetOcrResult
func goCBGetOcrResult(h C.GoCallbackHandle) *C.char {
	cb, buf := resolveCallback(h)
	if cb == nil || buf == nil {
		return nil
	}
	buf.mu.Lock()
	defer buf.mu.Unlock()
	return buf.replace(&buf.ocr, cb.GetOCRResult())
}

//export goCBGetLayoutResult
func goCBGetLayoutResult(h C.GoCallbackHandle) *C.char {
	cb, buf := resolveCallback(h)
	if cb == nil || buf == nil {
		return nil
	}
	buf.mu.Lock()
	defer buf.mu.Unlock()
	return buf.replace(&buf.layout, cb.GetLayoutResult())
}

//export goCBGetTableResult
func goCBGetTableResult(h C.GoCallbackHandle) *C.char {
	cb, buf := resolveCallback(h)
	if cb == nil || buf == nil {
		return nil
	}
	buf.mu.Lock()
	defer buf.mu.Unlock()
	return buf.replace(&buf.tableRes, cb.GetTableResult())
}
