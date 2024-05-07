# go-extract-vmlinux

A tool to extract vmlinux/vmlinuz in golang

Inspired from the official script from linux [extract-vmlinux](https://raw.githubusercontent.com/torvalds/linux/master/scripts/extract-vmlinux)

## How to use

```go
package main

import (
  "log"
  extract "github.com/rawdaGastan/go-extract-vmlinux"
)

func main() {
  kernelImagePath := "<Enter your compressed kernel path>"

  // decompress kernel
  err := extract.TryDecompressKernel(kernelImagePath, kernelImagePath+"-decompressed")
  if err != nil {
    log.Fatal(err)
  }

  // compressed kernel
  err = extract.IsValidELFKernel(kernelImagePath + "-decompressed")
  if err != nil {
    log.Fatal(err)
  }
}
```

## How to test

```bash
make test
```
