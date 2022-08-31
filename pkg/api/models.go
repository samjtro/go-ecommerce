package api

import (
	"errors"
	"math"

	"github.com/Kamva/mgm"
	"github.com/dgrijalva/jwt-go"
)

// TokenClaims is used to create JWT token
type TokenClaims struct {
	UserID string
	jwt.StandardClaims
}

// User is used to store user information
type User struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string `json:"email"`
	Password         string `json:"-"`
	GoogleUserID     string `json:"-" bson:"google_user_id"`
}

// Product is a sellable element
type Product struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string  `json:"title"`
	Price            float64 `json:"price"`
}

func (p Product) validate() error {
	if len(p.Title) < 5 {
		return errors.New("the title has to have at leat 5 characters")
	}

	if p.Price != math.Round(p.Price*100)/100 {
		return errors.New("the price has to many decimal places")
	}

	return nil
}
