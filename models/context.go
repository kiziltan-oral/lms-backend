package models

import "github.com/gin-gonic/gin"

type Context struct {
	Token string `json:"t"`
}

func NewContext(c *gin.Context) *Context {
	return &Context{
		Token: c.GetString("usertoken"),
	}
}
