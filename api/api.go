package api

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type API struct {
	usersRepo    repo.UserRepository
	sessionsRepo repo.SessionsRepository
	products     repo.ProductRepository
	cartsRepo    repo.CartRepository
	mux          *http.ServeMux
}

type Page struct {
	File string
}

func (p Page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filepath := path.Join("views", p.File)
	path, err := template.ParseFiles(filepath)

	if err != nil {
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}
	err = path.Execute(w, p)
	if err != nil {
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	// TODO: answer here
}

func NewAPI(usersRepo repo.UserRepository, sessionsRepo repo.SessionsRepository, products repo.ProductRepository, cartsRepo repo.CartRepository) API {
	mux := http.NewServeMux()
	api := API{
		usersRepo,
		sessionsRepo,
		products,
		cartsRepo,
		mux,
	}

	index := Page{File: "index.html"}
	mux.Handle("/", api.Get(index))

	// Please create routing for:
	// - Register page with endpoint `/page/register`, GET method and render `register.html` file on views folder
	// - Login page with endpoint `/page/login`, GET method and render `login.html` file on views folder
	register := Page{File: "register.html"}
	mux.Handle("/page/register", api.Get(register))

	login := Page{File: "login.html"}
	mux.Handle("/page/login", api.Get(login))
	// TODO: answer here

	mux.Handle("/user/register", api.Post(http.HandlerFunc(api.Register)))
	mux.Handle("/user/login", api.Post(http.HandlerFunc(api.Login)))
	mux.Handle("/user/logout", api.Get(api.Auth(http.HandlerFunc(api.Logout))))

	mux.Handle("/user/img/profile", api.Get(api.Auth(http.HandlerFunc(api.ImgProfileView))))
	mux.Handle("/user/img/update-profile", api.Post(api.Auth(http.HandlerFunc(api.ImgProfileUpdate))))

	// Please create routing for endpoint `/cart/add` with POST method, Authentication and handle api.AddCart
	mux.Handle("/cart/add", api.Post((api.Auth(http.HandlerFunc(api.AddCart)))))
	// TODO: answer here

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}

// package api

// import (
// 	repo "a21hc3NpZ25tZW50/repository"
// 	"fmt"
// 	"html/template"
// 	"net/http"
// 	"path"
// )

// type API struct {
// 	usersRepo    repo.UserRepository
// 	sessionsRepo repo.SessionsRepository
// 	products     repo.ProductRepository
// 	cartsRepo    repo.CartRepository
// 	mux          *http.ServeMux
// }

// type Page struct {
// 	File string
// }

// func (p Page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	filepath := path.Join("views", p.File)
// 	// TODO: answer here
// 	template, err := template.ParseFiles(filepath)
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	template.Execute(w, filepath)

// }

// func NewAPI(usersRepo repo.UserRepository, sessionsRepo repo.SessionsRepository, products repo.ProductRepository, cartsRepo repo.CartRepository) API {
// 	mux := http.NewServeMux()
// 	api := API{
// 		usersRepo,
// 		sessionsRepo,
// 		products,
// 		cartsRepo,
// 		mux,
// 	}

// 	index := Page{File: "index.html"}
// 	mux.Handle("/", api.Get(index))

// 	// Please create routing for:
// 	// - Register page with endpoint `/page/register`, GET method and render `register.html` file on views folder
// 	// - Login page with endpoint `/page/login`, GET method and render `login.html` file on views folder
// 	// TODO: answer here

// 	register := Page{File: "register.html"}
// 	mux.Handle("/page/register", api.Get(register))

// 	login := Page{File: "login.html"}
// 	mux.Handle("/page/login", api.Get(login))

// 	mux.Handle("/user/register", api.Post(http.HandlerFunc(api.Register)))
// 	mux.Handle("/user/login", api.Post(http.HandlerFunc(api.Login)))
// 	mux.Handle("/user/logout", api.Get(api.Auth(http.HandlerFunc(api.Logout))))

// 	mux.Handle("/user/img/profile", api.Get(api.Auth(http.HandlerFunc(api.ImgProfileView))))
// 	mux.Handle("/user/img/update-profile", api.Post(api.Auth(http.HandlerFunc(api.ImgProfileUpdate))))

// 	// Please create routing for endpoint `/cart/add` with POST method, Authentication and handle api.AddCart
// 	// TODO: answer here
// 	mux.Handle("/cart/add", api.Post((api.Auth(http.HandlerFunc(api.AddCart)))))

// 	return api
// }

// func (api *API) Handler() *http.ServeMux {
// 	return api.mux
// }

// func (api *API) Start() {
// 	fmt.Println("starting web server at http://localhost:8080")
// 	http.ListenAndServe(":8080", api.Handler())
// }
