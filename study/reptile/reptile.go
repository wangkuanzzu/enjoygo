package reptile

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func SimpleReptile() {

	resp, err := http.Get("https://studygolang.com/static/pkgdoc/pkg/net_http.htm")
	if err != nil {
		fmt.Println(err)
		resp.Body.Close()
		return
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)

	re := regexp.MustCompile(`<meta name="private:description" content="(.*)">`)
	b := re.FindSubmatch(bytes)[1]

	fmt.Println(string(b))
}
