package admin

import(
	"log"
	"os"

	"github.com/zxcas321/ProfileGolang/config"
	"github.com/zxcas321/ProfileGolang/models"
	"golang.org/x/crypto/bcrypt"
)

func AdminBoostrap() {
	var count int64

	if err := config.DB.Model(&models.Admin{}).Count(&count).Error ; err != nil {
		log.Fatal("Failed to count admins:", err)
	}

	if count > 0 {
		log.Println("Admin already exists, bootstrap skipped")
		return
	}

	email := os.Getenv("BOOTSTRAP_ADMIN_EMAIL")
	password := os.Getenv("BOOTSTRAP_ADMIN_PASSWORD")

	if email == "" || password == "" {
		log.Println("Bootstrap admin skipped: ENV not set")
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash admin password")
	}

	admin := models.Admin{
		Email:    email,
		Password: string(hashed),
	}

	if err := config.DB.Create(&admin).Error; err != nil {
		log.Fatal("Failed to create bootstrap admin:", err)
	}

	log.Println("Bootstrap admin created:", email)
}