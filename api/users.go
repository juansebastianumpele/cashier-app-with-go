package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
	"path"
	"text/template"
	"time"

	"github.com/google/uuid"
)

func (api *API) Register(w http.ResponseWriter, r *http.Request) {
	// Read username and password request with FormValue.
	creds := model.Credentials{Password: r.FormValue("password"), Username: r.FormValue("username")}
	if creds.Password == "" || creds.Username == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Username or Password empty"})
		return
	}

	// TODO: replace this

	// Handle request if creds is empty send response code 400, and message "Username or Password empty"
	// TODO: answer here

	err := api.usersRepo.AddUser(creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	filepath := path.Join("views", "status.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	var data = map[string]string{"name": creds.Username, "message": "register success!"}
	err = tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
	}
}

func (api *API) Login(w http.ResponseWriter, r *http.Request) {
	// Read usernmae and password request with FormValue.
	creds := model.Credentials{Password: r.FormValue("password"), Username: r.FormValue("username")} // TODO: replace this

	if creds.Password == "" && creds.Username == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Username or Password empty"})
		return
	}
	// Handle request if creds is empty send response code 400, and message "Username or Password empty"
	// TODO: answer here

	err := api.usersRepo.LoginValid(creds)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}
	c := &http.Cookie{}
	c.Name = "session_token"
	c.Value = uuid.New().String()
	c.Expires = time.Now().Add(5 * time.Hour)
	http.SetCookie(w, c)

	// Generate Cookie with Name "session_token", Path "/", Value "uuid generated with github.com/google/uuid", Expires time to 5 Hour.
	// TODO: answer here

	session := model.Session{Token: c.Value, Username: creds.Username, Expiry: c.Expires} // TODO: replace this
	err = api.sessionsRepo.AddSessions(session)

	api.dashboardView(w, r)
}

func (api *API) Logout(w http.ResponseWriter, r *http.Request) {
	//Read session_token and get Value:
	// sessionToken := "" // TODO: replace this
	c, _ := r.Cookie("session_token")
	cookie := c.Value
	if cookie == "" {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "http: named cookie not present"})
		return
	}

	api.sessionsRepo.DeleteSessions(c.Value)

	cookies := &http.Cookie{}
	cookies.Name = "session_token"
	cookies.Value = ""
	cookies.Expires = time.Now().Add(5 * time.Hour)
	http.SetCookie(w, c)

	//Set Cookie name session_token value to empty and set expires time to Now:
	// TODO: answer here

	filepath := path.Join("views", "login.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
	}
}
