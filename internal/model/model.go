package model

import (
	"time"

	"database/sql"

	"github.com/google/uuid"
)

type Model struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"default:now()"`
	DeletedAt sql.NullTime
}

type Note struct {
	Model
	Title    string
	SubTitle string
	Text     string
}

type Asset struct {
	Model
	Name   string // e.g. bitcoin
	Symbol string // e.g. btc
	Type   string // e.g. crypto
}

type Price struct {
	Model
	AssetId    uint
	ExchangeId uint
	Datetime   time.Time
	Value      float64
}

type Exchange struct {
	Model
	Name   string // e.g. Uniswap V2
	Symbol string // e.g. univ2
	Type   string // e.g. dex
}

type User struct {
	Model
	Username string
	Password string
}

type ApiKey struct {
	Model
	UserId   uint
	Key      uuid.UUID `gorm:"type:uuid"`
	Duration time.Duration
}
