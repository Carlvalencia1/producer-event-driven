package domain

type IReservationRabbitqm interface {
	Save(resevation *Reservation) error
}
