package core

import (
	"log"
	"producer/src/reservations/application"
	"producer/src/reservations/infrastructure"
	"producer/src/core/middlewares"
	"github.com/gin-gonic/gin"
)

func IniciarRutas() {
    mysqlConn, err := GetDBPool()
    if err != nil {
        log.Fatalf("Error al obtener la conexión a la base de datos: %v", err)
    }

	rabbitmqCh, err := GetChannel()
	if err != nil {
        log.Fatalf("Error al obtener la conexión a la base de datos: %v", err)
    }

    mysqlRepository := infrastructure.NewMysqlRepository(mysqlConn.DB)
	rabbitqmRepository := infrastructure.NewRabbitRepository(rabbitmqCh.ch)

	createReservationUseCase := application.NewCreateReservationUseCase(rabbitqmRepository, mysqlRepository)
	createReservationController := infrastructure.NewCreateReservationController(createReservationUseCase)

	router := gin.Default()
	middleware := middlewares.NewCorsMiddleware()	
	router.Use(middleware)

	router.POST("/reservation", createReservationController.Execute)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("error al iniciar el servidor: %v", err)
	}
}