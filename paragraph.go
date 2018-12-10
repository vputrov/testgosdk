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

// Represents paragraph resource
type Paragraph struct {

	// Gets or sets the link to this resource.
	SelfUri *ResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// A list of links that originate from this document.
	Links []ResourceUri `json:"Links,omitempty"`

	MarginLeft float64 `json:"MarginLeft,omitempty"`

	MarginRight float64 `json:"MarginRight,omitempty"`

	SpaceBefore float64 `json:"SpaceBefore,omitempty"`

	SpaceAfter float64 `json:"SpaceAfter,omitempty"`

	SpaceWithin float64 `json:"SpaceWithin,omitempty"`

	Indent float64 `json:"Indent,omitempty"`

	Alignment TextAlignment `json:"Alignment,omitempty"`

	FontAlignment FontAlignment `json:"FontAlignment,omitempty"`

	DefaultTabSize float64 `json:"DefaultTabSize,omitempty"`

	Depth int32 `json:"Depth,omitempty"`

	BulletChar string `json:"BulletChar,omitempty"`

	BulletHeight float64 `json:"BulletHeight,omitempty"`

	BulletType BulletType `json:"BulletType,omitempty"`

	NumberedBulletStartWith int32 `json:"NumberedBulletStartWith,omitempty"`

	NumberedBulletStyle NumberedBulletStyle `json:"NumberedBulletStyle,omitempty"`

	HangingPunctuation int32 `json:"HangingPunctuation,omitempty"`

	EastAsianLineBreak int32 `json:"EastAsianLineBreak,omitempty"`

	LatinLineBreak int32 `json:"LatinLineBreak,omitempty"`

	RightToLeft int32 `json:"RightToLeft,omitempty"`

	PortionList []ResourceUriElement `json:"PortionList,omitempty"`
}
