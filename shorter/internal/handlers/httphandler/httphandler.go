package httphandler

import (
	"net/http"

	"github.com/bartam1/mobilfox/shorter/internal/core/domain"
	port "github.com/bartam1/mobilfox/shorter/internal/core/ports"
	"github.com/bartam1/mobilfox/shorter/pkg/errors/httperror"
	echo "github.com/labstack/echo/v4"
)

type Shorter struct {
	service port.Service
}

func (s Shorter) GetUrlsWidthHash(ctx echo.Context) error {
	urls, err := s.service.Queries.UrlsWidthHash.Do(ctx.Request().Context())
	if err != nil {
		return httperror.Internal(ctx, err)
	}
	return ctx.JSON(http.StatusOK, urls)

}
func (s Shorter) GetUrl(ctx echo.Context, hash string) error {
	url, err := s.service.Queries.Url.Do(ctx.Request().Context(), hash)
	if err != nil {
		return httperror.NotFound(ctx, err)
	}
	return ctx.JSON(http.StatusOK, url)
}
func (s Shorter) MakeUrlHash(ctx echo.Context) error {
	u := domain.MakeUrlHash{Url: ctx.Request().PostFormValue("Url")}
	url, err := s.service.Commands.MakeUrlHash.Do(ctx.Request().Context(), u)
	if err != nil {
		return httperror.Internal(ctx, err)
	}
	return ctx.JSON(http.StatusOK, url)
}
func (s Shorter) DeleteUrl(ctx echo.Context, hash string) error {
	err := s.service.Commands.DeleteUrl.Do(ctx.Request().Context(), hash)
	if err != nil {
		return httperror.NotFound(ctx, err)
	}
	return ctx.String(http.StatusOK, "Deleted")
}

func New(ser port.Service) *Shorter {
	return &Shorter{service: ser}
}
