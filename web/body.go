package web

type LoginCredentials struct {
	Id string `form:"id" json:"id" binding:"required"`
	Pw string `form:"pw" json:"pw" binding:"required"`
}

type RegisterCredentials struct {
	Id   string `form:"id" json:"id" binding:"required"`
	Pw   string `form:"pw" json:"pw" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
}
