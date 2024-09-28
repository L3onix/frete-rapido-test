package model

type Address struct {
	Zipcode string `json:"zipcode" binding:"required"`
}
