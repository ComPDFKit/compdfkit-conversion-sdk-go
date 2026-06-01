// Package compdf provides Go language bindings for the ComPDFKit Conversion SDK.
//
// It wraps the C SDK headers in compdf_sdk/golang/include/ via cgo and exposes an idiomatic
// Go API mirroring the Java/Python bindings. Typical usage:
//
//	if err := compdf.LicenseVerify(license, "", appID); err != nil {
//	    log.Fatal(err)
//	}
//	compdf.Initialize("./resources")
//	defer compdf.Release()
//
//	opts := compdf.NewWordOptions()
//	opts.EnableOCR = false
//	if err := compdf.StartPDFToWord("in.pdf", "", "./out", opts, nil); err != nil {
//	    log.Fatal(err)
//	}
package compdf
