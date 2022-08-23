package sitedb

import (
	"almcm.poscoict.com/scm/pme/curly-engine/database"
	. "almcm.poscoict.com/scm/pme/curly-engine/log"
	"gorm.io/gorm"
)

type SiteInfo struct {
	gorm.Model
	ButtonElem        string `gorm:"not null;"`
	ButtonClass       string
	DivContainerClass string `gorm:"not null;"`
	SplitElem         string `gorm:"not null;"`
	SplitElemClass    string `gorm:"not null;"`
	DivImageclass     string `gorm:"not null;"`
	ATitleclass       string `gorm:"not null;"`
	TitleElem         string `gorm:"not null;"`
	TitleClass        string `gorm:"not null;"`
	PriceElem         string `gorm:"not null;"`
	PreceClass        string `gorm:"not null;"`
	Url               string `gorm:"not null;"`
	Item              string `gorm:"not null;"`
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
		site.ButtonElem, site.ButtonClass, site.DivContainerClass, site.SplitElem, site.SplitElemClass, site.DivImageclass, site.ATitleclass, site.TitleElem, site.TitleClass, site.PriceElem, site.PreceClass, site.Url, site.Item, site.Model.CreatedAt.String(), site.Model.UpdatedAt.String())

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
