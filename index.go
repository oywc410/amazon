package main
 
import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"regexp"
	"time"
	"strconv"
)

var digitsRegexp = regexp.MustCompile(`dp/(.*)\/ref`)

var chanM chan int

func main() {
	t1 := time.Now()
	chanM = make(chan int, 100)

	for i:=1; i < 401; i++ {
		go getAllData("http://www.amazon.co.jp/b/?node=2421961051&page=" + strconv.Itoa(i), i)
	}

	j := 0

	LOOP:
	for {
		select {
		case <-chanM:
			j++
			if j == 400 {
				break LOOP
			}
		}
	}

	fmt.Println(time.Now().Sub(t1))
}

func getAllData(link string, i int) {



	fmt.Println(link)
	doc, err := goquery.NewDocument(link)


	if err != nil {
		fmt.Println(err)
	}

	j := 0

	doc.Find(".s-result-item").Each(func(i int, s *goquery.Selection) {
		j++
		linkHref, exists := s.Find("a").Attr("href")
		if !exists {
			fmt.Println("no href")
		}

		//fmt.Println(linkHref)
		all := digitsRegexp.FindStringSubmatch(linkHref)
		fmt.Println(all[1])
	})

	chanM <- i

}
