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

// formatter package provides Struct type and method of EdgeAttribute
package formatter

import (
	"encoding/json"
	"errors"
)

// EdgeAttribute has Value and information of Value
type EdgeAttribute struct {
	Name     string      `json:"name"`
	DataType string      `json:"dataType"`
	Value    interface{} `json:"value"`
}

// create Error Messgae : "Invalid EdgeAttribute - [name of EdgeAttribute]"
//
// @return created Error
func (attr *EdgeAttribute) errInvalidEdgeAttribute() error {
	return errors.New("Invalid EdgeAttribute - " + attr.Name)
}

// create Error Messgae : "[name of EdgeAttribute] - [name of EdgeAttribute] do not match"
//
// @param attrB pointer of EdgeAttribute
//
// @return created Error
func (attrA *EdgeAttribute) errNotMatchEdgeAttribute(attrB *EdgeAttribute) error {
	return errors.New(attrA.Name + " - " + attrB.Name + " do not match")
}

// convert type of value to string
//
// @return true if type of value is string, otherwise false
func (attr *EdgeAttribute) convertToString() error {
	switch attr.Value.(type) {
	case string:
		return nil
	}
	return attr.errInvalidEdgeAttribute()
}

// convert type of value to int
//
// @return nil if it is success to convert type of value, otherwise error
func (attr *EdgeAttribute) convertToInt() error {
	switch attr.Value.(type) {
	case int:
		return nil
	case float64:
		floatValue := attr.Value.(float64)
		if floatValue == float64(int(floatValue)) {
			attr.Value = int(floatValue)
			return nil
		}
	}
	return attr.errInvalidEdgeAttribute()
}

// convert type of value to float64
//
// @return nil if it is success to convert type of value, otherwise error
func (attr *EdgeAttribute) convertToFloat64() error {
	switch attr.Value.(type) {
	case float64:
		return nil
	}
	return attr.errInvalidEdgeAttribute()
}

// convert type of value to list of EdgeAttribute
//
// @return nil if it is success to convert type of value, otherwise error
func (attr *EdgeAttribute) convertToEdgeAttributeList() error {
	switch attr.Value.(type) {
	case []EdgeAttribute:
		for _, edgeAttr := range attr.Value.([]EdgeAttribute) {
			err := edgeAttr.ConvertToEdgeAttribute()
			if err != nil {
				return err
			}
		}
		return nil
	case []interface{}:
		jsonBytes, err := json.Marshal(attr.Value)
		if err != nil {
			return attr.errInvalidEdgeAttribute()
		}
		var edgeAttributeList []EdgeAttribute
		err = json.Unmarshal(jsonBytes, &edgeAttributeList)
		if err != nil {
			return attr.errInvalidEdgeAttribute()
		}
		for _, edgeAttr := range edgeAttributeList {
			err = edgeAttr.ConvertToEdgeAttribute()
			if err != nil {
				return err
			}
		}
		attr.Value = edgeAttributeList
		return nil
	}
	return attr.errInvalidEdgeAttribute()
}

// convert all the wrong type of value in EdgeAttribute
//
// @return nil if it is success to convert type of value, otherwise error
func (attr *EdgeAttribute) ConvertToEdgeAttribute() error {
	if attr.Name == "" {
		return attr.errInvalidEdgeAttribute()
	}
	var err error = nil
	switch attr.DataType {
	case "string":
		err = attr.convertToString()
	case "float64":
		err = attr.convertToFloat64()
	case "integer":
		err = attr.convertToInt()
	case "attributes":
		err = attr.convertToEdgeAttributeList()
	default:
		err = attr.errInvalidEdgeAttribute()
	}
	return err
}

// Compare two EdgeAttribute
//
// @param edgeAttr pointer of EdgeAttribute
//
// @return nil if two EdgeAttribute are the equal, otherwise error
func (attr *EdgeAttribute) CompareEdgeAttribute(edgeAttr *EdgeAttribute) error {
	if edgeAttr == nil {
		return attr.errInvalidEdgeAttribute()
	}
	if attr.Name != edgeAttr.Name || attr.DataType != edgeAttr.DataType {
		return attr.errNotMatchEdgeAttribute(edgeAttr)
	}
	if attr.DataType == "attributes" {
		if len(attr.Value.([]EdgeAttribute)) != len(edgeAttr.Value.([]EdgeAttribute)) {
			return attr.errNotMatchEdgeAttribute(edgeAttr)
		}
		for _, attrList_attr := range attr.Value.([]EdgeAttribute) {
			exist := false
			for _, edgeAttrList_attr := range edgeAttr.Value.([]EdgeAttribute) {
				if attrList_attr.CompareEdgeAttribute(&edgeAttrList_attr) != nil {
					exist = true
					break
				}
			}
			if exist == false {
				return attr.errNotMatchEdgeAttribute(edgeAttr)
			}
		}
	} else if attr.Value != edgeAttr.Value {
		return attr.errNotMatchEdgeAttribute(edgeAttr)
	}
	return nil
}
