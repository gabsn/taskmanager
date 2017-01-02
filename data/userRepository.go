package data

import (
    "github.com/gabsn/taskmanager/models"
    "golang.org/x/crypto/bcrypt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
    C *mgo.Collection
}

func (r *UserRepository) CreateUser(user *models.User) (err error) {
    user.Id = bson.NewObjectId()
    user.HashPassword, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    user.Password = ""
    err = r.C.Insert(&user)
    return
}

func (r *UserRepository) Login(user models.User) (u models.User, err error) {
    err = r.C.Find(bson.M{"email": user.Email}).One(&u)
    if err != nil {
        return
    }
    err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
    if err != nil {
        u = models.User{}
    }
    return
}
