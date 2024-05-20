package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func CheckJWT(c *fiber.Ctx) error {
	bearerToken := c.Cookies("jwt")

	if bearerToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "token not found",
		})
	}

	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		// Убедитесь, что метод подписи в токене - ожидаемый
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Неожиданный метод подписи: %v", token.Header["alg"])
		}

		// Здесь должен быть ваш секретный ключ
		return jwtSecret, nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token format",
		})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		//fmt.Println("Токен валиден!")
		//fmt.Println("Claims:", claims)
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token format",
		})
	}
	c.Locals("userID", claims["id"].(string))

	return c.Next()
}

type Claims struct {
	jwt.Claims
	ID string `json:"id"`
}

func GetClaimsFromToken(tokenString string, secretKey string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secretKey), nil

	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}
