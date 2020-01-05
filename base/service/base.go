package service

import (
	"context"
	"github.com/mazti/restless/base/dto"
	"github.com/mazti/restless/base/repository"
	sharedDto "github.com/mazti/restless/shared/dto"
)

type BaseService interface {
	Create(ctx context.Context, req dto.CreateBaseReq) (dto.BaseResp, error)
	Get(ctx context.Context, id string) (dto.BaseResp, error)
	List(ctx context.Context, offset int, limit int) (dto.ListBaseResp, error)
	Update(ctx context.Context, req dto.UpdateBaseReq) (dto.BaseResp, error)
	Delete(ctx context.Context, id string) error
}

func NewBaseService(baseRepo repository.BaseRepository, metaRepo repository.MetaRepository) (BaseService, error) {
	return &baseService{
		baseRepo: baseRepo,
		metaRepo: metaRepo,
	}, nil
}

type baseService struct {
	baseRepo repository.BaseRepository
	metaRepo repository.MetaRepository
}

func (h baseService) Create(ctx context.Context, req dto.CreateBaseReq) (base dto.BaseResp, err error) {
	resp, err := h.metaRepo.Create(ctx, req.ToMeta())
	if err != nil {
		return base, err
	}
	err = h.baseRepo.CreateSchema(resp.Schema)
	if err != nil {
		return base, err
	}
	return dto.NewBaseResp(resp, EncodeID)
}

func (h baseService) Get(ctx context.Context, id string) (resp dto.BaseResp, err error) {
	metaID, err := DecodeID(id)
	if err != nil {
		return resp, err
	}
	meta, err := h.metaRepo.Get(ctx, metaID)
	if err != nil {
		return resp, err
	}
	return dto.NewBaseResp(meta, EncodeID)
}

func (h baseService) List(ctx context.Context, offset int, limit int) (resp dto.ListBaseResp, err error) {
	items, err := h.metaRepo.List(ctx, offset, limit)
	if err != nil {
		return resp, err
	}
	count := len(items)
	resp = dto.ListBaseResp{Results: make([]dto.BaseResp, count)}
	for i, item := range items {
		baseResp, err := dto.NewBaseResp(item, EncodeID)
		if err != nil {
			return resp, err
		}
		resp.Results[i] = baseResp
	}

	total, _ := h.metaRepo.Count(ctx)
	resp.Metadata = sharedDto.ListMetadata{
		Count:  count,
		Offset: offset,
		Limit:  limit,
		Total:  total,
	}
	return resp, nil
}

func (h baseService) Update(ctx context.Context, req dto.UpdateBaseReq) (resp dto.BaseResp, err error) {
	meta, err := req.ToMeta(DecodeID)
	if err != nil {
		return resp, err
	}
	updatedMeta, err := h.metaRepo.Update(ctx, meta)
	if err != nil {
		return resp, err
	}
	return dto.NewBaseResp(updatedMeta, EncodeID)
}

func (h baseService) Delete(ctx context.Context, id string) error {
	metaID, err := DecodeID(id)
	if err != nil {
		return err
	}
	_, err = h.metaRepo.Delete(ctx, metaID)
	return err
}
