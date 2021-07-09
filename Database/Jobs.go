package database

import "fmt"

func GetJobs() ([]Job, error) {
	db, err := getDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	row, err := db.Query("select * from jobs")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var jobs []Job
	for row.Next() {
		var job Job
		errr := row.Scan(&job.Id, &job.Title, &job.Description, &job.Company_id, &job.User_id)
		if errr != nil {
			return nil, errr
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}
func GetJobsByCandidate(candidate_id int) ([]Job, error) {
	db, err := getDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	qry := fmt.Sprintf("select * from jobs where id not in (select job_id from applications where user_id=%d)", candidate_id)
	row, err := db.Query(qry)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var jobs []Job
	for row.Next() {
		var job Job
		errr := row.Scan(&job.Id, &job.Title, &job.Description, &job.Company_id, &job.User_id)
		if errr != nil {
			return nil, errr
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func GetJobsByRecruiter(recruiter_id int) ([]Job, error) {
	db, err := getDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	qry := fmt.Sprintf("select * from jobs where user_id=%d", recruiter_id)
	row, err := db.Query(qry)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var jobs []Job
	for row.Next() {
		var job Job
		errr := row.Scan(&job.Id, &job.Title, &job.Description, &job.Company_id, &job.User_id)
		if errr != nil {
			return nil, errr
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}
func PostJob(job Job) error {
	Db, err := getDb()
	if err != nil {
		return err
	}
	defer Db.Close()
	query := "insert into jobs(title,description,company_id,user_id) values(?,?,?,?)"
	inst, er := Db.Prepare(query)
	if er != nil {
		return er
	}
	defer inst.Close()
	_, errr := inst.Exec(job.Title, job.Description, job.Company_id, job.User_id)
	if errr != nil {
		return errr
	}
	return nil
}
func ApplyJob(job_id int, user_id int) (int64, error) {
	db, err := getDb()
	if err != nil {
		return -1, err
	}
	defer db.Close()
	qry := fmt.Sprintf("insert into applications(user_id,job_id) values(%d,%d)", user_id, job_id)
	res, err := db.Prepare(qry)
	if err != nil {
		return -1, err
	}
	defer res.Close()
	rr, err := res.Exec()
	if err != nil {
		return -1, err
	}
	val, err := rr.RowsAffected()
	if err != nil {
		return -1, err
	}
	return val, nil
}
func GetJob(job_id int) (Job, error) {
	var job Job
	db, err := getDb()
	if err != nil {
		return job, err
	}
	qry := fmt.Sprintf("select * from jobs where id=%d", job_id)
	fmt.Println(qry)
	res, err := db.Query(qry)
	if err != nil {
		return job, err
	}
	for res.Next() {
		err := res.Scan(&job.Id, &job.Title, &job.Description, &job.Company_id, &job.User_id)
		if err != nil {
			return job, err
		}
	}
	return job, nil
}

func UpdateJob(job *Job) (int64, error) {
	db, err := getDb()
	if err != nil {
		return -1, err
	}
	qry := fmt.Sprintf(`update jobs set title="%s" , description="%s" where id=%d and user_id=%d`, job.Title, job.Description, job.Id, job.User_id)
	fmt.Println(qry)
	res, err := db.Prepare(qry)
	if err != nil {
		return -1, err
	}
	r, err := res.Exec()
	if err != nil {
		return -1, err
	}
	val, err := r.RowsAffected()
	if err != nil {
		return -1, err
	}
	return val, nil
}
func DeleteJob(job_id int, user_id int) (int64, error) {
	db, err := getDb()
	if err != nil {
		return -1, err
	}
	var qry string
	if user_id == -1 {
		fmt.Println("Deleted By ADMIN")
		qry = fmt.Sprintf(`delete from jobs where id=%d`, job_id)
	} else {
		qry = fmt.Sprintf(`delete from jobs where id=%d and user_id=%d`, job_id, user_id)

	}
	// delete applications
	fmt.Println(qry)
	res, err := db.Prepare(qry)
	if err != nil {
		return -1, err
	}
	r, err := res.Exec()
	if err != nil {
		return -1, err
	}
	val, err := r.RowsAffected()
	if err != nil {
		return -1, err
	}
	if val >= 1 {
		qery := fmt.Sprintf(`delete from applications where job_id=%d`, job_id)
		fmt.Println(qery)
		ress, err := db.Prepare(qery)
		if err != nil {
			return -1, err
		}
		_, err = ress.Exec()
		if err != nil {
			return -1, err
		}
	}
	return val, nil
}
