package domain

type IReservationMysq interface {
	Save(resevation *Reservation) error
}
