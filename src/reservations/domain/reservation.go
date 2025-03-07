package domain

type Reservation struct {
	Id int32			`json:"id"`
	Name  string  		`json:"name"`
	Description string 	`json:"description"`
	Price int32 		`json:"price"`
	UserName string		`json:"userName"`
	UserCellphone string`json:"cellPhone"`
	DateReservation string `json:"dateReservation"`
}

func NewReservation(name string, description string, price int32, userName string, userCellphone string, dateReservation string) *Reservation {
	return &Reservation{Name: name, Description: description, Price: price, UserName: userName, UserCellphone: userCellphone, DateReservation: dateReservation}
}