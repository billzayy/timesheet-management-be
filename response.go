package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ResponseData struct {
	Result  any     `json:"result"`
	Success bool    `json:"success"`
	Error   *string `json:"error"`
}

type RespList struct {
	TotalCount int `json:"total_count"`
	Items      any `json:"items"`
}

func GetTokenId(c *gin.Context) uuid.UUID {
	value, ok := c.Get("token")
	var errStr string

	if !ok {
		errStr = ErrTokenNotFound.Error()

		c.JSON(http.StatusBadRequest, ResponseData{
			Result:  nil,
			Success: false,
			Error:   &errStr,
		})
		return uuid.Nil
	}

	var id uuid.UUID

	switch t := value.(type) {
	case string:
		id = uuid.MustParse(t)
	default:
		id = uuid.Nil
	}

	return id
}
