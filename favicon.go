package getfavicon

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	url2 "net/url"
)

func Get(rawURL string) (string, error) {
	_url, err := url2.Parse(rawURL)
	if err != nil {
		return "", err
	}

	defaultIcon := _url.Scheme + "://" + _url.Host + "/favicon.ico"
	if !testIcon(defaultIcon) {
		return FromUrl(rawURL)
	}

	return defaultIcon, err
}

func FromHtml(html string) (string, error) {
	htmlBuf := io.NopCloser(bytes.NewBuffer([]byte(html)))
	doc, err := goquery.NewDocumentFromReader(htmlBuf)
	if err != nil {
		return "", err
	}

	fav := ""

	doc.Find("head link[rel*=icon]").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		fav, _ = selection.Attr("href")
		return true
	})
	return fav, nil
}

func FromUrl(url string) (string, error) {
	html, err := loadURL(url)
	if err != nil {
		return "", err
	}
	fav, err := FromHtml(html)
	return fav, err
}

func loadURL(url string) (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func testIcon(url string) bool {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false
	}

	//if resp.Header["content_type"] image/x

	return true
}
