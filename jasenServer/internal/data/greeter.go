package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"

	"jasenServer/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	result := r.data.db.Create(g)
	if result.Error != nil {
		return nil, errors.InternalServer("GREETER_CREATE_ERROR", "用户创建失败")
	}
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(ctx context.Context, id int64) (*biz.Greeter, error) {
	var g []*biz.Greeter
	if err := r.data.db.Where("id=?", id).Find(&g).Error; err != nil {
		return nil, errors.NotFound("GREETER_NOT_FOUND", "用户不存在")
	}
	return g[0], nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}
