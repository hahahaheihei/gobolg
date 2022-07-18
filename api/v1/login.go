package v1

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"go.mod/middleware"
	"go.mod/model"
	"go.mod/utils/errmsg"
)

// Login 后台登陆
func Login(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var token string
	var code int

	formData, code = model.CheckLogin(formData.UserName, formData.Password)

	if code == errmsg.SUCCSE {
		setToken(c, formData)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    formData.UserName,
			"id":      formData.ID,
			"message": errmsg.GetErrMsg(code),
			"token":   token,
		})
	}

}

// LoginFront 前台登录
func LoginFront(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var code int

	formData, code = model.CheckLoginFront(formData.UserName, formData.Password)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    formData.UserName,
		"id":      formData.ID,
		"message": errmsg.GetErrMsg(code),
	})
}

// token生成函数
func setToken(c *gin.Context, user model.User) {
	j := middleware.NewJWT()
	claims := middleware.MyClaims{
		Username: user.UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "GinBlog",
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
			"token":   token,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    user.UserName,
		"id":      user.ID,
		"message": errmsg.GetErrMsg(200),
		"token":   token,
	})
	return
}