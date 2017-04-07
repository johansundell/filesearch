package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	http.HandleFunc("/", Myhandler)
	http.ListenAndServe(":8080", nil)
}

func Myhandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tmp := r.Form["pattern"]
	pattern := ""
	if len(tmp) > 0 {
		pattern = tmp[0]
	}
	if pattern == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	files, err := findMatches("/home/johan/downloads/", pattern)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(files)

	fmt.Fprintf(w, "ok")
}

func findMatches(path, pattern string) ([]string, error) {
	/*filepath.Walk("/home/johan/downloads/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Mode().IsRegular() {
			parts := strings.FieldsFunc(path, func(c rune) bool { return os.PathSeparator == c })
			file := parts[len(parts)-1]
			if strings.Contains(file, ".go") && strings.Contains(file, "main") {
				fmt.Println(path)
			}
		}
		return nil
	})*/

	files, err := filepath.Glob(path + pattern)
	return files, err
}
