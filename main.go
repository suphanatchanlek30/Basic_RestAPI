package main

import (
	"github.com/gofiber/fiber/v2"

	// นำเข้าแพ็กเกจ controllers และ services
	"fiber-restapi/controllers"
	"fiber-restapi/services"
)

func main() {
	// สร้างแอปพลิเคชัน Fiber ใหม่
	app := fiber.New()

	// กำหนดเส้นทางสำหรับคำขอ GET ที่ราก "/"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// สร้าง service และ controller สำหรับผู้ใช้
	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	// กำหนดเส้นทางสำหรับคำขอ POST ที่ "/register" เพื่อสมัครผู้ใช้ใหม่
	app.Post("/register", userController.Register)

	// กำหนดเส้นทางสำหรับคำขอ POST ที่ "/login" เพื่อเข้าสู่ระบบผู้ใช้
	app.Post("/login", userController.Login)

	// เริ่มต้นเซิร์ฟเวอร์ที่พอร์ต 3000
	app.Listen(":3000")
}
