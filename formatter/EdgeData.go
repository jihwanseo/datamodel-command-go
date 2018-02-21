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

// formatter package provides Struct type and method of EdgeData
package formatter

import (
	"errors"
)

// EdgeData has version and title of EdgeData and list of EdgeElement
type EdgeData struct {
	Version      string        `json:"version"`
	DataTitle    string        `json:"dataTitle"`
	EdgeElements []EdgeElement `json:"edgeElements"`
}

// create Error Messgae : "Invalid EdgeData - [DataTitle of EdgeData]"
//
// @return created Error
func (data *EdgeData) errInvalidEdgeData() error {
	return errors.New("Invalid EdgeData - " + data.DataTitle)
}

// create Error Messgae : "[DataTitle of EdgeData] - [DataTitle of EdgeData] do not match"
//
// @param dataB pointer of EdgeData
//
// @return created Error
func (dataA *EdgeData) errNotMatchEdgeData(dataB *EdgeData) error {
	return errors.New(dataA.DataTitle + " - " + dataB.DataTitle + " do not match")
}

// convert all the wrong type of value in EdgeData
//
// @return true if it is success to convert type of value, otherwise false
func (data *EdgeData) ConvertToEdgeData() error {
	if data.DataTitle == "" || data.Version == "" || data.EdgeElements == nil {
		return data.errInvalidEdgeData()
	}
	edgeElementList := data.EdgeElements
	for i := range edgeElementList {
		err := edgeElementList[i].ConvertToEdgeElement()
		if err != nil {
			return err
		}
	}
	return nil
}

// Compare two EdgeData
//
// @param edgeData pointer of EdgeData
//
// @return nil if two EdgeData are the equal, otherwise error
func (data *EdgeData) CompareEdgeData(edgeData *EdgeData) error {
	if edgeData == nil {
		return data.errInvalidEdgeData()
	}
	if edgeData.Version != data.Version || edgeData.DataTitle != data.DataTitle {
		return data.errNotMatchEdgeData(edgeData)
	}
	for _, edgeElement := range edgeData.EdgeElements {
		exist := false
		for _, elem := range data.EdgeElements {
			err := elem.CompareEdgeElement(&edgeElement)
			if err == nil {
				exist = true
			}
		}
		if exist == false {
			return data.errNotMatchEdgeData(edgeData)
		}
	}
	return nil
}
