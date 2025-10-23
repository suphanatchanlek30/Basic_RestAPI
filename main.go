package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// สร้างแอปพลิเคชัน Fiber ใหม่
	app := fiber.New()

	// กำหนดเส้นทางสำหรับคำขอ GET ที่ราก "/"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// เริ่มต้นเซิร์ฟเวอร์ที่พอร์ต 3000
	app.Listen(":3000")
}
