package static

import (
	"os"
	"path"

	"golang.org/x/tools/godoc/static"
)

func OutputResources(out string) error {
	for file, body := range static.Files {
		if path.Ext(file) == ".html" {
			continue
		}

		filepath := path.Join(out, file)
		dir := path.Dir(filepath)
		if dir != "" {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				return err
			}
		}

		f, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.WriteString(body)
		if err != nil {
			return err
		}
	}
	return nil
}
