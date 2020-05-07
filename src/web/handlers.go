package main

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request, parm httprouter.Params) {
	userName, err1 := r.Cookie("username")
	sessionID, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		p := &HomePage{Name: "joex"}
		template, err := template.ParseFiles("../templates/home.html")
		if err != nil {
			log.Printf("Error during parseing html files: %v", err)
			return
		}
		template.Execute(w, p)
		return
	}
	if len(userName.Value) != 0 && len(sessionID.Value) != 0 {
		http.Redirect(w, r, "/userhome", http.StatusFound)
		return
	}
}

func userHomeHandler(w http.ResponseWriter, r *http.Request, parm httprouter.Params) {

	userName, err1 := r.Cookie("username")
	_, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	fname := r.FormValue("username")
	var p *UserPage

	if len(userName.Value) != 0 {
		p = &UserPage{Name: userName.Value}
	} else if len(fname) != 0 {
		p = &UserPage{Name: fname}
	}
	t, err := template.ParseFiles("../templates/userhome.html")
	if err != nil {
		log.Printf("Error during parsing userhome: %v", err)
		return
	}
	t.Execute(w, p)
}

func apiHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Method != http.MethodPost {
		req, _ := json.Marshal(ErrorRequestNotRecognized)
		io.WriteString(w, string(req))
		return
	}
	res, _ := ioutil.ReadAll(r.Body)
	apiBody := &ApiBody{}
	if err := json.Unmarshal(res, apiBody); err != nil {
		re, _ := json.Marshal(ErrorRequestBodyParseFailed)
		io.WriteString(w, string(re))
		return
	}
	defer r.Body.Close()
	request(apiBody, w, r)
}

func proxyHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u, _ := url.Parse("http://127.0.0.1:9000/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}
