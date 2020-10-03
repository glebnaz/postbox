package secure

import (
	"github.com/brianvoe/sjwt"
)

//GenerateJWT return new jwt as string
//set payload as user
//use secret key as password
func GenerateJWT(user, pass string) string {
	cl := sjwt.New()
	cl.Set("user", user)
	return cl.Generate([]byte(pass))
}

//ValidateJWT check jwt token
//user==user in jwt,true else return false
func ValidateJWT(jwt string, user, pass string) bool {
	verified := sjwt.Verify(jwt, []byte(pass))
	if !verified {
		return false
	}
	cl, err := sjwt.Parse(jwt)
	if err != nil {
		return false
	}
	userFromToken, err := cl.Get("user")
	if userFromToken != user || err != nil {
		return false
	}
	return true
}
