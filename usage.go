package main

import(
"fmt"
)

func usage() {
  fmt.Printf("vmount v.%s (c) Lyderic Landry, London 2017\n", version)
  fmt.Println("Usage: vmount <action> <arguments>")
  fmt.Println("  -l, --list       list favorites")
  fmt.Println("  -m, --mount      mount all favorites")
  fmt.Println("  -d, --dismount   dismount all favorites")
  fmt.Println("  -d, --dismount # dismount favorite mounted on slot #")
  fmt.Println("  -e, --edit       edit favorites XML configuration")
  fmt.Println("  --version")
}
