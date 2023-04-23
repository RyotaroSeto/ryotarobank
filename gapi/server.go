package gapi

import (
	"fmt"
	db "ryotarobank/db/sqlc"
	"ryotarobank/pb"
	"ryotarobank/token"
	"ryotarobank/util"
)

type Server struct {
	pb.UnimplementedRyotaroBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
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
