package api

import (
    "blog/pkg/e"
    "blog/pkg/utils"
    "blog/models"

    "github.com/gin-gonic/gin"
    "github.com/astaxie/beego/validation"
    
    "log"
    "net/http"
)

type auth struct {
    Username string `json:"username" valid:"Required; MaxSize(50)"`
    Password string `json:"password" valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
    var loginAuth auth
    c.BindJSON(&loginAuth)

    valid := validation.Validation{}
    // a := auth{Username: username, Password: password}
    ok, _ := valid.Valid(&loginAuth)

    data := make(map[string]interface{})
    code := e.INVALID_PARAMS
    if ok {
        isExist := models.CheckAuth(loginAuth.Username, loginAuth.Password)
        if isExist {
            token, err := util.GenerateToken(loginAuth.Username, loginAuth.Password)
            if err != nil {
                code = e.ERROR_AUTH_TOKEN
            } else {
                data["token"] = token
                
                code = e.SUCCESS
            }

        } else {
            code = e.ERROR_AUTH
        }
    } else {
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
        // "test" : loginAuth.Username,
    })
}