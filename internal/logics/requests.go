package logics

import "github.com/glebnaz/postbox/internal/entities"

//UserReq request for users endpoints
type UserReq struct {
	IDs   []string        `json:"ids,omitempty"`
	Users []entities.User `json:"users"`
}

//UserResp response for users endpoints
type UserResp struct {
	Status string          `json:"status"`
	Users  []entities.User `json:"users,omitempty"`
	Error  string          `json:"error,omitempty"`
}
