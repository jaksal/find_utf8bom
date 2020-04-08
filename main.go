package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dimchansky/utfbom"
)

var path string

func init() {
	if len(os.Args) != 2 {
		panic("invalid find path!")
	}
	path = os.Args[1]
	//flag.StringVar(&path, "path", ".", "scan path")
	//flag.Parse()

	fmt.Println("path=", path)
}

func main() {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() == false {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			_, enc := utfbom.Skip(f)
			if enc != utfbom.Unknown {
				fmt.Printf("%s ==> %s\n", path, enc)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
