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

func TestEncodeEdgeAttributeToJsonString(t *testing.T) {
	testEdgeAttribute := testAttrString
	expectedString := `{"name":"stringAttr","dataType":"string","value":"stringValue"}`
	jsonString, err := EncodeEdgeAttributeToJsonString(&testEdgeAttribute)
	if err != nil {
		t.Error(err.Error())
	} else if jsonString != expectedString {
		t.Error(errNotMatch)
	}
}

func TestEncodeEdgeAttributeToJsonString_InvalidEdgeAttribute(t *testing.T) {
	testEdgeAttribute := EdgeAttribute{testAttrString.Name, testAttrString.DataType, testAttrFloat.Value}
	_, err := EncodeEdgeAttributeToJsonString(&testEdgeAttribute)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestEncodeEdgeAttributeToJsonString_nil(t *testing.T) {
	_, err := EncodeEdgeAttributeToJsonString(nil)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestEncodeEdgeElementToJsonString(t *testing.T) {
	testEdgeElement := testElemDepth
	expectedString := `{"elementTitle":"elementTitle","edgeAttributes":` +
		`[{"name":"IntegerAttr","dataType":"integer","value":123},` +
		`{"name":"ListAttr","dataType":"attributes","value":` +
		`[{"name":"stringAttr","dataType":"string","value":"stringValue"},` +
		`{"name":"FloatAttr","dataType":"float64","value":1.23}]}]}`
	jsonString, err := EncodeEdgeElementToJsonString(&testEdgeElement)
	if err != nil {
		t.Error(err.Error())
	} else if jsonString != expectedString {
		t.Error(errNotMatch)
	}
}

func TestEncodeEdgeElementToJsonString_InvalidEdgeAttribute(t *testing.T) {
	invalidAttrList := []EdgeAttribute{
		{testAttrString.Name, testAttrString.DataType, testAttrFloat.Value}}
	testEdgeElement := EdgeElement{testElemDepth.ElementTitle, invalidAttrList}
	_, err := EncodeEdgeElementToJsonString(&testEdgeElement)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestEncodeEdgeElementToJsonString_InvalidEdgeElement_ElementTitle_Empty(t *testing.T) {
	testEdgeElement := EdgeElement{"", testElemDepth.EdgeAttributes}
	_, err := EncodeEdgeElementToJsonString(&testEdgeElement)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestEncodeEdgeElementToJsonString_InvalidEdgeElement_EdgeAttributes_nil(t *testing.T) {
	testEdgeElement := EdgeElement{testElemDepth.ElementTitle, nil}
	_, err := EncodeEdgeElementToJsonString(&testEdgeElement)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestEncodeEdgeElementToJsonString_InvalidEdgeElement_EdgeAttributes_Empty(t *testing.T) {
	var edgeAttributeList []EdgeAttribute
	testEdgeElement := EdgeElement{testElemDepth.ElementTitle, edgeAttributeList}
	_, err := EncodeEdgeElementToJsonString(&testEdgeElement)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestEncodeEdgeElementToJsonString_nil(t *testing.T) {
	_, err := EncodeEdgeElementToJsonString(nil)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestEncodeEdgeDataToJsonString(t *testing.T) {
	testEdgeData := testData
	expectedString := `{"version":"version","dataTitle":"dataTitle","edgeElements":` +
		`[{"elementTitle":"elementTitle","edgeAttributes":` +
		`[{"name":"IntegerAttr","dataType":"integer","value":123},` +
		`{"name":"ListAttr","dataType":"attributes","value":` +
		`[{"name":"stringAttr","dataType":"string","value":"stringValue"},` +
		`{"name":"FloatAttr","dataType":"float64","value":1.23}]}]},` +
		`{"elementTitle":"elementTitle","edgeAttributes":` +
		`[{"name":"stringAttr","dataType":"string","value":"stringValue"},` +
		`{"name":"FloatAttr","dataType":"float64","value":1.23}]}]}`
	jsonString, err := EncodeEdgeDataToJsonString(&testEdgeData)
	if err != nil {
		t.Error(err.Error())
	} else if jsonString != expectedString {
		t.Error(errNotMatch)
	}
}

func TestEncodeEdgeDataToJsonString_InvalidEdgeElement(t *testing.T) {
	edgeElementList := []EdgeElement{{testElemDepth.ElementTitle, nil}}
	testEdgeData := EdgeData{testData.Version, testData.DataTitle, edgeElementList}
	_, err := EncodeEdgeDataToJsonString(&testEdgeData)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestEncodeEdgeDataToJsonString_InvalidEdgeData_Version_Empty(t *testing.T) {
	testEdgeData := EdgeData{"", testData.DataTitle, testData.EdgeElements}
	_, err := EncodeEdgeDataToJsonString(&testEdgeData)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestEncodeEdgeDataToJsonString_InvalidEdgeData_DataTitle_Empty(t *testing.T) {
	testEdgeData := EdgeData{testData.Version, "", testData.EdgeElements}
	_, err := EncodeEdgeDataToJsonString(&testEdgeData)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestEncodeEdgeDataToJsonString_InvalidEdgeData_EdgeElements_nil(t *testing.T) {
	testEdgeData := EdgeData{testData.Version, testData.DataTitle, nil}
	_, err := EncodeEdgeDataToJsonString(&testEdgeData)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestEncodeEdgeDataToJsonString_InvalidEdgeData_EdgeElements_Empty(t *testing.T) {
	var edgeElementList []EdgeElement
	testEdgeData := EdgeData{testData.Version, testData.DataTitle, edgeElementList}
	_, err := EncodeEdgeDataToJsonString(&testEdgeData)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestEncodeEdgeDataToJsonString_nil(t *testing.T) {
	_, err := EncodeEdgeDataToJsonString(nil)
	if err == nil {
		t.Error(errNegativeCase)
	}
}
