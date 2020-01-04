package service

import (
	"context"
	"github.com/mazti/restless/base/config"
	"github.com/mazti/restless/base/dto"
	"github.com/mazti/restless/base/repository"
	"github.com/mazti/restless/meta/pb"
	"google.golang.org/grpc"
)

type TableService interface {
	Create(ctx context.Context, req dto.CreateTableReq) error
}

func NewTableService(repo repository.BaseRepository, conf config.Config) (TableService, error) {
	conn, err := grpc.Dial(conf.Meta, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &tableService{
		repo:       repo,
		metaClient: pb.NewMetaClient(conn),
	}, nil
}

type tableService struct {
	repo       repository.BaseRepository
	metaClient pb.MetaClient
}

func (h tableService) Create(ctx context.Context, req dto.CreateTableReq) error {
	//resp, err := h.metaClient.Create(ctx, &pb.CreateMetaRequest{Name: req.Name})
	//if err != nil {
	//	return err
	//}
	//if resp == nil {
	//	return errors.New("create base meta fail")
	//}
	err := h.repo.CreateTable(req)
	if err != nil {
		return err
	}
	return err
}
