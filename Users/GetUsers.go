package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	Auth "restapi/package/Auth"

	MyDb "restapi/package/Database"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	claim, err := Auth.VerifyJWTAndGetClaim(r)
	if err != nil {
		// if err == http.ErrContentLength || err == jwt.ErrSignatureInvalid {
		// if use http.StatusSeeOther it will redirect to url with same method (LIKE GET OR POST)
		http.Error(w, "No Access", http.StatusUnauthorized)
		return
		// }
		// w.WriteHeader(http.StatusInternalServerError)
		// return
	}
	if claim.Role != "Admin" {
		http.Error(w, "No Access", http.StatusUnauthorized)
		return
	}
	err = r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	fmt.Println("FORM")
	var users []MyDb.User
	for key, val := range r.Form {
		fmt.Println("My key ->", key, "and Values ->", val)
		if key == "type" {
			switch types := val[0]; types {
			case "Candidate":
				users, err = MyDb.GetAllUsersByType("Candidate")
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				json.NewEncoder(w).Encode(users)
				return
			case "Recruiter":
				users, err = MyDb.GetAllUsersByType("Recruiter")
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				json.NewEncoder(w).Encode(users)
				return
			}
		}
		break
	}
	users, err = MyDb.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
