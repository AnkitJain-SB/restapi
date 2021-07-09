package jobs

import (
	"encoding/json"
	"fmt"
	"net/http"
	Auth "restapi/package/Auth"
	MyDb "restapi/package/Database"
)

func GetJobs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("######### GET JOBS STARTED ###########")
	claim, err := Auth.VerifyJWTAndGetClaim(r)
	if err != nil {
		http.Error(w, "No Access", http.StatusUnauthorized)
		return
	}
	var jobs []MyDb.Job
	if claim.Role == "Candidate" {
		fmt.Println("In Candidate")
		jobs, err = MyDb.GetJobsByCandidate(claim.User_ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if claim.Role == "Recruiter" {
		fmt.Println("In Recruiter")
		jobs, err = MyDb.GetJobsByRecruiter(claim.User_ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if claim.Role == "Admin" {
		fmt.Println("In Admin")
		jobs, err = MyDb.GetJobs()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.Header().Add("Content-type", "application-json")
	err = json.NewEncoder(w).Encode(jobs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
