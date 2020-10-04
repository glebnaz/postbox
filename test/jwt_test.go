package test

import (
	"github.com/glebnaz/postbox/internal/secure"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strings"
	"testing"
	"time"
)

const userTestValue = "glebnaz"

type tableTestJWT struct {
	user  string
	pass  string
	valid bool
}

//TestJWT testing include function for jwt access
func TestJWT(t *testing.T) {
	assert := assert.New(t)
	tableTests := map[string]tableTestJWT{
		"positive_jwt": {user: userTestValue, pass: generateRandomPass(), valid: true},
		"negative_jwt": {user: userTestValue, pass: generateRandomPass(), valid: false},
	}

	for nameTest, test := range tableTests {
		if test.valid {
			t.Run(nameTest, func(t *testing.T) {
				token := secure.GenerateJWT(test.user, test.pass)
				assert.True(secure.ValidateJWT(token, test.user, test.pass), "Token should bee valid")
			})
		} else {
			t.Run(nameTest, func(t *testing.T) {
				token := secure.GenerateJWT(test.user, test.pass)
				wrongPass := generateRandomPass()
				assert.False(secure.ValidateJWT(token, test.user, wrongPass), "Token should bee invalid")
			})
		}
	}
}

//generateRandomPass return random pass len eight
func generateRandomPass() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < 8; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
