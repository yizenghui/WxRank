package repository

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
	"time"

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

// Insert wechat.Article
func Insert(url string) (article wechat.Article, err error) {

	article, err = wechat.Find(url)
	return
}

//Hot ..
func Hot(limit, offset int) (articles []orm.Article, err error) {

	var a orm.Article
	articles = a.Hot(limit, offset)
	for key, article := range articles {
		articles[key].Cover = "http://pic3.readfollow.com/" + base64.URLEncoding.EncodeToString([]byte(article.Cover))
		// articles[key].URL = strings.Replace(article.URL, `http://`, "https://", -1)
		articles[key].URL = strings.Replace(article.URL, `#rd`, "&scene=27#wechat_redirect", 1)

	}

	return
}

//New ..
func New(limit, offset int) (articles []orm.Article, err error) {

	var a orm.Article
	articles = a.New(limit, offset)
	for key, article := range articles {
		articles[key].Cover = "http://pic3.readfollow.com/" + base64.URLEncoding.EncodeToString([]byte(article.Cover))
		// articles[key].URL = strings.Replace(article.URL, `http://`, "https://", -1)
		articles[key].URL = strings.Replace(article.URL, `#rd`, "&scene=27#wechat_redirect", 1)
		//
	}
	return
}

//Like ..
func Like(id int) (a orm.Article, err error) {

	// var a orm.Article
	a.GetArticleByID(id)

	if a.Title == "" {
		err = errors.New("内容异常")
		return
	}

	if a.Like < 1 {
		a.Like = 1
	} else {
		a.Like++
	}

	if a.ID != 0 {
		a.Rank = Rank(int(a.Like), int(a.Hate), a.PubAt.Unix())
		a.Save()
	}

	a.Cover = "http://pic3.readfollow.com/" + base64.URLEncoding.EncodeToString([]byte(a.Cover))

	return
}

//Hate ..
func Hate(id int) (a orm.Article, err error) {

	// var a orm.Article
	a.GetArticleByID(int(id))

	if a.Title == "" {
		err = errors.New("内容异常")
		return
	}

	if a.Hate < 1 {
		a.Hate = 1
	} else {
		a.Hate++
	}

	if a.ID != 0 {
		a.Rank = Rank(int(a.Like), int(a.Hate), a.PubAt.Unix())
		a.Save()
	}

	a.Cover = "http://pic3.readfollow.com/" + base64.URLEncoding.EncodeToString([]byte(a.Cover))

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

		if article.URL == "" {
			return errors.New("不支持该链接！")
		}

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

		i64, err := strconv.ParseInt(article.PubAt, 10, 64)
		if err != nil {
			// fmt.Println(err)
			return errors.New("时间转化失败")
		}
		a.PubAt = time.Unix(i64, 0)
		a.Save()
		// fmt.Println(a)
	}
	// }
	return
}

// Remove wechat.Article
func Remove(url string) (article wechat.Article, err error) {

	article, err = wechat.Find(url)
	return
}
