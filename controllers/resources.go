package controllers

import (
    "github.com/gabsn/taskmanager/models"
)

type(
    // POST /user/register
    UserResource struct {
        Data models.User `json:"data"`
    }
    // POST /user/login
    LoginResource struct {
        Data LoginModel `json:"data"`
    }
    // Response for authorized user POST /user/register
    AuthUserResource struct {
        Data AuthUserModel `json:"data"`
    }
    // Model for authentication
    LoginModel struct {
        Email string `json:"email"`
        Password string `json:"password"`
    }
    // Model for authorized user with access token
    AuthUserModel struct {
        User models.User `json:"user"`
        Token string `json:"token"`
    }
)
