#include "bridge.h"

#include "compdf_basictypes_c.h"
#include "compdf_common_c.h"
#include "compdf_conversion_c.h"

/* cgo-generated prototypes for the //export goCB* Go callbacks. */
#include "_cgo_export.h"

#include <stdlib.h>
#include <string.h>

#if defined(_WIN32)
#include <windows.h>
#endif

/* The C SDK's CProgress / CCancel function pointers do not carry user context,
 * so we use thread-local storage to associate the "active" Go callback handle
 * with the currently-running Start* call on the caller's OS thread.
 *
 * cgo guarantees that the C side of a cgo call executes on the same OS thread
 * as the calling goroutine for the duration of the call, and the SDK's
 * conversion entry points invoke progress/cancel synchronously from that
 * thread. This keeps concurrent conversions from different goroutines fully
 * isolated. */
#if defined(_MSC_VER)
#define BRIDGE_TLS __declspec(thread)
#else
#define BRIDGE_TLS __thread
#endif

static BRIDGE_TLS GoCallbackHandle g_active_handle = 0;

/* ---- Trampolines exposed to the C SDK ---- */
static void bridge_progress(int current, int total) {
    if (g_active_handle != 0) {
        goCBProgress(g_active_handle, current, total);
    }
}

static bool bridge_cancel(void) {
    if (g_active_handle != 0) {
        return goCBCancel(g_active_handle);
    }
    return false;
}

static bool bridge_ocr(const char* image_path) {
    if (g_active_handle != 0) {
        return goCBOcr(g_active_handle, (char*)image_path);
    }
    return false;
}

static bool bridge_layout(const char* image_path) {
    if (g_active_handle != 0) {
        return goCBLayout(g_active_handle, (char*)image_path);
    }
    return false;
}

static bool bridge_table(const char* image_path) {
    if (g_active_handle != 0) {
        return goCBTable(g_active_handle, (char*)image_path);
    }
    return false;
}

static const char* bridge_get_ocr_result(void) {
    if (g_active_handle != 0) {
        return goCBGetOcrResult(g_active_handle);
    }
    return NULL;
}

static const char* bridge_get_layout_result(void) {
    if (g_active_handle != 0) {
        return goCBGetLayoutResult(g_active_handle);
    }
    return NULL;
}

static const char* bridge_get_table_result(void) {
    if (g_active_handle != 0) {
        return goCBGetTableResult(g_active_handle);
    }
    return NULL;
}

/* ---- Path encoding helpers (UTF-8 in -> platform native string out) ---- */
#if defined(_WIN32)
static wchar_t* utf8_to_wide(const char* utf8) {
    if (utf8 == NULL) return NULL;
    int len = MultiByteToWideChar(CP_UTF8, 0, utf8, -1, NULL, 0);
    if (len <= 0) return NULL;
    wchar_t* buf = (wchar_t*)malloc((size_t)len * sizeof(wchar_t));
    if (!buf) return NULL;
    MultiByteToWideChar(CP_UTF8, 0, utf8, -1, buf, len);
    return buf;
}
#define FREE_PATH(p) free((void*)(p))
#else
#define FREE_PATH(p) ((void)0)
#endif

/* ---- Helper to fill the SDK's CConvertOption from the POD GoConvertOption ---- */
static void fill_convert_option(CConvertOption* dst, const GoConvertOption* src, COCRLanguage* lang_buf, int lang_buf_size) {
    memset(dst, 0, sizeof(*dst));
    dst->enable_ai_layout = src->enable_ai_layout;
    dst->enable_ai_table_recognition = src->enable_ai_table_recognition;
    dst->contain_image = src->contain_image;
    dst->contain_page_background_image = src->contain_page_background_image;
    dst->json_contain_table = src->json_contain_table;
    dst->contain_annotation = src->contain_annotation;
    dst->excel_all_content = src->excel_all_content;
    dst->excel_csv_format = src->excel_csv_format;
    dst->enable_ocr = src->enable_ocr;
    dst->transparent_text = src->transparent_text;
    dst->txt_table_format = src->txt_table_format;
    dst->image_path_enhance = src->image_path_enhance;
    dst->formula_to_image = src->formula_to_image;
    dst->auto_create_folder = src->auto_create_folder;
    dst->output_document_per_page = src->output_document_per_page;
    dst->image_scaling = src->image_scaling;
    dst->page_layout_mode = (CPageLayoutMode)src->page_layout_mode;
    dst->excel_worksheet_option = (CExcelWorksheetOption)src->excel_worksheet_option;
    dst->html_option = (CHtmlOption)src->html_option;
    dst->ocr_option = (COCROption)src->ocr_option;
    dst->image_color_mode = (CImageColorMode)src->image_color_mode;
    dst->image_type = (CImageType)src->image_type;

    strncpy(dst->font_name, src->font_name, sizeof(dst->font_name) - 1);
    dst->font_name[sizeof(dst->font_name) - 1] = '\0';
    strncpy(dst->page_ranges, src->page_ranges, sizeof(dst->page_ranges) - 1);
    dst->page_ranges[sizeof(dst->page_ranges) - 1] = '\0';

    int n = src->language_count;
    if (n > lang_buf_size) n = lang_buf_size;
    for (int i = 0; i < n; ++i) {
        lang_buf[i] = (COCRLanguage)src->languages[i];
    }
    dst->languages = lang_buf;
    dst->language_count = n;
    dst->enable_document_orientation_classification = src->enable_document_orientation_classification;
    dst->enable_document_dewarp = src->enable_document_dewarp;
}

/* Build a populated CConvertCallback. Optional handlers (ocr/layout/table
 * and their get_*_result counterparts) are wired only when the Go side
 * reports that the user actually overrode them via `callback_flags`. This
 * lets the native AI pipeline kick in when the caller didn't supply a real
 * implementation, instead of being short-circuited by a no-op default. */
static void build_callback(CConvertCallback* cb, int callback_flags) {
    memset(cb, 0, sizeof(*cb));
    cb->handle = NULL;
    cb->progress = bridge_progress;
    cb->cancel = bridge_cancel;
    cb->ocr               = (callback_flags & GO_CB_FLAG_OCR)              ? bridge_ocr               : NULL;
    cb->layout            = (callback_flags & GO_CB_FLAG_LAYOUT)           ? bridge_layout            : NULL;
    cb->table             = (callback_flags & GO_CB_FLAG_TABLE)            ? bridge_table             : NULL;
    cb->get_ocr_result    = (callback_flags & GO_CB_FLAG_GET_OCR_RESULT)   ? bridge_get_ocr_result    : NULL;
    cb->get_layout_result = (callback_flags & GO_CB_FLAG_GET_LAYOUT_RESULT)? bridge_get_layout_result : NULL;
    cb->get_table_result  = (callback_flags & GO_CB_FLAG_GET_TABLE_RESULT) ? bridge_get_table_result  : NULL;
}

/* Function pointer type matching every CPDF_StartPDFToXxx entry point. */
typedef CSDKErrorCode (CSDK_COMPDFKIT_CALL *StartFunc)(COMPDFKIT_STRING, COMPDFKIT_STRING, COMPDFKIT_STRING, CConvertOption, CConvertCallback*);

#define MAX_OCR_LANGUAGES 32

static int dispatch_convert(StartFunc fn, const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb_handle, int callback_flags) {
    CConvertOption sdk_opt;
    COCRLanguage lang_buf[MAX_OCR_LANGUAGES];
    fill_convert_option(&sdk_opt, options, lang_buf, MAX_OCR_LANGUAGES);

    CConvertCallback cb;
    build_callback(&cb, cb_handle ? callback_flags : 0);

    GoCallbackHandle previous = g_active_handle;
    g_active_handle = cb_handle;

#if defined(_WIN32)
    wchar_t* w_file = utf8_to_wide(file_path);
    wchar_t* w_pwd = utf8_to_wide(password ? password : "");
    wchar_t* w_out = utf8_to_wide(output_path);
    CSDKErrorCode rc = fn(w_file, w_pwd, w_out, sdk_opt, &cb);
    FREE_PATH(w_file);
    FREE_PATH(w_pwd);
    FREE_PATH(w_out);
#else
    CSDKErrorCode rc = fn(file_path, password ? password : "", output_path, sdk_opt, &cb);
#endif

    g_active_handle = previous;
    return (int)rc;
}

/* ===== Public API exposed to Go ===== */

int compdf_go_license_verify(const char* license, const char* device_id, const char* app_id) {
    return (int)CPDF_LicenseVerify(license ? license : "", device_id ? device_id : "", app_id ? app_id : "");
}

void compdf_go_initialize(const char* resource_path) {
#if defined(_WIN32)
    wchar_t* w_path = utf8_to_wide(resource_path ? resource_path : "");
    CPDF_Initialize(w_path);
    FREE_PATH(w_path);
#else
    CPDF_Initialize(resource_path ? resource_path : "");
#endif
}

void compdf_go_release(void) {
    CPDF_Release();
}

void compdf_go_set_logger(bool enable_info, bool enable_warning) {
    CPDF_SetLogger(enable_info, enable_warning);
}

int compdf_go_set_document_ai_model(const char* model_path, int gpu_id) {
#if defined(_WIN32)
    wchar_t* w_path = utf8_to_wide(model_path ? model_path : "");
    int rc = (int)CPDF_SetDocumentAIModel(w_path, gpu_id);
    FREE_PATH(w_path);
    return rc;
#else
    return (int)CPDF_SetDocumentAIModel(model_path ? model_path : "", gpu_id);
#endif
}

void compdf_go_release_document_ai_model(void) {
    CPDF_ReleaseDocumentAIModel();
}

void compdf_go_set_document_ai_model_count(int layout_model_count, int table_model_count) {
    CPDF_SetDocumentAIModelCount(layout_model_count, table_model_count);
}

int compdf_go_get_page_count(const char* file_path, const char* password) {
#if defined(_WIN32)
    wchar_t* w_file = utf8_to_wide(file_path ? file_path : "");
    wchar_t* w_pwd = utf8_to_wide(password ? password : "");
    int rc = CPDF_GetPageCount(w_file, w_pwd);
    FREE_PATH(w_file);
    FREE_PATH(w_pwd);
    return rc;
#else
    return CPDF_GetPageCount(file_path ? file_path : "", password ? password : "");
#endif
}

int compdf_go_get_remaining_page_quota(void) {
    return CPDF_GetRemainingPageQuota();
}

void compdf_go_get_version(char* version_buf) {
    CPDF_GetVersion(version_buf);
}

int compdf_go_pdf_to_word(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags) {
    return dispatch_convert(CPDF_StartPDFToWord, file_path, password, output_path, options, cb, callback_flags);
}
int compdf_go_pdf_to_excel(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags) {
    return dispatch_convert(CPDF_StartPDFToExcel, file_path, password, output_path, options, cb, callback_flags);
}
int compdf_go_pdf_to_ppt(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags) {
    return dispatch_convert(CPDF_StartPDFToPpt, file_path, password, output_path, options, cb, callback_flags);
}
int compdf_go_pdf_to_html(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags) {
    return dispatch_convert(CPDF_StartPDFToHtml, file_path, password, output_path, options, cb, callback_flags);
}
int compdf_go_pdf_to_image(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags) {
    return dispatch_convert(CPDF_StartPDFToImage, file_path, password, output_path, options, cb, callback_flags);
}
int compdf_go_pdf_to_markdown(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags) {
    return dispatch_convert(CPDF_StartPDFToMarkdown, file_path, password, output_path, options, cb, callback_flags);
}
int compdf_go_pdf_to_rtf(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags) {
    return dispatch_convert(CPDF_StartPDFToRtf, file_path, password, output_path, options, cb, callback_flags);
}
int compdf_go_pdf_to_txt(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags) {
    return dispatch_convert(CPDF_StartPDFToTxt, file_path, password, output_path, options, cb, callback_flags);
}
int compdf_go_pdf_to_json(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags) {
    return dispatch_convert(CPDF_StartPDFToJson, file_path, password, output_path, options, cb, callback_flags);
}
int compdf_go_pdf_to_searchable_pdf(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags) {
    return dispatch_convert(CPDF_StartPDFToSearchablePDF, file_path, password, output_path, options, cb, callback_flags);
}
int compdf_go_pdf_to_ofd(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags) {
    return dispatch_convert(CPDF_StartPDFToOfd, file_path, password, output_path, options, cb, callback_flags);
}
