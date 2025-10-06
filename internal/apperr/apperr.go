package apperr

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrGetLeaderboard     = errors.New("error get leaderboard")
	ErrInvalidIDFormat    = errors.New("invalid id format")
	ErrForbidden          = errors.New("forbidden")
	ErrGetStatus          = errors.New("error get user status")
	ErrReferrerNotFound   = errors.New("referrer not found")
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrTaskNotFound       = errors.New("task not found")
)
