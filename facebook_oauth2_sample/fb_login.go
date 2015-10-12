package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	clientID     = os.Getenv("ENV_FB_CLIENT_ID")
	clientSecret = os.Getenv("ENV_FB_CLIENT_SECRET")
	redirectURL  = os.Getenv("ENV_FB_REDIRECT_URL")
)

var fbConfig = &oauth2.Config{
	ClientID:     clientID,
	ClientSecret: clientSecret,
	RedirectURL:  redirectURL,
	Scopes:       []string{"email", "user_birthday", "user_location", "user_about_me"},
	Endpoint:     facebook.Endpoint,
}

type user struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Username string `json:"name"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	url := fbConfig.AuthCodeURL("")
	w.Write([]byte("<html><title>Golang Login Facebook Example</title> <body> <a href='" + url + "'><button>Login with Facebook!</button> </a> </body></html>"))
}

func FBLogin(w http.ResponseWriter, r *http.Request) {

	var userData user

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tok, err := fbConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	// handle err. You need to change this into something more robust
	// such as redirect back to home page with error message
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	response, err := http.Get("https://graph.facebook.com/me?access_token=" + tok.AccessToken)

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
	bday := userData.Birthday
	fbusername := userData.Username
	email := userData.Email

	w.Write([]byte(fmt.Sprintf("Username %s ID is %s and birthday is %s and email is %s<br>", fbusername, id, bday, email)))

	img := "https://graph.facebook.com/" + id + "/picture?width=180&height=180"

	w.Write([]byte("Photo is located at " + img + "<br>"))
	// see https://www.socketloop.com/tutorials/golang-download-file-example on how to save FB file to disk

	w.Write([]byte("<img src='" + img + "'>"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/FBLogin", FBLogin)
	http.ListenAndServe(":8080", mux)

}
