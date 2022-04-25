package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	. "../dataloaders"
	. "../models"
)

func Run(){
	http.HandleFunc("/",Handler)
	http.ListenAndServe(":80830,nil")
}

func Handler(w http.ResponseWriter,r *http.Request){
	//page object
	page := Page(ID:7,Name:"Users",Description:"User list",URI:"/users")
	//data loaders
	users := LoadUsers()
	interests := LoadInterest();
	interestMappings := LoadInterestMappings()
	//name
	var newUser []User
	for _ ,user := range users {
		for _,interestMapping := range interestMappings{
			if(user.ID == interestMapping.UserID){
				for _,interest := range interests {
					if interestMappings.InterestID == interest.ID {
						user.Interests = append(user.Interests,interest)
					}
				}
			}
		}
		newUser = append(newUser, user);
		
	}
	viewModel := UserViewModel{Page: page,Users: newUser}
	data,_ := json.Marshal(viewModel)
	w.Write([]byte(data))
}