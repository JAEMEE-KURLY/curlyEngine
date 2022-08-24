package utility

import (
    "almcm.poscoict.com/scm/pme/curly-engine/database/gorm/itemdb"
    . "almcm.poscoict.com/scm/pme/curly-engine/log"
    "context"
    "github.com/PuerkitoBio/goquery"
    "github.com/chromedp/chromedp"
    "math"
    libUrl "net/url"
    "regexp"
    "strconv"
    "strings"
    "time"
)

func GetScrawlingInfo(buttonElem string, buttonClass string, divContainerClass string, splitElem string,
    splitElemClass string, divImageclass string, aTitleclass string, titleElem string,
    titleClass string, priceElem string, priceClass string, url string, item string) error {
    Logd("start crawling")
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    var name []string
    var price []string
    var imageSrc []string

    realUrl := url + libUrl.QueryEscape(item)

    var data [5]string
    if err := chromedp.Run(ctx,

        chromedp.Navigate(realUrl),
        chromedp.Sleep(time.Second*1),
        chromedp.OuterHTML("html", &data[0], chromedp.ByQuery),
        chromedp.ActionFunc(func(ctx context.Context) error {
            for i := 1; i <= 4; i++ {
                // todo search button and  click
                if buttonElem == "button" {
                    chromedp.Click(buttonElem+"."+strings.Join(strings.Split(buttonClass, " "), "."), chromedp.ByQueryAll).Do(ctx)
                } else {
                    chromedp.Click("//"+buttonElem+"[text() = '"+strconv.Itoa(i+1)+"']", chromedp.BySearch).Do(ctx)
                }
                chromedp.Sleep(time.Second * 3).Do(ctx)
                chromedp.OuterHTML("html", &data[i], chromedp.ByQuery).Do(ctx)
            }
            return nil
        }),
        //chromedp.Click("#root > div > div.css-1di1x1r-container > div.css-oiwa5q-defaultStyle-gridRow-IntegratedSearch > div.mainWrap > div:nth-child(2) > div > div.pagination-js.css-dpcmyw-defaultStyle > button:nth-child(11)", chromedp.ByQueryAll),
        //chromedp.Sleep(time.Second*5),
        //chromedp.OuterHTML("html", &data2, chromedp.ByQuery),
    ); err != nil {
        Loge("chromedp run failed")
    }
    for i := 0; i < len(data); i++ {
        var doc, _ = goquery.NewDocumentFromReader(strings.NewReader(data[i]))
        temp0 := doc.Find("div")
        temp0.Each(func(i int, s0 *goquery.Selection) {
            class, _ := s0.Attr("class")
            if class == divContainerClass {
                temp := s0.Find(splitElem)
                temp.Each(func(i int, s1 *goquery.Selection) {
                    class2, _ := s1.Attr("class")
                    if class2 == splitElemClass {
                        temp1 := s1.Find("div")
                        temp1.Each(func(i int, sImage *goquery.Selection) {
                            classImage, _ := sImage.Attr("class")
                            if classImage == divImageclass {
                                tempImage := sImage.Find("img")
                                tempImage.Each(func(i int, sImage2 *goquery.Selection) {
                                    imgSrc, exists := sImage2.Attr("src")
                                    if exists {
                                        //fmt.Println(imgSrc)
                                        imageSrc = append(imageSrc, imgSrc)
                                    }
                                })
                            }
                        })

                        temp2 := s1.Find("a")
                        temp2.Each(func(i int, s2 *goquery.Selection) {
                            class3, _ := s2.Attr("class")
                            if class3 == aTitleclass {
                                temp3 := s2.Find(titleElem)
                                temp3.Each(func(i int, s3 *goquery.Selection) {
                                    class4, _ := s3.Attr("class")
                                    if class4 == titleClass {
                                        //fmt.Printf("%s\n", strings.TrimSpace(s3.Text()))
                                        name = append(name, s3.Text())
                                    }
                                })
                            }
                        })
                        temp3 := s1.Find(priceElem)
                        tempPrice := math.MaxInt
                        var tempPriceString string
                        temp3.Each(func(i int, s3 *goquery.Selection) {
                            class3, _ := s3.Attr("class")
                            if class3 == priceClass {
                                //fmt.Printf("%s\n", strings.Trim(s3.Text(), ","))
                                intVar, _ := strconv.Atoi(strings.Replace(s3.Text(), ",", "", -1))
                                if intVar <= 100 {
                                    return
                                }
                                if intVar < tempPrice {
                                    tempPrice = intVar
                                    tempPriceString = s3.Text()
                                }
                            }
                        })
                        //fmt.Printf("%s\n", tempPrice)
                        price = append(price, tempPriceString)
                    }
                })
            }
        })
    }

    tempUrl, err := libUrl.Parse(url)
    if err != nil {
        Loge("url parse failed")
    }
    parts := strings.Split(tempUrl.Hostname(), ".")
    //domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
    //fmt.Println(parts[len(parts)-2])

    currentTime := time.Now()
    currentDate := currentTime.Format("2006-01-02")

    for i := 0; i < len(name); i++ {
        tempPrice, _ := strconv.Atoi(strings.Replace(price[i], ",", "", -1))
        r, _ := regexp.Compile("(([0-9]*[.])?[0-9]+(g|kg|G|Kg|kG|KG))")

        weight := r.FindString(name[i])
        weight = strings.Replace(weight, "KG", "", -1)
        weight = strings.Replace(weight, "kG", "", -1)
        weight = strings.Replace(weight, "Kg", "", -1)
        weight = strings.Replace(weight, "kg", "", -1)
        weight = strings.Replace(weight, "G", "", -1)
        weight = strings.Replace(weight, "g", "", -1)
        tempWeight, _ := strconv.ParseFloat(weight, 64)

        r, _ = regexp.Compile("(([0-9]?[-])?[0-9]+(개|봉|입|과))")

        cnt := r.FindString(name[i])

        r, _ = regexp.Compile("(g|kg|G|Kg|kG|KG)")

        unit := r.FindString(name[i])

        r, _ = regexp.Compile("(([0-9]+[+][0-9]+))")

        bundle := r.FindString(name[i])

        temp := &itemdb.ItemInfo{
            Category: item,
            Date:     currentDate,
            Name:     name[i],
            Price:    tempPrice,
            SiteName: parts[1],
            Weight:   tempWeight,
            Cnt:      cnt,
            Unit:     unit,
            Bundle:   bundle,
            ImageSrc: imageSrc[i],
        }
        itemdb.InsertNewItem(temp)
        Logd("%v\n", temp)
    }

    return nil
}
