package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

//go:embed index.html bg.jpg
var fs embed.FS

const port = ":1988"

func ditch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	now := time.Now().Format("20060102T1504")
	f, err := os.OpenFile(fmt.Sprintf("signups.%s.txt", now), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ditch(err)
	defer f.Close()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFS(fs, "index.html")
		ditch(err)
		ditch(t.Execute(w, false))
	})
	http.HandleFunc("/email", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFS(fs, "index.html")
		ditch(err)
		ditch(t.Execute(w, true))
	})
	http.Handle("/bg.jpg", http.FileServer(http.FS(fs)))
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		ditch(r.ParseForm())
		email := r.PostFormValue("email")
		crsid := r.PostFormValue("crsid")
		log.Printf("Got response: crsid: %s, email: %s\n", crsid, email)
		if email != "" {
			io.WriteString(f, email)
			io.WriteString(f, "\r\n")
		}
		if crsid != "" {
			io.WriteString(f, crsid)
			io.WriteString(f, "@cam.ac.uk\r\n")
		}
		http.Redirect(w, r, "/", 302)
	})
	url := fmt.Sprintf("http://localhost%s", port)
	openbrowser(url)
	log.Printf("Listening on %s. Opening browser for you...", url)
	ditch(http.ListenAndServe(port, nil))
}
