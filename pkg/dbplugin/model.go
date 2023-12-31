package dbplugin

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type DB = gorm.DB

type M struct {
	ID        uint                  `gorm:"primarykey"`
	CreatedAt int                   `gorm:"autoUpdateTime:milli"`
	UpdatedAt int                   `gorm:"autoUpdateTime:milli"`
	DeletedAt soft_delete.DeletedAt `gorm:"index;softDelete:milli"`
}
