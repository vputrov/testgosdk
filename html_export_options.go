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

// Provides options that control how a presentation is saved in Html format.
type HtmlExportOptions struct {

	Format string `json:"Format,omitempty"`

	// Get or sets flag for save presentation as zip file
	SaveAsZip bool `json:"SaveAsZip,omitempty"`

	// Get or set name of subdirectory in zip-file for store external files
	SubDirectoryName string `json:"SubDirectoryName,omitempty"`

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	ShowHiddenSlides bool `json:"ShowHiddenSlides,omitempty"`

	// Returns or sets a value determining the quality of the JPEG images inside PDF document.
	JpegQuality int32 `json:"JpegQuality,omitempty"`

	// Represents the pictures compression level
	PicturesCompression PicturesCompression `json:"PicturesCompression,omitempty"`

	// A boolean flag indicates if the cropped parts remain as part of the document. If true the cropped  parts will removed, if false they will be serialized in the document (which can possible lead to a  larger file)
	DeletePicturesCroppedAreas bool `json:"DeletePicturesCroppedAreas,omitempty"`

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
}
