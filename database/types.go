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

type Moneybox struct {
	ID 					int 			`json:"id"`
	User					int 			`json:"user"`
	Target_sum 			int 			`json:"target_sum"`
	Current_sum 		int 			`json:"current_sum"`
	Auto_adj				bool			`json:"auto_adj"`
	Adj_amount			int 			`json:"adj_amount"`
	Desc 					string		`json:"desc"`
	Name					string		`json:"name"`
}
