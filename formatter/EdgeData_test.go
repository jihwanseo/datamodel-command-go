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

func TestConvertToEdgeData(t *testing.T) {
	testEdgeData := testData
	err := testEdgeData.ConvertToEdgeData()
	if err != nil {
		t.Error(err.Error())
	}
}

func TestConvertToEdgeData_Version_empty(t *testing.T) {
	testEdgeData := EdgeData{"", testData.DataTitle, testData.EdgeElements}
	err := testEdgeData.ConvertToEdgeData()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestConvertToEdgeData_DataTitle_empty(t *testing.T) {
	testEdgeData := EdgeData{testData.Version, "", testData.EdgeElements}
	err := testEdgeData.ConvertToEdgeData()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestConvertToEdgeData_EdgeElements_empty(t *testing.T) {
	var edgeElementList []EdgeElement
	testEdgeData := EdgeData{testData.Version, testData.DataTitle, edgeElementList}
	err := testEdgeData.ConvertToEdgeData()
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestCompareEdgeData(t *testing.T) {
	edgeData := testData
	testEdgeData := testData
	err := edgeData.CompareEdgeData(&testEdgeData)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestCompareEdgeData_nil(t *testing.T) {
	edgeData := testData
	err := edgeData.CompareEdgeData(nil)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestCompareEdgeData_Different_Version(t *testing.T) {
	edgeData := testData
	testEdgeData := EdgeData{"diff_version", testData.DataTitle, testData.EdgeElements}
	err := edgeData.CompareEdgeData(&testEdgeData)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestCompareEdgeData_Different_DataTitle(t *testing.T) {
	edgeData := testData
	testEdgeData := EdgeData{testData.Version, "diff_dataTitle", testData.EdgeElements}
	err := edgeData.CompareEdgeData(&testEdgeData)
	if err == nil {
		t.Error(errNegativeCase)
	}
}

func TestCompareEdgeData_Different_EdgeElement(t *testing.T) {
	edgeData := testData
	testEdgeElementList := []EdgeElement{{testElemDepth.ElementTitle, testElemDepth.EdgeAttributes}}
	testEdgeData := EdgeData{testData.Version, testData.Version, testEdgeElementList}
	err := edgeData.CompareEdgeData(&testEdgeData)
	if err == nil {
		t.Error(errNegativeCase)
	}
}
