package models

import (
	"github.com/hanherb/omdb-api/lib/services"
	"gorm.io/gorm"
)

type baseTable struct {
	query *gorm.DB
}

func (b *baseTable) initQuery() {
	if b.query == nil {
		b.query = services.DB
	}
}

func (b *baseTable) SetSession(session *gorm.DB) {
	b.query = session
}

func (b *baseTable) isFoundFromError(err error) bool {
	if err == nil {
		return true
	}

	if err == gorm.ErrRecordNotFound {
		return false
	}
	return false
}
