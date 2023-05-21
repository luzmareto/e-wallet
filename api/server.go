package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	dbConn "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/handler"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/service"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/token"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

type Server struct {
	config utils.Config
	// store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config utils.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %v", err)
	}

	server := &Server{
		config: config,
		// store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello I am okay",
		})
	})
	dbconn := dbConn.Connect(server.config)
	svc := service.New(dbconn)

	userHander := handler.NewUserHandler(svc)
	h := handler.New(userHander)

	userr := router.Group("/api/v1/users")
	{
		userr.GET("/", h.UserHandler.List)
		userr.GET("/:id", h.UserHandler.GetByID)
		userr.GET("/username/:username", h.UserHandler.GetByUsername)
		userr.POST("/", h.UserHandler.Register)
		userr.PATCH("/", h.UserHandler.Update)
	}

	server.router = router

}

func (server *Server) Start(addres string) error {
	return server.router.Run(addres)
}
