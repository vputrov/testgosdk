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

// Effect format
type EffectFormat struct {

	// blur effect
	Blur *BlurEffect `json:"Blur,omitempty"`

	// glow effect
	Glow *GlowEffect `json:"Glow,omitempty"`

	// inner shadow effect
	InnerShadow *InnerShadowEffect `json:"InnerShadow,omitempty"`

	// outer shadow effect
	OuterShadow *OuterShadowEffect `json:"OuterShadow,omitempty"`

	// preset shadow effect
	PresetShadow *PresetShadowEffect `json:"PresetShadow,omitempty"`

	// soft edge effect
	SoftEdge *SoftEdgeEffect `json:"SoftEdge,omitempty"`

	// reflection effect
	Reflection *ReflectionEffect `json:"Reflection,omitempty"`

	// fill overlay effect
	FillOverlay *FillOverlayEffect `json:"FillOverlay,omitempty"`
}
