#ifndef COMPDF_GO_BRIDGE_H
#define COMPDF_GO_BRIDGE_H

/* This bridge layer wraps the C SDK so the Go cgo code never has to deal
 * with the platform-specific string type (wchar_t* on Windows, char* elsewhere)
 * defined in compdf_basictypes_c.h. The bridge always takes/returns UTF-8 char*
 * pointers and converts them internally when targeting Windows. */

#include <stdbool.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

/* ----- Enums kept binary-compatible with the C SDK ----- */
typedef enum {
    GO_PAGE_LAYOUT_BOX = 0,
    GO_PAGE_LAYOUT_FLOW = 1
} GoPageLayoutMode;

typedef enum {
    GO_OCR_INVALID_CHAR = 0,
    GO_OCR_SCAN_PAGE = 1,
    GO_OCR_INVALID_CHAR_AND_SCAN = 2,
    GO_OCR_ALL = 3
} GoOCROption;

typedef enum {
    GO_IMAGE_COLOR = 0,
    GO_IMAGE_GRAY = 1,
    GO_IMAGE_BINARY = 2
} GoImageColorMode;

typedef enum {
    GO_IMG_JPG = 0,
    GO_IMG_JPEG = 1,
    GO_IMG_JPEG2K = 2,
    GO_IMG_PNG = 3,
    GO_IMG_BMP = 4,
    GO_IMG_TIFF = 5,
    GO_IMG_TGA = 6,
    GO_IMG_GIF = 7,
    GO_IMG_WEBP = 8
} GoImageType;

typedef enum {
    GO_EXCEL_FOR_TABLE = 0,
    GO_EXCEL_FOR_PAGE = 1,
    GO_EXCEL_FOR_DOCUMENT = 2
} GoExcelWorksheetOption;

typedef enum {
    GO_HTML_SINGLE_PAGE = 0,
    GO_HTML_SINGLE_PAGE_WITH_BOOKMARK = 1,
    GO_HTML_MULTI_PAGE = 2,
    GO_HTML_MULTI_PAGE_WITH_BOOKMARK = 3
} GoHtmlPageOption;

/* ----- Mirror of CConvertOption using POD-only fields. The Go side fills this
 *       in and the bridge translates it to CConvertOption. ----- */
typedef struct {
    bool enable_ai_layout;
    bool enable_ai_table_recognition;
    bool contain_image;
    bool contain_page_background_image;
    bool json_contain_table;
    bool contain_annotation;
    bool excel_all_content;
    bool excel_csv_format;
    bool enable_ocr;
    bool transparent_text;
    bool txt_table_format;
    bool image_path_enhance;
    bool formula_to_image;
    bool auto_create_folder;
    bool output_document_per_page;
    float image_scaling;
    int page_layout_mode;
    int excel_worksheet_option;
    int html_option;
    int ocr_option;
    int image_color_mode;
    int image_type;
    char font_name[256];
    char page_ranges[256];

    /* OCR languages: pointer to int array + length. The bridge copies into the
     * COCRLanguage array that the C SDK expects. */
    const int* languages;
    int language_count;
} GoConvertOption;

/* Opaque handle used to identify a Go-side callback in the registry. 0 means
 * "no callback". */
typedef uintptr_t GoCallbackHandle;

/* Bit flags that tell the bridge which optional callback methods were
 * actually overridden on the Go side. Methods left at their default
 * (no-op) BaseCallback implementation MUST be reported as unset, so the
 * bridge can leave the matching C SDK function pointers null and the
 * native code falls back to its internal AI pipeline instead of calling
 * out to a default that does nothing. Progress/cancel are always wired. */
#define GO_CB_FLAG_OCR              (1 << 0)
#define GO_CB_FLAG_LAYOUT           (1 << 1)
#define GO_CB_FLAG_TABLE            (1 << 2)
#define GO_CB_FLAG_GET_OCR_RESULT   (1 << 3)
#define GO_CB_FLAG_GET_LAYOUT_RESULT (1 << 4)
#define GO_CB_FLAG_GET_TABLE_RESULT (1 << 5)

/* ----- Common SDK functions ----- */
int  compdf_go_license_verify(const char* license, const char* device_id, const char* app_id);
void compdf_go_initialize(const char* resource_path);
void compdf_go_release(void);
void compdf_go_set_logger(bool enable_info, bool enable_warning);
int  compdf_go_set_document_ai_model(const char* model_path, int gpu_id);
void compdf_go_release_document_ai_model(void);
void compdf_go_set_document_ai_model_count(int layout_model_count, int table_model_count);
int  compdf_go_get_page_count(const char* file_path, const char* password);
int  compdf_go_get_remaining_page_quota(void);
/* version_buf must be at least 64 bytes. */
void compdf_go_get_version(char* version_buf);

/* ----- Conversion functions. All accept UTF-8 paths. ----- */
int compdf_go_pdf_to_word(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags);
int compdf_go_pdf_to_excel(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags);
int compdf_go_pdf_to_ppt(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags);
int compdf_go_pdf_to_html(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags);
int compdf_go_pdf_to_image(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags);
int compdf_go_pdf_to_markdown(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags);
int compdf_go_pdf_to_rtf(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags);
int compdf_go_pdf_to_txt(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags);
int compdf_go_pdf_to_json(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags);
int compdf_go_pdf_to_searchable_pdf(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags);
int compdf_go_pdf_to_ofd(const char* file_path, const char* password, const char* output_path, const GoConvertOption* options, GoCallbackHandle cb, int callback_flags);

/* The goCB* trampolines are implemented in Go and exposed by cgo via the
 * generated _cgo_export.h header. bridge.c includes that header directly so
 * we deliberately do NOT declare them here (avoids a duplicate-declaration
 * mismatch on the GoCallbackHandle type cgo synthesizes). */

#ifdef __cplusplus
}
#endif

#endif /* COMPDF_GO_BRIDGE_H */
