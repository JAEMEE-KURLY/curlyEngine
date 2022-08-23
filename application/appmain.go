package appmmain

import (
	httpserver "almcm.poscoict.com/scm/pme/curly-engine/communication/http"
	. "almcm.poscoict.com/scm/pme/curly-engine/configure"
	gormdb "almcm.poscoict.com/scm/pme/curly-engine/database/gorm"
	"almcm.poscoict.com/scm/pme/curly-engine/database/gorm/categorydb"
	"almcm.poscoict.com/scm/pme/curly-engine/database/gorm/itemdb"
	"almcm.poscoict.com/scm/pme/curly-engine/database/gorm/sitedb"
	utility "almcm.poscoict.com/scm/pme/curly-engine/library"
	. "almcm.poscoict.com/scm/pme/curly-engine/log"
	"sync"
)

func Run(conf *Values) {
	if conf.CurlyEngine.Enabled == false {
		Logi("CurlyEngine Application Disabled")
		return
	}
	Logi("CurlyEngine Application Started")

	wg := new(sync.WaitGroup)

	err := gormdb.InitSingletonDB()
	if err == nil {
		sitedb.InitSiteTable()
		itemdb.InitItemTable()
		categorydb.InitCategoryTable()
	}

	// HTTP
	if conf.Net.EnableHttp && conf.CurlyEngine.EnableHttpServer {
		Logi("http enabled")
		go httpserver.HttpServer(conf.CurlyEngine.HttpServerPort)
	}
	go utility.SchedulerDay(ScheduleCrawling)

	wg.Wait()
}
