/*
 * --------------------------------------------------------------------------------------------------------------------
 * <copyright company="Aspose">
 *   Copyright (c) 2018 Aspose.Slides for Cloud
 * </copyright>
 * <summary>
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 * 
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 * 
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 * </summary>
 * --------------------------------------------------------------------------------------------------------------------
 */

package asposeslidescloud

// Provides options that control how a presentation is saved in Pdf format.
type PdfExportOptions struct {

	Format string `json:"Format,omitempty"`

	// Specifies compression type to be used for all textual content in the document.
	TextCompression PdfTextCompression `json:"TextCompression,omitempty"`

	// Determines if all characters of font should be embedded or only used subset.
	EmbedFullFonts bool `json:"EmbedFullFonts,omitempty"`

	// Desired conformance level for generated PDF document.
	Compliance PdfCompliance `json:"Compliance,omitempty"`

	// Returns or sets a value determining resolution of images inside PDF document. Property affects on file size, time of export and image quality.The default value is 96.
	SufficientResolution float64 `json:"SufficientResolution,omitempty"`

	// Returns or sets a value determining the quality of the JPEG images inside PDF document.
	JpegQuality int32 `json:"JpegQuality,omitempty"`

	// True to draw black frame around each slide.
	DrawSlidesFrame bool `json:"DrawSlidesFrame,omitempty"`

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	ShowHiddenSlides bool `json:"ShowHiddenSlides,omitempty"`

	// True to convert all metafiles used in a presentation to the PNG images.
	SaveMetafilesAsPng bool `json:"SaveMetafilesAsPng,omitempty"`

	// Setting user password to protect the PDF document. 
	Password string `json:"Password,omitempty"`

	// Determines if Aspose.Slides will embed common fonts for ASCII (33..127 code range) text. Fonts for character codes greater than 127 are always embedded. Common fonts list includes PDF's base 14 fonts and additional user specified fonts.
	EmbedTrueTypeFontsForASCII bool `json:"EmbedTrueTypeFontsForASCII,omitempty"`

	// Returns or sets an array of user-defined names of font families which Aspose.Slides should consider common.
	AdditionalCommonFontFamilies []string `json:"AdditionalCommonFontFamilies,omitempty"`

	// Gets or sets the position of the notes on the page.
	NotesPosition NotesPositions `json:"NotesPosition,omitempty"`

	// Gets or sets the position of the comments on the page.
	CommentsPosition CommentsPositions `json:"CommentsPosition,omitempty"`

	// Gets or sets the width of the comment output area in pixels (Applies only if comments are displayed on the right).
	CommentsAreaWidth int32 `json:"CommentsAreaWidth,omitempty"`

	// Gets or sets the color of comments area (Applies only if comments are displayed on the right).
	CommentsAreaColor string `json:"CommentsAreaColor,omitempty"`

	// True if comments that have no author are displayed. (Applies only if comments are displayed).
	ShowCommentsByNoAuthor bool `json:"ShowCommentsByNoAuthor,omitempty"`

	// Image transparent color.
	ImageTransparentColor string `json:"ImageTransparentColor,omitempty"`

	// True to apply specified   to an image.
	ApplyImageTransparent bool `json:"ApplyImageTransparent,omitempty"`
}
