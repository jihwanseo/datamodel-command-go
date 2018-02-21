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

func TestDecodeJsonStringToEdgeAttribute(t *testing.T) {
	jsonString := `{"name":"stringAttr","dataType":"string","value":"stringValue"}`
	expectedEdgeAttribute := testAttrString
	edgeAttribute, err := DecodeJsonStringToEdgeAttribute(&jsonString)
	if err != nil {
		t.Error(err.Error())
	}
	err = edgeAttribute.CompareEdgeAttribute(&expectedEdgeAttribute)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestDecodeJsonStringToEdgeAttribute_InvaildJsonString_NotMatchedDataType(t *testing.T) {
	jsonString := `{"name":"stringAttr","dataType":"float64","value":"stringValue"}`
	_, err := DecodeJsonStringToEdgeAttribute(&jsonString)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestDecodeJsonStringToEdgeAttribute_InvaildJsonString_NotMathcedKey(t *testing.T) {
	jsonString := `{"name":"stringAttr","no_key":"string","value":"stringValue"}`
	_, err := DecodeJsonStringToEdgeAttribute(&jsonString)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestDecodeJsonStringToEdgeAttribute_InvaildJsonString_Empty(t *testing.T) {
	jsonString := ""
	_, err := DecodeJsonStringToEdgeAttribute(&jsonString)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestDecodeJsonStringToEdgeAttribute_InvaildJsonString_nil(t *testing.T) {
	_, err := DecodeJsonStringToEdgeAttribute(nil)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestDecodeJsonStringToEdgeElement(t *testing.T) {
	jsonString := `{"elementTitle":"elementTitle","edgeAttributes":` +
		`[{"name":"IntegerAttr","dataType":"integer","value":123},` +
		`{"name":"ListAttr","dataType":"attributes","value":` +
		`[{"name":"stringAttr","dataType":"string","value":"stringValue"},` +
		`{"name":"FloatAttr","dataType":"float64","value":1.23}]}]}`
	expectedEdgeElement := testElemDepth
	edgeElement, err := DecodeJsonStringToEdgeElement(&jsonString)
	if err != nil {
		t.Error(err.Error())
	}
	err = edgeElement.CompareEdgeElement(&expectedEdgeElement)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestDecodeJsonStringToEdgeElement_InvaildJsonString_NotMatchedKey(t *testing.T) {
	jsonString := `{"elementTitle":"elementTitle","no_key":` +
		`[{"name":"IntegerAttr","dataType":"integer","value":123},` +
		`{"name":"ListAttr","dataType":"attributes","value":` +
		`[{"name":"stringAttr","dataType":"string","value":"stringValue"},` +
		`{"name":"FloatAttr","dataType":"float64","value":1.23}]}]}`
	_, err := DecodeJsonStringToEdgeElement(&jsonString)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestDecodeJsonStringToEdgeElement_InvaildJsonString_ElementTitle_Empty(t *testing.T) {
	jsonString := `{"elementTitle":"","edgeAttributes":` +
		`[{"name":"IntegerAttr","dataType":"integer","value":123},` +
		`{"name":"ListAttr","dataType":"attributes","value":` +
		`[{"name":"stringAttr","dataType":"string","value":"stringValue"},` +
		`{"name":"FloatAttr","dataType":"float64","value":1.23}]}]}`
	_, err := DecodeJsonStringToEdgeElement(&jsonString)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestDecodeJsonStringToEdgeElement_InvaildJsonString_Empty(t *testing.T) {
	jsonString := ""
	_, err := DecodeJsonStringToEdgeElement(&jsonString)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestDecodeJsonStringToEdgeElement_InvaildJsonString_nil(t *testing.T) {
	_, err := DecodeJsonStringToEdgeElement(nil)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestDecodeJsonStringToEdgeData(t *testing.T) {
	jsonString := `{"version":"version","dataTitle":"dataTitle","edgeElements":` +
		`[{"elementTitle":"elementTitle","edgeAttributes":` +
		`[{"name":"IntegerAttr","dataType":"integer","value":123},` +
		`{"name":"ListAttr","dataType":"attributes","value":` +
		`[{"name":"stringAttr","dataType":"string","value":"stringValue"},` +
		`{"name":"FloatAttr","dataType":"float64","value":1.23}]}]},` +
		`{"elementTitle":"elementTitle","edgeAttributes":` +
		`[{"name":"stringAttr","dataType":"string","value":"stringValue"},` +
		`{"name":"FloatAttr","dataType":"float64","value":1.23}]}]}`

	expectedEdgeData := testData
	edgeData, err := DecodeJsonStringToEdgeData(&jsonString)
	if err != nil {
		t.Error(err.Error())
	}
	err = edgeData.CompareEdgeData(&expectedEdgeData)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestDecodeJsonStringToEdgeData_InvalidJsonString_NotMatchedKey(t *testing.T) {
	jsonString := `{"no_key":"version","dataTitle":"dataTitle","edgeElements":` +
		`[{"elementTitle":"elementTitle","edgeAttributes":` +
		`[{"name":"IntegerAttr","dataType":"integer","value":123},` +
		`{"name":"ListAttr","dataType":"attributes","value":` +
		`[{"name":"stringAttr","dataType":"string","value":"stringValue"},` +
		`{"name":"FloatAttr","dataType":"float64","value":1.23}]}]},` +
		`{"elementTitle":"elementTitle","edgeAttributes":` +
		`[{"name":"stringAttr","dataType":"string","value":"stringValue"},` +
		`{"name":"FloatAttr","dataType":"float64","value":1.23}]}]}`
	_, err := DecodeJsonStringToEdgeData(&jsonString)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestDecodeJsonStringToEdgeData_InvalidJsonString_Version_Empty(t *testing.T) {
	jsonString := `{"version":"","dataTitle":"dataTitle","edgeElements":` +
		`[{"elementTitle":"elementTitle","edgeAttributes":` +
		`[{"name":"IntegerAttr","dataType":"integer","value":123},` +
		`{"name":"ListAttr","dataType":"attributes","value":` +
		`[{"name":"stringAttr","dataType":"string","value":"stringValue"},` +
		`{"name":"FloatAttr","dataType":"float64","value":1.23}]}]},` +
		`{"elementTitle":"elementTitle","edgeAttributes":` +
		`[{"name":"stringAttr","dataType":"string","value":"stringValue"},` +
		`{"name":"FloatAttr","dataType":"float64","value":1.23}]}]}`
	_, err := DecodeJsonStringToEdgeData(&jsonString)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestDecodeJsonStringToEdgeData_InvalidJsonString_DataTitle_Empty(t *testing.T) {
	jsonString := `{"version":"version","edgeElements":` +
		`[{"elementTitle":"elementTitle","edgeAttributes":` +
		`[{"name":"IntegerAttr","dataType":"integer","value":123},` +
		`{"name":"ListAttr","dataType":"attributes","value":` +
		`[{"name":"stringAttr","dataType":"string","value":"stringValue"},` +
		`{"name":"FloatAttr","dataType":"float64","value":1.23}]}]},` +
		`{"elementTitle":"elementTitle","edgeAttributes":` +
		`[{"name":"stringAttr","dataType":"string","value":"stringValue"},` +
		`{"name":"FloatAttr","dataType":"float64","value":1.23}]}]}`
	_, err := DecodeJsonStringToEdgeData(&jsonString)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestDecodeJsonStringToEdgeData_InvalidJsonString_Empty(t *testing.T) {
	jsonString := ""
	_, err := DecodeJsonStringToEdgeData(&jsonString)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestDecodeJsonStringToEdgeData_InvalidJsonString_nil(t *testing.T) {
	_, err := DecodeJsonStringToEdgeData(nil)
	if err == nil {
		t.Error(errNegativeCase)
	}
}
