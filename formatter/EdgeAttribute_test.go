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

// create error test
func TestErrInvalidEdgeAttribute(t *testing.T) {
	testEdgeAttribute := testAttrString
	err := testEdgeAttribute.errInvalidEdgeAttribute()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

// convert to string test
func TestConvertToString(t *testing.T) {
	testEdgeAttribute := testAttrString
	err := testEdgeAttribute.convertToString()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestConvertToString_InvaliedValue_Integer(t *testing.T) {
	testEdgeAttribute := EdgeAttribute{testAttrString.Name, testAttrString.DataType, testAttrInteger.Value}
	err := testEdgeAttribute.convertToString()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestConvertToString_InvaliedValue_Float64(t *testing.T) {
	testEdgeAttribute := EdgeAttribute{testAttrString.Name, testAttrString.DataType, testAttrFloat.Value}
	err := testEdgeAttribute.convertToString()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestConvertToString_InvaliedValue_EdgeAttributeList(t *testing.T) {
	testEdgeAttribute := EdgeAttribute{testAttrString.Name, testAttrString.DataType, testAttrList.Value}
	err := testEdgeAttribute.convertToString()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

// convert to int test
func TestConvertToInt(t *testing.T) {
	testEdgeAttribute := testAttrInteger
	err := testEdgeAttribute.convertToInt()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestConvertToInt_InvaliedValue_String(t *testing.T) {
	testEdgeAttribute := EdgeAttribute{testAttrInteger.Name, testAttrInteger.DataType, testAttrString.Value}
	err := testEdgeAttribute.convertToInt()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestConvertToInt_InvaliedValue_Float64(t *testing.T) {
	testEdgeAttribute := EdgeAttribute{testAttrInteger.Name, testAttrInteger.DataType, testAttrFloat.Value}
	err := testEdgeAttribute.convertToInt()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestConvertToInt_InvaliedValue_EdgeAttributeList(t *testing.T) {
	testEdgeAttribute := EdgeAttribute{testAttrInteger.Name, testAttrInteger.DataType, testAttrList.Value}
	err := testEdgeAttribute.convertToInt()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

// convert to float64 test
func TestConvertToFloat64(t *testing.T) {
	testEdgeAttribute := testAttrFloat
	err := testEdgeAttribute.convertToFloat64()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestConvertToFloat64_InvaliedValue_String(t *testing.T) {
	testEdgeAttribute := EdgeAttribute{testAttrFloat.Name, testAttrFloat.DataType, testAttrString.Value}
	err := testEdgeAttribute.convertToFloat64()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestConvertToFloat64_InvaliedValue_Integer(t *testing.T) {
	testEdgeAttribute := EdgeAttribute{testAttrFloat.Name, testAttrFloat.DataType, testAttrInteger.Value}
	err := testEdgeAttribute.convertToFloat64()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestConvertToFloat64_InvaliedValue_EdgeAttributeList(t *testing.T) {
	testEdgeAttribute := EdgeAttribute{testAttrFloat.Name, testAttrFloat.DataType, testAttrList.Value}
	err := testEdgeAttribute.convertToFloat64()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

// convert to EdgeAttributeList test
func TestConvertToEdgeAttributeList(t *testing.T) {
	testEdgeAttribute := testAttrList
	err := testEdgeAttribute.convertToEdgeAttributeList()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestConvertToEdgeAttributeList_InterfaceList(t *testing.T) {
	testList := []interface{}{testAttrInteger, testAttrList}
	testEdgeAttribute := EdgeAttribute{testAttrList.Name, testAttrList.DataType, testList}
	err := testEdgeAttribute.ConvertToEdgeAttribute()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestConvertToEdgeAttributeList_InvaliedValue_String(t *testing.T) {
	testEdgeAttribute := EdgeAttribute{testAttrList.Name, testAttrList.DataType, testAttrString.Value}
	err := testEdgeAttribute.convertToEdgeAttributeList()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestConvertToEdgeAttributeList_InvaliedValue_Float64(t *testing.T) {
	testEdgeAttribute := EdgeAttribute{testAttrList.Name, testAttrList.DataType, testAttrFloat.Value}
	err := testEdgeAttribute.convertToEdgeAttributeList()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestConvertToEdgeAttributeList_InvaliedValue_Integer(t *testing.T) {
	testEdgeAttribute := EdgeAttribute{testAttrList.Name, testAttrList.DataType, testAttrInteger.Value}
	err := testEdgeAttribute.convertToEdgeAttributeList()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestConvertToEdgeAttributeList_Has_InvaliedEdgeAttribute(t *testing.T) {
	invalidedgeAttrList := []EdgeAttribute{
		testAttrString, {testAttrFloat.Name, testAttrFloat.DataType, testAttrString.Value}}
	testEdgeAttribute := EdgeAttribute{testAttrList.Name, testAttrList.DataType, invalidedgeAttrList}
	err := testEdgeAttribute.convertToEdgeAttributeList()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestConvertToEdgeAttribute_String(t *testing.T) {
	testEdgeAttribute := testAttrString
	err := testEdgeAttribute.ConvertToEdgeAttribute()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestConvertToEdgeAttribute_Integer(t *testing.T) {
	testEdgeAttribute := testAttrInteger
	err := testEdgeAttribute.ConvertToEdgeAttribute()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestConvertToEdgeAttribute_Float(t *testing.T) {
	testEdgeAttribute := testAttrFloat
	err := testEdgeAttribute.ConvertToEdgeAttribute()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestConvertToEdgeAttribute_Attributes(t *testing.T) {
	testEdgeAttribute := testAttrList
	err := testEdgeAttribute.ConvertToEdgeAttribute()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestConvertToEdgeAttribute_InvalidDataType_NoDataType(t *testing.T) {
	testEdgeAttribute := EdgeAttribute{testAttrString.Name, "NoType", testAttrString.Value}
	err := testEdgeAttribute.ConvertToEdgeAttribute()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestConvertToEdgeAttribute_InvalidName_Empty(t *testing.T) {
	testEdgeAttribute := EdgeAttribute{"", testAttrString.DataType, testAttrString.Value}
	err := testEdgeAttribute.ConvertToEdgeAttribute()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestCompareEdgeAttribute(t *testing.T) {
	edgeAttribute := testAttrList
	testEdgeAttribute := testAttrList
	err := edgeAttribute.CompareEdgeAttribute(&testEdgeAttribute)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCompareEdgeAttribute_nil(t *testing.T) {
	edgeAttribute := testAttrString
	err := edgeAttribute.CompareEdgeAttribute(nil)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestCompareEdgeAttribute_Different_Name(t *testing.T) {
	edgeAttribute := testAttrString
	testEdgeAttribute := EdgeAttribute{"diff_name", testAttrString.DataType, testAttrString.Value}
	err := edgeAttribute.CompareEdgeAttribute(&testEdgeAttribute)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestCompareEdgeAttribute_Different_DataType(t *testing.T) {
	edgeAttribute := testAttrInteger
	testEdgeAttribute := EdgeAttribute{testAttrInteger.Name, testAttrFloat.DataType, 123.0}
	err := edgeAttribute.CompareEdgeAttribute(&testEdgeAttribute)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestCompareEdgeAttribute_Different_Value(t *testing.T) {
	edgeAttribute := testAttrString
	testEdgeAttribute := EdgeAttribute{testAttrString.Name, testAttrString.DataType, "diff_value"}
	err := edgeAttribute.CompareEdgeAttribute(&testEdgeAttribute)
	if err == nil {
		t.Error(errNegativeCase)
	}
}
