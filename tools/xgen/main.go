package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		die("xgen FILE1.go FILE2.go ...")
	}

	fset := token.NewFileSet()
	mode := parser.ImportsOnly | parser.ParseComments
	for _, f := range os.Args[1:] {
		if !strings.HasSuffix(f, ".go") {
			fmt.Fprintln(os.Stderr, "ignoring:", f)
			continue
		}
		a, err := parser.ParseFile(fset, f, nil, mode)
		checkErr(err)
		pkg := a.Name.Name
		for _, c := range a.Comments {
			txt := c.Text()
			switch {
			case strings.HasPrefix(txt, "BaseAddr:"):
				mmio(pkg, f, txt)
			}
		}
	}
}
