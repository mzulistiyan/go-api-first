package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/mzulistiyan/go-api-first/app"
	"github.com/mzulistiyan/go-api-first/controller"
	"github.com/mzulistiyan/go-api-first/helper"
	"github.com/mzulistiyan/go-api-first/middleware"
	"github.com/mzulistiyan/go-api-first/repository"
	"github.com/mzulistiyan/go-api-first/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
