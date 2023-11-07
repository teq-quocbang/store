package grpc

import (
	"context"
	"encoding/json"

	"git.teqnological.asia/teq-go/teq-pkg/teqerror"

	"github.com/teq-quocbang/store/proto"
	"github.com/teq-quocbang/store/usecase/grpc"
	"github.com/teq-quocbang/store/util/myerror"
)

func (a *TeqService) GetExampleByID(ctx context.Context, req *proto.GetByIDRequest) (*proto.ExampleResponse, error) {
	myExample, err := a.UseCase.GRPC.GetByID(ctx, &grpc.GetByIDRequest{ID: req.GetId()})
	if err != nil {
		return nil, teqerror.ErrGRPC(err)
	}

	resp := &proto.ExampleResponse{}

	b, err := json.Marshal(myExample)
	if err != nil {
		return nil, myerror.ErrJSONMarshal(err)
	}

	err = json.Unmarshal(b, resp)
	if err != nil {
		return nil, myerror.ErrJSONUnmarshal(err)
	}

	return resp, nil
}
