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

// Represents a chart axis
type Axis struct {

	// True if the axis is visible
	IsVisible bool `json:"IsVisible"`

	// True if the axis has a visible title
	HasTitle bool `json:"HasTitle"`

	// Axis position
	Position AxisPositionType `json:"Position"`

	// The scaling value of the display units for the value axis
	DisplayUnit DisplayUnitType `json:"DisplayUnit"`

	// The smallest time unit that is represented on the date axis
	BaseUnitScale TimeUnitType `json:"BaseUnitScale"`

	// True the major unit of the axis is automatically assigned
	IsAutomaticMajorUnit bool `json:"IsAutomaticMajorUnit"`

	// The major units for the date or value axis
	MajorUnit float64 `json:"MajorUnit"`

	// The major unit scale for the date axis
	MajorUnitScale TimeUnitType `json:"MajorUnitScale"`

	// The type of major tick mark for the specified axis
	MajorTickMark TickMarkType `json:"MajorTickMark"`

	// True the minor unit of the axis is automatically assigned
	IsAutomaticMinorUnit bool `json:"IsAutomaticMinorUnit"`

	// The minor units for the date or value axis
	MinorUnit float64 `json:"MinorUnit"`

	// The minor unit scale for the date axis
	MinorUnitScale TimeUnitType `json:"MinorUnitScale"`

	// The type of minor tick mark for the specified axis
	MinorTickMark TickMarkType `json:"MinorTickMark"`

	// True if the max value is automatically assigned
	IsAutomaticMaxValue bool `json:"IsAutomaticMaxValue"`

	// The maximum value on the value axis
	MaxValue float64 `json:"MaxValue"`

	// True if the min value is automatically assigned
	IsAutomaticMinValue bool `json:"IsAutomaticMinValue"`

	// The minimum value on the value axis
	MinValue float64 `json:"MinValue"`

	// True if the value axis scale type is logarithmic
	IsLogarithmic bool `json:"IsLogarithmic"`

	// The logarithmic base. Default value is 10
	LogBase float64 `json:"LogBase"`

	// The type of the category axis
	CategoryAxisType CategoryAxisType `json:"CategoryAxisType"`

	// True if the value axis crosses the category axis between categories. This property applies only to category axes, and it doesn't apply to 3-D charts
	AxisBetweenCategories bool `json:"AxisBetweenCategories"`

	// The distance of labels from the axis. Applied to category or date axis. Value must be between 0% and 1000%.             
	LabelOffset int32 `json:"LabelOffset"`

	// True if MS PowerPoint plots data points from last to first
	IsPlotOrderReversed bool `json:"IsPlotOrderReversed"`

	// True if the format is linked to source data
	IsNumberFormatLinkedToSource bool `json:"IsNumberFormatLinkedToSource"`

	// the format string for the Axis Labels
	NumberFormat string `json:"NumberFormat,omitempty"`

	// The CrossType on the specified axis where the other axis crosses
	CrossType CrossesType `json:"CrossType"`

	// The point on the axis where the perpendicular axis crosses it
	CrossAt float64 `json:"CrossAt"`

	// True for automatic tick marks spacing value
	IsAutomaticTickMarksSpacing bool `json:"IsAutomaticTickMarksSpacing"`

	// Specifies how many tick marks shall be skipped before the next one shall be drawn. Applied to category or series axis.
	TickMarksSpacing int32 `json:"TickMarksSpacing"`

	// True for automatic tick label spacing value
	IsAutomaticTickLabelSpacing bool `json:"IsAutomaticTickLabelSpacing"`

	// Specifies how many tick labels to skip between label that is drawn.
	TickLabelSpacing int32 `json:"TickLabelSpacing"`

	// The position of tick-mark labels on the specified axis.
	TickLabelPosition TickLabelPositionType `json:"TickLabelPosition"`

	// Represents the rotation angle of tick labels.
	TickLabelRotationAngle float64 `json:"TickLabelRotationAngle"`

	// Get or sets the fill format.
	FillFormat *interface{} `json:"FillFormat,omitempty"`

	// Get or sets the effect format.
	EffectFormat *EffectFormat `json:"EffectFormat,omitempty"`

	// Get or sets the line format.
	LineFormat *LineFormat `json:"LineFormat,omitempty"`
}
