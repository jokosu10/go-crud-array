package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Article struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	Content     string `form:"content"`
}

var articles []Article

// method create
func createaArticle(c echo.Context) error {
	var article Article

	if err := c.Bind(&article); err != nil {
		return err
	}
	articles = append(articles, article)

	printArticles(articles)

	return c.NoContent(http.StatusCreated)
}

func printArticles(articles []Article) {
	for i, article := range articles {
		fmt.Printf("%d. %s\n", i+1, article.Title)
	}

	fmt.Printf("Total article: %d\n", len(articles))
}

func main() {
	e := echo.New()
	e.POST("/articles", createaArticle)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
