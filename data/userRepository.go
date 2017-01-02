package data

type UserRpository struct {
    C *mgo.Collection
}

func (r *UserRepository) CreateUser(user *models.User) error {
    user.id := bson.NewObjectId()
    user.HashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    user.password = ""
    err = r.C.insert(&user)
    return err
}

func (r *UserRepository) Login(user models.User) (u models.user, err error) {
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
