package main

import (
    "fmt"
    "os"
    "log"
    "runtime"
	"path/filepath"
)

func findAllFiles(extension string, root string) []string {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
            fmt.Println(err)
            return nil
        }

        if !info.IsDir() && filepath.Ext(path) == ".txt" {
            files = append(files, path)
        }

        return nil
	})

	if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        fmt.Println(file)
    }

	return files
}

func main() {
    currentOS := runtime.GOOS
    fmt.Println("Your current OS is", currentOS)
    var root string

    switch currentOS {
    case "darwin":
		root = "/"
    case "linux":
        root = "/"
    case "windows":
        root = os.Getenv("SystemDrive") + "\\"
    default:
        fmt.Println("Your current OS isn't supported by VanSourceGuard")
        os.Exit(1)
    }

    entries, err := os.ReadDir(root)
    if err != nil {
        log.Fatal(err)
    }

    for _, e := range entries {
        fmt.Println(e.Name())
    }
}
