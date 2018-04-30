S-JTSK to GPS ( WGS84 ) coordinates convertor
=============================================


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

