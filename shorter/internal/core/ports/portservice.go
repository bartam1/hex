package port

import (
	"github.com/bartam1/mobilfox/shorter/internal/core/services/command"
	"github.com/bartam1/mobilfox/shorter/internal/core/services/query"
)

type Commands struct {
	//MakeRandomUser command.CancelTrainingHandler
	MakeUrlHash command.MakeUrlHash
	DeleteUrl   command.DeleteUrl
}

type Queries struct {
	UrlsWidthHash query.UrlsWidthHash
	Url           query.Url
}

type Events struct {
}

type Service struct {
	Commands Commands
	Queries  Queries
	Events   Events
}
