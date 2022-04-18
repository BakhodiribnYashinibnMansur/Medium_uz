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

// type SignUpUserData struct {
// 	Id         int
// 	Email      string
// 	FirstName  string
// 	SecondName string
// 	City       string
// 	IsVerified bool
// 	Phone      string
// 	Ratings    string
// 	PostReads  int
// 	Posts      int
// 	IsStuff    bool
// }
