Command Data Model
=======================================
Command Data Model is for command of metadata.
It will get you a easy expression of your request or response data which has a lot of categories and attributes. 
And Also, It's good at expressing a complex data which has depths. because Attribute Category can contain another Attribute.
It provides some mandatory categories in below.

- version - version of the datamodel (default : edge-1.0)
- dataTitle - protocol or data title
- edgeElements - element unit like command of metadata.
- elementTitle - each element of edgeElements list can be expressed as operation of metadata.
- edgeAttributes - attribute list
- name - attribute name
- dataType - data type
- value - actual value of the data

## Basic Format ##

```
{
  "version" : A
  "dataTitle" : B
  "edgeElements" : [
     {
       "elementTitle" : C
       "edgeAttributes" : [
          {
             "name" : D
             "dataType" : E
             "value" : F
          },
          ...
                          ]
     },
     ...
                   ]
}
```

## Example for OPCUA Device Service ##

#### READ ####

  - Request
```
{
  "version" : "edge-1.0"
  "dataTitle" : "opcua"
  "edgeElements" : [
     { 
       "elementTitle" : "read"
       "edgeAttributes" : [
          {
             "name" : "value_descriptor"
             "dataType" : "string"
             "value" : "temperature"
          }
                          ]
     }
                   ]
}
```

  - Response
```
{
  "version" : "edge-1.0"
  "dataTitle" : "opcua"
  "edgeElements" : [
     {
       "elementTitle" : "read"
       "edgeAttributes" : [
          {
             "name" : "response_information"
             "dataType" : "attribute"
             "value" : [
                {
                    "name" : "value_descriptor"
                    "dataType" : "string"
                    "value" : "temperature"
                },
                {
                    "name" : "value"
                    "dataType" : "string"
                    "value" : "80F"
                }
          },
          {
             "name" : "result"
             "dataType" : "string"
             "value" : "edge_response_ok"
          }
                          ]
     }
                   ]
}
  ```

#### WRITE ####

  - Request
```
{
  "version" : "edge-1.0"
  "dataTitle" : "opcua"
  "edgeElements" : [
     { 
       "elementTitle" : "write"
       "edgeAttributes" : [
          {
             "name" : "value_descriptor"
             "dataType" : "string"
             "value" : "fan"
          },
          {
             "name" : "input_argument"
             "dataType" : "string"
             "value" : "off"
          }
                          ]
     }
                   ]
}
```

  - Response
```
{
  "version" : "edge-1.0"
  "dataTitle" : "opcua"
  "edgeElements" : [
     {
       "elementTitle" : "write"
       "edgeAttributes" : [
          {
             "name" : "value_descriptor"
             "dataType" : "string"
             "value" : "fan"
          },
          {
             "name" : "result"
             "dataType" : "string"
             "value" : "edge_response_ok"
          }
                          ]
     }
                   ]
}
```

#### SUBSCRIPTION (Monitoring) ####

  - Request
```
{
  "version" : "edge-1.0"
  "dataTitle" : "opcua"
  "edgeElements" : [
     { 
       "elementTitle" : "sub"
       "edgeAttributes" : [
          {
             "name" : "value_descriptor"
             "dataType" : "string"
             "value" : "temperature"
          },
          {
             "name" : "sampling_interval"
             "dataType" : "double"
             "value" : "1000.0"
          }
                          ]
     }
                   ]
}
```

  - Response
```
{
  "version" : "edge-1.0"
  "dataTitle" : "opcua"
  "edgeElements" : [
     {
       "elementTitle" : "sub"
       "edgeAttributes" : [
          {
             "name" : "response_information"
             "dataType" : "attribute"
             "value" : [
                {
                    "name" : "value_descriptor"
                    "dataType" : "string"
                    "value" : "temperature"
                },
                {
                    "name" : "revised_sample_interval"
                    "dataType" : "double"
                    "value" : "1000.0"
                }
          },
          {
             "name" : "result"
             "dataType" : "string"
             "value" : "edge_response_ok"
          }
                          ]
     }
                   ]
}
```
