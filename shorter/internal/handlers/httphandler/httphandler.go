package httphandler

import (
	"errors"
	"net/http"

	"fmt"

	"regexp"

	"github.com/bartam1/mobilfox/shorter/internal/core/domain"
	service "github.com/bartam1/mobilfox/shorter/internal/core/services"
	"github.com/bartam1/mobilfox/shorter/pkg/errors/httperror"
	echo "github.com/labstack/echo/v4"
)

type Shorter struct {
	service service.Shorter
}

func (s Shorter) GetUrlsWidthHash(ctx echo.Context) error {
	urls, err := s.service.Queries.UrlsWidthHash.Do(ctx.Request().Context())
	if err != nil {
		return httperror.RespondWithSlugError(err)
	}
	return ctx.JSON(http.StatusOK, urls)

}
func (s Shorter) GetUrl(ctx echo.Context, hash string) error {
	url, err := s.service.Queries.Url.Do(ctx.Request().Context(), hash)
	if err != nil {
		return httperror.RespondWithSlugError(err)
	}
	return ctx.JSON(http.StatusOK, url)
}
func (s Shorter) MakeUrlHash(ctx echo.Context) error {
	json := echo.Map{}
	if err := ctx.Bind(&json); err != nil {
		return httperror.BadRequest("Bad request", err)
	}

	str := fmt.Sprintf("%v", json["Url"])
	re := regexp.MustCompile(`^(http|https)://`)

	if !re.MatchString(str) {
		return httperror.BadRequest(str+" http:// https:// is missing)", errors.New(str+" http:// https:// is missing."))
	}
	u := domain.MakeUrlHash{Url: str}
	url, err := s.service.Commands.MakeUrlHash.Do(ctx.Request().Context(), u)
	if err != nil {
		return httperror.RespondWithSlugError(err)
	}
	return ctx.JSON(http.StatusOK, url)
}
func (s Shorter) DeleteUrl(ctx echo.Context, hash string) error {
	err := s.service.Commands.DeleteUrl.Do(ctx.Request().Context(), hash)
	if err != nil {
		return httperror.RespondWithSlugError(err)
	}
	return ctx.String(http.StatusOK, "")
}

func New(ser service.Shorter) *Shorter {
	return &Shorter{service: ser}
}
