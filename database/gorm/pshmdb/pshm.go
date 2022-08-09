package pshmdb

import (
    "strings"

    "almcm.poscoict.com/scm/pme/curly-engine/database/gorm"
    . "almcm.poscoict.com/scm/pme/curly-engine/log"
    "errors"
    "gorm.io/gorm"
)

const PshmTableName = "pshm"

type PshmScheme struct {
    gorm.Model
    Name     string `gorm:"not null; uniqueIndex; size:40"`
    Size     int    `gorm:"not null"`
    Sender   string `gorm:"not null"`
    Receiver string `gorm:"not null"`
    CsvFile  string
}

func (PshmScheme) TableName() string {
    return PshmTableName
}

func DropTable() {
    if gormdb.MainDB == nil {
        Loge("gormdb.MainDB is nil")
        return
    }
    gormdb.MainDB.Db.Exec("DROP TABLE " + PshmTableName)
}

func InitPshmTable() {
    if gormdb.MainDB == nil {
        Logd("gormdb.MainDB is nil")
        return
    }
    CreateTable()
    Logd("DB: Init PSHM Table")
}

func CreateTable() {
    var err error

    if gormdb.MainDB == nil {
        Logd("gormdb.MainDB is nil")
        return
    }

    err = gormdb.MainDB.Db.AutoMigrate(&PshmScheme{})
    if err != nil {
        Loge("Failed auto migrate table : %s", err)
        return
    }
}

func InsertNewShm(shm *PshmScheme) error {
    gormdb.MainDB.Db.Create(shm)
    return nil
}

func DeleteShm(name string) error {
    upper := strings.ToUpper(name)
    if upper == "ALL" {
        gormdb.MainDB.Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&PshmScheme{})
    } else {
        gormdb.MainDB.Db.Where("name = ?", name).Unscoped().Delete(&PshmScheme{})
    }
    return nil
}

func UpdateShm(shm *PshmScheme) error {
    gormdb.MainDB.Db.Where("name = ?", shm.Name).Updates(shm)

    if shm.CsvFile == "" {
        gormdb.MainDB.Db.Model(&PshmScheme{}).Where("name = ?", shm.Name).Update("csv_file", "")
    }
    return nil
}

func GetShmList(offset int, limit int, field string, order string, appName string) (retPshm []PshmScheme) {
    var pshm []PshmScheme
    var orderString string

    if len(field) > 0 {
        orderString = field + " "
    } else {
        orderString = "created_at "
    }
    if len(order) > 0 {
        orderString += order
    } else {
        orderString += "desc"
    }

    if len(appName) > 0 {
        if limit <= 0 {
            gormdb.MainDB.Db.Model(&PshmScheme{}).Where("sender like ? or receiver like ?", appName, appName).
                Offset(offset).Order(orderString).Find(&pshm)
        } else {
            gormdb.MainDB.Db.Model(&PshmScheme{}).Where("sender like ? or receiver like ?", appName, appName).
                Limit(limit).Offset(offset).Order(orderString).Find(&pshm)
        }
    } else {
        if limit <= 0 {
            gormdb.MainDB.Db.Model(&PshmScheme{}).Offset(offset).Order(orderString).Find(&pshm)
        } else {
            gormdb.MainDB.Db.Model(&PshmScheme{}).Limit(limit).Offset(offset).Order(orderString).Find(&pshm)
        }
    }

    for _, info := range pshm {
        retPshm = append(retPshm, info)
    }

    return retPshm
}

func GetShmTotalCount() int64 {
    var count int64
    gormdb.MainDB.Db.Find(&PshmScheme{}).Count(&count)
    return count
}

func GetShm(name string) (*PshmScheme, error) {
    if gormdb.MainDB == nil {
        Logd("gormdb.MainDB is nil")
        return nil, errors.New("gormdb.MainDB is nil")
    }
    var pshm PshmScheme
    result := gormdb.MainDB.Db.First(&pshm, "name = ?", name)
    if result.Error != nil {
        Logd("DB First result error: %s", result.Error)
        return nil, result.Error
    }

    return &pshm, nil
}

func GetShmFilename(filename string) (*PshmScheme, error) {
    if gormdb.MainDB == nil {
        Logd("gormdb.MainDB is nil")
        return nil, errors.New("gormdb.MainDB is nil")
    }
    var pshm PshmScheme
    result := gormdb.MainDB.Db.First(&pshm, "csv_file = ?", filename)
    if result.Error != nil {
        Logd("DB First result error: %s", result.Error)
        return nil, result.Error
    }

    return &pshm, nil
}