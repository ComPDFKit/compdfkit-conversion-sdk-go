package compdf

// convertOptions is the internal aggregate that maps 1:1 onto CConvertOption
// in compdf_basictypes_c.h. It is populated from a per-format Options struct
// by the corresponding Start* function. End users do not construct this type
// directly.
type convertOptions struct {
	EnableAILayout             bool
	EnableAITableRecognition   bool
	ContainImage               bool
	ContainPageBackgroundImage bool
	JSONContainTable           bool
	ContainAnnotation          bool
	ExcelAllContent            bool
	ExcelCSVFormat             bool
	EnableOCR                  bool
	TransparentText            bool
	TxtTableFormat             bool
	ImagePathEnhance           bool
	FormulaToImage             bool
	AutoCreateFolder           bool
	OutputDocumentPerPage      bool
	ImageScaling               float32
	PageLayoutMode             PageLayoutMode
	ExcelWorksheetOption       ExcelWorksheetOption
	HtmlPageOption             HtmlPageOption
	OCROption                  OCROption
	ImageColorMode             ImageColorMode
	ImageType                  ImageType
	FontName                   string
	PageRanges                 string
	Languages                  []OCRLanguage
}

func defaultConvertOptions() convertOptions {
	return convertOptions{
		EnableAILayout:             true,
		EnableAITableRecognition:   true,
		ContainImage:               true,
		ContainPageBackgroundImage: true,
		JSONContainTable:           true,
		ContainAnnotation:          true,
		TxtTableFormat:             true,
		AutoCreateFolder:           true,
		ImageScaling:               1.0,
		PageLayoutMode:             PageLayoutFlow,
		ExcelWorksheetOption:       ExcelForTable,
		HtmlPageOption:             HtmlSinglePage,
		OCROption:                  OCRAll,
		ImageColorMode:             ImageColor,
		ImageType:                  ImageJPG,
	}
}

// WordOptions configures PDF->Word conversion.
type WordOptions struct {
	ContainImage               bool
	ContainAnnotation          bool
	EnableAILayout             bool
	EnableAITableRecognition   bool
	FormulaToImage             bool
	EnableOCR                  bool
	LayoutMode                 PageLayoutMode
	PageRanges                 string
	ContainPageBackgroundImage bool
	OutputDocumentPerPage      bool
	OCROption                  OCROption
	Languages                  []OCRLanguage
}

// NewWordOptions returns WordOptions initialized with SDK defaults.
func NewWordOptions() WordOptions {
	return WordOptions{
		ContainImage:               true,
		ContainAnnotation:          true,
		EnableAILayout:             true,
		EnableAITableRecognition:   true,
		LayoutMode:                 PageLayoutFlow,
		ContainPageBackgroundImage: true,
		OCROption:                  OCRAll,
	}
}

func (o WordOptions) toConvert() convertOptions {
	c := defaultConvertOptions()
	c.ContainImage = o.ContainImage
	c.ContainAnnotation = o.ContainAnnotation
	c.EnableAILayout = o.EnableAILayout
	c.EnableAITableRecognition = o.EnableAITableRecognition
	c.FormulaToImage = o.FormulaToImage
	c.EnableOCR = o.EnableOCR
	c.PageLayoutMode = o.LayoutMode
	c.PageRanges = o.PageRanges
	c.ContainPageBackgroundImage = o.ContainPageBackgroundImage
	c.OutputDocumentPerPage = o.OutputDocumentPerPage
	c.OCROption = o.OCROption
	c.Languages = o.Languages
	return c
}

// ExcelOptions configures PDF->Excel conversion.
type ExcelOptions struct {
	ContainImage             bool
	ContainAnnotation        bool
	EnableAILayout           bool
	EnableAITableRecognition bool
	FormulaToImage           bool
	EnableOCR                bool
	PageRanges               string
	AllContent               bool
	CSVFormat                bool
	WorksheetOption          ExcelWorksheetOption
	AutoCreateFolder         bool
	OutputDocumentPerPage    bool
	OCROption                OCROption
	Languages                []OCRLanguage
}

// NewExcelOptions returns ExcelOptions initialized with SDK defaults.
func NewExcelOptions() ExcelOptions {
	return ExcelOptions{
		ContainImage:             true,
		ContainAnnotation:        true,
		EnableAILayout:           true,
		EnableAITableRecognition: true,
		WorksheetOption:          ExcelForTable,
		AutoCreateFolder:         true,
		OCROption:                OCRAll,
	}
}

func (o ExcelOptions) toConvert() convertOptions {
	c := defaultConvertOptions()
	c.ContainImage = o.ContainImage
	c.ContainAnnotation = o.ContainAnnotation
	c.EnableAILayout = o.EnableAILayout
	c.EnableAITableRecognition = o.EnableAITableRecognition
	c.FormulaToImage = o.FormulaToImage
	c.EnableOCR = o.EnableOCR
	c.ExcelAllContent = o.AllContent
	c.ExcelCSVFormat = o.CSVFormat
	c.PageRanges = o.PageRanges
	c.ExcelWorksheetOption = o.WorksheetOption
	c.AutoCreateFolder = o.AutoCreateFolder
	c.OutputDocumentPerPage = o.OutputDocumentPerPage
	c.OCROption = o.OCROption
	c.Languages = o.Languages
	return c
}

// PptOptions configures PDF->PPT conversion.
type PptOptions struct {
	ContainImage               bool
	ContainAnnotation          bool
	EnableAILayout             bool
	EnableAITableRecognition   bool
	FormulaToImage             bool
	EnableOCR                  bool
	PageRanges                 string
	ContainPageBackgroundImage bool
	OutputDocumentPerPage      bool
	OCROption                  OCROption
	Languages                  []OCRLanguage
}

// NewPptOptions returns PptOptions initialized with SDK defaults.
func NewPptOptions() PptOptions {
	return PptOptions{
		ContainImage:               true,
		ContainAnnotation:          true,
		EnableAILayout:             true,
		EnableAITableRecognition:   true,
		ContainPageBackgroundImage: true,
		OCROption:                  OCRAll,
	}
}

func (o PptOptions) toConvert() convertOptions {
	c := defaultConvertOptions()
	c.ContainImage = o.ContainImage
	c.ContainAnnotation = o.ContainAnnotation
	c.EnableAILayout = o.EnableAILayout
	c.EnableAITableRecognition = o.EnableAITableRecognition
	c.FormulaToImage = o.FormulaToImage
	c.EnableOCR = o.EnableOCR
	c.PageRanges = o.PageRanges
	c.ContainPageBackgroundImage = o.ContainPageBackgroundImage
	c.OutputDocumentPerPage = o.OutputDocumentPerPage
	c.OCROption = o.OCROption
	c.Languages = o.Languages
	return c
}

// HtmlOptions configures PDF->HTML conversion.
type HtmlOptions struct {
	ContainImage               bool
	ContainAnnotation          bool
	EnableAILayout             bool
	EnableAITableRecognition   bool
	FormulaToImage             bool
	EnableOCR                  bool
	LayoutMode                 PageLayoutMode
	PageRanges                 string
	PageOption                 HtmlPageOption
	ContainPageBackgroundImage bool
	OutputDocumentPerPage      bool
	OCROption                  OCROption
	Languages                  []OCRLanguage
}

// NewHtmlOptions returns HtmlOptions initialized with SDK defaults.
func NewHtmlOptions() HtmlOptions {
	return HtmlOptions{
		ContainImage:               true,
		ContainAnnotation:          true,
		EnableAILayout:             true,
		EnableAITableRecognition:   true,
		LayoutMode:                 PageLayoutFlow,
		PageOption:                 HtmlSinglePage,
		ContainPageBackgroundImage: true,
		OCROption:                  OCRAll,
	}
}

func (o HtmlOptions) toConvert() convertOptions {
	c := defaultConvertOptions()
	c.ContainImage = o.ContainImage
	c.ContainAnnotation = o.ContainAnnotation
	c.EnableAILayout = o.EnableAILayout
	c.EnableAITableRecognition = o.EnableAITableRecognition
	c.FormulaToImage = o.FormulaToImage
	c.EnableOCR = o.EnableOCR
	c.PageLayoutMode = o.LayoutMode
	c.PageRanges = o.PageRanges
	c.HtmlPageOption = o.PageOption
	c.ContainPageBackgroundImage = o.ContainPageBackgroundImage
	c.OutputDocumentPerPage = o.OutputDocumentPerPage
	c.OCROption = o.OCROption
	c.Languages = o.Languages
	return c
}

// ImageOptions configures PDF->Image conversion.
type ImageOptions struct {
	ImageType      ImageType
	ImageColorMode ImageColorMode
	ImageScaling   float32
	PathEnhance    bool
	PageRanges     string
}

// NewImageOptions returns ImageOptions initialized with SDK defaults.
func NewImageOptions() ImageOptions {
	return ImageOptions{
		ImageType:      ImageJPG,
		ImageColorMode: ImageColor,
		ImageScaling:   1.0,
	}
}

func (o ImageOptions) toConvert() convertOptions {
	c := defaultConvertOptions()
	c.ImageType = o.ImageType
	c.ImageColorMode = o.ImageColorMode
	c.ImageScaling = o.ImageScaling
	c.ImagePathEnhance = o.PathEnhance
	c.PageRanges = o.PageRanges
	c.FormulaToImage = false
	return c
}

// MarkdownOptions configures PDF->Markdown conversion.
type MarkdownOptions struct {
	ContainImage             bool
	ContainAnnotation        bool
	EnableAILayout           bool
	EnableAITableRecognition bool
	EnableOCR                bool
	PageRanges               string
	OutputDocumentPerPage    bool
	OCROption                OCROption
	Languages                []OCRLanguage
}

// NewMarkdownOptions returns MarkdownOptions initialized with SDK defaults.
func NewMarkdownOptions() MarkdownOptions {
	return MarkdownOptions{
		ContainImage:             true,
		ContainAnnotation:        true,
		EnableAILayout:           true,
		EnableAITableRecognition: true,
		OCROption:                OCRAll,
	}
}

func (o MarkdownOptions) toConvert() convertOptions {
	c := defaultConvertOptions()
	c.ContainImage = o.ContainImage
	c.ContainAnnotation = o.ContainAnnotation
	c.EnableAILayout = o.EnableAILayout
	c.EnableAITableRecognition = o.EnableAITableRecognition
	c.EnableOCR = o.EnableOCR
	c.PageRanges = o.PageRanges
	c.FormulaToImage = false
	c.OutputDocumentPerPage = o.OutputDocumentPerPage
	c.OCROption = o.OCROption
	c.Languages = o.Languages
	return c
}

// RtfOptions configures PDF->RTF conversion.
type RtfOptions struct {
	ContainImage               bool
	ContainAnnotation          bool
	EnableAILayout             bool
	EnableAITableRecognition   bool
	FormulaToImage             bool
	EnableOCR                  bool
	PageRanges                 string
	ContainPageBackgroundImage bool
	OutputDocumentPerPage      bool
	OCROption                  OCROption
	Languages                  []OCRLanguage
}

// NewRtfOptions returns RtfOptions initialized with SDK defaults.
func NewRtfOptions() RtfOptions {
	return RtfOptions{
		ContainImage:               true,
		ContainAnnotation:          true,
		EnableAILayout:             true,
		EnableAITableRecognition:   true,
		ContainPageBackgroundImage: true,
		OCROption:                  OCRAll,
	}
}

func (o RtfOptions) toConvert() convertOptions {
	c := defaultConvertOptions()
	c.ContainImage = o.ContainImage
	c.ContainAnnotation = o.ContainAnnotation
	c.EnableAILayout = o.EnableAILayout
	c.EnableAITableRecognition = o.EnableAITableRecognition
	c.FormulaToImage = o.FormulaToImage
	c.EnableOCR = o.EnableOCR
	c.PageRanges = o.PageRanges
	c.ContainPageBackgroundImage = o.ContainPageBackgroundImage
	c.OutputDocumentPerPage = o.OutputDocumentPerPage
	c.OCROption = o.OCROption
	c.Languages = o.Languages
	return c
}

// TxtOptions configures PDF->TXT conversion.
type TxtOptions struct {
	EnableAILayout           bool
	EnableAITableRecognition bool
	EnableOCR                bool
	PageRanges               string
	TableFormat              bool
	OutputDocumentPerPage    bool
	OCROption                OCROption
	Languages                []OCRLanguage
}

// NewTxtOptions returns TxtOptions initialized with SDK defaults.
func NewTxtOptions() TxtOptions {
	return TxtOptions{
		EnableAILayout:           true,
		EnableAITableRecognition: true,
		TableFormat:              true,
		OCROption:                OCRAll,
	}
}

func (o TxtOptions) toConvert() convertOptions {
	c := defaultConvertOptions()
	c.EnableAILayout = o.EnableAILayout
	c.EnableAITableRecognition = o.EnableAITableRecognition
	c.EnableOCR = o.EnableOCR
	c.TxtTableFormat = o.TableFormat
	c.PageRanges = o.PageRanges
	c.FormulaToImage = false
	c.OutputDocumentPerPage = o.OutputDocumentPerPage
	c.OCROption = o.OCROption
	c.Languages = o.Languages
	return c
}

// JsonOptions configures PDF->JSON conversion.
type JsonOptions struct {
	ContainImage             bool
	ContainAnnotation        bool
	EnableAILayout           bool
	EnableAITableRecognition bool
	EnableOCR                bool
	PageRanges               string
	ContainTable             bool
	OutputDocumentPerPage    bool
	OCROption                OCROption
	Languages                []OCRLanguage
}

// NewJsonOptions returns JsonOptions initialized with SDK defaults.
func NewJsonOptions() JsonOptions {
	return JsonOptions{
		ContainImage:             true,
		ContainAnnotation:        true,
		EnableAILayout:           true,
		EnableAITableRecognition: true,
		ContainTable:             true,
		OCROption:                OCRAll,
	}
}

func (o JsonOptions) toConvert() convertOptions {
	c := defaultConvertOptions()
	c.ContainImage = o.ContainImage
	c.ContainAnnotation = o.ContainAnnotation
	c.EnableAILayout = o.EnableAILayout
	c.EnableAITableRecognition = o.EnableAITableRecognition
	c.EnableOCR = o.EnableOCR
	c.JSONContainTable = o.ContainTable
	c.PageRanges = o.PageRanges
	c.FormulaToImage = false
	c.OutputDocumentPerPage = o.OutputDocumentPerPage
	c.OCROption = o.OCROption
	c.Languages = o.Languages
	return c
}

// SearchablePdfOptions configures PDF->SearchablePDF conversion.
type SearchablePdfOptions struct {
	ContainImage               bool
	EnableOCR                  bool
	EnableAITableRecognition   bool
	FormulaToImage             bool
	PageRanges                 string
	ContainPageBackgroundImage bool
	OutputDocumentPerPage      bool
	OCROption                  OCROption
	Languages                  []OCRLanguage
	TransparentText            bool
}

// NewSearchablePdfOptions returns SearchablePdfOptions initialized with SDK defaults.
func NewSearchablePdfOptions() SearchablePdfOptions {
	return SearchablePdfOptions{
		ContainImage:               true,
		EnableAITableRecognition:   true,
		ContainPageBackgroundImage: true,
		OCROption:                  OCRAll,
	}
}

func (o SearchablePdfOptions) toConvert() convertOptions {
	c := defaultConvertOptions()
	c.ContainImage = o.ContainImage
	c.EnableOCR = o.EnableOCR
	c.EnableAITableRecognition = o.EnableAITableRecognition
	c.TransparentText = o.TransparentText
	c.FormulaToImage = o.FormulaToImage
	c.PageRanges = o.PageRanges
	c.ContainPageBackgroundImage = o.ContainPageBackgroundImage
	c.OutputDocumentPerPage = o.OutputDocumentPerPage
	c.OCROption = o.OCROption
	c.Languages = o.Languages
	return c
}

// OfdOptions configures PDF->OFD conversion.
type OfdOptions struct {
	TransparentText          bool
	EnableOCR                bool
	EnableAITableRecognition bool
	PageRanges               string
	OutputDocumentPerPage    bool
	OCROption                OCROption
	Languages                []OCRLanguage
}

// NewOfdOptions returns OfdOptions initialized with SDK defaults.
func NewOfdOptions() OfdOptions {
	return OfdOptions{
		EnableAITableRecognition: true,
		OCROption:                OCRAll,
	}
}

func (o OfdOptions) toConvert() convertOptions {
	c := defaultConvertOptions()
	c.TransparentText = o.TransparentText
	c.EnableOCR = o.EnableOCR
	c.EnableAITableRecognition = o.EnableAITableRecognition
	c.PageRanges = o.PageRanges
	c.OutputDocumentPerPage = o.OutputDocumentPerPage
	c.OCROption = o.OCROption
	c.Languages = o.Languages
	return c
}
