package web

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/smilecs/yellowpages/config"
	"github.com/smilecs/yellowpages/models"
)

//
// func SocialLogin(w http.ResponseWriter, r *http.Request) {
// 	var user models.User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		log.Println(err)
// 	}
//
// 	err = user.Add(config.Get())
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	data := struct {
// 		Message string
// 	}{"login success"}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(data)
// }

//
// func AdminLoginOld(w http.ResponseWriter, r *http.Request) {
// 	var user models.AdminUser
// 	var result Admin
//
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		log.Println(err)
// 	}
//
// 	log.Println(user)
// 	collection := config.Get().Database.C("admin").With(config.Get().Database.Session.Copy())
//
// 	log.Println(user.Username)
// 	err = collection.Find(bson.M{"username": user.Username, "password": user.Password}).One(&result)
// 	if err != nil {
// 		log.Println(err)
// 	}
//
// 	data, _ := json.Marshal(result)
// 	w.Write(data)
// }

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	var providedUserDetails models.AdminUser
	var existingUserDetails models.AdminUser

	err := json.NewDecoder(r.Body).Decode(&providedUserDetails)
	if err != nil {
		log.Println(err)
	}

	log.Println(providedUserDetails)
	// collection := config.Get().Database.C(config.ADMINSCOLLECTION).With(config.Get().Database.Session.Copy())

	existingUserDetails, err = providedUserDetails.Get(config.Get())
	// err = collection.Find(bson.M{"username": providedUserDetails.Username}).One(&existingUserDetails)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		msg := make(map[string]string)
		msg["error"] = "User not found"
		json.NewEncoder(w).Encode(msg)

	}

	err = bcrypt.CompareHashAndPassword(existingUserDetails.PasswordHash, []byte(providedUserDetails.Password))
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		msg := make(map[string]string)
		msg["error"] = "Incorrect Password"
		json.NewEncoder(w).Encode(msg)
	}

	jwtResponse, err := GenerateAdminJWT(existingUserDetails)
	//log.Println(jwtResponse)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		msg := make(map[string]string)
		msg["error"] = "There was an issue generating your token. Please try again"
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(jwtResponse)
	if err != nil {
		log.Println(err)
	}

}

func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var providedUserDetails models.AdminUser

	err := json.NewDecoder(r.Body).Decode(&providedUserDetails)
	if err != nil {
		log.Println(err)
	}

	log.Println(providedUserDetails)

	w.Header().Set("Content-Type", "application/json")
	msg := make(map[string]string)
	err = providedUserDetails.Add(config.Get())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		msg["error"] = "Unable to create Admin at this time"
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusOK)
	msg["message"] = "Success Creating Admin"
	err = json.NewEncoder(w).Encode(msg)
	if err != nil {
		log.Println(err)
	}

}
