package api

import (
	"net/http"
	"strconv"

	db "github.com/adictya/AgoraLive-backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createStreamRequest struct {
	Channel   string `json:"Channel" binding:"required"`
	Thumbnail string `json:"Thumbnail" binding:"required"`
}

func (server *Server) createStream(ctx *gin.Context) {
	var req createStreamRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateStreamParams{
		Channel:   req.Channel,
		Thumbnail: req.Thumbnail,
	}

	stream, err := server.store.CreateStream(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, stream)
}

func (server *Server) checkHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "The server is running")
}

func (server *Server) listStreams(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("Limit"))
	limit32 := int32(limit)

	offset, _ := strconv.Atoi(ctx.Query("Offset"))
	offset32 := int32(offset)

	arg := db.ListStreamsParams{
		Limit:  limit32,
		Offset: offset32 - 1,
	}

	streams, err := server.store.ListStreams(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, streams)
}

func (server *Server) getThumbnail(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))

	var id32 int32 = int32(id)

	streams, err := server.store.GetThumbnail(ctx, id32)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, streams)

}

type deleteStreamRequest struct {
	Channel string `json:"Channel" binding:"required"`
}

func (server *Server) deleteStream(ctx *gin.Context) {

	var req deleteStreamRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	err := server.store.DeleteStream(ctx, req.Channel)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusNoContent)

}
