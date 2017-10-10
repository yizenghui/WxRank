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
	articles, _ := Hot(5, 0)
	fmt.Println(articles)
}

func Test_New(t *testing.T) {
	articles, _ := New(5, 0)
	fmt.Println(articles)
}

func Test_PostLink(t *testing.T) {
	// link := url.QueryEscape("https://mp.weixin.qq.com/s?__biz=MjM5NzgzNTUyNA==&mid=2650373560&idx=1&sn=84964bd5ce47084ae175806a4bf279da&chksm=bede3bd389a9b2c540644ba0161a5ece8f5c84b3088da9eccc3d8d0712e21cf31f6a637b4940#rd")
	link := url.QueryEscape("http://mp.weixin.qq.com/s?__biz=MzI4ODU0Nzk0NA==&amp;mid=2247484254&amp;idx=2&amp;sn=83cec43c784e699a827a3c9987c9171c&amp;chksm=ec3df659db4a7f4fa457a95daa77a07534e3172902ca27f39ed7a91d9170d260073ace202325&amp;mpshare=1&amp;scene=1&amp;srcid=1010Yx7oukwdLbpQ9QY31fwU#rd")
	doc, err := goquery.NewDocument(fmt.Sprintf("http://localhost:8888/fetch?url=%v", link))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc.Html())
}
