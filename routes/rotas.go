package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusRibas/api-go-gin/controllers"
)

func Handle() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos",controllers.NovoAluno)
	r.GET("/alunos/:id",controllers.Buscar)
	r.DELETE("/alunos/:id", controllers.Delete)
	r.PATCH("/alunos/:id", controllers.Editar)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaCPF)
	r.GET("/alunos/", controllers.BuscaRG)
	r.Run()
}
