package v1

import (
	"LiuYanXiBlog/pkg/app"
	"LiuYanXiBlog/pkg/errcode"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Tag struct{}

type Params struct {
	Name  string `form:"name" binding:"max=100"`
	State uint32 `form:"state,default=1" binding:"oneof=0 1"`
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}

func (t Tag) List(c *gin.Context) {

	var param Params
	response := app.NewResponse(c)
	valid, errs := app.BindAndVaild(c, &param)
	var errall []string
	for _, err := range errs {
		errall = append(errall, err.Error())
	}

	fmt.Println("111111111111", errall)
	if !valid {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	response.ToResponse(gin.H{})

	return
}

func (t Tag) Create(c *gin.Context) {}

func (t Tag) Update(c *gin.Context) {}

func (t Tag) Delete(c *gin.Context) {}
