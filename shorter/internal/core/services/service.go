package service

import (
	"github.com/bartam1/mobilfox/shorter/internal/core/services/command"
	"github.com/bartam1/mobilfox/shorter/internal/core/services/query"
)

type Commands struct {
	MakeUrlHash command.MakeUrlHash
	DeleteUrl   command.DeleteUrl
}

type Queries struct {
	UrlsWidthHash query.UrlsWidthHash
	Url           query.Url
}

type Events struct {
}

type Shorter struct {
	Commands Commands
	Queries  Queries
	Events   Events
}
