package api

import (
	db "github.com/adictya/AgoraLive-backend/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type Server struct{
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server{

	server := &Server{store:store }
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/createStream",server.createStream)
	router.GET("/checkHealth",server.checkHealth)
	router.GET("/listStreams",server.listStreams)
	router.GET("/getThumbnail",server.getThumbnail)
	router.DELETE("/deleteStream",server.deleteStream)

	server.router = router
	return server
}

func (server *Server) Start(address string)error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H{
	return gin.H{"Error": err.Error()}
}
