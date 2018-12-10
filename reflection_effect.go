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

// Represents reflection effect 
type ReflectionEffect struct {

	// direction
	Direction float64 `json:"Direction"`

	// fade direction
	FadeDirection float64 `json:"FadeDirection"`

	// distance
	Distance float64 `json:"Distance"`

	// blur radius
	BlurRadius float64 `json:"BlurRadius"`

	// scale horizontal
	ScaleHorizontal float64 `json:"ScaleHorizontal"`

	// scale vertical
	ScaleVertical float64 `json:"ScaleVertical"`

	// skew horizontal
	SkewHorizontal float64 `json:"SkewHorizontal"`

	// skew vertical
	SkewVertical float64 `json:"SkewVertical"`

	// start pos alpha
	StartPosAlpha float64 `json:"StartPosAlpha"`

	// end pos alpha
	EndPosAlpha float64 `json:"EndPosAlpha"`

	// start reflection opacity
	StartReflectionOpacity float64 `json:"StartReflectionOpacity"`

	// end reflection opacity
	EndReflectionOpacity float64 `json:"EndReflectionOpacity"`

	// rectangle alignment
	RectangleAlign RectangleAlignment `json:"RectangleAlign"`

	// true if the reflection should rotate with the shape when the shape is rotated
	RotateShadowWithShape bool `json:"RotateShadowWithShape"`
}
