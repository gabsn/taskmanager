package common

import (
    "log"
    "io/ioutil"
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
