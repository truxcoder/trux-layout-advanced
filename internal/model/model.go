package model

import (
	"github.com/truxcoder/trux-layout-advanced/pkg/sid"
	"gorm.io/gorm"
	"time"
)

var s *sid.SnowID

func init() {
	s = sid.NewSnowID()
}

type Base struct {
	ID        int64 `json:"id,string" gorm:"autoIncrement:false;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	var id int64
	if id, err = s.GenInt64(); err != nil {
		return
	}
	b.ID = id
	return
}
