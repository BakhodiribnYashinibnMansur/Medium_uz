package mediumModel

type SingUpUserJson struct {
	Id         int    `json:"-"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	City       string `json:"city"`
	IsVerified bool   `json:"is_verified"`
	Phone      string `json:"phone"`
	IsStuff    bool   `json:"is_stuff"`
}
