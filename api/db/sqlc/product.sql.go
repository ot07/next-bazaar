// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: product.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const countProducts = `-- name: CountProducts :one
SELECT count(*) FROM products
`

func (q *Queries) CountProducts(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countProducts)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (
  name,
  description,
  price,
  stock_quantity,
  category_id,
  seller_id,
  image_url
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
) RETURNING id, name, description, price, stock_quantity, category_id, seller_id, image_url, created_at
`

type CreateProductParams struct {
	Name          string         `json:"name"`
	Description   sql.NullString `json:"description"`
	Price         string         `json:"price"`
	StockQuantity int32          `json:"stock_quantity"`
	CategoryID    uuid.UUID      `json:"category_id"`
	SellerID      uuid.UUID      `json:"seller_id"`
	ImageUrl      sql.NullString `json:"image_url"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.StockQuantity,
		arg.CategoryID,
		arg.SellerID,
		arg.ImageUrl,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.StockQuantity,
		&i.CategoryID,
		&i.SellerID,
		&i.ImageUrl,
		&i.CreatedAt,
	)
	return i, err
}

const getProduct = `-- name: GetProduct :one
SELECT id, name, description, price, stock_quantity, category_id, seller_id, image_url, created_at FROM products
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, id uuid.UUID) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.StockQuantity,
		&i.CategoryID,
		&i.SellerID,
		&i.ImageUrl,
		&i.CreatedAt,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT id, name, description, price, stock_quantity, category_id, seller_id, image_url, created_at FROM products
ORDER BY created_at
LIMIT $1
OFFSET $2
`

type ListProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.StockQuantity,
			&i.CategoryID,
			&i.SellerID,
			&i.ImageUrl,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const truncateProductsTable = `-- name: TruncateProductsTable :exec
TRUNCATE TABLE products CASCADE
`

func (q *Queries) TruncateProductsTable(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, truncateProductsTable)
	return err
}
