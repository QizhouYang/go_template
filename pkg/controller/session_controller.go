package controller

import (
	"errors"
	"go_template/pkg/controller/kolog"
	"go_template/pkg/database"
	"go_template/pkg/dto"
	"go_template/pkg/model"
	"go_template/pkg/service"
	"time"

	"go_template/pkg/constant"

	"go_template/pkg/util/captcha"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12/context"
	"github.com/spf13/viper"
)

type SessionController struct {
	Ctx         *context.Context
	UserService service.UserService
}

func NewSessionController() *SessionController {
	return &SessionController{
		UserService: service.NewUserService(),
	}
}

func (s *SessionController) Post() (*dto.Profile, error) {
	aul := dto.LoginCredential{}
	if err := s.Ctx.ReadJSON(&aul); err != nil {
		return nil, err
	}
	if aul.CaptchaId != "" {
		err := captcha.VerifyCode(aul.CaptchaId, aul.Code)
		if err != nil {
			return nil, err
		}
	}
	var profile *dto.Profile
	switch aul.AuthMethod {
	case constant.AuthMethodJWT:
		p, err := s.checkSessionLogin(aul.Username, aul.Password, true)
		if err != nil {
			return nil, err
		}
		profile = p
	default:
		p, err := s.checkSessionLogin(aul.Username, aul.Password, false)
		if err != nil {
			return nil, err
		}
		profile = p
		sId := s.Ctx.GetCookie(constant.CookieNameForSessionID)
		if sId != "" {
			s.Ctx.RemoveCookie(constant.CookieNameForSessionID)
		}
		session := constant.Sess.Start(s.Ctx)
		session.Set(constant.SessionUserKey, profile)
	}

	go kolog.Save(aul.Username, constant.LOGIN, "-")

	return profile, nil
}

func (s *SessionController) Delete() error {
	session := constant.Sess.Start(s.Ctx)

	operator := ""
	mapxx := session.GetAll()
	if value, ok := mapxx[constant.SessionUserKey]; ok {
		if user, isUser := value.(*dto.Profile); isUser {
			operator = user.User.Name
		}
	}
	session.Delete(constant.SessionUserKey)

	go kolog.Save(operator, constant.LOGOUT, "-")
	return nil
}

func (s *SessionController) Get() (*dto.Profile, error) {
	session := constant.Sess.Start(s.Ctx)
	user := session.Get(constant.SessionUserKey)
	if user == nil {
		return nil, errors.New("SESSION_KEY_NOT_FOUND")
	}
	p, ok := user.(*dto.Profile)
	if !ok {
		return nil, errors.New("USER_PROFILE_INVALID")
	}
	// 重新查询用户,被删除的用户要退出登陆
	var mo model.User
	if err := database.DB.Model(model.User{}).Where(&model.User{ID: p.User.UserId}).Preload("CurrentProject").First(&mo).Error; err != nil {
		return nil, err
	}

	roles, err := getUserRole(&mo)
	if err != nil {
		return nil, err
	}
	p.User = toSessionUser(mo)
	p.User.Roles = roles
	session.Set(constant.SessionUserKey, p)
	return p, nil
}

func (s *SessionController) GetStatus() (*dto.SessionStatus, error) {
	session := constant.Sess.Start(s.Ctx)
	user := session.Get(constant.SessionUserKey)
	return &dto.SessionStatus{IsLogin: user != nil}, nil
}

func (s *SessionController) checkSessionLogin(username string, password string, jwt bool) (*dto.Profile, error) {
	u, err := s.UserService.UserAuth(username, password)
	if err != nil {
		return nil, err
	}
	resp := &dto.Profile{}
	roles, err := getUserRole(u)
	if err != nil {
		return nil, err
	}
	resp.User = toSessionUser(*u)
	resp.User.Roles = roles
	if jwt {
		token, err := createToken(resp.User)
		if err != nil {
			return nil, err
		}
		resp.Token = token
	}
	return resp, err
}

func toSessionUser(u model.User) dto.SessionUser {
	return dto.SessionUser{
		UserId:   u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Language: u.Language,
		IsActive: u.IsActive,
		IsAdmin:  u.IsAdmin,
		IsSuper:  u.IsSuper,
	}
}

func createToken(user dto.SessionUser) (string, error) {
	exp := viper.GetInt("jwt.exp")
	secretKey := []byte(viper.GetString("jwt.secret"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":           user.Name,
		"email":          user.Email,
		"userId":         user.UserId,
		"isActive":       user.IsActive,
		"language":       user.Language,
		"isAdmin":        user.IsAdmin,
		"currentProject": user.CurrentProject,
		"roles":          user.Roles,
		"iat":            time.Now().Unix(),
		"exp":            time.Now().Add(time.Minute * time.Duration(exp)).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func getUserRole(user *model.User) ([]string, error) {

	if user.IsAdmin {
		return []string{constant.RoleAdmin}, nil
	} else {
		return nil, nil
	}
}
