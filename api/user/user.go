package user

import (
	"errors"
	"net/http"
)

var ErrorNoUserId = errors.New("no user id provided")

func GetUserId(r *http.Request) (int64, error) {
	userIDValue := r.Context().Value("user_id")
	if userIDValue == nil {
		return 0, ErrorNoUserId
	}

	userID, ok := userIDValue.(int64)
	if !ok {
		return 0, ErrorNoUserId
	}
	return userID, nil
}
