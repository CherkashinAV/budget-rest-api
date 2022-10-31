package database

type User struct {
	ID 					int			`json:"id"`
	Name 					string		`json:"name"`
	Surname 				string		`json:"surname"`
	Available_funds 	int			`json:"available_funds"`
	Lang 					string		`json:"lang"`
	Theme 				string		`json:"theme"`
}	


type Transaction struct {
	ID						int        	`json:"id"`
	User 					int 			`json:"user"`
	Date					string		`json:"date"`
	Ante					int			`json:"ante"`
	Desc					string 		`json:"desc"`
}