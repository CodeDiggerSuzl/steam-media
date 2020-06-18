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

// HomePage template of home page html to replace
type HomePage struct {
	Name string
}

// UserPage to replace in user page files
type UserPage struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request, parm httprouter.Params) {
	userName, err1 := r.Cookie("username")
	sessionID, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		p := &HomePage{Name: "joex"}
		template, err := template.ParseFiles("./templates/home.html")
		if err != nil {
			log.Printf("Error during parseing html files: %v", err)
			return
		}
		// execute the files
		template.Execute(w, p)
		return
	}
	if len(userName.Value) != 0 && len(sessionID.Value) != 0 {
		http.Redirect(w, r, "/userhome", http.StatusFound)
		return
	}
}

func userHomeHandler(w http.ResponseWriter, r *http.Request, parm httprouter.Params) {
	// get username from cookie
	userName, err1 := r.Cookie("username")

	_, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		log.Printf("Error during get user info :%v,%v", err1, err2)
		// relative url
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
	t, err := template.ParseFiles("./templates/userhome.html")
	if err != nil {
		log.Printf("Error during parsing userhome: %v", err)
		return
	}
	t.Execute(w, p)
}

// request to api handler
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

// cross origin resource sharing, using proxy, avoid cross origin
func proxyHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// parsing the final url
	u, _ := url.Parse("http://127.0.0.1:9000/")
	// httputil to generate a new replace the origin url, will not change the headers
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}
