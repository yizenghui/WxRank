package repository

import (
	"encoding/base64"
	"fmt"

	"github.com/yizenghui/WxRank/orm"
	"github.com/yizenghui/sda/wechat"
)

func init() {
	orm.DB().AutoMigrate(&orm.Post{}, &orm.Article{})

}

// Find wechat.Article
func Find(url string) (article wechat.Article, err error) {

	article, err = wechat.Find(url)
	return
}

func update(url string) (article wechat.Article, err error) {
	article, err = wechat.Find(url)
	return
}

// Find wechat.Article
func Insert(url string) (article wechat.Article, err error) {

	article, err = wechat.Find(url)
	return
}

//Hot ..
func Hot(limit, offset int) (articles []orm.Article, err error) {

	var a orm.Article
	articles = a.Hot(limit, offset)
	for key, article := range articles {
		articles[key].Cover = "http://localhost:1323/" + base64.URLEncoding.EncodeToString([]byte(article.Cover))
	}

	return
}

//New ..
func New() (articles []orm.Article, err error) {

	var a orm.Article
	articles = a.New()
	for key, article := range articles {
		articles[key].Cover = "http://localhost:1323/" + base64.URLEncoding.EncodeToString([]byte(article.Cover))
	}
	return
}

// Post ... url
func Post(url string) (err error) {
	var post orm.Post
	post.GetPostByURL(url)
	// if post.State == 0 {
	var a orm.Article
	article, err := wechat.Find(url)
	if err == nil {
		post.ArticleURL = article.URL
		post.State = 1
		post.Save()
		a.GetArticleByURL(article.URL)
		a.AppID = article.AppID
		a.AppName = article.AppName
		a.Title = article.Title
		a.Intro = article.Intro
		a.Cover = article.Cover
		a.Author = article.Author
		a.PubAt = article.PubAt
		a.Save()
		fmt.Println(a)
	}
	// }
	return
}

// Find wechat.Article
func Remove(url string) (article wechat.Article, err error) {

	article, err = wechat.Find(url)
	return
}
