package repository

import (
	"encoding/json"
	"testing"

	"github.com/glebnaz/postbox/internal/entities"
	"github.com/glebnaz/postbox/internal/logics"
	"github.com/stretchr/testify/assert"
)

func TestUserRepoGet(t *testing.T) {
	//prepare data for test
	assert := assert.New(t)
	sess, data := PrepareTest(UserRepoType)
	var users []entities.User
	err := json.Unmarshal(data, &users)
	assert.NoError(err)
	repo := logics.InitUserRepo(sess)

	//insert all data to collection
	for _, user := range users {
		err = sess.Insert("postbox_users", user)
		assert.NoError(err)
	}

	var ids []string
	for _, user := range users {
		ids = append(ids, user.ID)
	}

	usersInDB, err := repo.Get(ids)
	assert.NoError(err)
	assert.Equal(usersInDB, users)

	err = sess.RemoveWithIDs("postbox_users", ids)
	assert.NoError(err)
}
