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

// Represents Table shape resource.
type Table struct {

	// Gets or sets the link to this resource.
	SelfUri *ResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// A list of links that originate from this document.
	Links []ResourceUri `json:"Links,omitempty"`

	// Shape type.
	Type_ string `json:"Type,omitempty"`

	// Combined shape type.
	ShapeType string `json:"ShapeType,omitempty"`

	// Gets or sets the name.
	Name string `json:"Name,omitempty"`

	// Gets or sets the width.
	Width float64 `json:"Width,omitempty"`

	// Gets or sets the height.
	Height float64 `json:"Height,omitempty"`

	// Gets or sets the alternative text.
	AlternativeText string `json:"AlternativeText,omitempty"`

	// Gets or sets a value indicating whether this  is hidden.
	Hidden bool `json:"Hidden,omitempty"`

	// Gets or sets the X
	X float64 `json:"X,omitempty"`

	// Gets or sets the Y.
	Y float64 `json:"Y,omitempty"`

	// Gets z-order position of shape
	ZOrderPosition int32 `json:"ZOrderPosition,omitempty"`

	// Gets or sets the link to shapes.
	Shapes *ResourceUriElement `json:"Shapes,omitempty"`

	// Gets or sets the fill format.
	FillFormat *interface{} `json:"FillFormat,omitempty"`

	// Gets or sets the effect format.
	EffectFormat *EffectFormat `json:"EffectFormat,omitempty"`

	// Gets or sets the line format.
	LineFormat *LineFormat `json:"LineFormat,omitempty"`

	// Builtin table style.
	Style TableStylePreset `json:"Style,omitempty"`

	// Rows.
	Rows []TableRow `json:"Rows,omitempty"`

	// Columns.
	Columns []TableColumn `json:"Columns,omitempty"`

	// Determines whether the first column of a table has to be drawn with a special formatting.
	FirstCol bool `json:"FirstCol,omitempty"`

	// Determines whether the first row of a table has to be drawn with a special formatting.
	FirstRow bool `json:"FirstRow,omitempty"`

	// Determines whether the even rows has to be drawn with a different formatting.
	HorizontalBanding bool `json:"HorizontalBanding,omitempty"`

	// Determines whether the last column of a table has to be drawn with a special formatting.
	LastCol bool `json:"LastCol,omitempty"`

	// Determines whether the last row of a table has to be drawn with a special formatting.
	LastRow bool `json:"LastRow,omitempty"`

	// Determines whether the table has right to left reading order.
	RightToLeft bool `json:"RightToLeft,omitempty"`

	// Determines whether the even columns has to be drawn with a different formatting.
	VerticalBanding bool `json:"VerticalBanding,omitempty"`
}
