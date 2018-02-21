/*******************************************************************************
 * Copyright 2018 Samsung Electronics All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *******************************************************************************/

package formatter

import (
	"testing"
)

func TestConvertToEdgeElement(t *testing.T) {
	testEdgeElement := testElemDepth
	err := testEdgeElement.ConvertToEdgeElement()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestConvertToEdgeElement_ElementTitle_empty(t *testing.T) {
	testEdgeElement := EdgeElement{"", testElemDepth.EdgeAttributes}
	err := testEdgeElement.ConvertToEdgeElement()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestConvertToEdgeElement_EdgeAttribute_empty(t *testing.T) {
	var edgeAttributeList []EdgeAttribute
	testEdgeElement := EdgeElement{testElemDepth.ElementTitle, edgeAttributeList}
	err := testEdgeElement.ConvertToEdgeElement()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestConvertToEdgeElement_EdgeAttributes_nil(t *testing.T) {
	testEdgeElement := EdgeElement{testElemDepth.ElementTitle, nil}
	err := testEdgeElement.ConvertToEdgeElement()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestCompareEdgeElement(t *testing.T) {
	edgeElement := testElemDepth
	testEdgeElement := testElemDepth
	err := edgeElement.CompareEdgeElement(&testEdgeElement)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCompareEdgeElement_nil(t *testing.T) {
	testEdgeElement := testElemDepth
	err := testEdgeElement.CompareEdgeElement(nil)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestCompareEdgeElement_Different_ElementTitle(t *testing.T) {
	edgeElement := testElemDepth
	testEdgeElement := EdgeElement{"diff_elementTitle", testElemDepth.EdgeAttributes}
	err := edgeElement.CompareEdgeElement(&testEdgeElement)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestCompareEdgeElement_Different_EdgeAttribute(t *testing.T) {
	edgeElement := testElemDepth
	testEdgeElement := EdgeElement{testElemDepth.ElementTitle, attrList}
	err := edgeElement.CompareEdgeElement(&testEdgeElement)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestConvertToEdgeElement_EdgeAttributes_Has_InvaliedEdgeAttribute(t *testing.T) {
	edgeAttributeList := []EdgeAttribute{
		{testAttrString.Name, testAttrString.DataType, testAttrFloat.Value}}
	testEdgeElement := EdgeElement{testElemDepth.ElementTitle, edgeAttributeList}
	err := testEdgeElement.ConvertToEdgeElement()
	if err == nil {
		t.Error(errNegativeCase)
	}
}
