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

// Represents one cell of table.
type TableCell struct {

	// Cell text.
	Text string `json:"Text,omitempty"`

	// The number of rows spanned by a merged cell.
	RowSpan int32 `json:"RowSpan"`

	// The number of columns spanned by a merged cell.
	ColSpan int32 `json:"ColSpan"`

	// The top margin of the cell.
	MarginTop float64 `json:"MarginTop"`

	// The right margin of the cell.
	MarginRight float64 `json:"MarginRight"`

	// The left margin of the cell.
	MarginLeft float64 `json:"MarginLeft"`

	// The bottom margin of the cell.
	MarginBottom float64 `json:"MarginBottom"`

	// Text anchor type.
	TextAnchorType TextAnchorType `json:"TextAnchorType"`

	// The type of vertical text.
	TextVerticalType TextVerticalType `json:"TextVerticalType"`

	// Fill properties set of the cell.
	FillFormat *interface{} `json:"FillFormat,omitempty"`

	// Line properties set for the top border of the cell.
	BorderTop *LineFormat `json:"BorderTop,omitempty"`

	// Line properties set for the right border of the cell.
	BorderRight *LineFormat `json:"BorderRight,omitempty"`

	// Line properties set for the left border of the cell.
	BorderLeft *LineFormat `json:"BorderLeft,omitempty"`

	// Line properties set for the bottom border of the cell.
	BorderBottom *LineFormat `json:"BorderBottom,omitempty"`

	// Line properties set for the diagonal up border of the cell.
	BorderDiagonalUp *LineFormat `json:"BorderDiagonalUp,omitempty"`

	// Line properties set for the diagonal down border of the cell.
	BorderDiagonalDown *LineFormat `json:"BorderDiagonalDown,omitempty"`
}
