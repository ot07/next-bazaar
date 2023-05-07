// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ID            uuid.UUID      `json:"id"`
	Name          string         `json:"name"`
	Description   sql.NullString `json:"description"`
	Price         string         `json:"price"`
	StockQuantity int32          `json:"stock_quantity"`
	CategoryID    uuid.UUID      `json:"category_id"`
	SellerID      uuid.UUID      `json:"seller_id"`
	ImageUrl      sql.NullString `json:"image_url"`
	CreatedAt     time.Time      `json:"created_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	SessionToken uuid.UUID `json:"session_token"`
	UserID       uuid.UUID `json:"user_id"`
	ExpiredAt    time.Time `json:"expired_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	ID                uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	HashedPassword    string    `json:"hashed_password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}
