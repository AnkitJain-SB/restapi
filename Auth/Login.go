package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	MyDb "restapi/package/Database"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET LOGIN METHOD STARTED")
	//checking is already login or not
	_, err := VerifyJWTAndGetClaim(r)
	if err == nil {
		fmt.Println("Already Authenticated, USER ID =")
		http.Error(w, "PLEASE LOGOUT FIRST", http.StatusForbidden)
		return
	}
	// decoding credentials
	var credentials Credentials
	error := json.NewDecoder(r.Body).Decode(&credentials)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	//checking credentials in database
	user, exist, err := MyDb.VerfiyCredentials(credentials.Email, credentials.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	if exist {
		token, err := getJWT(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Println("LOGIN SUCCESSFULLY")
		w.Write([]byte(token))
		return
	}
	// if credentials not valid
	fmt.Println("INVALID LOGIN DETAILS")
	w.Header().Set("Content-type", "application/json")
	http.Error(w, "INVALID CREDENTIALS", http.StatusBadRequest)
}

// err := r.ParseForm()
// if err != nil {
// 	log.Fatalln(err)
// }
// fmt.Println("FORM")
// for key, val := range r.Form {
// 	fmt.Println("My key ->", key, "and Values ->", val)
// }
// fmt.Println(r.URL)
