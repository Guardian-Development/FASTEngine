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
    "bytes"
	"fmt"
    "log"
    "os"
	
	"github.com/Guardian-Development/fastengine/pkg/engine"
    "github.com/Guardian-Development/fastengine/pkg/fast/template/loader"
)

func main() { 
    // load templates
    file, err := os.Open("file path to fasttemplates.xml")
    if err != nil {
        // handle file error
    }
    templateStore, err := loader.Load(file)
    if err != nil {
        // handle store load failure
    }

    // create engines from same store
    logger := log.New(os.Stdout, "engine: ", log.Ldate|log.Ltime|log.Lshortfile)
    fastEngine1 := engine.New(templateStore, logger)
    fastEngine2 := engine.New(templateStore, logger)
    fastEngine3 := engine.New(templateStore, logger)
}
```

# logging

# limitations

# project structure
