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

import (
	"testing"
)

/* SlidesApiServiceTests Delete presentation slide by its index.
   Test for SlidesApi.DeleteSlideByIndex method
*/
func TestDeleteSlideByIndex(t *testing.T) {
    request := createDeleteSlideByIndexRequest()
    e := initializeTest("DeleteSlideByIndex", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.DeleteSlideByIndex(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createDeleteSlideByIndexRequest() DeleteSlideByIndexRequest {
    var request DeleteSlideByIndexRequest
    request.name = createTestParamValue("DeleteSlideByIndex", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteSlideByIndex", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("DeleteSlideByIndex", "password", "string").(string)
    request.folder = createTestParamValue("DeleteSlideByIndex", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteSlideByIndex", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Delete presentation slide by its index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid name
*/
func TestDeleteSlideByIndexInvalidname(t *testing.T) {
    request := createDeleteSlideByIndexRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("DeleteSlideByIndex", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    assertError(t, "DeleteSlideByIndex", "name", r.Code, e)
}

/* SlidesApiServiceTests Delete presentation slide by its index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid slideIndex
*/
func TestDeleteSlideByIndexInvalidslideIndex(t *testing.T) {
    request := createDeleteSlideByIndexRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("DeleteSlideByIndex", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    assertError(t, "DeleteSlideByIndex", "slideIndex", r.Code, e)
}

/* SlidesApiServiceTests Delete presentation slide by its index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid password
*/
func TestDeleteSlideByIndexInvalidpassword(t *testing.T) {
    request := createDeleteSlideByIndexRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("DeleteSlideByIndex", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    assertError(t, "DeleteSlideByIndex", "password", r.Code, e)
}

/* SlidesApiServiceTests Delete presentation slide by its index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid folder
*/
func TestDeleteSlideByIndexInvalidfolder(t *testing.T) {
    request := createDeleteSlideByIndexRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("DeleteSlideByIndex", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    assertError(t, "DeleteSlideByIndex", "folder", r.Code, e)
}

/* SlidesApiServiceTests Delete presentation slide by its index.
   Test for SlidesApi.DeleteSlideByIndex method with invalid storage
*/
func TestDeleteSlideByIndexInvalidstorage(t *testing.T) {
    request := createDeleteSlideByIndexRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("DeleteSlideByIndex", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.DeleteSlideByIndex(request)
    assertError(t, "DeleteSlideByIndex", "storage", r.Code, e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method
*/
func TestDeleteSlidesCleanSlidesList(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()
    e := initializeTest("DeleteSlidesCleanSlidesList", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.DeleteSlidesCleanSlidesList(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createDeleteSlidesCleanSlidesListRequest() DeleteSlidesCleanSlidesListRequest {
    var request DeleteSlidesCleanSlidesListRequest
    request.name = createTestParamValue("DeleteSlidesCleanSlidesList", "name", "string").(string)
    request.slides = createTestParamValue("DeleteSlidesCleanSlidesList", "slides", "[]int32").([]int32)
    request.password = createTestParamValue("DeleteSlidesCleanSlidesList", "password", "string").(string)
    request.folder = createTestParamValue("DeleteSlidesCleanSlidesList", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteSlidesCleanSlidesList", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid name
*/
func TestDeleteSlidesCleanSlidesListInvalidname(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("DeleteSlidesCleanSlidesList", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    assertError(t, "DeleteSlidesCleanSlidesList", "name", r.Code, e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid slides
*/
func TestDeleteSlidesCleanSlidesListInvalidslides(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()
    request.slides = invalidizeTestParamValue(request.slides, "slides", "[]int32").([]int32)
    e := initializeTest("DeleteSlidesCleanSlidesList", "slides", request.slides)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    assertError(t, "DeleteSlidesCleanSlidesList", "slides", r.Code, e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid password
*/
func TestDeleteSlidesCleanSlidesListInvalidpassword(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("DeleteSlidesCleanSlidesList", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    assertError(t, "DeleteSlidesCleanSlidesList", "password", r.Code, e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid folder
*/
func TestDeleteSlidesCleanSlidesListInvalidfolder(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("DeleteSlidesCleanSlidesList", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    assertError(t, "DeleteSlidesCleanSlidesList", "folder", r.Code, e)
}

/* SlidesApiServiceTests Delete presentation slides.
   Test for SlidesApi.DeleteSlidesCleanSlidesList method with invalid storage
*/
func TestDeleteSlidesCleanSlidesListInvalidstorage(t *testing.T) {
    request := createDeleteSlidesCleanSlidesListRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("DeleteSlidesCleanSlidesList", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.DeleteSlidesCleanSlidesList(request)
    assertError(t, "DeleteSlidesCleanSlidesList", "storage", r.Code, e)
}

/* SlidesApiServiceTests Remove presentation slide background color.
   Test for SlidesApi.DeleteSlidesSlideBackground method
*/
func TestDeleteSlidesSlideBackground(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()
    e := initializeTest("DeleteSlidesSlideBackground", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.DeleteSlidesSlideBackground(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createDeleteSlidesSlideBackgroundRequest() DeleteSlidesSlideBackgroundRequest {
    var request DeleteSlidesSlideBackgroundRequest
    request.name = createTestParamValue("DeleteSlidesSlideBackground", "name", "string").(string)
    request.slideIndex = createTestParamValue("DeleteSlidesSlideBackground", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("DeleteSlidesSlideBackground", "password", "string").(string)
    request.folder = createTestParamValue("DeleteSlidesSlideBackground", "folder", "string").(string)
    request.storage = createTestParamValue("DeleteSlidesSlideBackground", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Remove presentation slide background color.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid name
*/
func TestDeleteSlidesSlideBackgroundInvalidname(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("DeleteSlidesSlideBackground", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    assertError(t, "DeleteSlidesSlideBackground", "name", r.Code, e)
}

/* SlidesApiServiceTests Remove presentation slide background color.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid slideIndex
*/
func TestDeleteSlidesSlideBackgroundInvalidslideIndex(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("DeleteSlidesSlideBackground", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    assertError(t, "DeleteSlidesSlideBackground", "slideIndex", r.Code, e)
}

/* SlidesApiServiceTests Remove presentation slide background color.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid password
*/
func TestDeleteSlidesSlideBackgroundInvalidpassword(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("DeleteSlidesSlideBackground", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    assertError(t, "DeleteSlidesSlideBackground", "password", r.Code, e)
}

/* SlidesApiServiceTests Remove presentation slide background color.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid folder
*/
func TestDeleteSlidesSlideBackgroundInvalidfolder(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("DeleteSlidesSlideBackground", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    assertError(t, "DeleteSlidesSlideBackground", "folder", r.Code, e)
}

/* SlidesApiServiceTests Remove presentation slide background color.
   Test for SlidesApi.DeleteSlidesSlideBackground method with invalid storage
*/
func TestDeleteSlidesSlideBackgroundInvalidstorage(t *testing.T) {
    request := createDeleteSlidesSlideBackgroundRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("DeleteSlidesSlideBackground", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.DeleteSlidesSlideBackground(request)
    assertError(t, "DeleteSlidesSlideBackground", "storage", r.Code, e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.GetSlideWithFormat method
*/
func TestGetSlideWithFormat(t *testing.T) {
    request := createGetSlideWithFormatRequest()
    e := initializeTest("GetSlideWithFormat", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.GetSlideWithFormat(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    assertBinaryResponse(r, t)
}

func createGetSlideWithFormatRequest() GetSlideWithFormatRequest {
    var request GetSlideWithFormatRequest
    request.name = createTestParamValue("GetSlideWithFormat", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlideWithFormat", "slideIndex", "int32").(int32)
    request.format = createTestParamValue("GetSlideWithFormat", "format", "string").(string)
    request.width = createTestParamValue("GetSlideWithFormat", "width", "int32").(int32)
    request.height = createTestParamValue("GetSlideWithFormat", "height", "int32").(int32)
    request.password = createTestParamValue("GetSlideWithFormat", "password", "string").(string)
    request.folder = createTestParamValue("GetSlideWithFormat", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlideWithFormat", "storage", "string").(string)
    request.outPath = createTestParamValue("GetSlideWithFormat", "outPath", "string").(string)
    request.fontsFolder = createTestParamValue("GetSlideWithFormat", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.GetSlideWithFormat method with invalid name
*/
func TestGetSlideWithFormatInvalidname(t *testing.T) {
    request := createGetSlideWithFormatRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlideWithFormat", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideWithFormat(request)
    assertError(t, "GetSlideWithFormat", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.GetSlideWithFormat method with invalid slideIndex
*/
func TestGetSlideWithFormatInvalidslideIndex(t *testing.T) {
    request := createGetSlideWithFormatRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetSlideWithFormat", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideWithFormat(request)
    assertError(t, "GetSlideWithFormat", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.GetSlideWithFormat method with invalid format
*/
func TestGetSlideWithFormatInvalidformat(t *testing.T) {
    request := createGetSlideWithFormatRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "string").(string)
    e := initializeTest("GetSlideWithFormat", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideWithFormat(request)
    assertError(t, "GetSlideWithFormat", "format", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.GetSlideWithFormat method with invalid width
*/
func TestGetSlideWithFormatInvalidwidth(t *testing.T) {
    request := createGetSlideWithFormatRequest()
    request.width = invalidizeTestParamValue(request.width, "width", "int32").(int32)
    e := initializeTest("GetSlideWithFormat", "width", request.width)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideWithFormat(request)
    assertError(t, "GetSlideWithFormat", "width", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.GetSlideWithFormat method with invalid height
*/
func TestGetSlideWithFormatInvalidheight(t *testing.T) {
    request := createGetSlideWithFormatRequest()
    request.height = invalidizeTestParamValue(request.height, "height", "int32").(int32)
    e := initializeTest("GetSlideWithFormat", "height", request.height)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideWithFormat(request)
    assertError(t, "GetSlideWithFormat", "height", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.GetSlideWithFormat method with invalid password
*/
func TestGetSlideWithFormatInvalidpassword(t *testing.T) {
    request := createGetSlideWithFormatRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlideWithFormat", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideWithFormat(request)
    assertError(t, "GetSlideWithFormat", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.GetSlideWithFormat method with invalid folder
*/
func TestGetSlideWithFormatInvalidfolder(t *testing.T) {
    request := createGetSlideWithFormatRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlideWithFormat", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideWithFormat(request)
    assertError(t, "GetSlideWithFormat", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.GetSlideWithFormat method with invalid storage
*/
func TestGetSlideWithFormatInvalidstorage(t *testing.T) {
    request := createGetSlideWithFormatRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlideWithFormat", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideWithFormat(request)
    assertError(t, "GetSlideWithFormat", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.GetSlideWithFormat method with invalid outPath
*/
func TestGetSlideWithFormatInvalidoutPath(t *testing.T) {
    request := createGetSlideWithFormatRequest()
    request.outPath = invalidizeTestParamValue(request.outPath, "outPath", "string").(string)
    e := initializeTest("GetSlideWithFormat", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideWithFormat(request)
    assertError(t, "GetSlideWithFormat", "outPath", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.GetSlideWithFormat method with invalid fontsFolder
*/
func TestGetSlideWithFormatInvalidfontsFolder(t *testing.T) {
    request := createGetSlideWithFormatRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)
    e := initializeTest("GetSlideWithFormat", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.GetSlideWithFormat(request)
    assertError(t, "GetSlideWithFormat", "fontsFolder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Read slide info.
   Test for SlidesApi.GetSlidesSlide method
*/
func TestGetSlidesSlide(t *testing.T) {
    request := createGetSlidesSlideRequest()
    e := initializeTest("GetSlidesSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.GetSlidesSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetSlidesSlideRequest() GetSlidesSlideRequest {
    var request GetSlidesSlideRequest
    request.name = createTestParamValue("GetSlidesSlide", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlidesSlide", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("GetSlidesSlide", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesSlide", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesSlide", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid name
*/
func TestGetSlidesSlideInvalidname(t *testing.T) {
    request := createGetSlidesSlideRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlidesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    assertError(t, "GetSlidesSlide", "name", r.Code, e)
}

/* SlidesApiServiceTests Read slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid slideIndex
*/
func TestGetSlidesSlideInvalidslideIndex(t *testing.T) {
    request := createGetSlidesSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetSlidesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    assertError(t, "GetSlidesSlide", "slideIndex", r.Code, e)
}

/* SlidesApiServiceTests Read slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid password
*/
func TestGetSlidesSlideInvalidpassword(t *testing.T) {
    request := createGetSlidesSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlidesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    assertError(t, "GetSlidesSlide", "password", r.Code, e)
}

/* SlidesApiServiceTests Read slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid folder
*/
func TestGetSlidesSlideInvalidfolder(t *testing.T) {
    request := createGetSlidesSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlidesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    assertError(t, "GetSlidesSlide", "folder", r.Code, e)
}

/* SlidesApiServiceTests Read slide info.
   Test for SlidesApi.GetSlidesSlide method with invalid storage
*/
func TestGetSlidesSlideInvalidstorage(t *testing.T) {
    request := createGetSlidesSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlidesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlide(request)
    assertError(t, "GetSlidesSlide", "storage", r.Code, e)
}

/* SlidesApiServiceTests Read presentation slide background color type.
   Test for SlidesApi.GetSlidesSlideBackground method
*/
func TestGetSlidesSlideBackground(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()
    e := initializeTest("GetSlidesSlideBackground", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.GetSlidesSlideBackground(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetSlidesSlideBackgroundRequest() GetSlidesSlideBackgroundRequest {
    var request GetSlidesSlideBackgroundRequest
    request.name = createTestParamValue("GetSlidesSlideBackground", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlidesSlideBackground", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("GetSlidesSlideBackground", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesSlideBackground", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesSlideBackground", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation slide background color type.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid name
*/
func TestGetSlidesSlideBackgroundInvalidname(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlidesSlideBackground", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    assertError(t, "GetSlidesSlideBackground", "name", r.Code, e)
}

/* SlidesApiServiceTests Read presentation slide background color type.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid slideIndex
*/
func TestGetSlidesSlideBackgroundInvalidslideIndex(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetSlidesSlideBackground", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    assertError(t, "GetSlidesSlideBackground", "slideIndex", r.Code, e)
}

/* SlidesApiServiceTests Read presentation slide background color type.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid password
*/
func TestGetSlidesSlideBackgroundInvalidpassword(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlidesSlideBackground", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    assertError(t, "GetSlidesSlideBackground", "password", r.Code, e)
}

/* SlidesApiServiceTests Read presentation slide background color type.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid folder
*/
func TestGetSlidesSlideBackgroundInvalidfolder(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlidesSlideBackground", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    assertError(t, "GetSlidesSlideBackground", "folder", r.Code, e)
}

/* SlidesApiServiceTests Read presentation slide background color type.
   Test for SlidesApi.GetSlidesSlideBackground method with invalid storage
*/
func TestGetSlidesSlideBackgroundInvalidstorage(t *testing.T) {
    request := createGetSlidesSlideBackgroundRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlidesSlideBackground", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlideBackground(request)
    assertError(t, "GetSlidesSlideBackground", "storage", r.Code, e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method
*/
func TestGetSlidesSlideComments(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()
    e := initializeTest("GetSlidesSlideComments", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.GetSlidesSlideComments(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetSlidesSlideCommentsRequest() GetSlidesSlideCommentsRequest {
    var request GetSlidesSlideCommentsRequest
    request.name = createTestParamValue("GetSlidesSlideComments", "name", "string").(string)
    request.slideIndex = createTestParamValue("GetSlidesSlideComments", "slideIndex", "int32").(int32)
    request.password = createTestParamValue("GetSlidesSlideComments", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesSlideComments", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesSlideComments", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid name
*/
func TestGetSlidesSlideCommentsInvalidname(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlidesSlideComments", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    assertError(t, "GetSlidesSlideComments", "name", r.Code, e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid slideIndex
*/
func TestGetSlidesSlideCommentsInvalidslideIndex(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("GetSlidesSlideComments", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    assertError(t, "GetSlidesSlideComments", "slideIndex", r.Code, e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid password
*/
func TestGetSlidesSlideCommentsInvalidpassword(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlidesSlideComments", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    assertError(t, "GetSlidesSlideComments", "password", r.Code, e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid folder
*/
func TestGetSlidesSlideCommentsInvalidfolder(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlidesSlideComments", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    assertError(t, "GetSlidesSlideComments", "folder", r.Code, e)
}

/* SlidesApiServiceTests Read presentation slide comments.
   Test for SlidesApi.GetSlidesSlideComments method with invalid storage
*/
func TestGetSlidesSlideCommentsInvalidstorage(t *testing.T) {
    request := createGetSlidesSlideCommentsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlidesSlideComments", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlideComments(request)
    assertError(t, "GetSlidesSlideComments", "storage", r.Code, e)
}

/* SlidesApiServiceTests Read presentation slides info.
   Test for SlidesApi.GetSlidesSlidesList method
*/
func TestGetSlidesSlidesList(t *testing.T) {
    request := createGetSlidesSlidesListRequest()
    e := initializeTest("GetSlidesSlidesList", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.GetSlidesSlidesList(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createGetSlidesSlidesListRequest() GetSlidesSlidesListRequest {
    var request GetSlidesSlidesListRequest
    request.name = createTestParamValue("GetSlidesSlidesList", "name", "string").(string)
    request.password = createTestParamValue("GetSlidesSlidesList", "password", "string").(string)
    request.folder = createTestParamValue("GetSlidesSlidesList", "folder", "string").(string)
    request.storage = createTestParamValue("GetSlidesSlidesList", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Read presentation slides info.
   Test for SlidesApi.GetSlidesSlidesList method with invalid name
*/
func TestGetSlidesSlidesListInvalidname(t *testing.T) {
    request := createGetSlidesSlidesListRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("GetSlidesSlidesList", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlidesList(request)
    assertError(t, "GetSlidesSlidesList", "name", r.Code, e)
}

/* SlidesApiServiceTests Read presentation slides info.
   Test for SlidesApi.GetSlidesSlidesList method with invalid password
*/
func TestGetSlidesSlidesListInvalidpassword(t *testing.T) {
    request := createGetSlidesSlidesListRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("GetSlidesSlidesList", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlidesList(request)
    assertError(t, "GetSlidesSlidesList", "password", r.Code, e)
}

/* SlidesApiServiceTests Read presentation slides info.
   Test for SlidesApi.GetSlidesSlidesList method with invalid folder
*/
func TestGetSlidesSlidesListInvalidfolder(t *testing.T) {
    request := createGetSlidesSlidesListRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("GetSlidesSlidesList", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlidesList(request)
    assertError(t, "GetSlidesSlidesList", "folder", r.Code, e)
}

/* SlidesApiServiceTests Read presentation slides info.
   Test for SlidesApi.GetSlidesSlidesList method with invalid storage
*/
func TestGetSlidesSlidesListInvalidstorage(t *testing.T) {
    request := createGetSlidesSlidesListRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("GetSlidesSlidesList", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.GetSlidesSlidesList(request)
    assertError(t, "GetSlidesSlidesList", "storage", r.Code, e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.PostSlideSaveAs method
*/
func TestPostSlideSaveAs(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    e := initializeTest("PostSlideSaveAs", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.PostSlideSaveAs(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    assertBinaryResponse(r, t)
}

func createPostSlideSaveAsRequest() PostSlideSaveAsRequest {
    var request PostSlideSaveAsRequest
    request.name = createTestParamValue("PostSlideSaveAs", "name", "string").(string)
    request.slideIndex = createTestParamValue("PostSlideSaveAs", "slideIndex", "int32").(int32)
    request.format = createTestParamValue("PostSlideSaveAs", "format", "string").(string)
    request.options = createTestParamValue("PostSlideSaveAs", "options", "interface{}").(interface{})
    request.width = createTestParamValue("PostSlideSaveAs", "width", "int32").(int32)
    request.height = createTestParamValue("PostSlideSaveAs", "height", "int32").(int32)
    request.password = createTestParamValue("PostSlideSaveAs", "password", "string").(string)
    request.folder = createTestParamValue("PostSlideSaveAs", "folder", "string").(string)
    request.storage = createTestParamValue("PostSlideSaveAs", "storage", "string").(string)
    request.outPath = createTestParamValue("PostSlideSaveAs", "outPath", "string").(string)
    request.fontsFolder = createTestParamValue("PostSlideSaveAs", "fontsFolder", "string").(string)
    return request
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.PostSlideSaveAs method with invalid name
*/
func TestPostSlideSaveAsInvalidname(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PostSlideSaveAs", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "name", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.PostSlideSaveAs method with invalid slideIndex
*/
func TestPostSlideSaveAsInvalidslideIndex(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PostSlideSaveAs", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "slideIndex", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.PostSlideSaveAs method with invalid format
*/
func TestPostSlideSaveAsInvalidformat(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.format = invalidizeTestParamValue(request.format, "format", "string").(string)
    e := initializeTest("PostSlideSaveAs", "format", request.format)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "format", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.PostSlideSaveAs method with invalid options
*/
func TestPostSlideSaveAsInvalidoptions(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.options = invalidizeTestParamValue(request.options, "options", "interface{}").(interface{})
    e := initializeTest("PostSlideSaveAs", "options", request.options)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "options", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.PostSlideSaveAs method with invalid width
*/
func TestPostSlideSaveAsInvalidwidth(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.width = invalidizeTestParamValue(request.width, "width", "int32").(int32)
    e := initializeTest("PostSlideSaveAs", "width", request.width)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "width", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.PostSlideSaveAs method with invalid height
*/
func TestPostSlideSaveAsInvalidheight(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.height = invalidizeTestParamValue(request.height, "height", "int32").(int32)
    e := initializeTest("PostSlideSaveAs", "height", request.height)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "height", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.PostSlideSaveAs method with invalid password
*/
func TestPostSlideSaveAsInvalidpassword(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PostSlideSaveAs", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "password", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.PostSlideSaveAs method with invalid folder
*/
func TestPostSlideSaveAsInvalidfolder(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PostSlideSaveAs", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "folder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.PostSlideSaveAs method with invalid storage
*/
func TestPostSlideSaveAsInvalidstorage(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PostSlideSaveAs", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "storage", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.PostSlideSaveAs method with invalid outPath
*/
func TestPostSlideSaveAsInvalidoutPath(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.outPath = invalidizeTestParamValue(request.outPath, "outPath", "string").(string)
    e := initializeTest("PostSlideSaveAs", "outPath", request.outPath)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "outPath", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Convert slide to some format.
   Test for SlidesApi.PostSlideSaveAs method with invalid fontsFolder
*/
func TestPostSlideSaveAsInvalidfontsFolder(t *testing.T) {
    request := createPostSlideSaveAsRequest()
    request.fontsFolder = invalidizeTestParamValue(request.fontsFolder, "fontsFolder", "string").(string)
    e := initializeTest("PostSlideSaveAs", "fontsFolder", request.fontsFolder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    _, r, e := getTestApiClient().SlidesApi.PostSlideSaveAs(request)
    assertError(t, "PostSlideSaveAs", "fontsFolder", int32(r.StatusCode), e)
}

/* SlidesApiServiceTests Reorder presentation slide position
   Test for SlidesApi.PostSlidesReorderPosition method
*/
func TestPostSlidesReorderPosition(t *testing.T) {
    request := createPostSlidesReorderPositionRequest()
    e := initializeTest("PostSlidesReorderPosition", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.PostSlidesReorderPosition(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPostSlidesReorderPositionRequest() PostSlidesReorderPositionRequest {
    var request PostSlidesReorderPositionRequest
    request.name = createTestParamValue("PostSlidesReorderPosition", "name", "string").(string)
    request.oldPosition = createTestParamValue("PostSlidesReorderPosition", "oldPosition", "int32").(int32)
    request.newPosition = createTestParamValue("PostSlidesReorderPosition", "newPosition", "int32").(int32)
    request.slideToCopy = createTestParamValue("PostSlidesReorderPosition", "slideToCopy", "int32").(int32)
    request.position = createTestParamValue("PostSlidesReorderPosition", "position", "int32").(int32)
    request.slideToClone = createTestParamValue("PostSlidesReorderPosition", "slideToClone", "int32").(int32)
    request.source = createTestParamValue("PostSlidesReorderPosition", "source", "string").(string)
    request.password = createTestParamValue("PostSlidesReorderPosition", "password", "string").(string)
    request.folder = createTestParamValue("PostSlidesReorderPosition", "folder", "string").(string)
    request.storage = createTestParamValue("PostSlidesReorderPosition", "storage", "string").(string)
    request.layoutAlias = createTestParamValue("PostSlidesReorderPosition", "layoutAlias", "string").(string)
    return request
}

/* SlidesApiServiceTests Reorder presentation slide position
   Test for SlidesApi.PostSlidesReorderPosition method with invalid name
*/
func TestPostSlidesReorderPositionInvalidname(t *testing.T) {
    request := createPostSlidesReorderPositionRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PostSlidesReorderPosition", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PostSlidesReorderPosition(request)
    assertError(t, "PostSlidesReorderPosition", "name", r.Code, e)
}

/* SlidesApiServiceTests Reorder presentation slide position
   Test for SlidesApi.PostSlidesReorderPosition method with invalid oldPosition
*/
func TestPostSlidesReorderPositionInvalidoldPosition(t *testing.T) {
    request := createPostSlidesReorderPositionRequest()
    request.oldPosition = invalidizeTestParamValue(request.oldPosition, "oldPosition", "int32").(int32)
    e := initializeTest("PostSlidesReorderPosition", "oldPosition", request.oldPosition)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PostSlidesReorderPosition(request)
    assertError(t, "PostSlidesReorderPosition", "oldPosition", r.Code, e)
}

/* SlidesApiServiceTests Reorder presentation slide position
   Test for SlidesApi.PostSlidesReorderPosition method with invalid newPosition
*/
func TestPostSlidesReorderPositionInvalidnewPosition(t *testing.T) {
    request := createPostSlidesReorderPositionRequest()
    request.newPosition = invalidizeTestParamValue(request.newPosition, "newPosition", "int32").(int32)
    e := initializeTest("PostSlidesReorderPosition", "newPosition", request.newPosition)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PostSlidesReorderPosition(request)
    assertError(t, "PostSlidesReorderPosition", "newPosition", r.Code, e)
}

/* SlidesApiServiceTests Reorder presentation slide position
   Test for SlidesApi.PostSlidesReorderPosition method with invalid slideToCopy
*/
func TestPostSlidesReorderPositionInvalidslideToCopy(t *testing.T) {
    request := createPostSlidesReorderPositionRequest()
    request.slideToCopy = invalidizeTestParamValue(request.slideToCopy, "slideToCopy", "int32").(int32)
    e := initializeTest("PostSlidesReorderPosition", "slideToCopy", request.slideToCopy)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PostSlidesReorderPosition(request)
    assertError(t, "PostSlidesReorderPosition", "slideToCopy", r.Code, e)
}

/* SlidesApiServiceTests Reorder presentation slide position
   Test for SlidesApi.PostSlidesReorderPosition method with invalid position
*/
func TestPostSlidesReorderPositionInvalidposition(t *testing.T) {
    request := createPostSlidesReorderPositionRequest()
    request.position = invalidizeTestParamValue(request.position, "position", "int32").(int32)
    e := initializeTest("PostSlidesReorderPosition", "position", request.position)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PostSlidesReorderPosition(request)
    assertError(t, "PostSlidesReorderPosition", "position", r.Code, e)
}

/* SlidesApiServiceTests Reorder presentation slide position
   Test for SlidesApi.PostSlidesReorderPosition method with invalid slideToClone
*/
func TestPostSlidesReorderPositionInvalidslideToClone(t *testing.T) {
    request := createPostSlidesReorderPositionRequest()
    request.slideToClone = invalidizeTestParamValue(request.slideToClone, "slideToClone", "int32").(int32)
    e := initializeTest("PostSlidesReorderPosition", "slideToClone", request.slideToClone)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PostSlidesReorderPosition(request)
    assertError(t, "PostSlidesReorderPosition", "slideToClone", r.Code, e)
}

/* SlidesApiServiceTests Reorder presentation slide position
   Test for SlidesApi.PostSlidesReorderPosition method with invalid source
*/
func TestPostSlidesReorderPositionInvalidsource(t *testing.T) {
    request := createPostSlidesReorderPositionRequest()
    request.source = invalidizeTestParamValue(request.source, "source", "string").(string)
    e := initializeTest("PostSlidesReorderPosition", "source", request.source)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PostSlidesReorderPosition(request)
    assertError(t, "PostSlidesReorderPosition", "source", r.Code, e)
}

/* SlidesApiServiceTests Reorder presentation slide position
   Test for SlidesApi.PostSlidesReorderPosition method with invalid password
*/
func TestPostSlidesReorderPositionInvalidpassword(t *testing.T) {
    request := createPostSlidesReorderPositionRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PostSlidesReorderPosition", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PostSlidesReorderPosition(request)
    assertError(t, "PostSlidesReorderPosition", "password", r.Code, e)
}

/* SlidesApiServiceTests Reorder presentation slide position
   Test for SlidesApi.PostSlidesReorderPosition method with invalid folder
*/
func TestPostSlidesReorderPositionInvalidfolder(t *testing.T) {
    request := createPostSlidesReorderPositionRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PostSlidesReorderPosition", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PostSlidesReorderPosition(request)
    assertError(t, "PostSlidesReorderPosition", "folder", r.Code, e)
}

/* SlidesApiServiceTests Reorder presentation slide position
   Test for SlidesApi.PostSlidesReorderPosition method with invalid storage
*/
func TestPostSlidesReorderPositionInvalidstorage(t *testing.T) {
    request := createPostSlidesReorderPositionRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PostSlidesReorderPosition", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PostSlidesReorderPosition(request)
    assertError(t, "PostSlidesReorderPosition", "storage", r.Code, e)
}

/* SlidesApiServiceTests Reorder presentation slide position
   Test for SlidesApi.PostSlidesReorderPosition method with invalid layoutAlias
*/
func TestPostSlidesReorderPositionInvalidlayoutAlias(t *testing.T) {
    request := createPostSlidesReorderPositionRequest()
    request.layoutAlias = invalidizeTestParamValue(request.layoutAlias, "layoutAlias", "string").(string)
    e := initializeTest("PostSlidesReorderPosition", "layoutAlias", request.layoutAlias)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PostSlidesReorderPosition(request)
    assertError(t, "PostSlidesReorderPosition", "layoutAlias", r.Code, e)
}

/* SlidesApiServiceTests Update slide properties.
   Test for SlidesApi.PutSlidesSlide method
*/
func TestPutSlidesSlide(t *testing.T) {
    request := createPutSlidesSlideRequest()
    e := initializeTest("PutSlidesSlide", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.PutSlidesSlide(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPutSlidesSlideRequest() PutSlidesSlideRequest {
    var request PutSlidesSlideRequest
    request.name = createTestParamValue("PutSlidesSlide", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutSlidesSlide", "slideIndex", "int32").(int32)
    request.slideDto = createTestParamValue("PutSlidesSlide", "slideDto", "Slide").(Slide)
    request.password = createTestParamValue("PutSlidesSlide", "password", "string").(string)
    request.folder = createTestParamValue("PutSlidesSlide", "folder", "string").(string)
    request.storage = createTestParamValue("PutSlidesSlide", "storage", "string").(string)
    return request
}

/* SlidesApiServiceTests Update slide properties.
   Test for SlidesApi.PutSlidesSlide method with invalid name
*/
func TestPutSlidesSlideInvalidname(t *testing.T) {
    request := createPutSlidesSlideRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PutSlidesSlide", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    assertError(t, "PutSlidesSlide", "name", r.Code, e)
}

/* SlidesApiServiceTests Update slide properties.
   Test for SlidesApi.PutSlidesSlide method with invalid slideIndex
*/
func TestPutSlidesSlideInvalidslideIndex(t *testing.T) {
    request := createPutSlidesSlideRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PutSlidesSlide", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    assertError(t, "PutSlidesSlide", "slideIndex", r.Code, e)
}

/* SlidesApiServiceTests Update slide properties.
   Test for SlidesApi.PutSlidesSlide method with invalid slideDto
*/
func TestPutSlidesSlideInvalidslideDto(t *testing.T) {
    request := createPutSlidesSlideRequest()
    request.slideDto = invalidizeTestParamValue(request.slideDto, "slideDto", "Slide").(Slide)
    e := initializeTest("PutSlidesSlide", "slideDto", request.slideDto)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    assertError(t, "PutSlidesSlide", "slideDto", r.Code, e)
}

/* SlidesApiServiceTests Update slide properties.
   Test for SlidesApi.PutSlidesSlide method with invalid password
*/
func TestPutSlidesSlideInvalidpassword(t *testing.T) {
    request := createPutSlidesSlideRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PutSlidesSlide", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    assertError(t, "PutSlidesSlide", "password", r.Code, e)
}

/* SlidesApiServiceTests Update slide properties.
   Test for SlidesApi.PutSlidesSlide method with invalid folder
*/
func TestPutSlidesSlideInvalidfolder(t *testing.T) {
    request := createPutSlidesSlideRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PutSlidesSlide", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    assertError(t, "PutSlidesSlide", "folder", r.Code, e)
}

/* SlidesApiServiceTests Update slide properties.
   Test for SlidesApi.PutSlidesSlide method with invalid storage
*/
func TestPutSlidesSlideInvalidstorage(t *testing.T) {
    request := createPutSlidesSlideRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PutSlidesSlide", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PutSlidesSlide(request)
    assertError(t, "PutSlidesSlide", "storage", r.Code, e)
}

/* SlidesApiServiceTests Set presentation slide background color.
   Test for SlidesApi.PutSlidesSlideBackground method
*/
func TestPutSlidesSlideBackground(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()
    e := initializeTest("PutSlidesSlideBackground", "", "")
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    c := getTestApiClient()
    r, _, e := c.SlidesApi.PutSlidesSlideBackground(request)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    if r.Code != 200 && r.Code != 201 {
       t.Errorf("Wrong response code: %d.", r.Code)
       return
    }
}

func createPutSlidesSlideBackgroundRequest() PutSlidesSlideBackgroundRequest {
    var request PutSlidesSlideBackgroundRequest
    request.name = createTestParamValue("PutSlidesSlideBackground", "name", "string").(string)
    request.slideIndex = createTestParamValue("PutSlidesSlideBackground", "slideIndex", "int32").(int32)
    request.background = createTestParamValue("PutSlidesSlideBackground", "background", "SlideBackground").(SlideBackground)
    request.folder = createTestParamValue("PutSlidesSlideBackground", "folder", "string").(string)
    request.password = createTestParamValue("PutSlidesSlideBackground", "password", "string").(string)
    request.storage = createTestParamValue("PutSlidesSlideBackground", "storage", "string").(string)
    request.color = createTestParamValue("PutSlidesSlideBackground", "color", "string").(string)
    return request
}

/* SlidesApiServiceTests Set presentation slide background color.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid name
*/
func TestPutSlidesSlideBackgroundInvalidname(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()
    request.name = invalidizeTestParamValue(request.name, "name", "string").(string)
    e := initializeTest("PutSlidesSlideBackground", "name", request.name)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    assertError(t, "PutSlidesSlideBackground", "name", r.Code, e)
}

/* SlidesApiServiceTests Set presentation slide background color.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid slideIndex
*/
func TestPutSlidesSlideBackgroundInvalidslideIndex(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()
    request.slideIndex = invalidizeTestParamValue(request.slideIndex, "slideIndex", "int32").(int32)
    e := initializeTest("PutSlidesSlideBackground", "slideIndex", request.slideIndex)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    assertError(t, "PutSlidesSlideBackground", "slideIndex", r.Code, e)
}

/* SlidesApiServiceTests Set presentation slide background color.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid background
*/
func TestPutSlidesSlideBackgroundInvalidbackground(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()
    request.background = invalidizeTestParamValue(request.background, "background", "SlideBackground").(SlideBackground)
    e := initializeTest("PutSlidesSlideBackground", "background", request.background)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    assertError(t, "PutSlidesSlideBackground", "background", r.Code, e)
}

/* SlidesApiServiceTests Set presentation slide background color.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid folder
*/
func TestPutSlidesSlideBackgroundInvalidfolder(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()
    request.folder = invalidizeTestParamValue(request.folder, "folder", "string").(string)
    e := initializeTest("PutSlidesSlideBackground", "folder", request.folder)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    assertError(t, "PutSlidesSlideBackground", "folder", r.Code, e)
}

/* SlidesApiServiceTests Set presentation slide background color.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid password
*/
func TestPutSlidesSlideBackgroundInvalidpassword(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()
    request.password = invalidizeTestParamValue(request.password, "password", "string").(string)
    e := initializeTest("PutSlidesSlideBackground", "password", request.password)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    assertError(t, "PutSlidesSlideBackground", "password", r.Code, e)
}

/* SlidesApiServiceTests Set presentation slide background color.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid storage
*/
func TestPutSlidesSlideBackgroundInvalidstorage(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()
    request.storage = invalidizeTestParamValue(request.storage, "storage", "string").(string)
    e := initializeTest("PutSlidesSlideBackground", "storage", request.storage)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    assertError(t, "PutSlidesSlideBackground", "storage", r.Code, e)
}

/* SlidesApiServiceTests Set presentation slide background color.
   Test for SlidesApi.PutSlidesSlideBackground method with invalid color
*/
func TestPutSlidesSlideBackgroundInvalidcolor(t *testing.T) {
    request := createPutSlidesSlideBackgroundRequest()
    request.color = invalidizeTestParamValue(request.color, "color", "string").(string)
    e := initializeTest("PutSlidesSlideBackground", "color", request.color)
    if e != nil {
       t.Errorf("Error: %v.", e)
       return
    }
    r, _, e := getTestApiClient().SlidesApi.PutSlidesSlideBackground(request)
    assertError(t, "PutSlidesSlideBackground", "color", r.Code, e)
}
