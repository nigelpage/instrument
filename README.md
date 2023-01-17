# instrument

instrument is a Go package for instrumenting your application.

## Installation

Use go get to include the package in your go module.

```bash
go get github.com/nigelpage/instrument
```

## Usage

```go
import {
    "fmt"
    ins "github.com/nigelpage/instrument"
}

# returns 'StructuredError'
var e = ins.NewStructuredError(Information, c, "new file created")

if e != nil {
    fmt.Println(e)
}
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)