package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	dbConn "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/handler"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/middleware"
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
	tokenMaker, err := token.NewJWTMaker(config.TokenSymetricKey)
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

	dbconn := dbConn.Connect(server.config)
	svc := service.New(dbconn)

	// initialte handler
	userHander := handler.NewUserHandler(svc)
	authHander := handler.NewAUthHandler(server.config, server.tokenMaker, svc)
	walletHandler := handler.NewWalletHandler(svc)
	storeHandler := handler.NewStoreHandler(svc)

	// initiaate main handler
	h := handler.New(
		authHander,
		userHander,
		walletHandler,
		storeHandler,
	)

	// init router
	router := gin.Default()
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello I am okay",
		})
	})

	router.POST("/users", h.UserHandler.Register)
	router.POST("/users/login", h.AuthHandler.LoginUser)
	router.POST("/token/renew", h.AuthHandler.RenewAccessToken)

	// user router
	user := router.Group("/api/v1/users", middleware.AuthMiddleware(server.tokenMaker))
	{
		user.GET("/", h.UserHandler.List)
		user.GET("/:id", h.UserHandler.GetByID)
		user.GET("/detail", h.UserHandler.GetByUsername)
		user.PATCH("/", h.UserHandler.Update)
	}

	// wallet route
	wallet := router.Group("/api/v1/wallets", middleware.AuthMiddleware(server.tokenMaker))
	{
		wallet.POST("/:id", h.WalletHandler.AddWalletBalance)
		wallet.POST("/", h.WalletHandler.CreateWallets)
		wallet.POST("/withdrawal", h.StoreHandler.WithdrawalTransactions)
		wallet.POST("/topups", h.StoreHandler.TopupTransactions)
		wallet.POST("/transfer", h.StoreHandler.TopupTransactions)
	}

	server.router = router

}

func (server *Server) Start(addres string) error {
	return server.router.Run(addres)
}
