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

	http.HandleFunc("/", HandleRoot)
	http.HandleFunc("/css/", HandleCSS)

	fmt.Printf("listening on %v\n", viper.GetString("port"))
	http.ListenAndServe(viper.GetString("port"), nil)
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	static := viper.GetString("static")

	t, err := template.ParseFiles(path.Join(static, "template/cv.template"))

	cv, err := ioutil.ReadFile(path.Join(static, "markdown/cv.md"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cv = blackfriday.Run(cv, blackfriday.WithNoExtensions())

	if err := t.Execute(w, struct {
		Content template.HTML
	}{template.HTML(string(cv))}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleCSS(w http.ResponseWriter, r *http.Request) {
	static := viper.GetString("static")

	w.Header().Add("Cache-Control", "no-cache")

	http.ServeFile(w, r, path.Join(static, r.URL.Path))
}
