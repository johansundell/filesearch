package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

var path = `\\10.103.165.20\341510Common$\OX2 Vindel\Kundservices\CST-Dokument\`
var pathScanner = `\\10.103.165.20\341510Common$\OX2 Vindel\Kundservices\Scanner\`
var router *mux.Router
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/cst/{fmid}/",
		Myhandler,
	},
	Route{
		"scanner",
		"GET",
		"/scanner/{fmid}/",
		ScannerHandler,
	},
}

func init() {
	//path = `\\10.103.165.20\341510Common$\OX2 Vindel\Kundservices\CST-Dokument\`
}

/*func main() {
	//findTest(pathScanner, "2929")
	//return
	router = NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr:    ":8081",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 90 * time.Second,
		ReadTimeout:  90 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}*/

func ScannerHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pattern := vars["fmid"]
	//fmt.Println(pattern)
	files, err := findTest(pathScanner, pattern)
	if err != nil {
		log.Println(err)
	}

	for _, pa := range files {
		_, f := filepath.Split(pa)
		fmt.Fprintf(w, "<a target=\"_blank\" href=\"file://"+pa[2:]+"\">"+f+"</a></br>")

	}
	//fmt.Fprint(w, "test")
}

func Myhandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	//fmt.Println(vars)
	pattern := vars["fmid"]

	//fmt.Println(path)
	files, err := findMatches(path, pattern)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(len(files))

	for _, pa := range files {
		_, f := filepath.Split(pa)
		fmt.Fprintf(w, "<a target=\"_blank\" href=\"file://"+pa[2:]+"\">"+f+"</a></br>")

	}
	//fmt.Fprint(w, "test")
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

	files, err := filepath.Glob(path + "*(FM ID " + pattern + ")*.pdf")

	return files, err
}

func findTest(path, pattern string) ([]string, error) {
	result := []string{}
	fi, err := ioutil.ReadDir(path)
	if err != nil {
		return result, err
	}

	for _, v := range fi {
		if v.IsDir() {
			//fmt.Println(v.Name())
			files, _ := findMatches(path+v.Name()+"\\", pattern)
			//fmt.Println(files)
			result = append(result, files...)
		}
	}
	//fmt.Println(result)
	return result, err
}

func findDeep(path, pattern string) ([]string, error) {
	result := []string{}
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		/*if !info.IsDir() && info.Mode().IsRegular() {
			parts := strings.FieldsFunc(path, func(c rune) bool { return os.PathSeparator == c })
			file := parts[len(parts)-1]
			if strings.Contains(file, ".pdf") && strings.Contains(file, pattern) {
				fmt.Println(path)
			}
			files, _ := findMatches()
		}*/
		if info.IsDir() {
			fmt.Println(path)
		}
		return nil
	})
	return result, nil
}
