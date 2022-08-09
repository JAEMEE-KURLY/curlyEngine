package utility

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/labstack/gommon/log"
    "net/http"
)

func getCrawlingInfo() error {
    // Request
    // 이마트
    url := "https://emart.ssg.com/search.ssg?target=all&query=%EA%B9%80%EC%B9%98"
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    // HTML 읽기
    html, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    // 현황판 파싱
    wrapper := html.Find("ul.liveNum")
    items := wrapper.Find("li")

    // items 순회하면서 출력
    items.Each(func(idx int, sel *goquery.Selection) {
        title := sel.Find("strong.tit").Text()
        value := sel.Find("span.num").Text()
        before := sel.Find("span.before").Text()

        fmt.Println(title, value, before)
    })

    // 홈플러스
    url := "https://emart.ssg.com/search.ssg?target=all&query=%EA%B9%80%EC%B9%98"
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    // HTML 읽기
    html, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    // 현황판 파싱
    wrapper := html.Find("ul.liveNum")
    items := wrapper.Find("li")

    // items 순회하면서 출력
    items.Each(func(idx int, sel *goquery.Selection) {
        title := sel.Find("strong.tit").Text()
        value := sel.Find("span.num").Text()
        before := sel.Find("span.before").Text()

        fmt.Println(title, value, before)
    })

    // 롯데
    url := "https://emart.ssg.com/search.ssg?target=all&query=%EA%B9%80%EC%B9%98"
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    // HTML 읽기
    html, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    // 현황판 파싱
    wrapper := html.Find("ul.liveNum")
    items := wrapper.Find("li")

    // items 순회하면서 출력
    items.Each(func(idx int, sel *goquery.Selection) {
        title := sel.Find("strong.tit").Text()
        value := sel.Find("span.num").Text()
        before := sel.Find("span.before").Text()

        fmt.Println(title, value, before)
    })
}
