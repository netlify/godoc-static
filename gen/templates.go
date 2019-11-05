package gen

import (
	"log"
	"text/template"

	mystatic "github.com/netlify/godoc-static/static"
	"golang.org/x/tools/godoc/static"
)

func (r *Renderer) readGodocTemplate(name string) *template.Template {
	file, ok := static.Files[name]
	if !ok {
		log.Fatalf("missing template file: %s", name)
	}

	return r.templateFrom(name, file)
}

func (r *Renderer) templateFrom(name, source string) *template.Template {
	t, err := template.New(name).Funcs(r.pres.FuncMap()).Parse(source)
	if err != nil {
		log.Fatal("readTemplate: ", err)
	}
	return t
}

func (r *Renderer) readTemplates() {
	r.pres.CallGraphHTML = r.readGodocTemplate("callgraph.html")
	r.pres.DirlistHTML = r.readGodocTemplate("dirlist.html")
	r.pres.ErrorHTML = r.readGodocTemplate("error.html")
	r.pres.ExampleHTML = r.readGodocTemplate("example.html")
	r.pres.ImplementsHTML = r.readGodocTemplate("implements.html")
	r.pres.MethodSetHTML = r.readGodocTemplate("methodset.html")
	r.pres.PackageHTML = r.readGodocTemplate("package.html")
	r.pres.PackageRootHTML = r.readGodocTemplate("packageroot.html")
	r.pres.SearchHTML = r.readGodocTemplate("search.html")
	r.pres.SearchDocHTML = r.readGodocTemplate("searchdoc.html")
	r.pres.SearchCodeHTML = r.readGodocTemplate("searchcode.html")
	r.pres.SearchTxtHTML = r.readGodocTemplate("searchtxt.html")
	r.pres.SearchDescXML = r.readGodocTemplate("opensearch.xml")

	//r.pres.GodocHTML = r.readGodocTemplate("godoc.html")
	r.pres.GodocHTML = r.templateFrom("godoc.html", mystatic.GodocTemplate)
}
