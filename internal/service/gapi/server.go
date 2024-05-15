package gapi

import (
	"fmt"
	db "simple_bank/db/sqlc"
	"simple_bank/internal/token"
	pb "simple_bank/proto/pb/proto"
	"simple_bank/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	pb.UnimplementedUserServiceServer
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymetriKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
