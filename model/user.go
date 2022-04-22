package model

type SingUpUserJson struct {
	Id         int    `json:"-"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Image      string `json:"image"`
	City       string `json:"city"`
	Phone      string `json:"phone"`
}
