package app

import (
	"LiuYanXiBlog/pkg/errcode"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToresponseList(list interface{}, totalRows int) {

	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error, c *gin.Context) {

	response := gin.H{
		"code": err.Code(),
		"msg":  err.Msg(),
	}

	datails := err.Details()
	if len(datails) > 0 {
		response["details"] = strings.Join(datails, ", ")
	}
	fmt.Println("--------------------------")
	fmt.Println(err.Code(), response)
	c.JSON(err.Code(), map[string]string{"message": "Bad Request"})
	//c.Ctx.JSON(err.Code(), map[string]string{"message": "Bad Request"})
	fmt.Println("00000000000000000")
}
