package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/linkedin"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	clientID     = os.Getenv("ENV_LINKEDIN_CLIENT_ID")
	clientSecret = os.Getenv("ENV_LINKEDIN_CLIENT_SECRET")
	redirectURL  = os.Getenv("ENV_LINKEDIN_REDIRECT_URL")
)

var lnConfig = &oauth2.Config{
	ClientID:     clientID,
	ClientSecret: clientSecret,
	RedirectURL:  redirectURL,
	Scopes:       []string{"r_basicprofile"},
	Endpoint:     linkedin.Endpoint,
}

type user struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Headline  string `json:"headline"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	url := lnConfig.AuthCodeURL("")
	w.Write([]byte("<html><title>Golang Login Linkedin Example</title> <body> <a href='" + url + "&state=DCEeFWf45A53sdfiif424'><button>Login with Linkedin!</button> </a> </body></html>"))
}

func LinkedinLogin(w http.ResponseWriter, r *http.Request) {

	var userData user

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Println(r)
	tok, err := lnConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	// handle err. You need to change this into something more robust
	// such as redirect back to home page with error message
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	fmt.Println(tok)
	response, err := http.Get("https://api.linkedin.com/v1/people/~?format=json&oauth2_access_token=" + tok.AccessToken)

	if err != nil {
		w.Write([]byte(err.Error()))
	}
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

	err = json.Unmarshal(body, &userData) // here!
	fmt.Println(userData)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	id := userData.Id
	firstname := userData.FirstName
	lastname := userData.LastName
	headline := userData.Headline

	w.Write([]byte(fmt.Sprintf("FirstName %s, LastName %s, ID is %s and headline is %s<br>", firstname, lastname, id, headline)))

}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/LinkedinLogin", LinkedinLogin)
	http.ListenAndServe(":8080", mux)

}
