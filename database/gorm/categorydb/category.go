package categorydb

import (
	"almcm.poscoict.com/scm/pme/curly-engine/database"
	gormdb "almcm.poscoict.com/scm/pme/curly-engine/database/gorm"
	. "almcm.poscoict.com/scm/pme/curly-engine/log"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

const CategoryInfoTableName = "category_info"

type CategoryInfo struct {
	gorm.Model
	CategoryName string `gorm:"not null;uniqueIndex""`
	TargetInfo   string `gorm:"not null";`
}

func (CategoryInfo) TableName() string {
	return CategoryInfoTableName
}

var DbInfo *database.DbInfo

func GetDbInfo() (dbInfo *database.DbInfo) {
	return DbInfo
}

func DropTableCategory() {
	if DbInfo == nil || DbInfo.Db == nil {
		Loge("Database is nil")
		return
	}

	DbInfo.Db.Exec("DROP TABLE " + CategoryInfoTableName)
}

func CreateTableCategory() {
	var err error

	if DbInfo == nil {
		DbInfo, err = database.ConnNewDbFromConfig()
		if err != nil {
			Loge("ConnNewDbFromConfig Error : %s", err)
			return
		}
	}

	err = DbInfo.Db.AutoMigrate(&CategoryInfo{})
	if err != nil {
		Loge("Failed auto migrate table : %s", err)
		return
	}
}
func GetCategoryList() (retCategory []CategoryInfo) {
	var category_item []CategoryInfo

	gormdb.MainDB.Db.Model(&CategoryInfo{}).Find(&category_item)

	for _, info := range category_item {
		retCategory = append(retCategory, info)
	}
	return retCategory
}
func GetInfoArr() []int {
	var target_info string
	DbInfo.Db.Table("category_info").Select("target_info").Find(&target_info)
	slice := strings.Split(target_info, ",")
	arr := make([]int, len(slice))
	for i := range arr {
		arr[i], _ = strconv.Atoi(slice[i])
	}
	//fmt.Println(arr)
	return arr
}

func FindByIdItemName(category_name string) *CategoryInfo {
	if DbInfo == nil {
		var err error
		DbInfo, err = database.ConnNewDbFromConfig()
		if err != nil {
			Loge("ConnNewDbFromConfig Error : %s", err)
			return nil
		}
	}
	var category CategoryInfo
	DbInfo.Db.First(&category, "category_name = ?", category.CategoryName)

	Logd("CategoryInfo category_info=%s, target_info=%s, createAt=%s, updateAt=%s",
		category.CategoryName, category.TargetInfo, category.Model.CreatedAt.String(), category.Model.UpdatedAt.String())

	return &category
}

func InitCategoryTable() {
	var err error
	DbInfo, err = database.ConnNewDbFromConfig()
	if err != nil {
		return
	}
	CreateTableCategory()
}

func InsertNewItem(item *CategoryInfo) error {
	gormdb.MainDB.Db.Create(item)
	return nil
}
