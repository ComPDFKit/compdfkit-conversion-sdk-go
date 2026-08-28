package compdf

import "fmt"

// ErrorCode mirrors CSDKErrorCode in compdf_basictypes_c.h.
type ErrorCode int

const (
	ErrSuccess                    ErrorCode = 0
	ErrCancel                     ErrorCode = 1
	ErrFile                       ErrorCode = 2
	ErrPDFPassword                ErrorCode = 3
	ErrPDFPage                    ErrorCode = 4
	ErrPDFFormat                  ErrorCode = 5
	ErrPDFSecurity                ErrorCode = 6
	ErrOutOfMemory                ErrorCode = 7
	ErrIO                         ErrorCode = 8
	ErrCompress                   ErrorCode = 9
	ErrLicenseInvalid             ErrorCode = 20
	ErrLicenseExpire              ErrorCode = 21
	ErrLicenseUnsupportedPlatform ErrorCode = 22
	ErrLicenseUnsupportedID       ErrorCode = 23
	ErrLicenseUnsupportedDevice   ErrorCode = 24
	ErrLicensePermissionDeny      ErrorCode = 25
	ErrLicenseUninitialized       ErrorCode = 26
	ErrLicenseIllegalAccess       ErrorCode = 27
	ErrLicenseFileReadFailed      ErrorCode = 28
	ErrLicenseOCRPermissionDeny   ErrorCode = 29
	ErrLicenseConcurrencyExceeded ErrorCode = 30
	ErrLicensePageLimitExceeded   ErrorCode = 31
	ErrLicenseQuotaCorrupted      ErrorCode = 32
	ErrNoTable                    ErrorCode = 40
	ErrOCRFailure                 ErrorCode = 41
	ErrConverting                 ErrorCode = 60
	ErrInvalidArg                 ErrorCode = 80
	ErrInvalidHandle              ErrorCode = 81
	ErrModelInvalidFormat         ErrorCode = 82
	ErrModelFunctionUnsupported   ErrorCode = 83
	ErrModelFormatUnsupported     ErrorCode = 84
	ErrModelSDKMismatch           ErrorCode = 85
	ErrImageDataEmpty             ErrorCode = 86
	ErrImageWHError               ErrorCode = 87
	ErrImageUnsupportedFormat     ErrorCode = 88
	ErrImageInvalid               ErrorCode = 89
	ErrExpire                     ErrorCode = 90
	ErrMissingArg                 ErrorCode = 91
	ErrLicenseUnsupportedAPI      ErrorCode = 92
	ErrLicenseMismatch            ErrorCode = 93
	ErrInvalidTable               ErrorCode = 94
	ErrUnsupportedFeature         ErrorCode = 95
	ErrUnknown                    ErrorCode = 100
)

var errorMessages = map[ErrorCode]string{
	ErrSuccess:                    "success",
	ErrCancel:                     "conversion canceled",
	ErrFile:                       "file not found or cannot be opened",
	ErrPDFPassword:                "invalid pdf password",
	ErrPDFPage:                    "pdf page failed to load",
	ErrPDFFormat:                  "pdf format invalid or corrupted",
	ErrPDFSecurity:                "pdf encrypted by unsupported security handler",
	ErrOutOfMemory:                "out of memory",
	ErrIO:                         "system i/o error",
	ErrCompress:                   "folder compression failed",
	ErrLicenseInvalid:             "license invalid",
	ErrLicenseExpire:              "license expired",
	ErrLicenseUnsupportedPlatform: "license does not support current platform",
	ErrLicenseUnsupportedID:       "license does not support application id",
	ErrLicenseUnsupportedDevice:   "license does not support device id",
	ErrLicensePermissionDeny:      "license missing function permission",
	ErrLicenseUninitialized:       "license not initialized",
	ErrLicenseIllegalAccess:       "illegal access to api",
	ErrLicenseFileReadFailed:      "failed to read license file",
	ErrLicenseOCRPermissionDeny:   "license lacks OCR permission",
	ErrLicenseConcurrencyExceeded: "concurrency exceeds license limit",
	ErrLicensePageLimitExceeded:   "pages exceed license limit",
	ErrLicenseQuotaCorrupted:      "quota persistence file tampered",
	ErrNoTable:                    "no tables found",
	ErrOCRFailure:                 "ocr recognition failed",
	ErrConverting:                 "another conversion task is running",
	ErrInvalidArg:                 "invalid argument",
	ErrInvalidHandle:              "invalid handle",
	ErrModelInvalidFormat:         "model file invalid",
	ErrModelFunctionUnsupported:   "model function not supported",
	ErrModelFormatUnsupported:     "model format not supported",
	ErrModelSDKMismatch:           "model incompatible with sdk version",
	ErrImageDataEmpty:             "image data empty",
	ErrImageWHError:               "image width/height invalid",
	ErrImageUnsupportedFormat:     "image format not supported",
	ErrImageInvalid:               "image invalid or corrupted",
	ErrExpire:                     "resource expired",
	ErrMissingArg:                 "required argument missing",
	ErrLicenseUnsupportedAPI:      "license does not support this api",
	ErrLicenseMismatch:            "license does not match device/module/version",
	ErrInvalidTable:               "invalid table data",
	ErrUnsupportedFeature:         "source document uses an unsupported feature",
	ErrUnknown:                    "unknown error",
}

func (c ErrorCode) Error() string {
	if msg, ok := errorMessages[c]; ok {
		return fmt.Sprintf("compdf: %s (code=%d)", msg, int(c))
	}
	return fmt.Sprintf("compdf: unknown error code %d", int(c))
}

// IsSuccess reports whether the code represents a successful operation.
func (c ErrorCode) IsSuccess() bool { return c == ErrSuccess }

func (c ErrorCode) toError() error {
	if c == ErrSuccess {
		return nil
	}
	return c
}

// OCRLanguage mirrors COCRLanguage.
type OCRLanguage int

const (
	OCRLangUnknown     OCRLanguage = 0
	OCRLangChinese     OCRLanguage = 1
	OCRLangChineseTrad OCRLanguage = 2
	OCRLangEnglish     OCRLanguage = 3
	OCRLangKorean      OCRLanguage = 4
	OCRLangJapanese    OCRLanguage = 5
	OCRLangLatin       OCRLanguage = 6
	OCRLangDevanagari  OCRLanguage = 7
	OCRLangCyrillic    OCRLanguage = 8
	OCRLangArabic      OCRLanguage = 9
	OCRLangTamil       OCRLanguage = 10
	OCRLangTelugu      OCRLanguage = 11
	OCRLangKannada     OCRLanguage = 12
	OCRLangThai        OCRLanguage = 13
	OCRLangGreek       OCRLanguage = 14
	OCRLangEslav       OCRLanguage = 15
	OCRLangAuto        OCRLanguage = 16
)

// PageLayoutMode mirrors CPageLayoutMode.
type PageLayoutMode int

const (
	PageLayoutBox  PageLayoutMode = 0
	PageLayoutFlow PageLayoutMode = 1
)

// OCROption mirrors COCROption.
type OCROption int

const (
	OCRInvalidCharacter           OCROption = 0
	OCRScanPage                   OCROption = 1
	OCRInvalidCharacterAndScanned OCROption = 2
	OCRAll                        OCROption = 3
)

// ImageColorMode mirrors CImageColorMode.
type ImageColorMode int

const (
	ImageColor  ImageColorMode = 0
	ImageGray   ImageColorMode = 1
	ImageBinary ImageColorMode = 2
)

// ImageType mirrors CImageType.
type ImageType int

const (
	ImageJPG    ImageType = 0
	ImageJPEG   ImageType = 1
	ImageJPEG2K ImageType = 2
	ImagePNG    ImageType = 3
	ImageBMP    ImageType = 4
	ImageTIFF   ImageType = 5
	ImageTGA    ImageType = 6
	ImageGIF    ImageType = 7
	ImageWEBP   ImageType = 8
)

// ExcelWorksheetOption mirrors CExcelWorksheetOption.
type ExcelWorksheetOption int

const (
	ExcelForTable    ExcelWorksheetOption = 0
	ExcelForPage     ExcelWorksheetOption = 1
	ExcelForDocument ExcelWorksheetOption = 2
)

// HtmlPageOption mirrors CHtmlOption.
type HtmlPageOption int

const (
	HtmlSinglePage             HtmlPageOption = 0
	HtmlSinglePageWithBookmark HtmlPageOption = 1
	HtmlMultiPage              HtmlPageOption = 2
	HtmlMultiPageWithBookmark  HtmlPageOption = 3
)
