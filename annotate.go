package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"
)

func annotate(filePath string) {
	files, err := visit(filePath)
	if err != nil {
		fmt.Printf("error walking the path %v\n", err)
		return
	}

	for _, file := range files {
		tempFile := "tempfile.go"

		outfile, err := os.Create(tempFile)
		if err != nil {
			panic(err)
		}

		fset := token.NewFileSet()
		dirs, err := parser.ParseDir(fset, filePath, nil, 0)

		if err != nil {
			panic(err)
		}

		for _, dir := range dirs {
			for _, file := range dir.Files {
				ast.Inspect(file, func(n ast.Node) bool {
					switch x := n.(type) {
					case *ast.TypeSpec:
						if x.Name.String() == file.Name.String() {
							v := x.Type.(*ast.StructType)
							for _, field := range v.Fields.List {
								for _, n := range field.Names {
									var typeNameBuf bytes.Buffer
									err := printer.Fprint(&typeNameBuf, fset, field.Type)
									if err != nil {
										log.Fatalf("failed printing %s", err)
									}
									fmt.Printf("field %+v has type %+v\n", n.Name, typeNameBuf.String())
								}
							}
						}

					}

					return true
				})
			}
		}

		_, err = outfile.WriteString("// == Schema Info\n//\n// Table name: line_items\n")

		defer outfile.Close()

		f, err := os.Open(file)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			outfile.WriteString(scanner.Text())
			outfile.WriteString("\n")
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}

		outfile.Sync()

		err = os.Rename(tempFile, file)
		if err != nil {
			panic(err)
		}
	}
}

func visit(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

type mapKey struct {
	Key    int
	Option string
}
