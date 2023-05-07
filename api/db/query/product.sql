-- name: CreateProduct :one
INSERT INTO products (
  name,
  description,
  price,
  stock_quantity,
  category_id,
  seller_id,
  image_url
) VALUES (
  sqlc.arg('name'),
  sqlc.narg('description'),
  sqlc.arg('price'),
  sqlc.arg('stock_quantity'),
  sqlc.arg('category_id'),
  sqlc.arg('seller_id'),
  sqlc.narg('image_url')
) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: TruncateProductsTable :exec
TRUNCATE TABLE products CASCADE;
