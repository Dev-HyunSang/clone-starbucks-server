package models

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

/*주의사항
Ex - `gorm:"type: varchar(10)"` type 이후 뛰어쓰기 있는 경우 적용되지 않음.*/

type Users struct {
	Pk int `gorm:"primaryKey; autoIncrement:true;"`
	// TODO: why not using to gorm type uuid? => MySQL type Checking PLZ
	Id             uuid.UUID    `gorm:"type:varchar(36); not null;"`
	Name           string       `gorm:"type:varchar(100); not null;"`
	Email          string       `gorm:"type:longtext; not null;"`
	Password       string       `gorm:"type:longtext; not null;"`
	PhoneNumber    string       `gorm:"type:varchar(100); not null;"`
	DisplayName    string       `gorm:"type:varchar(100); not null;"`
	Birthday       time.Time    `gorm:"type:datetime; not null;"`
	AllowMarketing sql.NullBool `gorm:"default:false; not null;"`
	CreatedAt      time.Time    `gorm:"autoCreateTime;"`
	UpdatedAt      time.Time    `gorm:"autoCreateTime;"`
	ActivatedAt    sql.NullTime `gorm:"autoCreateTime;"`
	DeletedAt      sql.NullTime
}
