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

// formatter package provides Struct type and method of EdgeElement
package formatter

import (
	"errors"
)

// EdgeElement has title of EdgeElement and list of EdgeAttribute
type EdgeElement struct {
	ElementTitle   string          `json:"elementTitle"`
	EdgeAttributes []EdgeAttribute `json:"edgeAttributes"`
}

// create Error Messgae : "Invalid EdgeElement - [name of EdgeElement]"
//
// @return created Error
func (elem *EdgeElement) errInvalidEdgeElement() error {
	return errors.New("Invalid EdgeElement - " + elem.ElementTitle)
}

// create Error Messgae : "[name of EdgeElement] - [name of EdgeElement] do not match"
//
// @param elemB pointer of EdgeElement
//
// @return created Error
func (elemA *EdgeElement) errNotMatchEdgeElement(elemB *EdgeElement) error {
	return errors.New(elemA.ElementTitle + " - " + elemB.ElementTitle + " do not match")
}

// convert all the wrong type of value in EdgeElement
//
// @return nil if it is success to convert type of value, otherwise error
func (elem *EdgeElement) ConvertToEdgeElement() error {
	if elem.ElementTitle == "" || elem.EdgeAttributes == nil {
		return elem.errInvalidEdgeElement()
	}
	edgeAttributeList := elem.EdgeAttributes
	for i := range edgeAttributeList {
		err := edgeAttributeList[i].ConvertToEdgeAttribute()
		if err != nil {
			return err
		}
	}
	return nil
}

// Compare two EdgeElement
//
// @param edgeElement pointer of EdgeElement
//
// @return nil if two EdgeElement are the equal, otherwise error
func (elem *EdgeElement) CompareEdgeElement(edgeElement *EdgeElement) error {
	if edgeElement == nil {
		return elem.errInvalidEdgeElement()
	}
	if edgeElement.ElementTitle != elem.ElementTitle || len(elem.EdgeAttributes) != len(edgeElement.EdgeAttributes) {
		return elem.errNotMatchEdgeElement(edgeElement)
	}
	for _, edgeAttribute := range edgeElement.EdgeAttributes {
		exist := false
		for _, attr := range elem.EdgeAttributes {
			if edgeAttribute.CompareEdgeAttribute(&attr) == nil {
				exist = true
				break
			}
		}
		if exist == false {
			return elem.errNotMatchEdgeElement(edgeElement)
		}
	}
	return nil
}
