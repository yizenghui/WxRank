package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yizenghui/WxRank/orm"
	"github.com/yizenghui/WxRank/repository"
)

// 接入微信接口服务
func echoWxCallbackHandler(c echo.Context) error {
	repository.WechatServe(c.Response().Writer, c.Request())
	var err error
	return err
}

// Template ..
type Template struct {
	templates *template.Template
}

// Render ..
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

//Home 主页
func Home(c echo.Context) error {
	articles, err := repository.New(10, 0)

	if err != nil {

	}

	type Data struct {
		Title    string
		Articles []orm.Article
	}

	data := Data{
		Title:    "tttiii",
		Articles: articles,
	}

	return c.Render(http.StatusOK, "home", data)
}

//Hot 热门数据
func Hot(c echo.Context) error {
	// fmt.Println(mapWhere)

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	if limit <= 0 || limit > 100 {
		limit = 10
	}
	// limit = 10
	if offset < 0 || offset > 500 {
		offset = 0
	}

	articles, err := repository.Hot(limit, offset)

	if err != nil {

	}

	return c.JSON(http.StatusOK, articles)
}

//New 最新数据
func New(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	if limit <= 0 || limit > 100 {
		limit = 10
	}
	// limit = 10
	if offset < 0 || offset > 500 {
		offset = 0
	}

	articles, err := repository.New(limit, offset)

	if err != nil {

	}

	return c.JSON(http.StatusOK, articles)
}

//Hate 不喜欢
func Hate(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	article, err := repository.Hate(id)

	if err != nil {

	}

	return c.JSON(http.StatusOK, article)
}

//Like 喜欢
func Like(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	article, err := repository.Like(id)

	if err != nil {

	}

	return c.JSON(http.StatusOK, article)
}

//Post 报料接口
func Post(c echo.Context) error {
	url := c.QueryParam("url")
	fmt.Println(url)
	if url != "" {
		repository.Post(url)
		return c.JSON(http.StatusOK, "Post")
	}
	return c.JSON(http.StatusOK, "0")
}

func main() {
	orm.DB().AutoMigrate(&orm.User{})

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.Use(middleware.CORS())

	// e.Pre(middleware.HTTPSRedirect())
	// e.Pre(middleware.HTTPSNonWWWRedirect())
	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", Home)

	e.GET("/post", Post)

	e.GET("/new", New)
	e.POST("/new", New)

	e.GET("/hate/:id", Hate)
	e.GET("/like/:id", Like)

	e.GET("/hot", Hot)

	e.File("/favicon.ico", "images/favicon.ico")

	e.Any("/wx_callback", echoWxCallbackHandler)

	// e.Static("/", "src")
	// Start server
	// e.Logger.Fatal(e.Start(":8888"))
	e.Logger.Fatal(e.Start(":8888"))
	// e.Logger.Fatal(e.StartAutoTLS(":443"))

}
