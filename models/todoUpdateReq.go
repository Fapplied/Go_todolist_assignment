package models

type TodoUpdateReq struct{
	Title string 
	Body string
	DueDate string
	Done bool
}