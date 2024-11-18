package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Estrutura para representar uma atividade
type Activity struct {
	ID          int       json:"id"
	Title       string    json:"title"
	Description string    json:"description"
	Date        time.Time json:"date"
}

var activities []Activity
var nextID int = 1

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// Rota principal para exibir as atividades agendadas
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"activities": activities})
	})

	// Rota para adicionar uma nova atividade
	r.GET("/add", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add.html", nil)
	})

	// Rota para processar a criação de uma nova atividade
	r.POST("/add", func(c *gin.Context) {
		var newActivity Activity
		newActivity.ID = nextID
		newActivity.Title = c.PostForm("title")
		newActivity.Description = c.PostForm("description")
		date, _ := time.Parse("2006-01-02", c.PostForm("date"))
		newActivity.Date = date

		activities = append(activities, newActivity)
		nextID++

		c.Redirect(http.StatusSeeOther, "/")
	})

	r.Run(":8080") // Inicia o servidor na porta 8080
}