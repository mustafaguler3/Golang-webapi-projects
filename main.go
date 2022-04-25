package main

import (
	"fmt"
	"net/http"
	helper "./helpers"
)

func main() {

	uName,email,pwd,pwdConfirm := "","","",""

	mux := http.NewServeMux()

	//signup
	mux.Handle("/signup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		/*for key, value := range r.Form {
			fmt.Printf("%s = %s\n", key, value)
		}*/
		uName = r.FormValue("username")
		email = r.FormValue("email")
		pwd = r.FormValue("password")
		pwdConfirm = r.FormValue("confirm")

		uNameCheck := helper.IsEmpty(uName)
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)
		pwdConfirmCheck := helper.IsEmpty(pwdConfirm)

		if uNameCheck | emailCheck | pwdCheck | pwdConfirmCheck {
			fmt.Fprintf(w,"Error code is -10")
			return
		}

		if pwd == pwdConfirm {
			//save user to database
			fmt.Fprintf(w,"Registration successfull")
		}else {
			fmt.Fprintf(w,"password info must be the same")
		}

	})

	//login
	mux.Handle("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		email = r.FormValue("email")
		pwd = r.FormValue("password")
	
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)

		if pwdCheck | emailCheck {
			fmt.Fprintf(w,"There is empty data.")
			return
		}
		dbPwd := "1234!*."
		dbEmail := "mustafa@hotmail.com"

		if email == dbEmail && pwd == dbPwd {
			fmt.Fprintf(w,"Login successfull")
		}else {
			fmt.Fprintf(w,"Login failed")
		}
	})

	http.ListenAndServe(":8083", mux)
}
