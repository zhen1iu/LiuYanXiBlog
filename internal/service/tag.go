package service

type CountTagRequst struct {
	Name string `form :"name" binding:"max=100"`
	State uint8 `form :"state,default = 1" binding: "oneof=0 1"
}