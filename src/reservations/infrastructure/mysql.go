package infrastructure

import (
	"producer/src/reservations/domain"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) *MysqlRepository {
	return &MysqlRepository{db: db}
}

func (repo *MysqlRepository) Save(reservation *domain.Reservation) error {

	query := "INSERT INTO reservations (name, description, price, userName, userCellphone, dateReservation) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := repo.db.Exec(query, reservation.Name, reservation.Description, reservation.Price, reservation.UserName, reservation.UserCellphone, reservation.DateReservation)
	if err != nil {
		return fmt.Errorf("error al guardar orden: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener el ID insertado: %v", err)
	}
	reservation.Id = int32(id)

	return nil
}