package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RBAC/src/model/user"
)

func (web *Web) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userName, password string
	if r.FormValue("user_name") != "" {
		userName = r.FormValue("user_name")
	}

	if r.FormValue("password") != "" {
		password = r.FormValue("password")
	}

	userData, err := user.Login(userName, password)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("User Signup success. UserId: %d", userData.ID)))
}

func (web *Web) SignupHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	r.Body.Close()
	var userRequest user.User

	if err = json.Unmarshal(body, &userRequest); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	userId, err := userRequest.Signup()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("User Signup success. UserId: %d", userId)))
}
