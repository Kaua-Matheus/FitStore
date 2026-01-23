package handler

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kaua-Matheus/fitstore/backend/model/entitie"
)

// User
func GetUserById(db *gorm.DB, id uuid.UUID) (entitie.User, error) {
	
	user := entitie.User{}
	result := db.Where("id_user = ?", id).Find(&user); if result.Error != nil {
		return entitie.User{}, fmt.Errorf("error trying to get the user");
	} else {
		return user, nil;
	}
}

func GetUserByLogin(db *gorm.DB, user_login string) (entitie.User, error) {
	
	user := entitie.User{}
	result := db.Where("user_login = ?", user_login).Find(&user); if result.Error != nil {
		return entitie.User{}, fmt.Errorf("error trying to get the user");
	} else {
		return user, nil;
	}
}

func AddUser(db *gorm.DB, user entitie.User) (error) {

	result := db.Create(&user); if result.Error != nil {
		return fmt.Errorf("error trying to add the user %s", result.Error);
	} else {
		return nil;
	}

}