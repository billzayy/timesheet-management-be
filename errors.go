package backend

import "errors"

/* Not found */
var ErrUserNotFound = errors.New("user not found")
var ErrPositionNotFound = errors.New("position not found")
var ErrLevelNotFound = errors.New("level not found")
var ErrUserTypeNotFound = errors.New("user type not found")
var ErrTokenNotFound = errors.New("token not found")
var ErrRoleNotFound = errors.New("role not found")

/* Empty Value */
var ErrEmptyName = errors.New("empty name")
