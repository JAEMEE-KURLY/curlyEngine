package configure

import (
    "fmt"
    "gopkg.in/ini.v1"
    "strings"
    "errors"
)

const (
    SECT_CORE     = "CORE"
    SECT_TIMER    = "TIMER"
    SECT_NET      = "NET"
    SECT_LOG      = "LOG"
    SECT_DATABASE = "DATABASE"

    SECT_CurlyEngine = "CurlyEngine" // CurlyEngine: insert your section name string
)

var (
    Config ConfigInfo
)

type ConfigInfo struct {
    File  *ini.File
    Path  string
    Value *Values
}

type Values struct {
    // Core
    Core *ValueCore
    Time *ValueTime
    Net  *ValueNet
    Log  *ValueLog
    Db   *ValueDatabase

    // Application
    CurlyEngine *ValueCurlyEngine // CurlyEngine: insert your structure variable
}

type ValueCore struct {
    sectionName string
}

type ValueTime struct {
    sectionName string

    IdleTimeout uint64

    SchedulerMinute bool
    SchedulerHour   bool
    SchedulerDay    bool
}

type ValueNet struct {
    sectionName string

    EnableHttp   bool
    EnableWebsocket bool
    EnableGrpc   bool
    EnableTcp    bool
    EnableSerial bool
    EnableMqueue bool
}

type ValueLog struct {
    sectionName string

    EnableDebug bool
    EnableInfo  bool
    EnableError bool
    LogFile     string
    MaxSize     int
    MaxBackups  int
    MaxAge      int
    LocalTime   bool
    Compress    bool
}

type ValueDatabase struct {
    sectionName string

    Type      string
    UserId    string
    Password  string
    IpAddress string
    Port      int
    DbName    string
    DbNameLog string
}

func NewValues() *Values {
    if Config.Value == nil {
        Config.Value = &Values{}

        Config.Value.Core = &ValueCore{}
        Config.Value.Time = &ValueTime{}
        Config.Value.Net = &ValueNet{}
        Config.Value.Log = &ValueLog{}
        Config.Value.Db = &ValueDatabase{}

        Config.Value.CurlyEngine = &ValueCurlyEngine{} // CurlyEngine: insert your new structure
    }
    return Config.Value
}

func GetConfig() *Values {
    return Config.Value
}

func (c *Values) chkCfgFile(filePath string) (*Values, error) {
    if Config.File == nil {
        var err error
        Config.File, err = ini.Load(filePath)
        if err != nil {
            return nil, err
        }
        Config.Path = filePath
        // fmt.Printf("Load INI file successed [%s]\n", filePath)
    }
    return c, nil
}

func (c *Values) GetValueALL(filePath string) (*Values, error) {
    var err error
    _, err = c.chkCfgFile(filePath)
    if err != nil {
        return nil, err
    }

    _, err = c.GetValueCore(filePath)
    if err != nil {
        return nil, err
    }
    _, err = c.GetValueTimer(filePath)
    if err != nil {
        return nil, err
    }
    _, err = c.GetValueNet(filePath)
    if err != nil {
        return nil, err
    }
    _, err = c.GetValueLog(filePath)
    if err != nil {
        return nil, err
    }
    _, err = c.GetValueDatabase(filePath)
    if err != nil {
        return nil, err
    }

    _, err = c.GetValueCurlyEngine(filePath) // CurlyEngine: insert your option parsing fuction
    if err != nil {
        return nil, err
    }

    return c, nil
}

func (c *Values) GetValueCore(filePath string) (*Values, error) {
    _, err := c.chkCfgFile(filePath)
    if err != nil {
        return nil, err
    }

    c.Core.sectionName = SECT_CORE

    // fmt.Printf("Core Config Values: %+v\n", c.Log)

    return c, nil
}

func (c *Values) GetValueTimer(filePath string) (*Values, error) {
    _, err := c.chkCfgFile(filePath)
    if err != nil {
        return nil, err
    }

    c.Time.sectionName = SECT_TIMER
    c.Time.IdleTimeout = Config.File.Section(SECT_TIMER).Key("IdleTimeout").MustUint64(10)
    c.Time.SchedulerMinute = Config.File.Section(SECT_TIMER).Key("SchedulerMinute").MustBool(false)
    c.Time.SchedulerHour = Config.File.Section(SECT_TIMER).Key("SchedulerHour").MustBool(false)
    c.Time.SchedulerDay = Config.File.Section(SECT_TIMER).Key("SchedulerDay").MustBool(false)
    // fmt.Printf("Time Config Values: %+v\n", c.Log)

    return c, nil
}

func (c *Values) GetValueNet(filePath string) (*Values, error) {
    _, err := c.chkCfgFile(filePath)
    if err != nil {
        return nil, err
    }

    c.Net.sectionName = SECT_NET
    c.Net.EnableHttp = Config.File.Section(SECT_NET).Key("EnableHttp").MustBool(false)
    c.Net.EnableWebsocket = Config.File.Section(SECT_NET).Key("EnableWebsocket").MustBool(false)
    c.Net.EnableGrpc = Config.File.Section(SECT_NET).Key("EnableGrpc").MustBool(false)
    c.Net.EnableTcp = Config.File.Section(SECT_NET).Key("EnableTcp").MustBool(false)
    c.Net.EnableSerial = Config.File.Section(SECT_NET).Key("EnableSerial").MustBool(false)
    c.Net.EnableMqueue = Config.File.Section(SECT_NET).Key("EnableMqueue").MustBool(false)

    // fmt.Printf("Net Config Values: %+v\n", c.Log)

    return c, nil
}

func (c *Values) GetValueLog(filePath string) (*Values, error) {
    _, err := c.chkCfgFile(filePath)
    if err != nil {
        return nil, err
    }

    c.Log.sectionName = SECT_LOG
    c.Log.EnableDebug = Config.File.Section(SECT_LOG).Key("EnableDebug").MustBool(true)
    c.Log.EnableInfo = Config.File.Section(SECT_LOG).Key("EnableInfo").MustBool(true)
    c.Log.EnableError = Config.File.Section(SECT_LOG).Key("EnableError").MustBool(true)
    c.Log.LogFile = Config.File.Section(SECT_LOG).Key("LogFile").MustString("sample.log")
    c.Log.MaxSize = Config.File.Section(SECT_LOG).Key("MaxSize").MustInt(100)
    c.Log.MaxBackups = Config.File.Section(SECT_LOG).Key("MaxBackups").MustInt(30)
    c.Log.MaxAge = Config.File.Section(SECT_LOG).Key("MaxAge").MustInt(30)
    c.Log.LocalTime = Config.File.Section(SECT_LOG).Key("LocalTime").MustBool(true)
    c.Log.Compress = Config.File.Section(SECT_LOG).Key("Compress").MustBool(false)

    // fmt.Printf("Log Config Values: %+v\n", c.Log)

    return c, nil
}

func (c *Values) GetValueDatabase(filePath string) (*Values, error) {
    _, err := c.chkCfgFile(filePath)
    if err != nil {
        return nil, err
    }

    c.Db.sectionName = SECT_DATABASE
    c.Db.Type = Config.File.Section(SECT_DATABASE).Key("Type").MustString("mysql")
    c.Db.UserId = Config.File.Section(SECT_DATABASE).Key("UserId").MustString("posgo")
    c.Db.Password = Config.File.Section(SECT_DATABASE).Key("Password").MustString("posgo")
    c.Db.IpAddress = Config.File.Section(SECT_DATABASE).Key("IpAddress").MustString("127.0.0.1")
    c.Db.Port = Config.File.Section(SECT_DATABASE).Key("Port").MustInt(-1)
    c.Db.DbName = Config.File.Section(SECT_DATABASE).Key("DbName").MustString("posgo")
    c.Db.DbNameLog = Config.File.Section(SECT_DATABASE).Key("DbNameLog").MustString("log")

    // fmt.Printf("Dayabase Config Values: %+v\n", c.Db)

    return c, nil
}

func (c *Values) PrintValues(sectName string) (prtStr string) {
    isAll := false

    sectName = strings.ToUpper(sectName)
    switch sectName {
    case SECT_CORE:
        prtStr += fmt.Sprintf(" * %s Config Values: [%+v]\n", sectName, c.Core)
    case SECT_TIMER:
        prtStr += fmt.Sprintf(" * %s Config Values: [%+v]\n", sectName, c.Time)
    case SECT_LOG:
        prtStr += fmt.Sprintf(" * %s Config Values: [%+v]\n", sectName, c.Log)
    case SECT_NET:
        prtStr += fmt.Sprintf(" * %s Config Values: [%+v]\n", sectName, c.Net)
    case SECT_DATABASE:
        prtStr += fmt.Sprintf(" * %s Config Values: [%+v]\n", sectName, c.Db)

    case SECT_CurlyEngine: // application configure value
        prtStr += fmt.Sprintf(" * %s Config Values: [%+v]\n", sectName, c.CurlyEngine)
    default:
        isAll = true
    }

    if isAll {
        prtStr += fmt.Sprintf(" * %s Config Values: [%+v]\n", SECT_CORE, c.Core)
        prtStr += fmt.Sprintf(" * %s Config Values: [%+v]\n", SECT_TIMER, c.Time)
        prtStr += fmt.Sprintf(" * %s Config Values: [%+v]\n", SECT_LOG, c.Log)
        prtStr += fmt.Sprintf(" * %s Config Values: [%+v]\n", SECT_NET, c.Net)
        prtStr += fmt.Sprintf(" * %s Config Values: [%+v]\n", SECT_DATABASE, c.Db)
        prtStr += fmt.Sprintf(" * %s Config Values: [%+v]\n", SECT_CurlyEngine, c.CurlyEngine) // application configure value
    }
    return
}

// SetConfigValue 설정값을 변경한다
func (c *Values) SetConfigValue(section string, key string, value string) error {
    if Config.File == nil {
        return errors.New("error: Config file is nil")
    }
    Config.File.Section(section).Key(key).SetValue(value)
    return nil
}

// SaveConfigFile 변경한 설정값을 파일로 저장한다
func (c *Values) SaveConfigFile() error {
    err := Config.File.SaveTo(Config.Path)
    if err != nil {
        return err
    }
    return nil
}
