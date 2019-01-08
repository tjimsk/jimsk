package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/russross/blackfriday"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("jimsk starting...")

	viper.SetDefault("port", ":80")
	viper.SetDefault("static", "../static")
	viper.BindEnv("port", "GOPORT")
	viper.BindEnv("static", "GOSTATIC")

	http.HandleFunc("/", HandleCV)
	http.HandleFunc("/css/", HandleCSS)

	fmt.Printf("listening on %v\n", viper.GetString("port"))
	http.ListenAndServe(viper.GetString("port"), nil)
}

func HandleCV(w http.ResponseWriter, r *http.Request) {
	static := viper.GetString("static")

	t, err := template.ParseFiles(
		path.Join(static, "template/cv.template"),
	)

	cv, err := ioutil.ReadFile(path.Join(static, "markdown/cv.md"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cv = blackfriday.Run(cv, blackfriday.WithNoExtensions())

	footer, err := ioutil.ReadFile(path.Join(static, "markdown/footer.md"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	footer = blackfriday.Run(footer, blackfriday.WithNoExtensions())

	data := struct {
		Content template.HTML
		Footer  template.HTML
	}{
		template.HTML(string(cv)),
		template.HTML(string(footer)),
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleCSS(w http.ResponseWriter, r *http.Request) {
	static := viper.GetString("static")

	w.Header().Add("Cache-Control", "no-cache")

	http.ServeFile(w, r, path.Join(static, r.URL.Path))
}
