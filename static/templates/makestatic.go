// Adapted from https://github.com/golang/tools/blob/master/godoc/static/gen.go
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"unicode"

	"github.com/netlify/godoc-static/static"
)

func main() {
	if err := makestatic(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func makestatic() error {
	buf := new(bytes.Buffer)
	fmt.Fprint(buf, "package static\n\n")
	for name, fn := range static.Files {
		b, err := ioutil.ReadFile(fn)
		if err != nil {
			return err
		}
		fmt.Fprintf(buf, "var %s string = ", name)
		appendQuote(buf, b)
		fmt.Fprintf(buf, "\n\n")
	}
	output, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("files.go", output, 0666)
	if err != nil {
		return fmt.Errorf("error while writing file: %v", err)
	}
	return nil
}

// appendQuote is like strconv.AppendQuote, but we avoid the latter
// because it changes when Unicode evolves
func appendQuote(out *bytes.Buffer, data []byte) {
	out.WriteByte('"')
	for _, b := range data {
		if b == '\\' || b == '"' {
			out.WriteByte('\\')
			out.WriteByte(b)
		} else if b <= unicode.MaxASCII && unicode.IsPrint(rune(b)) && !unicode.IsSpace(rune(b)) {
			out.WriteByte(b)
		} else {
			fmt.Fprintf(out, "\\x%02x", b)
		}
	}
	out.WriteByte('"')
}
