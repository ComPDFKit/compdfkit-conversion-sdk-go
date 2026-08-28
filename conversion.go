package compdf

/*
#include <stdlib.h>
#include "bridge.h"
*/
import "C"

import (
	"runtime"
	"unsafe"
)

// converter is the function-pointer type bound by each Start* function. Keeps
// the per-format wrappers free of duplication.
type converter func(file, pwd, out *C.char, opt *C.GoConvertOption, cb C.GoCallbackHandle, flags C.int) C.int

// runConvert builds the C-side option struct, registers the callback if any,
// invokes the underlying SDK function and translates the result code.
func runConvert(fn converter, filePath, password, outputPath string, opts convertOptions, cb ConvertCallback) error {
	cFile := C.CString(filePath)
	cPwd := C.CString(password)
	cOut := C.CString(outputPath)
	defer C.free(unsafe.Pointer(cFile))
	defer C.free(unsafe.Pointer(cPwd))
	defer C.free(unsafe.Pointer(cOut))

	var cOpt C.GoConvertOption
	cOpt.enable_ai_layout = C.bool(opts.EnableAILayout)
	cOpt.enable_ai_table_recognition = C.bool(opts.EnableAITableRecognition)
	cOpt.contain_image = C.bool(opts.ContainImage)
	cOpt.contain_page_background_image = C.bool(opts.ContainPageBackgroundImage)
	cOpt.json_contain_table = C.bool(opts.JSONContainTable)
	cOpt.contain_annotation = C.bool(opts.ContainAnnotation)
	cOpt.excel_all_content = C.bool(opts.ExcelAllContent)
	cOpt.excel_csv_format = C.bool(opts.ExcelCSVFormat)
	cOpt.enable_ocr = C.bool(opts.EnableOCR)
	cOpt.transparent_text = C.bool(opts.TransparentText)
	cOpt.txt_table_format = C.bool(opts.TxtTableFormat)
	cOpt.image_path_enhance = C.bool(opts.ImagePathEnhance)
	cOpt.formula_to_image = C.bool(opts.FormulaToImage)
	cOpt.auto_create_folder = C.bool(opts.AutoCreateFolder)
	cOpt.output_document_per_page = C.bool(opts.OutputDocumentPerPage)
	cOpt.image_scaling = C.float(opts.ImageScaling)
	cOpt.page_layout_mode = C.int(opts.PageLayoutMode)
	cOpt.excel_worksheet_option = C.int(opts.ExcelWorksheetOption)
	cOpt.html_option = C.int(opts.HtmlPageOption)
	cOpt.ocr_option = C.int(opts.OCROption)
	cOpt.image_color_mode = C.int(opts.ImageColorMode)
	cOpt.image_type = C.int(opts.ImageType)

	writeFixedString(&cOpt.font_name[0], len(cOpt.font_name), opts.FontName)
	writeFixedString(&cOpt.page_ranges[0], len(cOpt.page_ranges), opts.PageRanges)

	var pinner runtime.Pinner
	defer pinner.Unpin()

	var langs []C.int
	if len(opts.Languages) > 0 {
		langs = make([]C.int, len(opts.Languages))
		for i, l := range opts.Languages {
			langs[i] = C.int(l)
		}
		pinner.Pin(&langs[0])
		cOpt.languages = (*C.int)(unsafe.Pointer(&langs[0]))
		cOpt.language_count = C.int(len(langs))
	}
	cOpt.enable_document_orientation_classification = C.bool(opts.EnableDocumentOrientationClassification)
	cOpt.enable_document_dewarp = C.bool(opts.EnableDocumentDewarp)

	handle, cleanup := registerCallback(cb)
	defer cleanup()

	rc := fn(cFile, cPwd, cOut, &cOpt, handle, overrideFlags(cb))
	return ErrorCode(rc).toError()
}

func writeFixedString(buf *C.char, capacity int, s string) {
	dst := unsafe.Slice((*byte)(unsafe.Pointer(buf)), capacity)
	for i := range dst {
		dst[i] = 0
	}
	n := len(s)
	if n > capacity-1 {
		n = capacity - 1
	}
	copy(dst[:n], s)
}

// StartPDFToWord converts a PDF document to a Word document.
func StartPDFToWord(filePath, password, outputPath string, opts WordOptions, cb ConvertCallback) error {
	return runConvert(func(f, p, o *C.char, opt *C.GoConvertOption, h C.GoCallbackHandle, fl C.int) C.int {
		return C.compdf_go_pdf_to_word(f, p, o, opt, h, fl)
	}, filePath, password, outputPath, opts.toConvert(), cb)
}

// StartPDFToExcel converts a PDF document to an Excel document.
func StartPDFToExcel(filePath, password, outputPath string, opts ExcelOptions, cb ConvertCallback) error {
	return runConvert(func(f, p, o *C.char, opt *C.GoConvertOption, h C.GoCallbackHandle, fl C.int) C.int {
		return C.compdf_go_pdf_to_excel(f, p, o, opt, h, fl)
	}, filePath, password, outputPath, opts.toConvert(), cb)
}

// StartPDFToPpt converts a PDF document to a PowerPoint presentation.
func StartPDFToPpt(filePath, password, outputPath string, opts PptOptions, cb ConvertCallback) error {
	return runConvert(func(f, p, o *C.char, opt *C.GoConvertOption, h C.GoCallbackHandle, fl C.int) C.int {
		return C.compdf_go_pdf_to_ppt(f, p, o, opt, h, fl)
	}, filePath, password, outputPath, opts.toConvert(), cb)
}

// StartPDFToHtml converts a PDF document to HTML.
func StartPDFToHtml(filePath, password, outputPath string, opts HtmlOptions, cb ConvertCallback) error {
	return runConvert(func(f, p, o *C.char, opt *C.GoConvertOption, h C.GoCallbackHandle, fl C.int) C.int {
		return C.compdf_go_pdf_to_html(f, p, o, opt, h, fl)
	}, filePath, password, outputPath, opts.toConvert(), cb)
}

// StartPDFToImage converts each PDF page to an image.
func StartPDFToImage(filePath, password, outputPath string, opts ImageOptions, cb ConvertCallback) error {
	return runConvert(func(f, p, o *C.char, opt *C.GoConvertOption, h C.GoCallbackHandle, fl C.int) C.int {
		return C.compdf_go_pdf_to_image(f, p, o, opt, h, fl)
	}, filePath, password, outputPath, opts.toConvert(), cb)
}

// StartPDFToMarkdown converts a PDF document to Markdown.
func StartPDFToMarkdown(filePath, password, outputPath string, opts MarkdownOptions, cb ConvertCallback) error {
	return runConvert(func(f, p, o *C.char, opt *C.GoConvertOption, h C.GoCallbackHandle, fl C.int) C.int {
		return C.compdf_go_pdf_to_markdown(f, p, o, opt, h, fl)
	}, filePath, password, outputPath, opts.toConvert(), cb)
}

// StartPDFToRtf converts a PDF document to RTF.
func StartPDFToRtf(filePath, password, outputPath string, opts RtfOptions, cb ConvertCallback) error {
	return runConvert(func(f, p, o *C.char, opt *C.GoConvertOption, h C.GoCallbackHandle, fl C.int) C.int {
		return C.compdf_go_pdf_to_rtf(f, p, o, opt, h, fl)
	}, filePath, password, outputPath, opts.toConvert(), cb)
}

// StartPDFToTxt converts a PDF document to plain text.
func StartPDFToTxt(filePath, password, outputPath string, opts TxtOptions, cb ConvertCallback) error {
	return runConvert(func(f, p, o *C.char, opt *C.GoConvertOption, h C.GoCallbackHandle, fl C.int) C.int {
		return C.compdf_go_pdf_to_txt(f, p, o, opt, h, fl)
	}, filePath, password, outputPath, opts.toConvert(), cb)
}

// StartPDFToJson converts a PDF document to structured JSON.
func StartPDFToJson(filePath, password, outputPath string, opts JsonOptions, cb ConvertCallback) error {
	return runConvert(func(f, p, o *C.char, opt *C.GoConvertOption, h C.GoCallbackHandle, fl C.int) C.int {
		return C.compdf_go_pdf_to_json(f, p, o, opt, h, fl)
	}, filePath, password, outputPath, opts.toConvert(), cb)
}

// StartPDFToSearchablePdf converts a PDF document to a searchable PDF.
func StartPDFToSearchablePdf(filePath, password, outputPath string, opts SearchablePdfOptions, cb ConvertCallback) error {
	return runConvert(func(f, p, o *C.char, opt *C.GoConvertOption, h C.GoCallbackHandle, fl C.int) C.int {
		return C.compdf_go_pdf_to_searchable_pdf(f, p, o, opt, h, fl)
	}, filePath, password, outputPath, opts.toConvert(), cb)
}

// StartPDFToOfd converts a PDF document to OFD.
func StartPDFToOfd(filePath, password, outputPath string, opts OfdOptions, cb ConvertCallback) error {
	return runConvert(func(f, p, o *C.char, opt *C.GoConvertOption, h C.GoCallbackHandle, fl C.int) C.int {
		return C.compdf_go_pdf_to_ofd(f, p, o, opt, h, fl)
	}, filePath, password, outputPath, opts.toConvert(), cb)
}
