package rest_session

import (
	"github.com/adwip/aj-teknik-backend-admin/common-lib/metadata"
	"github.com/labstack/echo/v5"
)

func (s *restSession) ErrorHandler(ctx *echo.Context, err error) {

	/*
		byteData, errExtraction := json.Marshal(ctx.QueryParams())
		if errExtraction != nil {
			s.log.Fatal(fmt.Sprintf("[\x1b[%dm%s\x1b[0m] %s \n", 31, "FATAL", errExtraction.Error()))
		}

		if ctx.Request().Method != http.MethodGet {
			body, errExtraction := io.ReadAll(ctx.Request().Body)
			if err != nil {
				s.log.Fatal(fmt.Sprintf("[\x1b[%dm%s\x1b[0m] %s \n", 31, "FATAL", errExtraction.Error()))
			}
			payload = string(body)
		}
	*/
	s.writeRestLog(err, ctx.Request().Method, ctx.Request().Header.Get(metadata.XRequestId), ctx.Path())
}
