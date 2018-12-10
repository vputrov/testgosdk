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

// Represents portion resource
type Portion struct {

	// Gets or sets the link to this resource.
	SelfUri *ResourceUri `json:"SelfUri,omitempty"`

	AlternateLinks []ResourceUri `json:"AlternateLinks,omitempty"`

	// A list of links that originate from this document.
	Links []ResourceUri `json:"Links,omitempty"`

	Text string `json:"Text,omitempty"`

	FontBold int32 `json:"FontBold,omitempty"`

	FontItalic int32 `json:"FontItalic,omitempty"`

	FontUnderline TextUnderlineType `json:"FontUnderline,omitempty"`

	StrikethroughType TextStrikethroughType `json:"StrikethroughType,omitempty"`

	TextCapType TextCapType `json:"TextCapType,omitempty"`

	Escapement float64 `json:"Escapement,omitempty"`

	Spacing float64 `json:"Spacing,omitempty"`

	FontColor string `json:"FontColor,omitempty"`

	HighlightColor string `json:"HighlightColor,omitempty"`

	FontHeight float64 `json:"FontHeight,omitempty"`

	NormaliseHeight int32 `json:"NormaliseHeight,omitempty"`

	ProofDisabled int32 `json:"ProofDisabled,omitempty"`

	SmartTagClean bool `json:"SmartTagClean,omitempty"`

	KerningMinimalSize float64 `json:"KerningMinimalSize,omitempty"`

	Kumimoji int32 `json:"Kumimoji,omitempty"`

	LanguageId string `json:"LanguageId,omitempty"`

	AlternativeLanguageId string `json:"AlternativeLanguageId,omitempty"`

	IsHardUnderlineFill int32 `json:"IsHardUnderlineFill,omitempty"`

	IsHardUnderlineLine int32 `json:"IsHardUnderlineLine,omitempty"`

	FillFormat *interface{} `json:"FillFormat,omitempty"`

	EffectFormat *EffectFormat `json:"EffectFormat,omitempty"`

	LineFormat *LineFormat `json:"LineFormat,omitempty"`

	UnderlineFillFormat *interface{} `json:"UnderlineFillFormat,omitempty"`

	UnderlineLineFormat *LineFormat `json:"UnderlineLineFormat,omitempty"`
}
