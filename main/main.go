package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/djherbis/times"
	"github.com/russross/blackfriday"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("jimsk starting...")

	viper.SetEnvPrefix("app")

	viper.SetDefault("port", ":80")
	viper.SetDefault("static", "../static")

	viper.BindEnv("port")
	viper.BindEnv("static")

	http.HandleFunc("/", HandleMarkdown)
	http.HandleFunc("/assets/", HandleAssets)
	http.HandleFunc("/index.css", HandleCSS)
	http.HandleFunc("/index.js", HandleJS)

	log.Printf("port=%v", viper.GetString("port"))
	log.Printf("static=%v", viper.GetString("static"))

	http.ListenAndServe(viper.GetString("port"), nil)
}

func HandleMarkdown(w http.ResponseWriter, r *http.Request) {
	log.Println("requesting", r.URL.Path)

	f, err := markdownFile(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	renderResponse(w, f)
}

func HandleAssets(w http.ResponseWriter, r *http.Request) {
	log.Println("requesting asset", r.URL.Path)
	http.ServeFile(w, r, path.Join(viper.GetString("static"), r.URL.Path))
}

func markdownFile(r *http.Request) (md string, err error) {
	root := viper.GetString("static")
	segs := []string{}

	for _, p := range strings.Split(r.URL.Path, "/") {
		if len(p) == 0 {
			continue
		}

		lookup := path.Join(root, strings.Join(segs, "/"), p)

		// check lookup as directory
		if _, err := os.Stat(lookup); err == nil {
			segs = append(segs, p)
			log.Printf("\tsegs=%v", segs)
			continue
		}

		// check lookup as markdown file
		if _, err := os.Stat(fmt.Sprintf("%v.md", lookup)); err == nil {
			segs = append(segs, p)
			md = path.Join(root, strings.Join(segs, "/")+".md")
			log.Printf("\tsegs=%v;md=%v", segs, md)
			return md, nil
		}

		// fallback: lookup is no longer valid
		md = path.Join(root, strings.Join(segs, "/"), "index.md")
		log.Printf("\tsegs=%v;p=%v (404)", segs, p)

		return md, fmt.Errorf("resource just doesn't exist")
	}

	// loop ended on a directory so try returning the index.md inside that directory
	if len(segs) > 0 {
		md = path.Join(root, strings.Join(segs, "/"), "index.md")

		if _, err := os.Stat(md); err == nil {
			log.Printf("\tsegs=%v;=>md=%v", segs, md)
			return md, nil
		}
	}

	// fallback: return resume
	md = path.Join(root, "resume/index.md")
	log.Printf("\tsegs=%v;=>md=%v", segs, md)
	return
}

func renderResponse(w http.ResponseWriter, markdownFile string) {
	log.Println("\t\trendering", markdownFile)

	root := viper.GetString("static")

	t, err := template.ParseFiles(
		path.Join(root, "index.html"),
		path.Join(root, "header.html"),
		path.Join(root, "footer.html"),
	)

	data := struct {
		Header       template.HTML
		Content      template.HTML
		DateModified string
		Footer       template.HTML
	}{}

	// read content markdown
	content, err := ioutil.ReadFile(markdownFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	data.Content = template.HTML(blackfriday.Run(content, blackfriday.WithNoExtensions()))

	// get content modification date
	// fi, err := os.Lstat(markdownFile)
	ti, err := times.Stat(markdownFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	data.DateModified = ti.ModTime().Format("01/02/06")

	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleCSS(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-Control", "no-cache")
	http.ServeFile(w, r, path.Join(viper.GetString("static"), "index.css"))
}

func HandleJS(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-Control", "no-cache")
	http.ServeFile(w, r, path.Join(viper.GetString("static"), "index.js"))
}
