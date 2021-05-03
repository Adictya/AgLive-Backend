package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/adictya/AgoraLive-backend/api"
	db "github.com/adictya/AgoraLive-backend/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbdriver = "postgres"
	dbsource = "postgresql://root:secret@localhost:5432/root?sslmode=disable"
	serverAddress = "0.0.0.0:8080"

	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "secret"
	dbname   = "root"
 )

func main(){

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

		log.Print(psqlInfo)

	// conn, err := sql.Open("postgres", psqlInfo)

	conn ,err := sql.Open(dbdriver,dbsource)

	if err != nil {
		log.Fatal("Cannot connect to db :", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil{
		log.Fatal("Failed to start server : ",err)
	}


}
