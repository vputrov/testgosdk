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

type SvgExportOptions struct {

	Format string `json:"Format,omitempty"`

	// Determines whether the text on a slide will be saved as graphics.
	VectorizeText bool `json:"VectorizeText,omitempty"`

	// Returns or sets the lower resolution limit for metafile rasterization.
	MetafileRasterizationDpi int32 `json:"MetafileRasterizationDpi,omitempty"`

	// Determines whether the 3D text is disabled in SVG.
	Disable3DText bool `json:"Disable3DText,omitempty"`

	// Disables splitting FromCornerX and FromCenter gradients.
	DisableGradientSplit bool `json:"DisableGradientSplit,omitempty"`

	// SVG 1.1 lacks ability to define insets for markers. Aspose.Slides SVG writing engine has workaround for that problem: it crops end of line with arrow, so, line doesn't overlap markers. This option switches off such behavior.
	DisableLineEndCropping bool `json:"DisableLineEndCropping,omitempty"`

	// Determines JPEG encoding quality.
	JpegQuality int32 `json:"JpegQuality,omitempty"`

	// Represents the pictures compression level
	PicturesCompression PicturesCompression `json:"PicturesCompression,omitempty"`

	// A boolean flag indicates if the cropped parts remain as part of the document. If true the cropped  parts will removed, if false they will be serialized in the document (which can possible lead to a  larger file)
	DeletePicturesCroppedAreas bool `json:"DeletePicturesCroppedAreas,omitempty"`

	// Determines a way of handling externally loaded fonts.
	ExternalFontsHandling ExternalFontsHandling `json:"ExternalFontsHandling,omitempty"`
}
