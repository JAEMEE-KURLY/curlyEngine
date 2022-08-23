package itemdb

import (
	"almcm.poscoict.com/scm/pme/curly-engine/database"
	. "almcm.poscoict.com/scm/pme/curly-engine/log"
	"gorm.io/gorm"
)

type ItemInfo struct {
	gorm.Model
	Name     string `gorm:"not null; uniqueIndex; size:20"`
	Price    string `gorm:"not null";`
	SiteName string `gorm:"not null";`
	Weight   string `gorm:"not null";`
	Category string `gorm:"not null";`
	Cnt      string
	Unit     string
	Bundle   string
	ImageSrc string
	Size     string
}

var DbInfo *database.DbInfo

func GetDbInfo() (dbInfo *database.DbInfo) {
	return DbInfo
}

func DropTableItem() {
	if DbInfo == nil || DbInfo.Db == nil {
		Loge("Database is nil")
		return
	}

	DbInfo.Db.Exec("DROP TABLE item_infos")
}

func CreateTableItem() {
	var err error

	if DbInfo == nil {
		DbInfo, err = database.ConnNewDbFromConfig()
		if err != nil {
			Loge("ConnNewDbFromConfig Error : %s", err)
			return
		}
	}

	err = DbInfo.Db.AutoMigrate(&ItemInfo{})
	if err != nil {
		Loge("Failed auto migrate table : %s", err)
		return
	}
}

func InsertUser(item_info *ItemInfo) {
	DbInfo.Db.Create(item_info)
}

func FindByIdItemName(item_name string) *ItemInfo {
	if DbInfo == nil {
		var err error
		DbInfo, err = database.ConnNewDbFromConfig()
		if err != nil {
			Loge("ConnNewDbFromConfig Error : %s", err)
			return nil
		}
	}
	var item ItemInfo
	DbInfo.Db.First(&item, "name = ?", item)

	Logd("ItemInfo category=%s, name=%s, price=%s, site_name=%s, weight=%s, cnt=%s, unit=%s, bundle=%s, imagesrc=%s, size=%s, createAt=%s, updateAt=%s",
		item.Category, item.Name, item.Price, item.SiteName, item.Weight, item.Cnt, item.Unit, item.Bundle, item.ImageSrc, item.Size, item.Model.CreatedAt.String(), item.Model.UpdatedAt.String())

	return &item
}

func InitItemTable() {
	var err error
	DbInfo, err = database.ConnNewDbFromConfig()
	if err != nil {
		return
	}
	CreateTableItem()
}
