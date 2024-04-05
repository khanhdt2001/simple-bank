package api

import (
	"fmt"
	db "simple_bank/db/sqlc"
	"simple_bank/internal/token"
	"simple_bank/internal/validation"
	"simple_bank/util"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{store: store}

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymetriKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validation.ValidCurrency)
	}
	server.setUpRouter()
	server.tokenMaker = tokenMaker
	server.config = config
	return server, nil
}

func (server *Server) setUpRouter() {
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
