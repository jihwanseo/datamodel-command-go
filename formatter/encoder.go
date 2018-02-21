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

// formatter package provides encode function for EdgeData, EdgeElement, EdgeAttribute
package formatter

import (
	"encoding/json"
	"errors"
)

var errNil error = errors.New("nil pointer error")

// Encode EdgeAttribute To JsonString
//
// @param edgeAttribute EdgeAttribute to be encoded
// @param jsonString    result of encoding
//
// @return JsonString and nil if it is succeed to encode EdgeAttribute to JsonString, otherwise error
func EncodeEdgeAttributeToJsonString(edgeAttribute *EdgeAttribute) (string, error) {
	jsonString := ""
	if edgeAttribute == nil {
		return jsonString, errNil
	}
	err := edgeAttribute.ConvertToEdgeAttribute()
	if err == nil {
		jsonBytes, err := json.Marshal(edgeAttribute)
		if err == nil {
			jsonString = string(jsonBytes)
		}
	}
	return jsonString, err
}

// Encode EdgeElement To JsonString
//
// @param edgeElement EdgeElement to be encoded
// @param jsonString  result of encoding
//
// @return JsonString and nil if it is succeed to encode EdgeElement to JsonString, otherwise error
func EncodeEdgeElementToJsonString(edgeElement *EdgeElement) (string, error) {
	jsonString := ""
	if edgeElement == nil {
		return jsonString, errNil
	}
	err := edgeElement.ConvertToEdgeElement()
	if err == nil {
		jsonBytes, err := json.Marshal(edgeElement)
		if err == nil {
			jsonString = string(jsonBytes)
		}
	}
	return jsonString, err
}

// Encode EdgeData To JsonString
//
// @param edgeData   EdgeData to be encoded
// @param jsonString result of encoding
//
// @return JsonString and nil if it is succeed to encode EdgeData to JsonString, otherwise error
func EncodeEdgeDataToJsonString(edgeData *EdgeData) (string, error) {
	jsonString := ""
	if edgeData == nil {
		return jsonString, errNil
	}
	err := edgeData.ConvertToEdgeData()
	if err == nil {
		jsonBytes, err := json.Marshal(edgeData)
		if err == nil {
			jsonString = string(jsonBytes)
		}
	}
	return jsonString, err
}
