// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

import (
	"fmt"
	"log"
	"net/url"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func Test_Post(t *testing.T) {
	Post("https://mp.weixin.qq.com/s/e_0BhJ0vOvIqt610MjgKLA")
	Post("https://mp.weixin.qq.com/s/lu9yU_Wvd8yQs6alZ7OBqg")
	Post("https://mp.weixin.qq.com/s/widHAGN7zOymivIUihkuTw")
	Post("https://mp.weixin.qq.com/s/wXOhpZHdzlHlf5jgC54AaA")
}

func Test_Hot(t *testing.T) {
	articles, _ := Hot()
	fmt.Println(articles)
}

func Test_New(t *testing.T) {
	articles, _ := New()
	fmt.Println(articles)
}

func Test_PostLink(t *testing.T) {
	link := url.QueryEscape("https://mp.weixin.qq.com/s?__biz=MjM5NzgzNTUyNA==&mid=2650373560&idx=1&sn=84964bd5ce47084ae175806a4bf279da&chksm=bede3bd389a9b2c540644ba0161a5ece8f5c84b3088da9eccc3d8d0712e21cf31f6a637b4940#rd")
	doc, err := goquery.NewDocument(fmt.Sprintf("http://localhost:8888/post?url=%v", link))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc.Html())
}
