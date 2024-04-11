package service

import (
	"fmt"
	"mezink-assignment/config"
	"mezink-assignment/infra"
	"mezink-assignment/internal/model"

	"gorm.io/gorm/clause"
)

type IRecordService interface {
	CreateRecord(record model.Record) (err error)
	GetRecordByID(id int) (record model.Record, err error)
	GetRecordsByFilter(filter model.RecordFilterRequest) (records model.Records, err error)
	UpsertRecord(record model.Record) (err error)
	DeleteRecord(id int) (err error)
}

type RecordService struct {
	Config config.Config
	Infra  *infra.Connection
}

func (u *RecordService) CreateRecord(record model.Record) (err error) {
	return u.Infra.PG.DB.Table(record.TableName()).Create(&record).Error
}

func (u *RecordService) GetRecordByID(id int) (record model.Record, err error) {
	err = u.Infra.PG.DB.
		Table(record.TableName()).
		Select(record.SelectedFields()).
		Where("id = ?", id).
		Take(&record).Error

	return
}

func (u *RecordService) GetRecordsByFilter(filter model.RecordFilterRequest) (records model.Records, err error) {
	mdl := model.Record{}
	db := u.Infra.PG.DB.
		Table(mdl.TableName()).
		Select(mdl.SelectedFields())

	if filter.StartDate.Valid {
		db.Where("created_at >= ?", fmt.Sprintf("%s 00:00:00", filter.StartDate.ValueOrZero()))
	}
	if filter.EndDate.Valid {
		db.Where("created_at <= ?", fmt.Sprintf("%s 23:59:59", filter.EndDate.ValueOrZero()))
	}
	if filter.MinCount.Valid {
		db.Where("(SELECT SUM(m) FROM UNNEST(marks) m) >= ?", filter.MinCount.ValueOrZero())
	}
	if filter.MaxCount.Valid {
		db.Where("(SELECT SUM(m) FROM UNNEST(marks) m) <= ?", filter.MaxCount.ValueOrZero())
	}

	err = db.Find(&records).Error

	return
}

func (u *RecordService) UpsertRecord(record model.Record) (err error) {
	return u.Infra.PG.DB.
		Table(record.TableName()).
		Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&record).Error
}

func (u *RecordService) DeleteRecord(id int) (err error) {
	return u.Infra.PG.DB.
		Table(model.Record{}.TableName()).
		Where("id = ?", id).
		Delete(model.Record{}).Error
}
