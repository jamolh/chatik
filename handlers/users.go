package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jamolh/chatik/db"
	"github.com/jamolh/chatik/models"
	"github.com/julienschmidt/httprouter"
)

func Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		response = models.Response{
			Code: 200,
		}
		user models.User
		err  error
	)
	defer response.Send(w, r)

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Code = http.StatusBadRequest
		response.Send(w, r)
		return
	}
	if !db.CheckUserExists(r.Context(), user.Username) {
		response.Code = 404
		response.Message = "user by this username not exists"
		return
	}

	dbUser, err := db.GetUser(r.Context(), user.Username)
	if err != nil {
		response.Code = 500
		log.Printf("handlers:Register db.GetUser by username:%v error:%v\n", user.Username, err)
		return
	}
	if user.Password != dbUser.Password {
		response.Code = 400
		response.Message = "wrong user password"
		return
	}
	dbUser.Password = ""
	response.Payload = dbUser
}
