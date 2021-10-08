package app

import (
	"net/http"
	"os"
	"strconv"

	"github.com/abdullahkaraman/go-todo-api/app/models"
	"github.com/abdullahkaraman/go-todo-api/app/utils"
	"github.com/abdullahkaraman/go-todo-api/config"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	router *mux.Router
	db     models.Datastore
}

func (app *App) Start(conf *config.Config) {
	db, err := models.InitDB(conf.DBConfig)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	app.db = db
	app.router = mux.NewRouter()
	app.initRouters()
	app.run(":8000")
}

func (app *App) initRouters() {
	app.router.HandleFunc("/", app.status).Methods("Get")
	app.router.HandleFunc("/todo", app.listTodos).Methods("Get")
	app.router.HandleFunc("/todo/{id:[0-9]+}", app.getTodo).Methods("Get")
}

func (app *App) run(addr string) {
	loggedRouter := handlers.LoggingHandler(os.Stdout, app.router)
	http.ListenAndServe(addr, loggedRouter)
}

func (app *App) listTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := app.db.AllTodos()
	if err != nil {
		utils.ServerError(w)
		return
	}

	utils.RespondJson(w, http.StatusOK, todos)
}

func (app *App) getTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.BadRequest(w, "ID must be an int")
	}

	todo, err := app.db.GetTodo(id)
	if err != nil {
		utils.ServerError(w)
		return
	}

	utils.RespondJson(w, http.StatusOK, todo)
}

func (app *App) status(w http.ResponseWriter, r *http.Request) {
	utils.RespondJson(w, http.StatusOK, "API is up and working!")
}
