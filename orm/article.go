package orm

import "time"

// Article has and belongs to many languages, use `Article_languages` as join table
type Article struct {
	ID        uint `gorm:"primary_key"`
	Title     string
	Author    string
	AppName   string
	AppID     string
	Cover     string
	Intro     string
	PubAt     time.Time
	Like      int64   `gorm:"type:default(0)"`
	Hate      int64   `gorm:"type:default(0)"`
	URL       string  `gorm:"type:varchar(100);unique_index"`
	Rank      float64 `sql:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// GetArticleByID 获取Article
func (article *Article) GetArticleByID(id int) {
	DB().First(article, id)
}

// GetArticleByURL 通过url获取Article 如果没有的话进行初始化 (注：此url由文章详细页获得)
func (article *Article) GetArticleByURL(url string) {
	DB().Where(Article{URL: url}).FirstOrCreate(article)
}

// Save 保存用户信息
func (article *Article) Save() {
	DB().Save(&article)
}

// Hot 热门
func (article *Article) Hot(limit, offset int) (articles []Article) {
	DB().Offset(offset).Limit(limit).Order("rank DESC").Find(&articles)
	return
}

// New 最新
func (article *Article) New(limit, offset int) (articles []Article) {
	DB().Offset(offset).Limit(limit).Order("id DESC").Find(&articles)
	return
}
