package common

import (
    "time"
    "gopkg.in/mgo.v2"
)

var session *mgo.Session

func GetSession() *mgo.Session {
    if session == nil {
        createDbSession()
    }
    return session
}

func createDbSession() {
    var err error
    session, err = mgo.DialWithInfo(*mgo.DialInfo{
        Addrs: []string{AppConfig.MongoDBHost},
        Username: AppConfig.DBUser,
        Password: AppConfig.DBPwd,
        Timeout: 30 * time.Second,
    })
    if err != nil {
        log.Fatal("[GetSession]:", err)
    }
}

func addIndexes() {
    var err error
    userIndex := mgo.Index{
        Key: []string{"email"},
        Unique: true,
        Background: true,
        Sparse: true,
    }
    taskIndex := mgo.Index{
        Key: []string{"createdby"},
        Unique: false,
        Background: true,
        Sparse: true,
    }
    noteIndex := mgo.Index{
        Key: []string{"taskid"},
        Unique: false,
        Background: true,
        Sparse: true,
    }
    session := GetSession().Copy()
    defer session.Close()
    db := session.DB(AppConfig.Database)
    userCol := db.C("users")
    taskCol := db.C("tasks")
    noteCol := db.C("notes")
    err = userCol.EnsureIndex(userIndex)
    if err != nil {
        log.Fatal(err)
    }
    err = taskCol.EnsureIndex(taskIndex)
    if err != nil {
        log.Fatal(err)
    }
    err = noteCol.EnsureIndex(noteIndex)
    if err != nil {
        log.Fatal(err)
    }
}
