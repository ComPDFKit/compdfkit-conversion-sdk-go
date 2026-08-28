package compdf

import "testing"

func TestDocumentPreprocessingOptionsAreMapped(t *testing.T) {
	tests := []struct {
		name    string
		convert func() convertOptions
	}{
		{"word", func() convertOptions {
			o := NewWordOptions()
			o.EnableDocumentOrientationClassification = true
			o.EnableDocumentDewarp = true
			return o.toConvert()
		}},
		{"excel", func() convertOptions {
			o := NewExcelOptions()
			o.EnableDocumentOrientationClassification = true
			o.EnableDocumentDewarp = true
			return o.toConvert()
		}},
		{"ppt", func() convertOptions {
			o := NewPptOptions()
			o.EnableDocumentOrientationClassification = true
			o.EnableDocumentDewarp = true
			return o.toConvert()
		}},
		{"html", func() convertOptions {
			o := NewHtmlOptions()
			o.EnableDocumentOrientationClassification = true
			o.EnableDocumentDewarp = true
			return o.toConvert()
		}},
		{"markdown", func() convertOptions {
			o := NewMarkdownOptions()
			o.EnableDocumentOrientationClassification = true
			o.EnableDocumentDewarp = true
			return o.toConvert()
		}},
		{"rtf", func() convertOptions {
			o := NewRtfOptions()
			o.EnableDocumentOrientationClassification = true
			o.EnableDocumentDewarp = true
			return o.toConvert()
		}},
		{"txt", func() convertOptions {
			o := NewTxtOptions()
			o.EnableDocumentOrientationClassification = true
			o.EnableDocumentDewarp = true
			return o.toConvert()
		}},
		{"json", func() convertOptions {
			o := NewJsonOptions()
			o.EnableDocumentOrientationClassification = true
			o.EnableDocumentDewarp = true
			return o.toConvert()
		}},
		{"searchable pdf", func() convertOptions {
			o := NewSearchablePdfOptions()
			o.EnableDocumentOrientationClassification = true
			o.EnableDocumentDewarp = true
			return o.toConvert()
		}},
		{"ofd", func() convertOptions {
			o := NewOfdOptions()
			o.EnableDocumentOrientationClassification = true
			o.EnableDocumentDewarp = true
			return o.toConvert()
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.convert()
			if !got.EnableDocumentOrientationClassification {
				t.Error("EnableDocumentOrientationClassification was not mapped")
			}
			if !got.EnableDocumentDewarp {
				t.Error("EnableDocumentDewarp was not mapped")
			}
		})
	}
}

func TestConvertOptionDefaultsMatchSDKHeader(t *testing.T) {
	got := defaultConvertOptions()
	if !got.TransparentText {
		t.Error("TransparentText default = false, want true")
	}
	if !got.FormulaToImage {
		t.Error("FormulaToImage default = false, want true")
	}
	if got.ImageScaling != 4.0 {
		t.Errorf("ImageScaling default = %v, want 4.0", got.ImageScaling)
	}
	if got.EnableDocumentOrientationClassification {
		t.Error("EnableDocumentOrientationClassification default = true, want false")
	}
	if got.EnableDocumentDewarp {
		t.Error("EnableDocumentDewarp default = true, want false")
	}

	if imageScaling := NewImageOptions().ImageScaling; imageScaling != 4.0 {
		t.Errorf("NewImageOptions().ImageScaling = %v, want 4.0", imageScaling)
	}
	if !NewWordOptions().FormulaToImage || !NewExcelOptions().FormulaToImage ||
		!NewPptOptions().FormulaToImage || !NewHtmlOptions().FormulaToImage ||
		!NewRtfOptions().FormulaToImage || !NewSearchablePdfOptions().FormulaToImage {
		t.Error("a format-specific FormulaToImage default does not match the SDK header")
	}
	if !NewSearchablePdfOptions().TransparentText || !NewOfdOptions().TransparentText {
		t.Error("a format-specific TransparentText default does not match the SDK header")
	}
}

func TestUnsupportedFeatureError(t *testing.T) {
	const want = "compdf: source document uses an unsupported feature (code=95)"
	if got := ErrUnsupportedFeature.Error(); got != want {
		t.Fatalf("ErrUnsupportedFeature.Error() = %q, want %q", got, want)
	}
}
