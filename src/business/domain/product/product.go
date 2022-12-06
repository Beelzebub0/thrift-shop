package product

import (
	"context"

	"github.com/Beelzebub0/thrift-shop/src/business/entity"
	sql "github.com/Beelzebub0/thrift-shop/src/lib/database"
)

type Interface interface {
	Get(ctx context.Context, params entity.ProductParams) ([]entity.Product, error)
	GetList(ctx context.Context, params entity.ProductParams) (entity.Product, error)
}

type product struct {
	db sql.Interface
}

func Init(db sql.Interface) Interface {
	p := &product{
		db: db,
	}
	return p
}

func (p *product) Get(ctx context.Context, params entity.ProductParams) ([]entity.Product, error)

func (p *product) GetList(ctx context.Context, params entity.ProductParams) (entity.Product, error)
