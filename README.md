# go-gsettings

global settings ( like os.env but different )

use in your package:

```go
package p1

import (
    gsettings "github.com/9glt/go-gsettings"
)

func P1() {
    value, err := gsettings.Get("app-prefix-key")
    if err != nil {
        panic(err)
    }
    println("got: ",value)
}
```

and use in your main app to set global variables to your packages:

```go
package main

import (
    // "github.com/your/package/p1"
    gsettings "github.com/9glt/go-gsettings"
)

func main() {
    gsettings.Set("app-prefix-key", "db://")
}
```
