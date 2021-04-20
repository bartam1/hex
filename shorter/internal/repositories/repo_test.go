package repo_test

import (
	"context"
	"testing"

	"github.com/bartam1/mobilfox/shorter/internal/core/domain"
	port "github.com/bartam1/mobilfox/shorter/internal/core/ports"
	"github.com/bartam1/mobilfox/shorter/internal/repositories/memrepo"
	"github.com/bartam1/mobilfox/shorter/internal/repositories/psqlrepo"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

const (
	API_URL = "postgres://postgres:almafa@localhost:5432/db"
)

func TestRepo(t *testing.T) {
	repositories := createRepositories(t)
	for i := range repositories {
		r := repositories[i]

		t.Run(r.Name, func(t *testing.T) {
			t.Parallel()

			t.Run("testMake_GetUrlHash", func(t *testing.T) {
				t.Parallel()
				testMake_GetUrlHash(t, r.Repository)
			})
		})
	}
}

func testMake_GetUrlHash(t *testing.T, r port.RepoIf) {
	testCases := []domain.UrlHash{
		{Url: "http://almakorte.hu", Hash: "12345678"},
		{Url: "http://amrte.hu", Hash: "12345178"},
		{Url: "http://akorte.hu", Hash: "11345678"},
		{Url: "http://almorte.hu", Hash: "12345618"},
		{Url: "http://almakrte.hu", Hash: "12345671"},
	}

	for _, tc := range testCases {
		_, err := r.MakeUrlHash(context.Background(), tc)
		require.NoError(t, err)
		r, err := r.GetUrl(context.Background(), tc.Hash)
		require.NoError(t, err)
		if !cmp.Equal(r, tc) {
			t.Errorf("Not equal")
		}
	}

}

type Repository struct {
	Name       string
	Repository port.RepoIf
}

func createRepositories(t *testing.T) []Repository {
	return []Repository{
		{
			Name:       "Psql",
			Repository: newPsqlRepo(t, context.Background()),
		},
		{
			Name:       "Mem",
			Repository: µ(memrepo.New())[0].(port.RepoIf),
		},
	}
}

func newPsqlRepo(t *testing.T, ctx context.Context) psqlrepo.PostgresqlDB {
	c, err := psqlrepo.New(ctx, API_URL)
	require.NoError(t, err)
	return c
}

func µ(a ...interface{}) []interface{} {
	return a
}
