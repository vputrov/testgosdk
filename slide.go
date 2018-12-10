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

type Slide struct {

	// Gets or sets the link to this resource.
	SelfUri *ResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// A list of links that originate from this document.
	Links []ResourceUri `json:"Links,omitempty"`

	// Gets or sets the width.
	Width float64 `json:"Width,omitempty"`

	// Gets or sets the height.
	Height float64 `json:"Height,omitempty"`

	// Specifies if shapes of the master slide should be shown on the slide. True by default.
	ShowMasterShapes bool `json:"ShowMasterShapes,omitempty"`

	// Gets or sets the  link to the layout slide.
	LayoutSlide *ResourceUriElement `json:"LayoutSlide,omitempty"`

	// Gets or sets the  link to list of top-level shapes.
	Shapes *ResourceUriElement `json:"Shapes,omitempty"`

	// Gets or sets the link to theme.
	Theme *ResourceUriElement `json:"Theme,omitempty"`

	// Gets or sets the  link to placeholders.
	Placeholders *ResourceUriElement `json:"Placeholders,omitempty"`

	// Gets or sets the link to images.
	Images *ResourceUriElement `json:"Images,omitempty"`

	// Gets or sets the link to comments.
	Comments *ResourceUriElement `json:"Comments,omitempty"`

	// Get or sets the link to slide's background
	Background *ResourceUriElement `json:"Background,omitempty"`

	// Get or sets the link to notes slide.
	NotesSlide *ResourceUriElement `json:"NotesSlide,omitempty"`
}
