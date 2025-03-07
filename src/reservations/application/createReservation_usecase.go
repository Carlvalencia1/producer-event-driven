package application

import (
	"producer/src/reservations/domain"
	"log"
)

type CreateReservationUseCase struct {
	rabbitqmRepository domain.IReservationRabbitqm
	mysqlRepository domain.IReservationMysq
}

func NewCreateReservationUseCase(rabbitqmRepository domain.IReservationRabbitqm, mysqlRepository domain.IReservationMysq) *CreateReservationUseCase {
	return &CreateReservationUseCase{rabbitqmRepository: rabbitqmRepository, mysqlRepository: mysqlRepository}
}

func (usecase *CreateReservationUseCase) SetOrder(mysqlRepository domain.IReservationMysq, rabbitqmRepository domain.IReservationRabbitqm) {
	usecase.mysqlRepository = mysqlRepository
	usecase.rabbitqmRepository = rabbitqmRepository
}

func (usecase *CreateReservationUseCase) Run(reservation *domain.Reservation) error {
	err := usecase.mysqlRepository.Save(reservation)
	errSendMessage := usecase.rabbitqmRepository.Save(reservation)

	if err != nil || errSendMessage != nil {
		log.Panicf("error to send message %s", err); 
		return err
	}
	
	return errSendMessage; 
}
