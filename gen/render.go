package gen

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path"

	"golang.org/x/tools/godoc"
	"golang.org/x/tools/godoc/static"
	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/mapfs"
)

type Renderer struct {
	ModPath    string
	TargetPath string

	corpus *godoc.Corpus
	pres   *godoc.Presentation
}

func NewRenderer(path, mod, target string) *Renderer {
	fs := vfs.NewNameSpace()
	fs.Bind("/lib/godoc", mapfs.New(static.Files), "/", vfs.BindReplace)

	srcFS := vfs.OS(path)
	fs.Bind("/src/"+mod, srcFS, "/", vfs.BindReplace)

	corpus := godoc.NewCorpus(fs)
	corpus.Init()
	pres := godoc.NewPresentation(corpus)

	r := &Renderer{
		ModPath:    mod,
		TargetPath: target,
		corpus:     corpus,
		pres:       pres,
	}
	r.readTemplates()
	return r
}

func (r *Renderer) Generate(pkgPath string) error {
	req, err := r.makeRequest(pkgPath)
	if err != nil {
		return err
	}

	rec := httptest.NewRecorder()
	r.pres.ServeHTTP(rec, req)

	res := rec.Result()
	status := res.StatusCode
	if status != http.StatusOK {
		log.Printf("error: %s", string(rec.Body.Bytes()))
		return fmt.Errorf("invalid status code: %d", status)
	}

	targetDir := path.Join(r.TargetPath, pkgPath)
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(path.Join(targetDir, "index.html"))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, rec.Result().Body)
	return err
}
