#ifndef COMPDFKIT_BASICTYPES_C_H
#define COMPDFKIT_BASICTYPES_C_H

#ifdef __cplusplus
extern "C" {
#endif

#include <stdbool.h>

#include "com_config_macros.h"

#if defined(_WIN32)
#define COMPDFKIT_STRING const wchar_t*
#define CSDK_COMPDFKIT_CALL __stdcall
#else
#define COMPDFKIT_STRING const char*
#define CSDK_COMPDFKIT_CALL
#endif  // _WIN32

/// \brief Type of the currently verified license.
typedef enum _CLicenseType {
    e_CLicenseUnknown = 0,
    e_CLicenseTrial = 1,
    e_CLicenseFormal = 2
} CLicenseType;

/// \brief Error code.
typedef enum _CSDKErrorCode {
    /// Success, and no error occurs.
    e_CErrSuccess = 0,
    /// Conversion process was canceled.
    e_CErrCancel = 1,
    /// File cannot be found or could not be opened.
    e_CErrFile = 2,
    /// Invalid password. Usually, this error may occur when loading a PDF document with password. When meet this, user should load document again with correct password.
    e_CErrPDFPassword = 3,
    /// PDF page failed to load.
    e_CErrPDFPage = 4,
    /// Format is invalid. For files, this may also mean that file is corrupted.
    e_CErrPDFFormat = 5,
    /// PDF document is encrypted by some unsupported security handler.
    e_CErrPDFSecurity = 6,
    /// Out-of-memory error occurs.
    e_CErrOutOfMemory = 7,
    /// System I/O error.
    e_CErrIO = 8,
    /// Folder compression failed.
    e_CErrCompress = 9,

    /// The license is invalid
    e_CErrLicenseInvalid = 20,
    /// The license has expired
    e_CErrLicenseExpire = 21,
    /// The license does not support the current platform
    e_CErrLicenseUnsupportedPlatform = 22,
    /// The license does not support the application id
    e_CErrLicenseUnsupportedID = 23,
    /// The license does not support the device id
    e_CErrLicenseUnsupportedDevice = 24,
    /// The license does not have the function permission
    e_CErrLicensePermissionDeny = 25,
    /// License has not been initialized
    e_CErrLicenseUninitialized = 26,
    /// Illegal access to the API interface
    e_CErrLicenseIllegalAccess = 27,
    /// Failed to read license file
    e_CErrLicenseFileReadFailed = 28,
    /// The license does not have OCR permissions.
    e_CErrLicenseOCRPermissionDeny = 29,
    /// The current conversion concurrency exceeds the license limit.
    e_CErrLicenseConcurrencyExceeded = 30,
    /// The current conversion pages exceed the license limit.
    e_CErrLicensePageLimitExceeded = 31,
    /// The quota persistence file has been tampered with.
    e_CErrLicenseQuotaCorrupted = 32,

    /// No tables found in the source file.
    e_CErrNoTable = 40,
    /// Failed to call OCR recognition
    e_CErrOCRFailure = 41,

    /// Currently executing conversion task.
    e_CErrConverting = 60,

    /// Invalid argument provided.
    e_CErrInvalidArg = 80,
    /// The handle is invalid, possibly uninitialized or already released.
    e_CErrInvalidHandle = 81,
    /// The model file format is invalid or corrupted.
    e_CErrModelInvalidFormat = 82,
    /// The model does not support the requested function or operation.
    e_CErrModelFunctionUnsupported = 83,
    /// The model format is not supported by this system or SDK.
    e_CErrModelFormatUnsupported = 84,
    /// The model is incompatible with the current SDK version.
    e_CErrModelSDKMismatch = 85,
    /// The image data is empty or not provided.
    e_CErrImageDataEmpty = 86,
    /// The image width or height is invalid (e.g., less than or equal to zero).
    e_CErrImageWHError = 87,
    /// The image format is not supported.
    e_CErrImageUnsupportedFormat = 88,
    /// The image is invalid or corrupted.
    e_CErrImageInvalid = 89,
    /// The resource (e.g., license, token) has expired.
    e_CErrExpire = 90,
    /// A required argument is missing.
    e_CErrMissingArg = 91,
    /// The current license does not support calling this API.
    e_CErrLicenseUnsupportedAPI = 92,
    /// The license does not match the device, module, or version.
    e_CErrLicenseMismatch = 93,
    /// The table data is invalid or null.
    e_CErrInvalidTable = 94,

    /// The source document uses a feature this SDK cannot fully reproduce,
    /// or the requested output cannot carry the input's content. A usable
    /// output may still have been written: Word to PDF normally still
    /// produces the PDF and reports this in its optional diagnostics.
    /// Distinct from e_CErrIO: nothing failed on the disk or in the writer.
    e_CErrUnsupportedFeature = 95,

    /// Unknown error
    e_CErrUnknown = 100
}CSDKErrorCode;

/// \brief OCR language.
typedef enum _COCRLanguage{
    /// Unknown language.
    e_CUNKNOWN = 0,
    /// Chinese (Simplified).
    e_CCHINESE,
    /// Chinese (Traditional).
    e_CCHINESE_TRA,
    /// English.
    e_CENGLISH,
    /// Korean.
    e_CKOREAN,
    /// Japanese.
    e_CJAPANESE,
    /// Latin.
    e_CLATIN,
    /// Devanagari.
    e_CDEVANAGARI,
    /// Cyrillic.
    e_CCYRILLIC,
    /// Arabic.
    e_CARABIC,
    /// Tamil.
    e_CTAMIL,
    /// Telugu.
    e_CTELUGU,
    /// Kannada.
    e_CKANNADA,
    /// THAI.
    e_CTHAI,
    /// GREEK.
    e_CGREEK,
    /// Eslav.
    e_CEslav,
    /// Automatically select language.
    e_CAUTO
} COCRLanguage;

/// \brief Page layout mode.
typedef enum _CPageLayoutMode{
    /// Box mode.
    e_PageLayoutModeBox,
    /// Flowing mode.
    e_PageLayoutModeFlow
} CPageLayoutMode;

/// \brief OCR option.
typedef enum _COCROption
{
    /// Recognize illegal characters in PDF documents.
    e_CInvalidCharacter,
    /// Recognize scanned pages in PDF documents.
    e_CScanPage,
    /// Recognize illegal characters and scanned pages in PDF documents.
    e_CInvalidCharacterAndScanPage,
    /// Recognize all characters on all pages.
    e_CAll,
} COCROption;

/// \brief DocumentAI table recognition model profile.
typedef enum _CTableRecognitionAccuracy
{
    /// Prioritize recognition accuracy; may use more resources.
    e_CHighAccuracy = 0,
    /// Balanced recognition accuracy, resource usage, and throughput.
    e_CBalanced = 1,
} CTableRecognitionAccuracy;

/// \brief Image color mode.
typedef enum _CImageColorMode
{
    /// Color mode.
    e_CColor,
    /// Gray mode.
    e_CGray,
    /// Binary mode.
    e_CBinary,
} CImageColorMode;

/// \brief Image type.
typedef enum _CImageType
{
    /// .jpg.
    e_CJPG,
    /// .jpeg.
    e_CJPEG,
    /// .jp2.
    e_CJPEG2000,
    /// .png.
    e_CPNG,
    /// .bmp.
    e_CBMP,
    /// .tiff.
    e_CTIFF,
    /// .tga.
    e_CTGA,
    /// .gif.
    e_CGIF,
    /// .webp
    e_CWEBP,
} CImageType;

/// \brief Excel Worksheet option.
typedef enum _CExcelWorksheetOption
{
    /// A worksheet to contain only one table.
    e_CForTable,
    /// A worksheet to contain table for PDF Page.
    e_CForPage,
    /// A worksheet to contain table for PDF Document.
    e_CForDocument
} CExcelWorksheetOption;

/// \brief Html option.
typedef enum _CHtmlOption
{
    /// Convert the entire PDF file into a single HTML file.
    e_CSinglePage,
    /// Convert the PDF file into a single HTML file with an outline for navigation at the beginning of the HTML page.
    e_CSinglePageWithBookmark,
    /// Convert the PDF file into multiple HTML files.
    e_CMultiPage,
    /// Convert the PDF file into multiple HTML files. Each HTML file corresponds to a PDF page, and users can navigate to the next HTML file via a link at the bottom of the HTML page.
    e_CMultiPageWithBookmark,
} CHtmlOption;

/// \brief How the PDF bookmarks (the outline shown in a reader's navigation
/// panel) are generated when converting a Word document to PDF.
typedef enum _CWordPdfBookmarkMode {
    /// Do not generate bookmarks.
    e_CWordPdfBookmarkNone = 0,
    /// One bookmark per heading, that is per paragraph with an outline level,
    /// nested by level.
    e_CWordPdfBookmarkHeadings,
    /// One bookmark per Word bookmark, named after it.
    e_CWordPdfBookmarkWordBookmarks,
} CWordPdfBookmarkMode;

/// \brief Convert options.
typedef struct _CConvertOption{
    /// Whether to enable AI layout analysis during conversion.
    bool enable_ai_layout;

    /// Whether to enable AI table recognition during conversion.
    bool enable_ai_table_recognition;

    /// Whether to contain images when converting,which takes effect only when enable_ocr is false.
    bool contain_image;

    /// Whether to contain page background images when converting,which takes effect only when enable_ocr is true.
    bool contain_page_background_image;

    /// Whether to contain table when convert pdf to json.
    bool json_contain_table;

    /// Whether to contain annotations when converting.
    bool contain_annotation;

    /// Whether to set contain all pdf content to the xlsx file.
    bool excel_all_content;

    /// Whether to set to save the table in csv format.
    bool excel_csv_format;

    /// Whether to use OCR.
    bool enable_ocr;

    /// Whether to make the text transparent in the output document.
    bool transparent_text;

    /// Whether format table when convert pdf to txt.
    bool txt_table_format;

    /// Whether to enhance the image path.
    bool image_path_enhance;

    /// Whether to convert formulas to images.
    bool formula_to_image;

    /// Whether to automatically create folder when exporting multiple documents.
    bool auto_create_folder;

    /// whether to output one document per page.
    bool output_document_per_page;

    /// Number of OCR languages specified in the languages array.
    int language_count;

    /// Specify the image scaling of the image file.
    float image_scaling;

    /// Specify the layout mode.
    CPageLayoutMode page_layout_mode;

    /// Specify the Excel worksheet option.
    CExcelWorksheetOption excel_worksheet_option;

    /// Specify the Html option.
    CHtmlOption html_option;

    /// Specify the OCR option.
    COCROption ocr_option;

    /// Specify the image color mode of the image file.
    CImageColorMode image_color_mode;

    /// Specify the image type of the image file.
    CImageType image_type;

    /// Specify the fontname to convert, e.g. "Arial".
    char font_name[256];

    /// Specify the pages to convert, e.g. "1-3,5,7-9"
    char page_ranges[256];

    /// Specify the OCR languages: an array of language_count entries, at most 32 of
    /// which are used. If OCR is enabled and no language is given, English is used.
    COCRLanguage* languages;

    /// Whether to classify and correct document orientation before OCR image recognition.
    /// Takes effect only when enable_ocr is true; requires the document_process.model package.
    bool enable_document_orientation_classification;

    /// Whether to apply DocumentAI dewarp before OCR image recognition.
    /// Takes effect only when enable_ocr is true; requires the document_process.model package.
    bool enable_document_dewarp;

    /// Whether to size the output page from the image's own resolution, matching it
    /// to the nearest standard paper. Only takes effect when the input is an image
    /// file; when disabled, one image pixel becomes one page point.
    bool enable_auto_page_size;

    /// Specify the output page width in points (1 pt = 1/72 inch), or 0 to not
    /// specify one. Set together with custom_page_height to choose the page
    /// outright, which takes precedence over enable_auto_page_size. The size is
    /// used exactly as given, orientation included, and the image is scaled to
    /// fit inside it without cropping. Only takes effect when the input is an
    /// image file. For example A4 is 595.28 x 841.89, JIS B5 is 515.91 x 728.50.
    float custom_page_width;

    /// Specify the output page height in points. See custom_page_width.
    float custom_page_height;

    /// Table recognition accuracy. CPDF_CONVERT_OPTION_DEFAULT sets e_CBalanced;
    /// a zero-initialized structure gives e_CHighAccuracy. e_CHighAccuracy requires
    /// the table_legacy.model package; if the model of the requested accuracy is not
    /// loaded, the conversion still succeeds, but tables are recognized
    /// with the rule-based method instead, so the result differs.
    CTableRecognitionAccuracy table_recognition_accuracy;

    /// How the PDF bookmarks are generated. Only takes effect for Word to PDF.
    CWordPdfBookmarkMode word_pdf_bookmark_mode;
} CConvertOption;

#define CPDF_CONVERT_OPTION_DEFAULT { \
    true, true, true, true, true, true, false, false, false, true, true, false, true, true, false, \
    0, 4.0f, e_PageLayoutModeFlow, e_CForTable, e_CSinglePage, e_CAll, e_CColor, e_CJPG, \
    {0}, {0}, 0, false, false, false, 0.0f, 0.0f, e_CBalanced, e_CWordPdfBookmarkNone \
}

#ifdef __cplusplus
}
#endif

#endif // COMPDFKIT_BASICTYPES_C_H
