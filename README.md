# cfg

A Go library to work with JSON configuration files

## Usage

```json
{
    "hello": "world",
    "welcome": {
        "bonjour": ["pays", "monde", "univers"],
    },
    "dang": 11
}
```
Define your struct. You can add the tag `cfg:"optionnal"` to any field except for numeric and boolean types as Go has default value for them, we can't check if the field is missing or not.
```go
type Conf struct {
    Hello   string
    Welcome struct {
        Bonjour     []string
        AuRevoir    []string `cfg:"optional"`
    }
    Dang    int
}
```
And simply:
```go
package main

import "gopkg.in/Gamoloco/cfg.v1"

func main() {
    var conf Conf
    err := cfg.Load("path/to/file.json", &conf)
}
```
