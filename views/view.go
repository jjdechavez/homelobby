package views

import (
	"net/http"
	"path/filepath"
	"text/template"
)

var LayoutDir string = "views/layouts"

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func InitView(r *http.Request, layout string, files ...string) *View {
	hxRequest := r.Header.Get("HX-Request")
	isHxRequest := hxRequest == "true"

	files = append(layoutFiles(), files...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	defaultLayout := layout
	if isHxRequest {
		defaultLayout = "body"
	}

	return &View{
		Template: t,
		Layout:   defaultLayout,
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*html")
	if err != nil {
		panic(err)
	}
	return files
}
