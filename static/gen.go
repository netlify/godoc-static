package static

//go:generate go run templates/makestatic.go

var Files = map[string]string{
	"GodocTemplate": "templates/godoc.html",
}
