package gen

import (
	"net/http"
	"path"
)

func (r *Renderer) makeRequest(pkgPath string) (*http.Request, error) {
	url := path.Join("/pkg", r.ModPath, pkgPath) + "/"
	return http.NewRequest(http.MethodGet, url, nil)
}

func (r *Renderer) GenerateAll(pkgPath, targetDir string) error {
	rootReq, err := r.makeRequest("")
	if err != nil {
		return err
	}

	if err := r.Generate(""); err != nil {
		return err
	}

	info := r.pres.GetPkgPageInfo(path.Join("/src", r.ModPath), r.ModPath, r.pres.GetPageInfoMode(rootReq))
	for _, dir := range info.Dirs.List {
		err := r.Generate(dir.Path)
		if err != nil {
			return err
		}
	}
	return nil
}
