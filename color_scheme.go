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

// Slide's color scheme DTO
type ColorScheme struct {

	// Gets or sets the link to this resource.
	SelfUri *ResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// A list of links that originate from this document.
	Links []ResourceUri `json:"Links,omitempty"`

	Accent1 string `json:"Accent1,omitempty"`

	Accent2 string `json:"Accent2,omitempty"`

	Accent3 string `json:"Accent3,omitempty"`

	Accent4 string `json:"Accent4,omitempty"`

	Accent5 string `json:"Accent5,omitempty"`

	Accent6 string `json:"Accent6,omitempty"`

	Dark1 string `json:"Dark1,omitempty"`

	Dark2 string `json:"Dark2,omitempty"`

	FollowedHyperlink string `json:"FollowedHyperlink,omitempty"`

	Hyperlink string `json:"Hyperlink,omitempty"`

	Light1 string `json:"Light1,omitempty"`

	Light2 string `json:"Light2,omitempty"`
}
