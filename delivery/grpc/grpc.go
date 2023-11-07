package grpc

import (
	"github.com/teq-quocbang/store/proto"
	"github.com/teq-quocbang/store/usecase"
)

type TeqService struct {
	proto.UnimplementedTeqServiceServer
	UseCase *usecase.UseCase
}
