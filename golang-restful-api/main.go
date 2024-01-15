package main

import (
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/repository"
	"golang-restful-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	caetgoryController := controller.NewCategoryController(categoryService)
	router := httprouter.New()

	router.GET("/api/categories", caetgoryController.FindAll)
	router.GET("/api/categories/:categoryId", caetgoryController.FindById)
	router.POST("/api/categories", caetgoryController.Create)
	router.PUT("/api/categories/:categoryId", caetgoryController.Update)
	router.DELETE("/api/categories/:categoryId", caetgoryController.Delete)

}