# fastengine

[![Build Status](https://travis-ci.org/Guardian-Development/FASTEngine.svg?branch=master)](https://travis-ci.org/Guardian-Development/FASTEngine)

A FAST engine implementation in go. This package is based on the fast specification 1.1. 

# how to use

Given you have your fast templates in an xml based format, you can initialise the fast engine like so:

```go
package main 

import (
    "bytes"
    "fmt"
    "log"
    "os"
    
    "github.com/Guardian-Development/fastengine/pkg/engine"
)

func main() { 
    // create engine
    logger := log.New(os.Stdout, "engine: ", log.Ldate|log.Ltime|log.Lshortfile)
    fastEngine, err := engine.NewFromTemplateFile("file path to fasttemplates.xml", logger)
    if err != nil {
        // handle load failure
    }
    
    // fast encoded message to read
    message := bytes.NewBuffer([]byte{192, 1, 144, 138, 139})
    
    // read message
    fixMessage, err := fastEngine.Deserialise(message)
    
    if err != nil {
    	// handle problem reading message
    }
    
    fmt.Printf("%v", fixMessage)
}
```

The engine is not thread safe, as it has to use a dictionary of previous values based on your templates, to decode the fast messages. Therefore, it is recommended you initialise multiple engines for different feeds. 

If you wish to only load the templates once, you can use the following code to initialise multiple engines from the same template store:

```go
package main 

import (
    "log"
    "os"
    
    "github.com/Guardian-Development/fastengine/pkg/engine"
    "github.com/Guardian-Development/fastengine/pkg/fast/template/loader"
)

func main() { 
    logger := log.New(os.Stdout, "engine: ", log.Ldate|log.Ltime|log.Lshortfile)
    
    // load templates
    file, err := os.Open("file path to fasttemplates.xml")
    if err != nil {
        // handle file error
    }
    templateStore, err := loader.Load(file, logger)
    if err != nil {
        // handle store load failure
    }

    // create engines from same store
    fastEngine1 := engine.New(templateStore, logger)
    fastEngine2 := engine.New(templateStore, logger)
    fastEngine3 := engine.New(templateStore, logger)
}
```

## example message decoding

Given the following message template:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<templates xmlns="http://www.fixprotocol.org/ns/fast/td/1.1">
    <template name="MDHeartbeat_144" id="144" dictionary="144" xmlns="http://www.fixprotocol.org/ns/fast/td/1.1">
        <string name="ApplVerID" id="1128">
            <constant value="9"/>
        </string>
        <string name="MsgType" id="35">
            <constant value="0"/>
        </string>
        <uInt32 name="MsgSeqNum" id="34"/>
        <uInt64 name="SendingTime" id="52"/>
    </template>
</templates>
```

and the corresponding message to decode arrives: 

```go
/*
    Message format:
    11000000           pmap
    00000001 10010000  template 144
    10001010           34 = 10
    10001011           52 = 11
*/
message := bytes.NewBuffer([]byte{192, 1, 144, 138, 139})
```

the engine can be used like so, to decode the message:
```go
package main 

import (
    "bytes"
    "fmt"
    "log"
    "os"
    
    "github.com/Guardian-Development/fastengine/pkg/engine"
)

func main() { 
    // create engine
    logger := log.New(os.Stdout, "engine: ", log.Ldate|log.Ltime|log.Lshortfile)
    fastEngine, err := engine.NewFromTemplateFile("ExampleTemplate.xml", logger)
    if err != nil {
        panic("unable to load templates, stopping application")
    }
    
    // read message
    fixMessage, err := fastEngine.Deserialise(message)
    
    if err != nil {
    	panic("unable to read message, stopping application")
    }
    
    fmt.Printf("%v", fixMessage)
    // this produces => 1128=9|35=0|34=10|52=11|
}
```

explaining fix message result:
- template 144 used: we read pmap as a single byte (11000000). As there is a 1 at position 2, we know the template id is encoded. We decode this as 00000001 10010000 (remove stop bit => 0000001 0010000 => 00000010010000 => 144). We then use this template to decode the rest of the message.
- 1128=9: this is not encoded in the message as this is a required constant operation in the template, so the constant value is returned.
- 35=0: this is not encoded in message as this is a required constant operation in the template, so the constant value is returned.
- 34=10: this is encoded in the message and we read byte 10001010 as 10 (remove stop bit => 0001010 => 10).
- 52=11: this is encoded in the message and we read byte 10001011 as 11 (remove stop bit => 0001011 => 11).

# logging

The package aims to provide minimal logging overhead or rely on opinionated dependencies. The library will only log when an error occurs, and will provide information as to why the error has occurred before returning an appropriate error to the user application.

# limitations

This library only provides a global dictionary, that is reset between every message. This is due to the main use cases of fast being to send messages over Multicast, where relying on a previous value can become difficult. If there is interest on having more dictionary support, the library can easily be extended to support this.

# project structure

```
pkg
 ┣ engine
 ┃ ┣ engine.go : contains the main application entry point. This loads templates using the template_loader.go to create a store, then uses templates in store to decode messages.
 ┣ fast
 ┃ ┣ decoder
 ┃ ┃ ┣ decoder.go : provides the binary level decoder logic for reading fast values
 ┃ ┣ dictionary
 ┃ ┃ ┗ dictionary.go : provides a key value store for previous values
 ┃ ┣ errors
 ┃ ┃ ┗ errors.go : provides error messages based on the fast 1.1 spec
 ┃ ┣ field
 ┃ ┃ ┣ fieldasciistring
 ┃ ┃ ┃ ┣ field.go : contains logic for decoding ascii strings
 ┃ ┃ ┣ fieldbytevector
 ┃ ┃ ┃ ┣ field.go : contains logic for decoding byte vectors
 ┃ ┃ ┣ fielddecimal
 ┃ ┃ ┃ ┣ field.go : contains logic for decoding decimals
 ┃ ┃ ┣ fieldint32
 ┃ ┃ ┃ ┣ field.go : contains logic for decoding int32
 ┃ ┃ ┣ fieldint64
 ┃ ┃ ┃ ┣ field.go : contains logic for decoding int64
 ┃ ┃ ┣ fieldsequence
 ┃ ┃ ┃ ┣ field.go : contains logic for decoding sequences
 ┃ ┃ ┣ fielduint32
 ┃ ┃ ┃ ┣ field.go : contains logic for decoding uint32
 ┃ ┃ ┣ fielduint64
 ┃ ┃ ┃ ┣ field.go : contains logic for decoding uint64
 ┃ ┃ ┣ fieldunicodestring
 ┃ ┃ ┃ ┣ field.go : contains logic for decoding unicode strings
 ┃ ┃ ┗ properties
 ┃ ┃ ┃ ┗ properties.go : contains properties all fast fields must have (id for example)
 ┃ ┣ header
 ┃ ┃ ┗ message_header.go : contains logic for reading the header of a fast message in order to get its pmap and template to use for decoding
 ┃ ┣ operation
 ┃ ┃ ┗ operation.go : contains logic for all operation that can be applied to fields
 ┃ ┣ presencemap
 ┃ ┃ ┣ presence_map.go : contains logic for interrogating a presence map
 ┃ ┣ template
 ┃ ┃ ┣ loader
 ┃ ┃ ┃ ┣ converter
 ┃ ┃ ┃ ┃ ┣ value_converter.go : converts strings found in xml templates to their correct values 
 ┃ ┃ ┃ ┣ loadasciistring
 ┃ ┃ ┃ ┃ ┗ loader.go : loads asciistring from xml
 ┃ ┃ ┃ ┣ loadbytevector
 ┃ ┃ ┃ ┃ ┗ loader.go : loads bytevector from xml
 ┃ ┃ ┃ ┣ loaddecimal
 ┃ ┃ ┃ ┃ ┗ loader.go : loads decimal from xml
 ┃ ┃ ┃ ┣ loadint32
 ┃ ┃ ┃ ┃ ┗ loader.go : loads int32 from xml
 ┃ ┃ ┃ ┣ loadint64
 ┃ ┃ ┃ ┃ ┗ loader.go : loads int64 from xml
 ┃ ┃ ┃ ┣ loadproperties
 ┃ ┃ ┃ ┃ ┗ loader.go : loads common properties for all fields from xml
 ┃ ┃ ┃ ┣ loaduint32
 ┃ ┃ ┃ ┃ ┗ loader.go : loads uint32 from xml
 ┃ ┃ ┃ ┣ loaduint64
 ┃ ┃ ┃ ┃ ┗ loader.go : loads uint64 from xml
 ┃ ┃ ┃ ┣ loadunicodestring
 ┃ ┃ ┃ ┃ ┗ loader.go : loads unicodestring from xml
 ┃ ┃ ┃ ┣ template_loader.go : reads the xml templates, identifies the type of each element (uint32, int32 etc) then uses the appropriate loader to load the field
 ┃ ┃ ┣ store
 ┃ ┃ ┃ ┗ template_store.go : represents a loaded set of templates that can be used to decode messages
 ┃ ┃ ┗ structure
 ┃ ┃ ┃ ┗ structure.go : contains constants for xml tags
 ┃ ┗ value
 ┃ ┃ ┗ value.go : represents a fast value read from byte buffer (decoders read into these types)
 ┗ fix
 ┃ ┗ fix.go : represents a fix value (the engine returns these types, fields decode from fast values to fix values using operations)
```

## building from source

In order to build this project, running all tests, the following command is used from the root of the project: 

```bash
docker build --no-cache -t fastengine .
```
