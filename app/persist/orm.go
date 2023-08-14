package persist

import (
	"continential/app/models"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbEngine string, dsn string) (*DBORM, error) {
	sqlDB, err := sql.Open(dbEngine, dsn)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	return &DBORM{
		DB: gormDB,
	}, err
}

func (db *DBORM) GetAllConnections() (contents []models.Content, err error) {
	return contents, db.Find(&contents).Error
}
