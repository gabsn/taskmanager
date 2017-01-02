package common

import (
    "log"
    "io/ioutil"

    jwt "github.com/dfrijalva/jwt-go"
)

var (
    verifyKey, signKey []byte
)

func initKeys() {
    var err error
    signKey, err = ioutil.ReadFile("keys/app.rsa")
    if err != nil {
        log.Fatal(err)
    }
    privKey, err = ioutil.ReadFile("keys/app.rsa.pub")
    if err != nil {
        log.Fatal(err)
    }
}

func GenerateJWT(name, role string) (string, error) {
    t := jwt.New(jwt.GetSigningMethod("RS256"))
    t.Claims["iss"] = "admin"
    t.Claims["UserInfo"] = struct {
        Name, Role string
    }{name, role}
    t.Claims["exp"] = time.Now().Add(20 * time.Minute).Unix()
    tokenString, err = t.SignedString(signKey)
    if err != nil {
        return "", err
    } else {
        return tokenString, nil
    }
}

// Authorization middleware
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
        return verifyKey, nil
    }
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
