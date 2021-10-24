package models

import (
	"context"
	"time"

	"github.com/hanherb/omdb-api/lib/variables"
)

type baseLogger struct {
	baseTable
}

func (baseLogger) TableName() string {
	return variables.TableName.Logger
}

type Logger struct {
	baseLogger

	ID              int32     `gorm:"column:id" json:"id"`
	SearchWord      string    `gorm:"column:search_word" json:"search_word"`
	Pagination      int32     `gorm:"column:pagination" json:"pagination"`
	Status          string    `gorm:"column:status" json:"status"`
	DatetimeCreated time.Time `gorm:"column:datetime_created" json:"datetime_created"`
}

type Loggers struct {
	baseLogger
	data []*Logger
}

func (u *Logger) FilterByID(id int32) {
	u.initQuery()
	u.query = u.query.Where("id = ?", id)
}

func (u *Logger) Insert(ctx context.Context) error {
	u.initQuery()
	u.DatetimeCreated = time.Now()

	return u.query.WithContext(ctx).Select(
		"search_word",
		"pagination",
		"status",
		"datetime_created").Create(u).Error
}

func (u *Logger) Update(ctx context.Context) error {
	u.initQuery()
	return u.query.WithContext(ctx).Save(u).Error
}

func (u *Logger) Delete(ctx context.Context) error {
	u.initQuery()
	return u.query.WithContext(ctx).Delete(u).Error
}

func (u *Logger) Get(ctx context.Context) (found bool, err error) {
	u.initQuery()
	err = u.query.WithContext(ctx).First(u).Error
	return u.isFoundFromError(err), err
}

// ==================================================================================

func (u *Loggers) Get(ctx context.Context) error {
	u.initQuery()
	return u.query.WithContext(ctx).Find(&u.data).Error
}

func (u *Loggers) Data() []*Logger {
	return u.data
}
