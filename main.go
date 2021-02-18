package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Article struct {
	Title       string `form:"title" json:"title"`
	Description string `form:"description" json:"description"`
	Content     string `form:"content" json:"content"`
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

// method print article in console
func printArticles(articles []Article) {
	for i, article := range articles {
		fmt.Printf("%d. %s\n", i+1, article.Title)
	}

	fmt.Printf("Total article: %d\n", len(articles))
}

// method print article
func showArticle(c echo.Context) error {
	articleId, err := strconv.Atoi(c.Param("id"))

	if len(articles) < articleId {
		return c.NoContent(http.StatusNotFound)
	}

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, c.JSON(http.StatusOK, articles[articleId]))
}

func main() {
	articles = make([]Article, 0)
	e := echo.New()
	e.POST("/articles", createaArticle)
	e.GET("/articles", func(c echo.Context) error {
		return c.JSON(http.StatusOK, articles)
	})
	e.GET("/articles/:id", showArticle)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
