package psqlrepo

import (
	"context"

	"strings"

	"github.com/bartam1/mobilfox/shorter/internal/core/domain"
	"github.com/bartam1/mobilfox/shorter/pkg/errors/exterror"
	pgx "github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type postgresqlDB struct {
	dbclient *pgx.Conn
}

func (d postgresqlDB) GetUrlsWidthHash(ctx context.Context) (us []domain.UrlHash, err error) {
	rows, err := d.dbclient.Query(ctx, "select * from URLHASH")
	if err != nil {
		return []domain.UrlHash{}, exterror.NewRepo(err, func() { logrus.Error(err) })
	}
	var usrow domain.UrlHash
	us = make([]domain.UrlHash, 0)
	for rows.Next() {
		rows.Scan(&usrow.Hash, &usrow.Url)
		us = append(us, usrow)
	}
	return us, nil
}
func (d postgresqlDB) GetUrl(ctx context.Context, hash string) (u domain.UrlHash, err error) {
	var usrow domain.UrlHash
	e := d.dbclient.QueryRow(ctx, "select Url, Hash from URLHASH where Hash=$1", hash).Scan(&usrow.Url, &usrow.Hash)

	if e != nil {
		return domain.UrlHash{}, exterror.NewRepoSlug(e, "There is no url with that hash!", func() { logrus.Error(e) })
	}
	return domain.UrlHash{Url: usrow.Url, Hash: usrow.Hash}, nil
}
func (d postgresqlDB) MakeUrlHash(ctx context.Context, mu domain.UrlHash) (u domain.UrlHash, err error) {
	_, e := d.dbclient.Exec(ctx, "INSERT INTO URLHASH (Url, Hash) VALUES ($1,  $2)", mu.Url, mu.Hash)
	if e != nil {
		if !strings.Contains(e.Error(), "duplicate") {
			return domain.UrlHash{}, exterror.NewRepo(e, func() { logrus.Error(e) })
		}
		return domain.UrlHash{}, nil
	}
	return mu, nil
}
func (d postgresqlDB) DeleteUrl(ctx context.Context, hash string) error {
	_, e := d.dbclient.Exec(ctx, "delete from URLHASH where hash=$1", hash)
	if e != nil {
		return exterror.NewRepoSlug(e, "There is no url with that hash!", func() { logrus.Error(e) })
	}
	return nil
}
func New(ctx context.Context, strConn string) (postgresqlDB, error) {
	connConfig, err := pgx.ParseConfig(strConn)
	if err != nil {
		return postgresqlDB{nil}, err
	}

	conn, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		return postgresqlDB{nil}, err
	}
	return postgresqlDB{conn}, nil
}
