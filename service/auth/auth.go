package auth

import (
	"net/http"

	"web_template/ecode"
	"web_template/model"

	"github.com/labstack/echo/v4"
)

func (s *Service) KeyAuthValidator(session string, c echo.Context) (bool, error) {
	//cookie, err := c.Cookie("cookie")
	//if err != nil {
	//	return false, json(c, ecode.CookieErr)
	//}
	//fmt.Println("cookie:", cookie.String())

	userId, err := s.dao.GetUserIdBySession(session)
	if err != nil {
		return false, json(c, ecode.Unauthorized)
	}
	if userId == 0 {
		return false, json(c, ecode.Unauthorized)
	}
	s.dao.ExpireUserIdBySession(session)
	return true, nil
}

func json(c echo.Context, err error) error {
	codes := ecode.AnalyseError(err)
	d := &model.ReturnData{
		Code:    codes.Code(),
		Message: codes.Message(),
		Data:    nil,
	}
	return c.JSON(http.StatusOK, d)
}
