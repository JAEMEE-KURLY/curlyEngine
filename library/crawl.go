package utility

import (
	. "almcm.poscoict.com/scm/pme/curly-engine/log"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

var countMap map[int]string

func GetCrawlingInfo() error {
	// Request
	// 이마트
	url := "https://emart.ssg.com/search.ssg?target=all&query=%EA%B9%80%EC%B9%98"
	resp, err := http.Get(url)
	if err != nil {
		Loge("failed to get http response")
	}
	defer resp.Body.Close()

	// HTML 읽기
	html, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		Loge("failed to exec goquery")
	}

	html.Find("#item_unit_1000017881155 > div.cunit_info > div.cunit_md.notranslate > div > a > em.tx_ko").Each(func(i int, s *goquery.Selection) {
		//class, _ := s.Attr("class")
		Logd("%s", s.Text())
		//divSeletctor, ok := s.Attr("class")
		//if ok {
		//	Logd("%s", divSeletctor)
		//}
		//divSeletctor, ok := s.Find("div")
		//if ok {
		//	//for p := s.Parent(); p.Size() > 0 && !ok; p = p.Parent() {
		//	//    classes, ok = p.Attr("class")
		//	//}
		//	Logd("%s", classes)
		//}
		//Logd("Link #%d:\ntext: %s\nclass: %s\n\n", i, s.Text(), classes)
	})

	return nil
}
