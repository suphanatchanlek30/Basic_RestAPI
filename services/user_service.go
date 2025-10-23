// services/user_service.go

package services

import (
	// นำเข้าแพ็กเกจ models
	"fiber-restapi/models"
	"fmt"
	"time"
)

// UserService โครงสร้างข้อมูลสำหรับบริการผู้ใช้
type UserService struct {
	users  []models.User // สไลซ์เก็บข้อมูลผู้ใช้ จำลองฐานข้อมูล
	nextID int           // สร้างเลข id ถัดไป รหัสผู้ใช้ถัดไป
}

// NewUserService ฟังก์ชันสร้างบริการผู้ใช้ใหม่
// คืนค่าตัวชี้ไปยัง UserService ที่สร้างขึ้นใหม่
func NewUserService() *UserService {
	// คืนค่าตัวชี้ไปยัง UserService ที่มีสไลซ์ผู้ใช้ว่างเปล่าและ nextID เริ่มต้นที่ 1
	return &UserService{
		// กำหนดค่าเริ่มต้นของ users และ nextID
		users:  []models.User{},
		nextID: 1,
	}
}

// สร้างฟังก์ชัน Register เพื่อสร้างผู้ใช้ใหม่
// อธิบาย func (s *UserService) Register(user models.User) (models.User, error)
// (s *UserService) หมายถึง ฟังก์ชัน Register เป็นเมธอดของโครงสร้าง UserService
// user models.User คือ พารามิเตอร์ที่รับเข้ามาเป็นโครงสร้าง User
// (models.User, error) คือ ค่าที่ฟังก์ชันนี้จะคืนค่าออกมา เป็นโครงสร้าง User และข้อผิดพลาด (ถ้ามี)
func (s *UserService) Register(user models.User) (models.User, error) {
	// ตรวจสอบว่าชื่อผู้ใช้มีอยู่แล้วหรือไม่
	for _, u := range s.users {
		// ถ้าชื่อผู้ใช้หรืออีเมลมีอยู่แล้ว ให้คืนค่าข้อผิดพลาด
		if u.Username == user.Username {
			return models.User{}, fmt.Errorf("username %s already exists", user.Username)
		}

		// ตรวจสอบอีเมลซ้ํา
		if u.Email == user.Email {
			return models.User{}, fmt.Errorf("email %s already exists", user.Email)
		}
	}

	// ถ้าไม่มีชื่อผู้ใช้ซ้ํา ให้สร้างผู้ใช้ใหม่
	user.ID = s.nextID
	// เพิ่ม nextID เพื่อเตรียมพร้อมสำหรับผู้ใช้ถัดไป
	s.nextID++
	// ตั้งค่าวันที่สร้างและอัปเดต
	user.CreatedAt = time.Now().Format(time.RFC3339) // ใช้รูปแบบ RFC3339 สำหรับวันที่ และเวลา เช่น "2023-10-05T14:48:00Z"
	user.UpdatedAt = time.Now().Format(time.RFC3339) // ใช้รูปแบบ RFC3339 สำหรับวันที่ และเวลา เช่น "2023-10-05T14:48:00Z"

	// กำหนดค่าเริ่มต้นให้กับบทบาทผู้ใช้
	if user.Role == "" {
		user.Role = "user"
	}

	// เพิ่มผู้ใช้ใหม่ลงในสไลซ์ผู้ใช้
	s.users = append(s.users, user)

	return user, nil
}

// สร้างฟังก์ชัน Login เพื่อยืนยันตัวตนของผู้ใช้
// อธิบาย func (s *UserService) Login(username, password string) (models.User, error)
// (s *UserService) หมายถึง ฟังก์ชัน Login เป็นเมธอดของโครงสร้าง UserService
// username, password string คือ พารามิเตอร์ที่รับเข้ามาเป็นชื่อผู้ใช้และรหัสผ่าน
// (models.User, error) คือ ค่าที่ฟังก์ชันนี้จะคืนค่าออกมา เป็นโครงสร้าง User และข้อผิดพลาด (ถ้ามี)
func (s *UserService) Login(username, password string) (models.User, error) {
	// ค้นหาผู้ใช้ตามชื่อผู้ใช้และรหัสผ่าน
	for _, u := range s.users {
		if u.Username == username && u.Password == password {
			return u, nil // คืนค่าผู้ใช้ถ้าพบ
		}
		// ยังไม่เพิ่ม token
	}

	return models.User{}, fmt.Errorf("invalid username or password") // คืนค่าข้อผิดพลาดถ้าไม่พบ
}
