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

var errNegativeCase string = "It was an error case, but it couldn't find error"
var errNotMatch string = "Result and expected result do not match"

var testAttrString EdgeAttribute = EdgeAttribute{"stringAttr", "string", "stringValue"}
var testAttrFloat EdgeAttribute = EdgeAttribute{"FloatAttr", "float64", 1.23}
var testAttrInteger EdgeAttribute = EdgeAttribute{"IntegerAttr", "integer", 123}
var attrList []EdgeAttribute = []EdgeAttribute{testAttrString, testAttrFloat}
var testAttrList EdgeAttribute = EdgeAttribute{"ListAttr", "attributes", attrList}

var edgeAttrList []EdgeAttribute = []EdgeAttribute{testAttrInteger, testAttrList}
var testElemDepth EdgeElement = EdgeElement{"elementTitle", edgeAttrList}

var testElem EdgeElement = EdgeElement{"elementTitle", attrList}
var edgeElemList []EdgeElement = []EdgeElement{testElemDepth, testElem}
var testData EdgeData = EdgeData{"version", "dataTitle", edgeElemList}
