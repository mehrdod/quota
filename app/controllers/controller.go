package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	_WRONG_PARAMS = 4000

	_INTERNAL_SERVER_ERR = 5000
)

type errorCode struct {
	HttpCode int    `json:"http_code"`
	msg      string `json:"msg"`
}

var errorCodes map[int]errorCode

func init() {
	errorCodes = make(map[int]errorCode)

	// 4xxx
	errorCodes[_WRONG_PARAMS] = errorCode{http.StatusBadRequest, "неверные параметры"}

	// 5xxx
	errorCodes[_INTERNAL_SERVER_ERR] = errorCode{http.StatusInternalServerError, "Внутреняя ошибка сервера"}
}

func replyError(c *gin.Context, code int) {
	c.AbortWithStatusJSON(errorCodes[code].HttpCode, gin.H{"msg": errorCodes[code].msg})
}
