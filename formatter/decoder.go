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

// formatter package provides decode function for EdgeData, EdgeElement, EdgeAttribute
package formatter

import (
	"encoding/json"
	"errors"
)

var errJsonStringNull error = errors.New("jsonString is nil")

// Decode JsonString To EdgeAttribute
//
// @param jsonString    Json String to be decoded
// @param edgeAttribute result of decoding
//
// @return EdgeAttribute and nil if it is succeed to decode JsonString to EdgeAttribute, otherwise nil
func DecodeJsonStringToEdgeAttribute(jsonString *string) (EdgeAttribute, error) {
	var edgeAttribute EdgeAttribute
	if jsonString == nil {
		return edgeAttribute, errJsonStringNull
	}
	err := json.Unmarshal([]byte(*jsonString), &edgeAttribute)
	if err == nil {
		err = edgeAttribute.ConvertToEdgeAttribute()
	}
	return edgeAttribute, err
}

// Decode JsonString To EdgeElement
//
// @param jsonString  Json String to be decoded
// @param edgeElement result of decoding
//
// @return EdgeElement and nil if it is succeed to decode JsonString to EdgeElement, otherwise error
func DecodeJsonStringToEdgeElement(jsonString *string) (EdgeElement, error) {
	var edgeElement EdgeElement
	if jsonString == nil {
		return edgeElement, errJsonStringNull
	}
	err := json.Unmarshal([]byte(*jsonString), &edgeElement)
	if err == nil {
		err = edgeElement.ConvertToEdgeElement()
	}
	return edgeElement, err
}

// Decode JsonString To EdgeData
//
// @param jsonString Json String to be decoded
// @param edgeData   result of decoding
//
// @return EdgeData and nil if it is succeed to decode JsonString to EdgeData, otherwise error
func DecodeJsonStringToEdgeData(jsonString *string) (EdgeData, error) {
	var edgeData EdgeData
	if jsonString == nil {
		return edgeData, errJsonStringNull
	}
	err := json.Unmarshal([]byte(*jsonString), &edgeData)
	if err == nil {
		err = edgeData.ConvertToEdgeData()
	}
	return edgeData, err
}
