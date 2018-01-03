Command Data Model (go-lang version)
=======================================

Command Data Model is for command of metadata. 
It will get you a easy expression of your data to send to micro service. And Also, It's good at expressing all data. because edgeAttributes can contain another edgeAttribute object.

## Prerequisites ##
- Go compiler
   - Version : 1.9.2
   - [How to install](https://golang.org/doc/install)
   
- Set `GOPATH` to `Project Path` to use datamodel-command-go library

## Installation ##
Enter as follows in linux command line
```
$ cd $GOPATH/src
$ git clone https://github.sec.samsung.net/RS7-EdgeComputing/datamodel-command-go.git
$ GOOS=linux go install ./datamodel-command-go/formatter/
```
As a result, `datamodel-command-go/formatter.a` is generated in `$GOPATH/pkg/[$GOOS_$GOARCH]`

## How to Use ##

To use the datamodel-command-go libraries you first need to import the libraries into your project :
```
import "datamodel-command-go/formatter"
```
You can Use data models(EdgeAttribute, EdgeElement, EdgeData) as below :
```
attr := formatter.EdgeAttribute{"[name]", "[dataType]", [value]}
attr.ConvertToEdgeAttribute()
```
You can Use encoder and decoder as below :
```
jsonString, err := formatter.EncodeEdgeAttributeToJsonString(&EdgeAttribute)
EdgeAttribute, err := formatter.EncodeEdgeAttributeToJsonString(&jsonString)
```

## Unit test and code coverage report ##
1. GOTO : `~/datamodel-command-go`
2. Run the script `$ ./unittest.sh`
3. Open `coverall.html` in web browser
