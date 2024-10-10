package controllers

import (
	driver "backendgo/Driver"
	structmodels "backendgo/StructModels"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user structmodels.NewUser

	//Recogemos el cuerpo de la petición
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil { //Si hay un error en el decode
		http.Error(w, "Invalid payload", 400) //Si falla se devuelve un 400 bad request
	}

	//Comprobamos que no existe el usuario

	if driver.UserExists(user.Username) {
		http.Error(w, "User already exists", 409)
		return
	}

	//llamamos a la base de datos
	code := driver.CreateUser(user.Username, user.Email, user.Password)
	if code == 1 {
		http.Error(w, "Couldnt create user, read LOG", 404)
		return
	}

}
