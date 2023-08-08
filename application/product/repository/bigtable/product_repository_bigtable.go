package repository

import (
	"context"
	"smart-recommendation/domain"
	"smart-recommendation/internal/errorhandler"

	"cloud.google.com/go/bigtable"
)

// struct to store the connection that will be used
type ProductRepository struct {
	conn *bigtable.Client
}

// init new repository product
func NewProductRepository(conn *bigtable.Client) *ProductRepository {
	return &ProductRepository{
		conn: conn,
	}
}

func (obj *ProductRepository) Get(key string) (result domain.Product, err errorhandler.ErrorData) {
	tbl := obj.conn.Open("userID")
	ctx := context.Background()
	row, errData := tbl.ReadRow(ctx, key)
	if errData != nil {
		err = errorhandler.ErrorRepo(errData)
		return
	}
	for _, cells := range row {
		for _, cell := range cells {
			if cell.Column == "rekomendasi:norek" {
				result.Norek = string(cell.Value)
			} else if cell.Column == "rekomendasi:rekomendation" {
				result.Recomendation = string(cell.Value)
			}
		}
	}
	return
}

func (obj *ProductRepository) Insert(column string, columnTwo string, data string, datatwo string) (err errorhandler.ErrorData) {
	tbl := obj.conn.Open("userID")
	ctx := context.Background()
	mutation := bigtable.NewMutation()
	mutation.Set("rekomendasi", column, bigtable.Now(), []byte(data))
	mutation.Set("rekomendasi", columnTwo, bigtable.Now(), []byte(datatwo))
	errData := tbl.Apply(ctx, data, mutation)
	if errData != nil {
		err = errorhandler.ErrorRepo(errData)
	}
	return
}
