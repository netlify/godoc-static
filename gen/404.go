package gen

import (
	"os"
	"path"

	"golang.org/x/tools/godoc"
)

func (r *Renderer) Make404(targetPath string) error {
	dir := path.Join(r.TargetPath, path.Dir(targetPath))
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(path.Join(r.TargetPath, targetPath))
	if err != nil {
		return err
	}
	defer f.Close()

	page := &godoc.Page{
		Tabtitle: "Not found",
		Body: []byte(`
			<p>Resource could not be found.</p>
			<p><a href="/">Back to the package listing</a></p>
		`),
	}
	return r.pres.GodocHTML.Execute(f, page)
}
