package main

import (
	"database/sql"
	"log"

	"github.com/adictya/AgoraLive-backend/api"
	db "github.com/adictya/AgoraLive-backend/db/sqlc"
	"github.com/adictya/AgoraLive-backend/util"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "secret"
	dbname   = "root"
 )

func main(){

	config, err := util.LoadConfig(".")
	if err!= nil{
		log.Fatal("cannot load config:", err)
	}

	conn ,err := sql.Open(config.DBDriver,config.DBSource)

	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	store := db.NewStore(conn)
	server,err := api.NewServer(config,store)

	err = server.Start(config.ServerAddress)

	if err != nil{
		log.Fatal("Failed to start server : ",err)
	}


}
