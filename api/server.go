package api

import (
	"fmt"

	db "github.com/adictya/AgoraLive-backend/db/sqlc"
	"github.com/adictya/AgoraLive-backend/token"
	"github.com/adictya/AgoraLive-backend/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct{
	config util.Config
	store *db.Store
	tokenMaker token.Maker
	router *gin.Engine
}

func NewServer(config util.Config ,store *db.Store) (*Server, error){

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err!= nil{
		return nil, fmt.Errorf("cannot create token maker: &w",err)
	}



	server := &Server{
		config: config,
		store:store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/createStream",server.createStream)
	router.GET("/checkHealth",server.checkHealth)
	router.GET("/listStreams",server.listStreams)
	router.GET("/getThumbnail",server.getThumbnail)
	router.DELETE("/deleteStream",server.deleteStream)
	router.POST("/createUser",server.createUser)

	server.router = router
	return server,nil
}

func (server *Server) Start(address string)error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H{
	return gin.H{"Error": err.Error()}
}
