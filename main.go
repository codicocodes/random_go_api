package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/codicocodes/go-api-test/api_utils"
	"github.com/codicocodes/go-api-test/db_utils"
	"github.com/codicocodes/go-api-test/users"
)

func main() {
	var db = db_utils.Connect()
	var app = App{db: db}
	http.HandleFunc("/", app.ListUsers)
	app.Run()
}

type App struct {
	db *sql.DB
}

func (a App) Run() {
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (a App) ListUsers(res http.ResponseWriter, req *http.Request) {
	s := users.GetUserService(a.db)
	users := s.List()
	usersJson, err := json.Marshal(&users)
	if err != nil {
		api_utils.ReturnJsonResponse(res, http.StatusInternalServerError, []byte(err.Error()))
		return
	}
	api_utils.ReturnJsonResponse(res, http.StatusOK, usersJson)
}
