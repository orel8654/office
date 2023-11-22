package main

import (
	"fmt"
	"log"
	"net"
	"office/internal/config"
	"office/internal/delivery/office"
	"office/internal/repo/postgres"
	"office/pkg/office/pkg/prots"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

func main() {
	if err := run(":8002"); err != nil {
		log.Fatal(err)
	}
}

func run(serv string) error {
	//INIT CONF DB
	configDB, err := config.NewConfig("./config/database.yaml")
	if err != nil {
		return err
	}
	//INIT DB CONN
	s := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		configDB.Username, configDB.Password, configDB.Database, configDB.Host, configDB.Port,
	)
	db, err := sqlx.Connect("postgres", s)
	if err != nil {
		return err
	}
	myRepo := postgres.NewRepo(db)

	//GRPC USERS
	list, err := net.Listen("tcp", serv)
	if err != nil {
		return err
	}
	serverRegistration := grpc.NewServer()
	newOffice := office.NewOffice(myRepo)
	prots.RegisterOfficeServiceServer(serverRegistration, newOffice)

	return serverRegistration.Serve(list)
}
