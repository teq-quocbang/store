package grpc

import (
	"context"
	"encoding/json"

	"git.teqnological.asia/teq-go/teq-pkg/teqerror"

	"github.com/teq-quocbang/course-register/usecase/grpc"
	"github.com/teq-quocbang/course-register/util/myerror"
	"github.com/teq-quocbang/store/proto"
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
