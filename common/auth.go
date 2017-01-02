package common

import (
    "log"
    "time"
    "io/ioutil"
    "net/http"

    jwt "github.com/dgrijalva/jwt-go"
    "github.com/dgrijalva/jwt-go/request"
)

var (
    verifyKey, signKey []byte
)

type UserInfo struct {
    Name, Role string
}

func initKeys() {
    var err error
    signKey, err = ioutil.ReadFile("keys/app.rsa")
    if err != nil {
        log.Fatal(err)
    }
    verifyKey, err = ioutil.ReadFile("keys/app.rsa.pub")
    if err != nil {
        log.Fatal(err)
    }
}

func GenerateJWT(name, role string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "iss": "admin",
        "UserInfo": UserInfo{name, role},
        "exp": time.Now().Add(20 * time.Minute).Unix(),
    })
    tokenString, err := token.SignedString(signKey)
    if err != nil {
        return "", err
    } else {
        return tokenString, nil
    }
}

// Authorization middleware
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
        return verifyKey, nil
    })
    if err != nil {
        switch err.(type) {
        case *jwt.ValidationError:
            vErr := err.(*jwt.ValidationError)
            switch vErr.Errors {
            case jwt.ValidationErrorExpired:
                DisplayAppError(w, err, "Access Token is expired, get a new Token", 401)
                return
            default:
                DisplayAppError(w, err, "Error wile parsing the Access Token", 500)
                return
            }
        default:
            DisplayAppError(w, err, "Error while parsing the Access Token", 500)
            return
        }
    } else if token.Valid {
        next(w,r)
    } else {
        DisplayAppError(w, err, "Invalid Access Token", 401)
    }
}
