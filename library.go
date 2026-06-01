package compdf

/*
#include <stdlib.h>
#include "bridge.h"
*/
import "C"

import (
	"unsafe"
)

// LicenseVerify verifies the SDK license. Pass an empty deviceID to let the
// SDK resolve a device id automatically (or supply your own).
func LicenseVerify(license, deviceID, appID string) error {
	cLicense := C.CString(license)
	cDevice := C.CString(deviceID)
	cApp := C.CString(appID)
	defer C.free(unsafe.Pointer(cLicense))
	defer C.free(unsafe.Pointer(cDevice))
	defer C.free(unsafe.Pointer(cApp))
	rc := C.compdf_go_license_verify(cLicense, cDevice, cApp)
	return ErrorCode(rc).toError()
}

// Initialize prepares the SDK with the given resource directory. Must be
// called once before any conversion API and after a successful LicenseVerify.
func Initialize(resourcePath string) {
	cPath := C.CString(resourcePath)
	defer C.free(unsafe.Pointer(cPath))
	C.compdf_go_initialize(cPath)
}

// Release tears down all SDK resources. Pair with Initialize.
func Release() {
	C.compdf_go_release()
}

// SetLogger toggles info/warning logging emitted by the SDK.
func SetLogger(enableInfo, enableWarning bool) {
	C.compdf_go_set_logger(C.bool(enableInfo), C.bool(enableWarning))
}

// SetDocumentAIModel installs the document AI model used for OCR/layout/table
// recognition. gpuID = -1 disables GPU acceleration.
func SetDocumentAIModel(modelPath string, gpuID int) error {
	cPath := C.CString(modelPath)
	defer C.free(unsafe.Pointer(cPath))
	rc := C.compdf_go_set_document_ai_model(cPath, C.int(gpuID))
	return ErrorCode(rc).toError()
}

// ReleaseDocumentAIModel frees memory held by the AI model. Call
// SetDocumentAIModel again before reusing AI features.
func ReleaseDocumentAIModel() {
	C.compdf_go_release_document_ai_model()
}

// SetDocumentAIModelCount tunes the number of layout/table-recognition model
// instances loaded.
func SetDocumentAIModelCount(layoutModelCount, tableModelCount int) {
	C.compdf_go_set_document_ai_model_count(C.int(layoutModelCount), C.int(tableModelCount))
}

// GetPageCount returns the page count of filePath, or a negative value on
// failure. Pass an empty password for non-encrypted documents.
func GetPageCount(filePath, password string) int {
	cFile := C.CString(filePath)
	cPwd := C.CString(password)
	defer C.free(unsafe.Pointer(cFile))
	defer C.free(unsafe.Pointer(cPwd))
	return int(C.compdf_go_get_page_count(cFile, cPwd))
}

// GetRemainingPageQuota returns the remaining page quota under the active
// license. See the C header for the special negative values.
func GetRemainingPageQuota() int {
	return int(C.compdf_go_get_remaining_page_quota())
}

// GetVersion returns the SDK version string.
func GetVersion() string {
	buf := make([]C.char, 64)
	C.compdf_go_get_version(&buf[0])
	return C.GoString(&buf[0])
}
