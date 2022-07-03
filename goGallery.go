package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/exp/slices"
)

var Router *mux.Router

var Root string
var Port string
var DirsStyle string
var ImagesStyle string

const xmlHeader = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<?xml-stylesheet type=\"text/xsl\" href=\"%s\"?>\n<%s>\n%s\n</%s>"

var images = []string{"png", "jpg", "jpeg", "gif", "bmp"}
var videos = []string{"mp4", "mov", "avi"}

type File struct {
	Path string
	Type string
}

type Directory struct {
	Name  string
	Count int
}

func main() {
	Root = os.Getenv("GALLERYROOT")
	if Root == "" {
		panic("ERROR: GALLERYROOT=\"\"")
	}
	Port = os.Getenv("GALLERYPORT")
	if Port == "" {
		panic("ERROR: GALLERYPORT=\"\"")
	}
	DirsStyle = os.Getenv("GALLERYDIRSSTYLE")
	if DirsStyle == "" {
		panic("ERROR: GALLERYDIRSSTYLE=\"\"")
	}
	ImagesStyle = os.Getenv("GALLERYIMAGESSTYLE")
	if ImagesStyle == "" {
		panic("ERROR: GALLERYIMAGESSTYLE=\"\"")
	}

	Router = mux.NewRouter()
	Router.HandleFunc("/", dirList)
	Router.HandleFunc("/{dir}", imageList)

	http.ListenAndServe(Port, Router)
}

func dirList(w http.ResponseWriter, r *http.Request) {
	dirs, err := ioutil.ReadDir(Root)
	if err != nil {
		panic(err)
	}

	l := make([]Directory, 0)
	for _, d := range dirs {
		if d.IsDir() {
			x, err := ioutil.ReadDir(filepath.Join(Root, d.Name()))
			if err != nil {
				panic(err)
			}
			l = append(l, Directory{Name: d.Name(), Count: len(x)})
		}
	}

	data, _ := xml.MarshalIndent(l, "", " ")
	w.Write([]byte(fmt.Sprintf(xmlHeader, DirsStyle, "dirs", string(data), "dirs")))
}

func imageList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data, _ := xml.MarshalIndent(getImages(filepath.Join(Root, vars["dir"]), vars["dir"]), "", " ")
	w.Write([]byte(fmt.Sprintf(xmlHeader, ImagesStyle, "images", string(data), "images")))
}

func getImages(path string, dir string) []File {
	l := make([]File, 0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if !f.IsDir() {
			ftype := "unknown"
			split := strings.Split(f.Name(), ".")
			if slices.Contains(images, split[len(split)-1]) {
				ftype = "image"
			}
			if slices.Contains(videos, split[len(split)-1]) {
				ftype = "video"
			}
			l = append(l, File{Path: fmt.Sprintf("%s/%s", dir, f.Name()), Type: ftype})
		}
	}
	return l
}