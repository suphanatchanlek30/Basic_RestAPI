package controllers

import (
	"fiber-restapi/models"
	"fiber-restapi/services"

	"github.com/gofiber/fiber/v2"
)

// สร้าง UserController เป็น struct โครงสร้างข้อมูลสำหรับตัวควบคุมผู้ใช้
type UserController struct {
	userService *services.UserService // ตัวชี้ไปยังบริการผู้ใช้
}

// สร้าง NewUserController ฟังก์ชันสร้างตัวควบคุมผู้ใช้ใหม่
// คืนค่าตัวชี้ไปยัง UserController ที่สร้างขึ้นใหม่
func NewUserController(userService *services.UserService) *UserController {
	// คืนค่าตัวชี้ไปยัง UserController ที่มี userService ที่ส่งเข้ามา
	return &UserController{
		// กําหนดค่าเริ่มต้นของ userService
		userService: userService,
	}
}

// สร้างฟังก์ชัน Register เพื่อจัดการคำขอสมัครผู้ใช้ใหม่
func (crtl *UserController) Register(c *fiber.Ctx) error {

	// สร้างตัวแปร user เพื่อเก็บข้อมูลผู้ใช้ที่รับเข้ามา
	user := new(models.User)

	// เช็คว่ามีข้อมูลผู้ใช้ใน request body หรือไม่
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "ข้อมูลไม่ถูกต้อง",
			"success": false,
			"error":   "Invalid request",
		})
	}

	// เรียกใช้บริการผู้ใช้เพื่อสมัครผู้ใช้ใหม่
	registerUser, err := crtl.userService.Register(*user)

	// ถ้ามีข้อผิดพลาดในการสมัครผู้ใช้
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "สมัครผู้ใช้ไม่สำเร็จ",
			"success": false,
			"error":   err.Error(),
		})
	}

	// ถ้าสมัครผู้ใช้สําเร็จ
	return c.Status(200).JSON(fiber.Map{
		"message": "สมัครผู้ใช้สําเร็จ",
		"success": true,
		"user": fiber.Map{
			"id":         registerUser.ID,
			"username":   registerUser.Username,
			"email":      registerUser.Email,
			"full_name":  registerUser.FullName,
			"role":       registerUser.Role,
			"created_at": registerUser.CreatedAt,
			"updated_at": registerUser.UpdatedAt,
		},
	})
}

// สร้างฟังก์ชัน Login เพื่อจัดการคำขอเข้าสู่ระบบผู้ใช้
func (crtl *UserController) Login(c *fiber.Ctx) error {
	// สร้างตัวแปร user เพื่อเก็บข้อมูลผู้ใช้ที่รับเข้ามา
	user := new(models.User)

	// เช็คว่ามีข้อมูลผู้ใช้ใน request body หรือไม่
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "ข้อมูลไม่ถูกต้อง",
			"success": false,
			"error":   "Invalid request",
		})
	}

	// เรียกใช้บริการผู้ใช้เพื่อเข้าสู่ระบบ
	loggedInUser, err := crtl.userService.Login(user.Username, user.Password)

	// ถ้าเข้าสู่ระบบไม่สําเร็จ
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "เข้าสู่ระบบไม่สําเร็จ",
			"success": false,
			"error":   err.Error(),
		})
	}

	// ถ้าเข้าสู่ระบบสําเร็จ
	return c.Status(200).JSON(fiber.Map{
		"message": "เข้าสู่ระบบสําเร็จ",
		"success": true,
		"user": fiber.Map{
			"id":         loggedInUser.ID,
			"username":   loggedInUser.Username,
			"email":      loggedInUser.Email,
			"full_name":  loggedInUser.FullName,
			"role":       loggedInUser.Role,
			"created_at": loggedInUser.CreatedAt,
			"updated_at": loggedInUser.UpdatedAt,
		},
	})
}
