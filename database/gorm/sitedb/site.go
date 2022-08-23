package sitedb

import (
	"almcm.poscoict.com/scm/pme/curly-engine/database"
	. "almcm.poscoict.com/scm/pme/curly-engine/log"
	"gorm.io/gorm"
)

type SiteInfo struct {
	gorm.Model
	buttonElem        string `gorm:"not null;"`
	buttonClass       string
	divContainerClass string `gorm:"not null;"`
	splitElem         string `gorm:"not null;"`
	splitElemClass    string `gorm:"not null;"`
	divImageclass     string `gorm:"not null;"`
	aTitleclass       string `gorm:"not null;"`
	titleElem         string `gorm:"not null;"`
	titleClass        string `gorm:"not null;"`
	priceElem         string `gorm:"not null;"`
	preceClass        string `gorm:"not null;"`
	url               string `gorm:"not null;"`
	item              string `gorm:"not null;"`
}

var DbInfo *database.DbInfo

func GetDbInfo() (dbInfo *database.DbInfo) {
	return DbInfo
}

func DropTableSite() {
	if DbInfo == nil || DbInfo.Db == nil {
		Loge("Database is nil")
		return
	}

	DbInfo.Db.Exec("DROP TABLE site_infos")
}

func CreateTableSite() {
	var err error

	if DbInfo == nil {
		DbInfo, err = database.ConnNewDbFromConfig()
		if err != nil {
			Loge("ConnNewDbFromConfig Error : %s", err)
			return
		}
	}

	err = DbInfo.Db.AutoMigrate(&SiteInfo{})
	if err != nil {
		Loge("Failed auto migrate table : %s", err)
		return
	}
}

func InsertSite(site_info *SiteInfo) {
	DbInfo.Db.Create(site_info)
}

func FindByIdSiteName(site_name string) *SiteInfo {
	if DbInfo == nil {
		var err error
		DbInfo, err = database.ConnNewDbFromConfig()
		if err != nil {
			Loge("ConnNewDbFromConfig Error : %s", err)
			return nil
		}
	}
	var site SiteInfo
	DbInfo.Db.First(&site, "site_name = ?", site_name)

	Logd("SiteInfobuttonElem= %s, buttonClass= %s, divContainerClass= %s, splitElem= %s, splitElemClass= %s, divImageclass= %s, aTitleclass= %s, titleElem= %s, titleClass= %s, priceElem= %s, preceClass= %s, url= %s, item              ",
		site.buttonElem, site.buttonClass, site.divContainerClass, site.splitElem, site.splitElemClass, site.divImageclass, site.aTitleclass, site.titleElem, site.titleClass, site.priceElem, site.preceClass, site.url, site.item, site.Model.CreatedAt.String(), site.Model.UpdatedAt.String())

	return &site
}

func InitSiteTable() {
	var err error
	DbInfo, err = database.ConnNewDbFromConfig()
	if err != nil {
		return
	}
	CreateTableSite()
}
