S-JTSK to GPS ( WGS84 ) coordinates convertor
=============================================

[![Go Report Card](https://goreportcard.com/badge/github.com/encero/sjtsk2gps)](https://goreportcard.com/report/github.com/encero/sjtsk2gps) [![GoDoc](https://godoc.org/github.com/encero/sjtsk2gps?status.svg)](https://godoc.org/github.com/encero/sjtsk2gps)

Usage

```
package main

import (
	"github.com/encero/sjtsk2gps"
)
func main() {
	var KrovakX, KrovakY, height float64
	
	lat, lon, height := sjtsk2gps.Convert(KrovakX, KrovakY, height)
}
```

