// models/user.go

// Package models
package models

// User โครงสร้างข้อมูลสำหรับผู้ใช้
type User struct {
	ID        int    `json:"id"`         // รหัสผู้ใช้
	Username  string `json:"username"`   // ชื่อผู้ใช้
	Password  string `json:"password"`   // รหัสผ่าน
	Email     string `json:"email"`      // อีเมลผู้ใช้
	FullName  string `json:"full_name"`  // ชื่อเต็มของผู้ใช้
	Role      string `json:"role"`       // บทบาทของผู้ใช้ (เช่น admin, user)
	CreatedAt string `json:"created_at"` // วันที่สร้างผู้ใช้
	UpdatedAt string `json:"updated_at"` // วันที่อัปเดตผู้ใช้ล่าสุด
}
