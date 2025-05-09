package main

import (
	"Go-Template/internal/auth/domain"
	"Go-Template/internal/auth/handler"
	"Go-Template/internal/auth/repository"
	"Go-Template/internal/auth/usecase"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// โหลดไฟล์ .env หลายไฟล์
	err := godotenv.Load("config/.Database.env", "config/.Secret.env", "config/.Setting.env")
	if err != nil {
		log.Fatalf("Error loading .env files")
	}

	// แปลงค่า DB_PORT เป็นตัวเลข
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid port number: %s", os.Getenv("DB_PORT"))
	}

	// สร้าง DSN
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.User{})

	app := fiber.New()

	// สร้าง repository, usecase และ handler
	repo := repository.NewAuthGormRepository(db)
	uc := usecase.NewAuthUsecase(repo)
	handler := handler.NewAuthHandler(uc)

	// กำหนด routes พื้นฐาน
	handler.AuthRoutes(app)

	// ตัวอย่าง routes ที่ใช้ middleware
	// app.Get("/admin", middlewares.CheckRole("admin"), func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 		"message": "Welcome Admin!",
	// 	})
	// })

	// app.Get("/user", middlewares.CheckRole("user"), func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 		"message": "Welcome User!",
	// 	})
	// })

	app.Listen(":3000")
}
