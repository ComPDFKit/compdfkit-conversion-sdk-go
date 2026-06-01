# ComPDF Conversion SDK for Go

# 1. Overview

ComPDF Conversion SDK is a high-performance library designed for extracting and transforming the data within your PDF files, such as text, images, tables, links, and annotations, into various file formats. Our Conversion SDK retains the original document layout and the properties of the file data, ensuring a seamless document conversion experience.

Effortlessly integrate the ComPDF Conversion SDK into your projects in just a few steps, and enable the following file format conversions:

- Convert PDF to Word (.docx)
- Convert PDF to Excel (.xlsx)
- Convert PDF to PowerPoint (.pptx)
- Convert PDF to HTML (.html)
- Convert PDF to CSV (.csv)
- Convert PDF to Image (.png, .jpg, .jpeg, .jpeg2000, .bmp, .tiff, .tga, .gif, .webp)
- Convert PDF to Plain Text (.txt)
- Convert PDF to Rich Text Format (.rtf)

- Convert PDF to Searchable PDF (.pdf)


- Convert PDF to OFD (.ofd)

- Convert PDF to Structured Data (.json)
- Convert PDF to Markdown (.md)


Harness the power of ComPDF Conversion SDK and take your document processing to the next level with intelligent tools designed for accuracy and efficiency. To enhance your format conversion results, ComPDF offers AI-powered document tools with the following capabilities:


- Optical Character Recognition (OCR)


- Layout Analysis


- Table Recognition



## 1.1 Why ComPDF Conversion SDK

- Mature Technology

  With years of technology accumulation, we have established a complete mechanism of product iteration to offer a continuous guarantee for product competitiveness.

- Complete PDF and Format Conversion Functionalities

  Our comprehensive features can meet diverse needs and are easy for our customers to use without training costs.

- High-quality Service

  professional service and technical support to quickly respond to users' feedback through onsite service or remote support like telephone, email, etc.

- Independent Intellectual Property Rights

  Our technology is independent and compliant with ISO, helping enterprises conduct international business without considering copyright risks.

## 1.2 ComPDF Conversion SDK

The ComPDF Conversion SDK is designed to facilitate the effortless conversion of PDF files into various other formats, all while preserving the original layout and formatting of the documents. In this guide, we will explore the ComPDF Conversion SDK and demonstrate how to utilize it within your Go projects.

## 1.3 License & Trial

The ComPDF Conversion SDK is a commercial SDK that requires a license to grant developers the right to develop and distribute their applications. In development mode, each license is only valid for one device ID. ComPDF provides flexible licensing models, please contact [our sale's team](mailto:support@compdf.com) for more information. However, even if you have a license, it is prohibited to distribute any documents, sample code, or source code of the ComPDF Conversion SDK to any third parties.

If you do not have a License, please feel free to contact the [ComPDF Team](https://www.compdf.com/contact-sales) to obtain a License to try the ComPDF Conversion SDK.
# 2. Getting Started

## 2.1 Requirements
Before starting, please make sure that you have already met the following prerequisites.

### 2.1.1 Get ComPDF License Key

ComPDF provides two types of license key: 30-day free trial license, and commercial license.

#### How to Get Free Trial License

[Contact our sales team](https://www.compdf.com/contact-sales) and we'll send you a 30-day free trial license for ComPDF Conversion SDK.

#### How to Get Commercial License

ComPDF Conversion SDK is a commercial SDK that requires a license for application release. Any documents, sample code, or source code distribution from the released package of ComPDF to any third party is prohibited.

**Contact Sales**

To get commercial license for ComPDF Conversion SDK, feel free to [contact our sales team](https://www.compdf.com/contact-sales).

For Go Conversion SDK, the commercial license must be bound to your developer device ID ([How to find the developer device ID](https://www.compdf.com/faq/how-to-find-the-device-id)), and each license is only valid for one device ID in development mode.


### 2.1.2 Download Conversion SDK

[Contact us](https://www.compdf.com/contact-sales) to obtain the ComPDF Go Conversion SDK.


### 2.1.3 System Requirements

| Platform | System Requirements | Development Environment | Notes |
| :------ | :------ | :------ | :------ |
| Windows | - Windows 7 or higher. | Go 1.21 or higher, with CGO enabled and a C compiler (e.g., MinGW-w64) available. | - |
| macOS | - Mac OS 11.0 or higher. | Go 1.21 or higher, with CGO enabled and Xcode Command Line Tools installed. | - |
| Linux | - Linux x64.<br/>- GLIBC 2.31 or higher. | Go 1.21 or higher, with CGO enabled and a C compiler (e.g., gcc) available. | - |

## 2.2 SDK Package Structure

The ComPDF Conversion SDK for Go should contain at least the following:

- ***"include"*** - C SDK header files used by CGO.
- ***"lib"*** - Platform-specific dynamic libraries (e.g., `compdfkit_conversion.dll`, `libcompdfkit_conversion.so`, `libcompdfkit_conversion.dylib`).

> **Note:** After building your Go application, the native dynamic libraries under `lib/<os>/<arch>/` (e.g., `libDocumentAI.so.*`, `libonnxruntime.so.*`, `libopencv_world.so.*` on Linux, or the corresponding `.dll` / `.dylib` files on Windows and macOS) are **not** statically linked into the resulting binary. You must make sure they can be located by the dynamic loader at runtime, otherwise the application will fail to start with a "library not found" error. Common approaches:
>
> - **Windows:** copy the `.dll` files next to your executable, or add the directory containing them to the `PATH` environment variable.
> - **Linux:** add the directory containing the `.so` files to `LD_LIBRARY_PATH`, install them into a system path such as `/usr/local/lib` (and run `ldconfig`), or build your binary with an embedded `rpath`/`runpath` (e.g., `-Wl,-rpath,$ORIGIN/lib` via `CGO_LDFLAGS`).
> - **macOS:** add the directory containing the `.dylib` files to `DYLD_LIBRARY_PATH`, or embed an `rpath` (e.g., `-Wl,-rpath,@executable_path/lib` via `CGO_LDFLAGS`) and ship the libraries alongside your binary.
>
> For redistribution, it is recommended to package the native libraries together with your executable and use a relative `rpath`/`PATH` so the application works out of the box on end-user machines.

### 2.2.1 Apply the License Key
If you don't have a license key, please check out [how to obtain a license key](/guides/conversion-sdk/go/requirements).

ComPDF Conversion SDK currently supports offline authentication to verify license keys.

*Learn about:*

[*What is the authentication mechanism of ComPDF's license?*](https://www.compdf.com/faq/authentication-mechanism-of-compdfkit-license)

#### Copy the License Key Path

Accurately obtaining the license key is crucial for the application of the license.

1. In the email you received, locate the XML file containing the license key.
2. Open the XML file, and determine the license type based on the `<type>` field. If `<type>online</type>` is present, it indicates an online license. If `<type>offline</type>` is present or if the field is absent, it indicates an offline license.

**Online License**:

```xml
<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<license version="1">
    <platform>windows</platform>
    <starttime>xxxxxxxx</starttime>
    <endtime>xxxxxxxx</endtime>
    <type>online</type>
    <key>LICENSE_KEY</key>
</license>

```

**Offline License**:

```xml
<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<license version="1">
    <platform>windows</platform>
    <starttime>xxxxxxxx</starttime>
    <endtime>xxxxxxxx</endtime>
    <key>LICENSE_KEY</key>
</license>

```

#### Apply the License Key
You can perform offline authentication using the following method:

Before using the classes and methods of the ComPDF Conversion SDK in your project, you need to initialize the SDK with a valid license key. If you don’t have a license key, feel free to contact the [ComPDF team](https://www.compdf.com/contact-sales) to obtain one.

```go
import compdf "github.com/ComPDFKit/compdfkit-conversion-sdk-go/v4"

err := compdf.LicenseVerify("LICENSE_KEY_PATH", "DEVICE_ID", "APP_ID")
```

# 3. Conversion Guides

ComPDF Conversion SDK allows developers to use a simple API to convert PDF to commonly used file
formats such as Word, Excel, PowerPoint, HTML, CSV, PNG, JPEG, RTF, TXT, Searchable PDF, OFD, JSON, and Markdown. It provides a wide range of customized conversion
options, such as whether to include images or annotations in PDF documents, whether to enable OCR, whether to enable layout analysis, and more.
## 3.1 Initialize Library Resources

#### **Overview**

Initialize the necessary file and memory resources required by the ComPDF Conversion SDK.

#### **Notes**

- You must initialize SDK resources before calling any conversion interface.

- When using OCR, Layout Analysis, PDF to Searchable PDF, or PDF to OFD, make sure the font resources and DocumentAI model resources in the `resource` directory are available.


#### **Example**

```go
compdf.Initialize("path/to/resource")
```
## 3.2 Set DocumentAI Model

#### **Overview**

Before using OCR, Layout Analysis, Table Recognition you need to set the DocumentAI model path first.

`SetDocumentAIModel` supports an optional `gpuID` parameter used to specify the GPU device index for the AI model. When `gpuID` is `-1`, GPU acceleration is disabled.

#### **Set AI Model Instance Count**

If you need to control the number of Layout Analysis and Table Recognition model instances, call the corresponding interface to set the model instance counts.


#### **Use Your Own AI Engine (SDK v4.1.0+)**

This option is available only in SDK v4.1.0 or later. You can skip `SetDocumentAIModel` and plug in your own OCR, Layout Analysis, or Table Recognition engine through the callbacks on `ConvertCallback`. See [3.11 Use Custom AI Models via Callbacks](#311-use-custom-ai-models-via-callbacks) for details.


#### **Sample**

```go
modelPath := "***"
compdf.SetDocumentAIModelCount(1, 1)
compdf.SetDocumentAIModel(modelPath, 0)
```
## 3.3 Get Conversion Progress

ComPDF Conversion SDK obtains the conversion progress through callback interfaces. The following example demonstrates how to get the conversion progress while performing a PDF to Word task.


```go
callback := compdf.NewProgressCallback(func(current, total int) {
    fmt.Printf("Progress: %d / %d\n", current, total)
})
```
## 3.4 Cancel Conversion Task

ComPDF Conversion SDK supports interrupting your ongoing conversion task at any time. The following example demonstrates how to interrupt your ongoing conversion task.


```go
type myCallback struct {
    compdf.BaseCallback
}

func (c *myCallback) IsCancelled() bool {
    return false
}

callback := &myCallback{}
```
## 3.5 Select Page Range for Conversion

ComPDF Conversion SDK supports converting a specified page range. When an empty string is passed, all pages will be converted. If the page range exceeds one page, you can also choose to enable the `OutputDocumentPerPage` option to output each PDF page as a separate file. The following example demonstrates how to specify a page range when performing a conversion task:

```go
options := compdf.NewWordOptions()
options.OutputDocumentPerPage = false
options.PageRanges = "1-3,5,7-9"
```
## 3.6 Conversion Options: Contain Image & Annotation

#### **Overview**

In the process of converting PDF documents into various formats, ComPDF Conversion SDK offers two additional options for users: one option to determine whether images are included in the generated document, and another to decide if annotations from the PDF file are to be retained.

- With the "Include Images" option enabled, ComPDF Conversion SDK will extract the images from the PDF document and embed them in the corresponding pages and positions in the output file. For areas with overlapping images, ComPDF Conversion SDK merges these images into one and embeds it into the exact location on the corresponding page of the output file.

- When the "Include Annotations" option is selected, most annotations are converted into raster images and embedded at the respective positions within your document. However, certain types of annotations, such as highlights, underlines, strikeouts, and squiggly, are converted into their respective formatting equivalents in the converted Word, PPT, and HTML documents, and are marked over the corresponding text. It is important to note that the conversion won't be 100% accurate in every instance.

In the ComPDF Conversion SDK, the options for including images and annotations are commonly used in the following format conversions:

- PDF to Word
- PDF to Excel
- PDF to PowerPoint
- PDF to HTML
- PDF to RTF
- Extract PDF to JSON
- Extract PDF to Markdown

#### **About Text Markup Annotation**

- **Highlight:** When converting PDF to Word format keeping the highlight markups, it is important to note that Microsoft Word only supports 15 highlighter colors. To approximate the original document's appearance as closely as possible, text in the Word document will be marked with a text background color that matches the color of the original document's highlight annotation. For conversions to Microsoft PPT format, the native highlighting feature within the format is used to mark the text. In the case of converting to HTML format, a unique `<span>` tag is created for the marked text, and the background style is set to match the color of the corresponding annotation in the original document.

- **Underlines & Squiggly:** When converting PDF to Word or PPT formats keeping the underline and squiggly markups, the marked text will be marked with the same style in Microsoft Office. When converted to HTML format, the marked text will be styled to display the same effect. However, if a paragraph of text in the original document is marked by both underline and squiggly, then the text will only be marked with one type (Because squiggly is actually a type of underline in Word, PPT, and HTML formats).
  
- **Strikeout:** When converting strikeout markups to Word and PowerPoint formats, the marked text will be added with a strikeout natively supported by Microsoft Office. However, in these two file formats, the color of the strikeout itself cannot be the same as that in the original PDF document because the strikeout color in Word and PPT will only change according to the color of the marked text font itself. When converted to HTML format, the same strikeout color as the original document will be displayed.
  

#### **Sample**

This Sample demonstrates how to use the ComPDF Conversion SDK to convert a PDF document to a Word document with the selected options: Include images and annotations.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

wordOptions := compdf.NewWordOptions()
wordOptions.ContainImage = true
wordOptions.ContainAnnotation = true
wordOptions.EnableAILayout = true
wordOptions.EnableOCR = false

err := compdf.StartPDFToWord(inputFilePath, password, outputFileName, wordOptions, nil)
```
## 3.7 Page Layout Mode

In certain formats, the page layout mode plays a key role in the quality of the converted document. ComPDF Conversion SDK supports two layout modes: Flow Layout and Box Layout.

- **Flow Layout:** This layout uses paragraph indentations, columns, and tab positions to adjust the content. Its main advantage is flexibility; content can flow automatically as the document is edited, and it adapts to various screen sizes on different devices. This layout also supports structured maintenance and can implement consistent global formatting through style templates (e.g., titles, body text). Common use cases include documents that are frequently modified, such as reports, manuals, and dynamic tables.
- **Box Layout:** Based on the PDF's "digital paper" model, this layout accurately positions every element (text, images, tables) on the page using a coordinate system (e.g., text is positioned 5 cm from the top and 3 cm from the left). The main advantage is high-precision rendering, which ensures consistency across different platforms. This layout is particularly useful for documents requiring precise reproduction, such as contracts, design drafts, and academic papers.

In the ComPDF Conversion SDK, page layout modes are commonly used in the following format conversions:

- PDF to Word
- PDF to HTML

#### **Sample**

This example demonstrates how to convert a PDF document to Word with Flow Layout and Box Layout:

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

wordOptions := compdf.NewWordOptions()
wordOptions.LayoutMode = compdf.PageLayoutFlow
err := compdf.StartPDFToWord(inputFilePath, password, outputFileName, wordOptions, nil)

wordOptions.LayoutMode = compdf.PageLayoutBox
err = compdf.StartPDFToWord(inputFilePath, password, outputFileName, wordOptions, nil)
```
## 3.8 OCR

#### **Overview**

OCR (Optical Character Recognition) is the process of converting images of typed, handwritten, or printed text into machine-encoded text.

OCR is commonly used for text recognition and extraction from the following types of documents:

- Non-editable scanned PDF files
- Photographs of documents.
- Scene photos such as advertising layouts, signboards, etc.
- Identification cards, passports, vehicle license plates, and other official plates.
- Invoices, bills, receipts, and other financial documents.

The following features support OCR:

- PDF to Word
- PDF to Excel
- PDF to PowerPoint (PPT)
- PDF to HTML
- PDF to Rich Text Format (RTF)
- PDF to Text (TXT)
- PDF to CSV
- PDF to Searchable PDF
- PDF to OFD
- Extract PDF to JSON
- Extract PDF to Markdown


OCR Language Support of ComPDF Conversion SDK:

| Script / Notes | Language (Native)                           | Language (In English)       |
| -------------- | ------------------------------ | ---------------------- |
| Latn; American | English                        | English               |
| Latn; Canadian | Français canadien              | French                |
| Hans/Hant      | 中文简体                       | Chinese (Simplified)  |
| Hans/Hant      | 中文繁体                       | Chinese (Traditional) |
| Jpan           | 日本語                         | Japanese              |
| Kore           | 한국어                         | Korean                |
| Latn           | Deutsch                        | German              |
| Latn           | Српски (латиница)              | Serbian (latin)       |
| Latn           | Occitan, lenga d'òc, provençal | Occitan               |
| Latn           | Dansk                          | Danish                |
| Latn           | Italiano                       | Italian               |
| Latn; European | Español                        | Spanish               |
| Latn; European | Português (Portugal)           | Portuguese            |
| Latn           | Te reo Māori                   | Maori                 |
| Latn           | Bahasa Melayu                  | Malay                 |
| Latn           | Malti                          | Maltese               |
| Latn           | Nederlands                     | Dutch                 |
| Latn; Bokmål   | Norsk                          | Norwegian             |
| Latn           | Polski                         | Polish                |
| Latn           | Română                         | Romanian              |
| Latn           | Slovenčina                     | Slovak                |
| Latn           | Slovenščina                    | Slovenian             |
| Latn           | shqip                          | Albanian              |
| Latn           | Svenska                        | Swedish               |
| Latn           | Swahili                        | Swahili               |
| Latn           | Wikang Tagalog                 | Tagalog               |
| Latn           | Türkçe                         | Turkish               |
| Latn           | oʻzbekcha                      | Uzbek                 |
| Latn           | Tiếng Việt                     | Vietnamese            |
| Latn           | Afrikaans                      | Afrikaans             |
| Latn           | Azərbaycan                     | Azerbaijani           |
| Latn           | Bosanski                       | Bosnian               |
| Latn           | Čeština                        | Czech                 |
| Latn           | Cymraeg                        | Welsh                 |
| Latn           | Eesti keel                     | Estonian              |
| Latn           | Gaeilge                        | Irish                 |
| Latn           | Hrvatski                       | Croatian              |
| Latn           | Magyar                         | Hungarian             |
| Latn           | Bahasa Indonesia               | Indonesian            |
| Latn           | Íslenska                       | Icelandic             |
| Latn           | Kurdî                          | Kurdish               |
| Latn           | Lietuvių                       | Lithuanian            |
| Latn           | Latviešu                       | Latvian               |

#### **Set OCR Language**

In the current mainline version, OCR languages should be passed through the `Languages` for each conversion task, rather than through a separate global interface.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

wordOptions := compdf.NewWordOptions()
wordOptions.EnableOCR = true
wordOptions.Languages = []compdf.OCRLanguage{compdf.OCRLangEnglish}
err := compdf.StartPDFToWord(inputFilePath, password, outputFileName, wordOptions, nil)
```

#### **OCR Options**
Different OCR options can be selected according to actual needs. Below are the currently supported OCR options.

- **OCRInvalidCharacter:** Recognizes invalid or garbled characters in the PDF document through OCR, while normal characters are not processed by OCR.
- **OCRScanPage:** Recognizes scanned pages in the PDF document through OCR, while editable pages are not processed by OCR.
- **OCRInvalidCharacterAndScanned:** Recognizes both invalid characters and scanned pages in the PDF document through OCR.
- **OCRAll:** Recognizes all pages and characters in the PDF document through OCR.

#### **Preserve Page Background**
When OCR is enabled, you can choose whether to enable the `ContainPageBackgroundImage` option. If this option is enabled, the original page background image of the PDF will be preserved. If it is disabled, the image result detected during page layout analysis will be retained.

#### **Notice**

- The quality of the OCR result depends on the quality of the input image. If the input image has a low resolution, the OCR result quality will be affected. A good rule of thumb is that the more pixels in the character shapes, the better. If the character bounding box is smaller than 20x20 pixels, OCR quality will drop exponentially. The ideal image is a grayscale image with a resolution around 300 DPI.
- When performing OCR, make sure the OCR language setting matches the language in the PDF document to achieve the best OCR conversion quality.
- OCR functionality is currently not supported on operating systems lower than Windows 10.


#### **Convert Images to Other Document Formats**

The OCR function also supports converting input images into Word, Excel, PPT, HTML, CSV, RTF, TXT, JSON, and other formats. This sample demonstrates how to use the ComPDF OCR function to convert image files to a DOCX file.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

wordOptions := compdf.NewWordOptions()
wordOptions.EnableOCR = true
wordOptions.Languages = []compdf.OCRLanguage{compdf.OCRLangEnglish}
err := compdf.StartPDFToWord(inputFilePath, password, outputFileName, wordOptions, nil)
```

#### **Sample**

This Sample demonstrates how to use the ComPDF OCR function to convert a PDF to DOCX file.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

wordOptions := compdf.NewWordOptions()
wordOptions.EnableOCR = true
wordOptions.Languages = []compdf.OCRLanguage{compdf.OCRLangEnglish}
err := compdf.StartPDFToWord(inputFilePath, password, outputFileName, wordOptions, nil)
```
## 3.9 Layout Analysis

#### **Overview**

Layout analysis is the process of leveraging Artificial Intelligence (AI) technology to parse and understand the structure of a document's layout. Its primary goal is to extract text, images, tables, layers, and other data from the input documents.

Layout analysis has several common use cases, including:

- **Intelligent recognition of tables within PDF documents**: This feature is particularly useful for analyzing company financial statements, invoices, bank statements, experimental data, medical test reports, and more.
- **Smart extraction of text, images, or tables from PDF documents through layout analysis**: This functionality greatly aids in the analysis and extraction of information from identification cards, receipts, licenses, documents, ancient books, and other various types of files.

Features that support Layout Analysis:

- PDF to Word
- PDF to Excel
- PDF to PowerPoint (PPT)
- PDF to HTML
- PDF to RTF
- PDF to TXT
- PDF to CSV
- Extract PDF to JSON
- Extract PDF to Markdown

#### **Notice**

- You need to load the DocumentAI model (see [3.2 Set DocumentAI Model](#32-set-documentai-model)) before using layout analysis, **or** plug in your own layout model through the callbacks described in [3.11 Use Custom AI Models via Callbacks](#311-use-custom-ai-models-via-callbacks).
- When OCR is enabled, layout analysis is automatically enabled.
- AI table recognition is a separate stage controlled by `EnableAITableRecognition`. See [3.10 Table Recognition](#310-table-recognition).

#### **Sample**

This Sample demonstrates how to use the ComPDF OCR function to convert PDF to DOCX file.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

wordOptions := compdf.NewWordOptions()
wordOptions.EnableAILayout = true
err := compdf.StartPDFToWord(inputFilePath, password, outputFileName, wordOptions, nil)
```
## 3.10 Table Recognition

#### **Overview**

Table Recognition reconstructs the internal structure of tables detected during layout analysis, including rows, columns, merged cells, and cell boundaries, so that the converted document preserves the original tabular semantics instead of producing a flat grid of text fragments.

It is controlled by the independent option `EnableAITableRecognition`, which is **enabled by default**. The table model is only invoked for table regions reported by layout analysis whose detection confidence is below the trusted threshold; high-confidence native PDF tables bypass the model to save inference time.

Typical scenarios that benefit from Table Recognition:

- **Borderless or partially bordered tables**, where ruling lines alone cannot describe the structure.
- **Tables with merged header cells, multi-row headers, or spanning cells**, such as financial statements, lab reports, and invoices.
- **Scanned tables processed by OCR**, where geometric reconstruction is required before cell-level data extraction.

Features that support Table Recognition:

- PDF to Word
- PDF to Excel
- PDF to PowerPoint (PPT)
- PDF to HTML
- PDF to RTF
- PDF to CSV
- Extract PDF to JSON
- Extract PDF to Markdown

#### **Notice**

- Table Recognition runs only when layout analysis is active (i.e. `EnableAILayout` is enabled, or implicitly when `EnableOCR` is enabled).
- You need to load the DocumentAI model before using Table Recognition, **or** plug in your own table model via the callbacks described in [3.11 Use Custom AI Models via Callbacks](#311-use-custom-ai-models-via-callbacks).
- Disabling `EnableAITableRecognition` turns the table model off entirely; detected table regions then fall back to geometric reconstruction from the underlying page objects.

#### **Sample**

This sample demonstrates how to convert a PDF to a Word document with Table Recognition enabled.

```go
compdf.SetDocumentAIModel("path/documentai.model", -1)

wordOptions := compdf.NewWordOptions()
wordOptions.EnableAILayout = true
wordOptions.EnableAITableRecognition = true // enabled by default in 4.1
err := compdf.StartPDFToWord("input.pdf", "password", "path/output.docx", wordOptions, nil)
```
## 3.11 Use Custom AI Models via Callbacks

#### **Overview**

Starting with **SDK v4.1.0**, ComPDF Conversion SDK exposes a callback-based extension point that lets you plug in your own AI inference engine for OCR, Layout Analysis, and Table Recognition. Instead of relying on the built-in DocumentAI model loaded by `SetDocumentAIModel`, you can:

- Run inference with any model or runtime you choose (e.g. your in-house engine, PaddleOCR, a cloud OCR API).
- Return the result to the SDK as a JSON string with a well-defined schema.

When the relevant callback pair is registered on `ConvertCallback`, the SDK skips its built-in model invocation for that capability and consumes your JSON output instead. If a pair is left unset, the SDK falls back to the built-in DocumentAI model (when available).

#### **Callback Pairs**

Each AI capability uses **two callbacks**: a *trigger* (invoked by the SDK with the path to a page image saved as PNG in a temp directory) and a *result getter* (invoked by the SDK immediately afterwards to retrieve the JSON string).

| Capability         | Trigger field / setter             | Result getter field / setter                | Triggered when                                                                              |
| ------------------ | ---------------------------------- | ------------------------------------------- | ------------------------------------------------------------------------------------------- |
| OCR                | `ocr`                              | `get_ocr_result`                            | OCR is enabled                                                                              |
| Layout Analysis    | `layout`                           | `get_layout_result`                         | layout analysis is enabled (or implicitly when OCR is enabled)                              |
| Table Recognition  | `table`                            | `get_table_result`                          | table recognition is enabled **and** a table region is detected by layout analysis          |

Rules:

- The trigger receives a UTF-8 path to a PNG file. Return `true` if your inference succeeded, `false` to make the SDK ignore the result for that page.
- The getter must return a UTF-8 JSON string. The SDK copies the string into an internal buffer before consuming it.
- Both callbacks for a capability must be set together. If only one is provided, the SDK falls back to the built-in path.
- Coordinates in your JSON must be in the **pixel space of the image the trigger received** (top-left origin, X right, Y down).
- Confidence filtering: OCR spans with `confidence < 0.1` and layout objects with `confidence < 0.45` are discarded by the SDK.
- When all three capabilities you need are covered by your own callbacks, `SetDocumentAIModel` does not have to be called.

#### **Sample**

```go
// Embed BaseCallback so you only override the methods you need.
type customAICallback struct {
    compdf.BaseCallback
    ocrJSON    string
    layoutJSON string
    tableJSON  string
}

func (c *customAICallback) OnOCR(imagePath string) bool {
    // Run your OCR engine on `imagePath`, cache the JSON result.
    c.ocrJSON = ""
    return true
}
func (c *customAICallback) OnLayout(imagePath string) bool {
    c.layoutJSON = ""
    return true
}
func (c *customAICallback) OnTable(imagePath string) bool {
    c.tableJSON = ""
    return true
}

func (c *customAICallback) GetOCRResult() string    { return c.ocrJSON }
func (c *customAICallback) GetLayoutResult() string { return c.layoutJSON }
func (c *customAICallback) GetTableResult() string  { return c.tableJSON }

callback := &customAICallback{}
wordOptions := compdf.NewWordOptions()
wordOptions.EnableOCR = true
wordOptions.EnableAILayout = true
wordOptions.EnableAITableRecognition = true
err := compdf.StartPDFToWord("input.pdf", "password", "path/output.docx", wordOptions, callback)
```

#### **JSON Schemas**

The expected JSON schema for each capability, including field names, coordinate space (pixel space of the image the trigger received), and value ranges, is the same across all language bindings. See the C++ developer guide chapter [Use Custom AI Models via Callbacks](../developer_guide_c++/developer_guide_c++.md#311-use-custom-ai-models-via-callbacks).
## 3.12 Output Font Option

#### **Overview**

In some output formats, you can set the preferred font name to unify the default font style in the output document.

#### **Supported Formats**

The `FontName` option currently applies to the following formats:

- PDF to Word
- PDF to Excel
- PDF to PowerPoint

#### **Example**

The following example demonstrates how to set the preferred font name for the output document.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

wordOptions := compdf.NewWordOptions()
wordOptions.FontName = "Arial"
err := compdf.StartPDFToWord(inputFilePath, password, outputFileName, wordOptions, nil)
```
## 3.13 Convert PDF to Word

#### **Overview**

Converting PDF to Word is an operation that converts a PDF file into an editable Word file. By converting PDF to Word, you can easily edit, modify, insert, or delete text and images, and adjust layout and properties.

#### **Layout differences**

- **Flow Layout:** Ideal for editing, with content dynamically adapting to different positions as you edit. However, a Word file may display differently because of incompatibilities across software or application versions. This makes it unsuitable for highly precise documents such as certificates or formal electronic records.
  
- **Fixed Layout:** Ensures a stable, uniform appearance and print quality across all devices. The content and formatting are locked upon creation, making alterations difficult without affecting the overall layout. It's preferred for formal documentation such as business reports and official electronic records.

#### **Sample**

This sample demonstrates how to convert from a PDF to DOCX file.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

wordOptions := compdf.NewWordOptions()
err := compdf.StartPDFToWord(inputFilePath, password, outputFileName, wordOptions, nil)
```

#### **Convert Formulas to Images**

When a document contains complex formulas and you want to preserve visual consistency in the output document, you can enable the `FormulaToImage` option.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

wordOptions := compdf.NewWordOptions()
wordOptions.FormulaToImage = true
err := compdf.StartPDFToWord(inputFilePath, password, outputFileName, wordOptions, nil)
```
## 3.14 Convert PDF to Excel

#### **Overview**

ComPDF Conversion SDK supports converting PDF documents to Microsoft Excel format (.xlsx). By extracting, parsing, and importing data from PDF into Excel, users can further edit, analyze, or share Excel files. This feature helps increase productivity, reduce manual entry errors, and simplify complex document processing tasks.

#### **Set the content options for Excel**

When converting PDF files to Excel files, you need to pay attention to the settings of the following options, which will directly affect the content written to the Excel file.

- Content options:

  If you set the `AllContent` option, the converted XLSX file will contain all content in the PDF.

- Worksheet options:

| Option | Description |
| --- | --- |
| `ExcelForTable` | Create one sheet per table. |
| `ExcelForPage` | Create one sheet per page. |
| `ExcelForDocument` | Create one sheet for the entire document. |


#### **Sample**

This sample demonstrates how to convert from a PDF to XLSX file.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

excelOptions := compdf.NewExcelOptions()
err := compdf.StartPDFToExcel(inputFilePath, password, outputFileName, excelOptions, nil)
```
## 3.15 Convert PDF to PowerPoint

#### **Overview**

ComPDF Conversion SDK provides the function of converting PDF files to PowerPoint files and restoring the layout and format of the original document, which can meet the needs of users for the presentation and editing of document content in Microsoft PowerPoint.

#### **Sample**

This sample demonstrates how to convert from a PDF to PPTX file.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

pptOptions := compdf.NewPPTOptions()
err := compdf.StartPDFToPPT(inputFilePath, password, outputFileName, pptOptions, nil)
```
## 3.16 Convert PDF to HTML

#### **Overview**

ComPDF Conversion SDK provides the PDF to HTML function, which can convert PDF files to HTML files while maintaining the layout and format of the original document, allowing users to browse and view the document on Web.

#### **Notice**

When converting PDF to HTML format, ComPDF Conversion SDK provides the following four options to create HTML files:

| Option | Description |
| --- | --- |
| `HtmlSinglePage` | Convert the entire PDF file into a single HTML file, where all PDF pages are connected in sequence by page number, displayed on the same HTML page. |
| `HtmlSinglePageWithBookmark` | Convert the PDF file into a single HTML file with all PDF page content on the same page in sequential order. This option also adds a bookmark navigation bar to the HTML page for quick navigation to a specific page. |
| `HtmlMultiPage` | Convert the PDF file into multiple HTML files. Each HTML file corresponds to a PDF page, and users can navigate to the next page via a link at the bottom of the page. |
| `HtmlMultiPageWithBookmark` | Convert the PDF file into multiple HTML files, each corresponding to a PDF page. An outline HTML page is provided for navigation, allowing users to jump to the corresponding HTML page by clicking the outline. |


#### **Sample**

This sample demonstrates how to convert from a PDF to HTML file.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

htmlOptions := compdf.NewHtmlOptions()
err := compdf.StartPDFToHtml(inputFilePath, password, outputFileName, htmlOptions, nil)
```
## 3.17 Convert PDF to CSV

#### **Overview**

ComPDF Conversion SDK supports converting PDF documents to CSV (Comma-Separated Values). Converting PDF to CSV is a common need, usually used to extract tabular or structured data from PDF documents and convert it into CSV files.

#### **Set Whether to Automatically Create Folders**
When multiple CSV files may be output, you can control whether to automatically create folders to store the CSV files by setting the `AutoCreateFolder` option. When this option is enabled, a folder with the same name as the output file will be automatically created in the output path to store the CSV files.

#### **Sample**

This sample demonstrates how to convert from a PDF to CSV file.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

csvOptions := compdf.NewCsvOptions()
err := compdf.StartPDFToCsv(inputFilePath, password, outputFileName, csvOptions, nil)
```
## 3.18 Convert PDF to Image

#### **Overview**

ComPDF Conversion SDK provides an API for converting PDFs to images. Integrate ComPDF Conversion SDK to your apps to convert PDF into images easily. 

#### **Setting Image Formats**

In ComPDF Conversion SDK, supported image formats include:

- JPG
- JPEG
- JPEG2000
- PNG
- BMP
- TIFF
- TGA
- GIF
- WEBP

#### **Setting Image Color Modes**

Supported image color modes in ComPDF Conversion SDK include:

- **Color:** Color mode, where the image effect is consistent with the original PDF page.
- **Gray:** Grayscale mode.
- **Binary:** Black and white mode, which applies binarization to the original effect.

#### **Setting Image Scaling**

The SDK supports setting image scaling. The default scaling is 1.0, which maintains the original PDF page size. If you want to double the image size, you can set `ImageScaling` to 2.0; similarly, to reduce the image size by half, set `ImageScaling` to 0.5.

#### **Enhancing Image Path Display**

The SDK supports an option called `ImagePathEnhance` for enhancing the display of image paths. This option can be enabled when you want to enhance the display effect of paths within the PDF page.

#### **Notice**

- A higher `ImageScaling` value results in images with higher resolution, but it also increases memory usage and slows down the conversion.
- A higher `ImageScaling` value does not necessarily equate to higher clarity; the clarity also depends on the original image resolution in the document.

#### **Sample**

This sample demonstrates how to convert from a PDF to Image file.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

imageOptions := compdf.NewImageOptions()
err := compdf.StartPDFToImage(inputFilePath, password, outputFileName, imageOptions, nil)
```
## 3.19 Convert PDF to RTF

#### **Overview**

RTF is a popular text format that can retain the format and style data of the text, and it is convenient for most text readers to read and write. Integrate ComPDF Conversion SDK to convert PDFs to RTF files now.

#### **Sample**

This sample demonstrates how to convert from a PDF to RTF file.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

rtfOptions := compdf.NewRtfOptions()
err := compdf.StartPDFToRtf(inputFilePath, password, outputFileName, rtfOptions, nil)
```
## 3.20 Convert PDF to TXT

#### **Overview**

When you need to extract the text content in the PDF file, in order for data analysis, text mining, information retrieval, etc. Using ComPDF Conversion SDK, you can easily extract the text in the PDF into the TXT file.

#### **Preserving Table Format**

The SDK supports an option called `TxtTableFormat` that preserves the table format when writing the TXT file, meaning that the original table structure is maintained. It is generally recommended to enable this option, especially for data extraction scenarios.

#### **Sample**

This sample demonstrates how to convert from a PDF to TXT file.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

txtOptions := compdf.NewTxtOptions()
err := compdf.StartPDFToTxt(inputFilePath, password, outputFileName, txtOptions, nil)
```
## 3.21 Convert PDF to Searchable PDF

#### **Overview**

To make a searchable PDF by adding invisible text to an image based PDF such as a scanned document using OCR.

#### **Set Transparent Text Layer**

When outputting a searchable PDF, you can use the following option to control the hidden text layer:

- `TransparentText`: Whether to output a transparent text layer.

#### **Sample**

Full code sample which shows how to use the ComPDF OCR module on scanned documents in multiple languages.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

pdfOptions := compdf.NewSearchablePdfOptions()
pdfOptions.EnableOCR = true
err := compdf.StartPDFToSearchablePDF(inputFilePath, password, outputFileName, pdfOptions, nil)
```
## 3.22 Convert PDF to OFD

#### **Overview**

ComPDF Conversion SDK supports converting PDF documents to OFD documents. Similar to searchable PDF, OFD conversion also supports OCR, page background preservation, and transparent text layers.

#### **Sample**

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

ofdOptions := compdf.NewOfdOptions()
ofdOptions.EnableOCR = true
ofdOptions.Languages = []compdf.OCRLanguage{compdf.OCRLangEnglish}
err := compdf.StartPDFToOfd(inputFilePath, password, outputFileName, ofdOptions, nil)
```
## 3.23 Release library resources

**Overview**

Release the files and memory resources occupied by ComPDF Conversion SDK.

**Notice**

- After calling this interface to release library resources, the ComPDF Conversion SDK will no longer function properly and must be reloaded.

- When you only want to release the resources occupied by the AI model rather than all resources used by the ComPDF Conversion SDK, you can achieve this by calling the `ReleaseDocumentAIModel` interface.


**Sample**

```go
compdf.ReleaseDocumentAIModel()
compdf.Release()
```
# 4. Data Extraction Guides

Unleash the Power of Data with ComPDF Conversion SDK's Data Extraction to detect, recognize, analyze, and
extract the PDF text, image, table, etc.
## 4.1 Extract PDF to JSON

#### **Overview**

Extract text, tables, and images from PDF documents to a JSON file.

#### **Standard table and non-standard table**

Commonly, tables can be divided into two categories: standard tables and non-standard tables. The specific definitions are as follows:

- Standard table: The table border and the inner lines of the table are complete and clear. There is no need to manually add table lines to divide the table content.

- **Non-Standard Tables:** Tables lacking borders or clear inner lines, requiring manual additions of table lines to separate contents.

#### **Table Extraction Option**

ComPDF Conversion SDK supports the option `ContainTable`. When enabled, table content is extracted from PDFs together with table structure; otherwise, table content is treated as regular text.

#### **Notice**

- Without enabling AI layout analysis or OCR options, tables in the original PDF cannot be extracted. It is recommended to enable AI layout analysis or OCR for high-precision table recognition.

#### **Sample**

Full sample code which illustrates the text extraction capabilities.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

jsonOptions := compdf.NewJsonOptions()
err := compdf.StartPDFToJson(inputFilePath, password, outputFileName, jsonOptions, nil)
```
## 4.2 Extract PDF to Markdown

#### **Overview**

Extract text, tables and images from PDF documents to Markdown file.

#### **Sample**

Full code sample which shows how to convert from a PDF to Markdown file.

```go
inputFilePath := "***"
password := "***"
outputFileName := "***"

markdownOptions := compdf.NewMarkdownOptions()
err := compdf.StartPDFToMarkdown(inputFilePath, password, outputFileName, markdownOptions, nil)
```
# 5. Support
## 5.1 FAQ


- Does OCR work on x86 architecture?

  Currently, OCR only works on x64 architecture.
## 5.2 Contact Us

Thanks for your interest in ComPDF Conversion SDK, an easy-to-use and powerful development solution. If you encounter technical questions or bug issues when using ComPDF Conversion SDK, please submit a problem report to the [ComPDF team](mailto:support@compdf.com). The following information will help us solve your problem more efficiently:

- ComPDF Conversion SDK product and version.
- Your operating system and IDE version.
- Detailed descriptions of the problem.
- Any other related information, such as an error screenshot.

#### **Contact Information**

- Home link: [https://www.compdf.com](https://www.compdf.com/)

- Technical Support: https://www.compdf.com/support
- Email: [support@compdf.com](mailto:support@compdf.com)

Thanks,
The ComPDF Team