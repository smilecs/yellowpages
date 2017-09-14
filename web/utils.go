package web

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/smilecs/yellowpages/config"
	"github.com/smilecs/yellowpages/models"
)

func Setup(w http.ResponseWriter, r *http.Request) {
	adminUser := models.AdminUser{
		Name:     "Admin",
		Username: "admin",
		Password: "p@ssw0rd",
		Email:    "contact@calabarpages.com",
		Type:     "superadmin",
	}
	err := adminUser.Add(config.Get())
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte("success adding admin user"))
}

//
// //Userget reads the json web token(JWT) content from context and marshals it ito a user struct,
// func Userget(r *http.Request) (models.User, error) {
// 	u := r.Context().Value("User")
//
// 	user := models.User{}
// 	if u != nil {
// 		err := mapstructure.Decode(u, &user)
//
// 		if err != nil {
// 			return user, err
// 		}
// 		return user, nil
// 	}
// 	return user, nil
// }

//Turn user details into a hasked token that can be used to recognize the user in the future.
// func GenerateJWT(user models.User) (map[string]interface{}, error) {
// 	claims := jwt.MapClaims{}
//
// 	msg := make(map[string]interface{})
//
// 	// set our claims
// 	claims["User"] = user
// 	claims["Name"] = user.Name
//
// 	// set the expire time
//
// 	claims["exp"] = time.Now().Add(time.Hour * 24 * 30 * 12).Unix() //24 hours inn a day, in 30 days * 12 months = 1 year in milliseconds
//
// 	// create a signer for rsa 256
// 	t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
//
// 	pub, err := jwt.ParseRSAPrivateKeyFromPEM(config.Get().Encryption.Private)
// 	if err != nil {
// 		return msg, err
// 	}
// 	tokenString, err := t.SignedString(pub)
//
// 	if err != nil {
// 		return msg, err
// 	}
//
// 	// msg:= make(map[string]interface{})
// 	msg["User"] = user
// 	msg["Message"] = "Token successfully generated"
// 	msg["Token"] = tokenString
//
// 	return msg, nil
//
// }

//Turn user details into a hasked token that can be used to recognize the user in the future.
func GenerateAdminJWT(user models.AdminUser) (map[string]interface{}, error) {
	msg := make(map[string]interface{})

	claims := jwt.MapClaims{}
	// set our claims
	claims["User"] = user
	claims["Name"] = user.Name

	// set the expire time

	claims["exp"] = time.Now().Add(time.Hour * 24 * 30 * 12).Unix() //24 hours inn a day, in 30 days * 12 months = 1 year in milliseconds

	// create a signer for rsa 256
	t := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)

	pub, err := jwt.ParseRSAPrivateKeyFromPEM(config.Get().Encryption.Private)
	if err != nil {
		return msg, err
	}
	tokenString, err := t.SignedString(pub)

	if err != nil {
		return msg, err
	}

	// msg:= make(map[string]interface{})
	msg["User"] = user
	msg["Message"] = "Token successfully generated"
	msg["Token"] = tokenString

	return msg, nil

}
