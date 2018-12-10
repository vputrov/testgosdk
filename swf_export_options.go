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

// Provides options that control how a presentation is saved in SWF format.
type SwfExportOptions struct {

	Format string `json:"Format,omitempty"`

	// Specifies whether the generated document should include hidden slides or not. Default is false. 
	ShowHiddenSlides bool `json:"ShowHiddenSlides,omitempty"`

	// Specifies whether the generated SWF document should be compressed or not. Default is true. 
	Compressed bool `json:"Compressed,omitempty"`

	// Specifies whether the generated SWF document should include the integrated document viewer or not. Default is true. 
	ViewerIncluded bool `json:"ViewerIncluded,omitempty"`

	// Specifies whether border around pages should be shown. Default is true. 
	ShowPageBorder bool `json:"ShowPageBorder,omitempty"`

	// Show/hide fullscreen button. Can be overridden in flashvars. Default is true. 
	ShowFullScreen bool `json:"ShowFullScreen,omitempty"`

	// Show/hide page stepper. Can be overridden in flashvars. Default is true. 
	ShowPageStepper bool `json:"ShowPageStepper,omitempty"`

	// Show/hide search section. Can be overridden in flashvars. Default is true. 
	ShowSearch bool `json:"ShowSearch,omitempty"`

	// Show/hide whole top pane. Can be overridden in flashvars. Default is true. 
	ShowTopPane bool `json:"ShowTopPane,omitempty"`

	// Show/hide bottom pane. Can be overridden in flashvars. Default is true. 
	ShowBottomPane bool `json:"ShowBottomPane,omitempty"`

	// Show/hide left pane. Can be overridden in flashvars. Default is true. 
	ShowLeftPane bool `json:"ShowLeftPane,omitempty"`

	// Start with opened left pane. Can be overridden in flashvars. Default is false. 
	StartOpenLeftPane bool `json:"StartOpenLeftPane,omitempty"`

	// Enable/disable context menu. Default is true. 
	EnableContextMenu bool `json:"EnableContextMenu,omitempty"`

	// Image that will be displayed as logo in the top right corner of the viewer. The image data is a base 64 string. Image should be 32x64 pixels PNG image, otherwise logo can be displayed improperly. 
	LogoImage string `json:"LogoImage,omitempty"`

	// Gets or sets the full hyperlink address for a logo. Has an effect only if a  is specified. 
	LogoLink string `json:"LogoLink,omitempty"`

	// Specifies the quality of JPEG images. Default is 95.
	JpegQuality int32 `json:"JpegQuality,omitempty"`

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
