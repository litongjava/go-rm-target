package main

import (
  "fmt"
  "os"
  "path/filepath"
)

func main() {
  deleteTaget()
}

func deleteTaget() {
  if len(os.Args) < 2 {
    fmt.Println("Usage: go-rm-target [path]")
    return
  }
  path := os.Args[1]
  fmt.Println("Path:", path)

  err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
    stat, _ := os.Stat(path)
    if stat == nil {
      return nil
    }
    if err != nil {
      fmt.Println("error", err)
    }
    if info.IsDir() {
      folderName := info.Name()
      if folderName == "target" {
        err := os.RemoveAll(path)
        if err != nil {
          fmt.Println("delete faild", path)
        } else {
          fmt.Println("deleted", path)
          return nil
        }
      }
    }
    return nil
  })

  if err != nil {
    fmt.Println("Error:", err)
  }

  fmt.Println("Done")
}
