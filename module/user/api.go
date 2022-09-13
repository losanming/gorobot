package user

import (
	"errors"
	"example.com/m/module/db"
	"example.com/m/utils/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Controller struct{}

// 注册路由
func (self Controller) RegisterRoute(r *gin.RouterGroup) {
	r.POST("/user/WebLogin", response.API(self.WebLogin))       // 后台 - 账号密码登录
	r.POST("/user/WebRegister", response.API(self.WebRegister)) // 后台 - 账号密码登录
}

func (self *Controller) WebLogin(c *gin.Context) (data map[string]interface{}, err error) {
	var user db.User
	var user_login db.UserLogin
	user_login.UserName = c.PostForm("user_name")
	user_login.PassWord = c.PostForm("pass_word")

	if user_login.UserName == "" || user_login.PassWord == "" {
		return nil, errors.New("用户名或密码为空")
	}
	user.UserName = user_login.UserName
	result, err := user.FindUserInfoByUserName()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return data, errors.New("用户名错误")
		} else {
			return data, err
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(result.PassWord), []byte(user_login.PassWord))
	if err != nil {
		return nil, errors.New("密码错误")
	}
	data = make(map[string]interface{})
	data["msg"] = "ok"
	data["user_name"] = result.UserName
	data["user_id"] = result.Id
	return data, err
}

func (self *Controller) WebRegister(c *gin.Context) (data map[string]interface{}, err error) {
	var user_register db.UserRegister
	var user db.User
	user_register.UserName = c.PostForm("user_name")
	user_register.PassWord = c.PostForm("pass_word")

	if user_register.UserName == "" || user_register.PassWord == "" {
		return nil, errors.New("用户名或密码为空")
	}
	user.UserName = user_register.UserName
	hash, err := bcrypt.GenerateFromPassword([]byte(user_register.PassWord), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.PassWord = string(hash)
	err = user.Create()
	if err != nil {
		return data, err
	}
	data = make(map[string]interface{})
	data["msg"] = "ok"
	data["user_id"] = user.Id
	return data, err
}
