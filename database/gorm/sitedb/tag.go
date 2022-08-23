package sitedb

import (
	"almcm.poscoict.com/scm/pme/curly-engine/database"
	gormdb "almcm.poscoict.com/scm/pme/curly-engine/database/gorm"
	"almcm.poscoict.com/scm/pme/curly-engine/database/gorm/categorydb"
	. "almcm.poscoict.com/scm/pme/curly-engine/log"
	"gorm.io/gorm"
)

const SiteInfoTableName = "tag_info"

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
	PriceClass        string `gorm:"not null;"`
	Url               string `gorm:"not null;"`
	SiteName          string `gorm:"not null;"`
}

type AllInfo struct {
	CategoryName      string
	SiteName          string `gorm:"not null;"`
	ButtonElem        string `gorm:"not null;"`
	ButtonClass       string
	DivContainerClass string `gorm:"not null;"`
	SplitElem         string `gorm:"not null;"`
	SplitElemClass    string `gorm:"not null;"`
	DivImageclass     string `gorm:"not null;"`
	ATitleclass       string
	TitleElem         string `gorm:"not null;"`
	TitleClass        string
	PriceElem         string `gorm:"not null;"`
	PriceClass        string `gorm:"not null;"`
	Url               string `gorm:"not null;"`
}

func (SiteInfo) TableName() string {
	return SiteInfoTableName
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

	DbInfo.Db.Exec("DROP TABLE " + SiteInfoTableName)
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
func GetCrawlInfoList() (retInfo []AllInfo) {
	arr := categorydb.GetCategoryList()
	cnt := categorydb.GetInfoArr()
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(cnt); j++ {
			temp := GetInfo(cnt[j])
			retInfo = append(retInfo, AllInfo{
				CategoryName:      arr[i].CategoryName,
				SiteName:          temp.SiteName,
				ButtonElem:        temp.ButtonElem,
				ButtonClass:       temp.ButtonClass,
				DivContainerClass: temp.DivContainerClass,
				SplitElem:         temp.SplitElem,
				SplitElemClass:    temp.SplitElemClass,
				DivImageclass:     temp.DivImageclass,
				ATitleclass:       temp.ATitleclass,
				TitleElem:         temp.TitleElem,
				TitleClass:        temp.TitleClass,
				PriceElem:         temp.PriceElem,
				PriceClass:        temp.PriceClass,
				Url:               temp.Url,
			})
		}
	}

	return retInfo
}
func GetSiteInfoList() (retSite []SiteInfo) {
	var site []SiteInfo

	gormdb.MainDB.Db.Model(&SiteInfo{}).Find(&site)

	for _, info := range site {
		retSite = append(retSite, info)
	}
	//fmt.Printf("%v", retSite)
	return retSite
}
func GetInfo(id int) SiteInfo {
	var info SiteInfo
	DbInfo.Db.Model(&SiteInfo{}).Where("id=?", id).Scan(&info)
	//fmt.Printf("%v", info)
	return info
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

	Logd("SiteInfo buttonElem= %s, buttonClass= %s, divContainerClass= %s, splitElem= %s, splitElemClass= %s, divImageclass= %s, aTitleclass= %s, titleElem= %s, titleClass= %s, priceElem= %s, preceClass= %s, url= %s, site_name=%s, createAt=%s, updateAt=%s",
		site.ButtonElem, site.ButtonClass, site.DivContainerClass, site.SplitElem, site.SplitElemClass, site.DivImageclass, site.ATitleclass, site.TitleElem, site.TitleClass, site.PriceElem, site.PriceClass, site.Url, site.SiteName, site.Model.CreatedAt.String(), site.Model.UpdatedAt.String())

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
