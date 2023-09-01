package main

import (
	"fmt"
	"github.com/trumanwong/cryptogo"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func generateCookie() string {
	mUID := cryptogo.MD5ToUpper(time.Now().Format(time.DateTime))
	gUID := cryptogo.MD5ToUpper(mUID)
	dOB := time.Now().Format("20060102150405")
	sID := cryptogo.MD5ToUpper(dOB)
	result := "MUID=" + mUID + "; SNRHOP=I=&TS=; SRCHD=AF=NOFORM; SRCHHPGUSR=SRCHLANG=zh-Hans; SRCHUID=V=2&GUID=" + gUID + "&dmnchg=1; SRCHUSR=DOB=" + dOB + "; SUID=M; _EDGE_S=F=1&SID=" + sID + "; _EDGE_V=1; _SS=SID=" + sID + "; MUIDB=" + mUID
	return result
}

func main() {
	link := "https://cn.bing.com/search?q=" + url.QueryEscape("今天深圳天气")
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, link, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", generateCookie())
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.200")
	req.Header.Add("Referer", "https://developer.microsoft.com/")
	req.Header.Add("Sec-Ch-Ua", "\"Not/A)Brand\";v=\"99\", \"Microsoft Edge\";v=\"115\", \"Chromium\";v=\"115\"")
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Add("Sec-Ch-Ua-Platform-Version", "\"15.0.0\"")
	req.Header.Add("Sec-Ch-Ua-Arch", "\"x64\"")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "cross-site")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(body))
}
