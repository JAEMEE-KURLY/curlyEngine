package appmmain

import (
	"almcm.poscoict.com/scm/pme/curly-engine/database/gorm/sitedb"
	utility "almcm.poscoict.com/scm/pme/curly-engine/library"
	. "almcm.poscoict.com/scm/pme/curly-engine/log"
)

func ScheduleCrawling() {
	Logd("Crawling start from Scheduler")
	////Add Logic
	siteList := sitedb.GetCrawlInfoList()
	Logd("%v", siteList)
	for _, site := range siteList {
		go utility.GetScrawlingInfo(site.ButtonElem, site.ButtonClass, site.DivContainerClass,
			site.SplitElem, site.SplitElemClass, site.DivImageclass, site.ATitleclass,
			site.TitleElem, site.TitleClass, site.PriceElem, site.PriceClass,
			site.Url, site.CategoryName)
	}
}
