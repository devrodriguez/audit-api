package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/devrodriguez/audit-api/models"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
)

// SignIn retorna un token de autenticacion
func SignIn(gCtx *gin.Context) {
	var response models.Response
	var user models.User

	// == VALIDATE USER AND PASSWORD ==
	// Get user data
	if err := gCtx.BindJSON(&user); err != nil {
		response.Message = "Error binding data"
		response.Error = err.Error()
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	if !ValidateUserAuth(&user) {
		response.Message = "Wrong user or password"

		gCtx.JSON(http.StatusOK, response)
		return
	}

	// == CREATE JWT TOKEN ==
	token, err := CreateToken(gCtx.Request)

	log.Println(string(token))

	if err != nil {
		response.Message = "Autenticacion fallida"
		response.Error = err.Error()

		gCtx.JSON(http.StatusOK, response)
		return
	}

	response.Data = gin.H{"token": string(token)}
	gCtx.JSON(http.StatusOK, response)
}

// Login valida el token de Authorization
func Login(gCtx *gin.Context) {
	var response models.Response
	req := gCtx.Request

	err := VerifyToken(req)

	if err != nil {
		response.Message = "¡Usuario no autorizado!"
		response.Error = err.Error()

		gCtx.JSON(http.StatusOK, response)
		return
	}

	response.Message = "Welcome"
	gCtx.JSON(http.StatusOK, response)
}

func CreateToken(r *http.Request) (string, error) {
	var secret = os.Getenv("JWT_SECRET")
	var hs = jwt.NewHS256([]byte(secret))
	now := time.Now()

	expiration, _ := time.ParseDuration(os.Getenv("EXPIRATION"))

	payload := models.JwtPayload{
		Payload: jwt.Payload{
			ExpirationTime: jwt.NumericDate(now.Add(expiration)),
			NotBefore:      jwt.NumericDate(now),
			IssuedAt:       jwt.NumericDate(now),
		},
	}

	token, err := jwt.Sign(payload, hs)

	log.Println(err)

	if err != nil {
		return "", err
	}

	return string(token), nil
}

func VerifyToken(r *http.Request) error {
	os.Setenv("JWT_SECRET", "dev1986")
	var secret = jwt.NewHS256([]byte(os.Getenv("JWT_SECRET")))
	var payload models.JwtPayload
	now := time.Now()

	token := []byte(r.Header.Get("Authorization"))

	expValidator := jwt.ExpirationTimeValidator(now)
	nbfValidator := jwt.NotBeforeValidator(now)
	validatePayload := jwt.ValidatePayload(&payload.Payload, expValidator, nbfValidator)

	hd, err := jwt.Verify(token, secret, &payload, validatePayload)

	log.Println(hd)

	if err != nil {
		return err
	}

	return nil
}

func ValidateUserAuth(user *models.User) bool {
	log.Println(user)
	if user.Name == "john" && user.Password == "12345" {
		return true
	}

	return false
}
