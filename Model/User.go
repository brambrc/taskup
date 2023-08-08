package Model

import (
	"fmt"
	"os"
	"taskup/Database"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required" gorm:"unique"`
	Name     string `form:"name" json:"name" binding:"required"`
	Role     string `form:"role" json:"role" binding:"required"`
}

type Claims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func (u *User) BeforeSave(*gorm.DB) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Save() (*User, error) {
	var err error

	// check duplicate username
	err = Database.Database.Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func GetAuthenticatedID(tokenString string) (uint, error) {

	//get token from headers

	var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

	onlyToken := tokenString[len("Bearer "):]

	fmt.Println("onlyToken", onlyToken)
	token, _ := jwt.Parse(onlyToken, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})

	var claims = token.Claims.(jwt.MapClaims)
	var id = claims["id"]

	fmt.Println("claims id", claims["id"])

	return uint(id.(float64)), nil

}

func (u *User) Update(id uint) (*User, error) {
	var err error

	updates := map[string]interface{}{
		"email": u.Email,
		"name":  u.Name,
		"role":  u.Role,
	}

	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		updates["password"] = string(hashedPassword)
	}

	if err = Database.Database.Model(&User{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByEmail(email string) (User, error) {
	var user User
	err := Database.Database.Where("email = ?", email).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
