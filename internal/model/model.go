package model

import (
	"time"

	"database/sql"
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
	Asset    uint
	Exchange uint
	Datetime time.Time
	Value    float64
}

type Exchange struct {
	Model
	Name   string // e.g. Uniswap V2
	Symbol string // e.g. univ2
	Type   string // e.g. dex
}
