package compdf

// ConvertCallback receives progress, cancellation, OCR, layout, and table callbacks.
// All methods have safe default implementations via BaseCallback so user callbacks
// only need to implement the methods they actually use.
type ConvertCallback interface {
	// OnProgress is called periodically with the current page (1-based) and
	// total page count of the conversion job.
	OnProgress(currentPage, totalPage int)

	// IsCancelled returns true to request the SDK to cancel the running job.
	IsCancelled() bool

	// OnOCR is called by the SDK when it has written a temporary PNG to
	// imagePath and is requesting external OCR. Return true on success.
	OnOCR(imagePath string) bool

	// OnLayout is called by the SDK to request external layout analysis on
	// the given image. Return true on success.
	OnLayout(imagePath string) bool

	// OnTable is called by the SDK to request external table recognition on
	// the given image. Return true on success.
	OnTable(imagePath string) bool

	// GetOCRResult should return the latest OCR result as a JSON string.
	GetOCRResult() string

	// GetLayoutResult should return the latest layout result as a JSON string.
	GetLayoutResult() string

	// GetTableResult should return the latest table result as a JSON string.
	GetTableResult() string
}

// BaseCallback provides no-op implementations of every ConvertCallback method.
// Embed it in your own callback type and override only the methods you need:
//
//	type myCB struct{ compdf.BaseCallback }
//	func (c *myCB) OnProgress(cur, total int) { ... }
type BaseCallback struct{}

func (BaseCallback) OnProgress(currentPage, totalPage int) {}
func (BaseCallback) IsCancelled() bool                     { return false }
func (BaseCallback) OnOCR(imagePath string) bool           { return false }
func (BaseCallback) OnLayout(imagePath string) bool        { return false }
func (BaseCallback) OnTable(imagePath string) bool         { return false }
func (BaseCallback) GetOCRResult() string                  { return "" }
func (BaseCallback) GetLayoutResult() string               { return "" }
func (BaseCallback) GetTableResult() string                { return "" }

// ProgressFunc is a convenience adapter so callers can pass a plain function
// when they only care about progress reporting.
type ProgressFunc func(currentPage, totalPage int)

type progressOnlyCallback struct {
	BaseCallback
	fn ProgressFunc
}

func (p *progressOnlyCallback) OnProgress(c, t int) {
	if p.fn != nil {
		p.fn(c, t)
	}
}

// NewProgressCallback wraps a ProgressFunc as a ConvertCallback. Pass nil to
// the Start* functions if you don't want any callback at all.
func NewProgressCallback(fn ProgressFunc) ConvertCallback {
	return &progressOnlyCallback{fn: fn}
}
