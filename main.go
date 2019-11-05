package main

import (
	"flag"
	"log"
	"os"
	"path"

	"github.com/netlify/godoc-static/gen"
	"github.com/netlify/godoc-static/static"
)

const out = "dist"

func main() {
	var srcPath string
	flag.StringVar(&srcPath, "path", "", "path to the `source`")
	var mod string
	flag.StringVar(&mod, "mod", "", "path of the `module`")
	flag.Parse()

	if srcPath == "" || mod == "" {
		flag.PrintDefaults()
		os.Exit(2)
	}

	render := gen.NewRenderer(srcPath, mod, out)
	if err := render.GenerateAll("/", out); err != nil {
		log.Fatal(err)
	}

	if err := static.OutputResources(path.Join(out, "lib/godoc")); err != nil {
		log.Fatal(err)
	}

	if err := render.Make404("404.html"); err != nil {
		log.Fatal(err)
	}
}
