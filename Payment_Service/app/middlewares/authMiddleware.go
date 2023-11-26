package middlewares

import (
	"fmt"
	"strings"

	"github.com/Altair1618/Tubes_IF4031_03/Payment_Service/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func AuthMiddleware(c *fiber.Ctx) error {
	pubKeyRaw := viper.Get("RSA_PUBLIC_KEY").(string)
	pubKeyRSAForm, err := utils.DecodePublicKey(pubKeyRaw)
	if err != nil {
		responseBody := utils.ResponseBody{
			Code:    401,
			Message: "Unauthorized",
		}
		return utils.CreateResponseBody(c, responseBody)

	}

	authorizationHeader := c.Get("authorization")

	authorizationHeaderArray := strings.Split(authorizationHeader, " ")
	tokenType := authorizationHeaderArray[0]

	if len(authorizationHeaderArray) != 2 || tokenType != "Bearer" {
		responseBody := utils.ResponseBody{
			Code:    401,
			Message: "Unauthorized",
		}
		return utils.CreateResponseBody(c, responseBody)
	}

	token := authorizationHeaderArray[1]

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return pubKeyRSAForm, nil
	})
	if err != nil {
		responseBody := utils.ResponseBody{
			Code:    401,
			Message: "Unauthorized",
		}
		return utils.CreateResponseBody(c, responseBody)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		responseBody := utils.ResponseBody{
			Code:    401,
			Message: "Unauthorized",
		}
		return utils.CreateResponseBody(c, responseBody)
	}

	userData := make(map[string]string)
	userData["username"] = claims["username"].(string)
	userData["userId"] = claims["userId"].(string)

	c.Locals("userInfo", userData)

	return c.Next()
}
