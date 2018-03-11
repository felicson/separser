package separser

import (
	"errors"
	"net/url"
	"strings"
)

var (
	ErrUnknownSe = errors.New("Unknown search engine")
	ErrWrongHost = errors.New("Wrong Host")
)

type se int

type seP struct {
	se    se
	param string
}

const (
	Yandex se = iota + 1
	Google
	Mail
	Bing
	Rambler
	Myprom
	Techserver
	Webalta
	Nigma
)

var seMap = map[string]seP{

	"Yandex":     seP{Yandex, "text"},
	"Google":     seP{Google, "q"},
	"Mail":       seP{Mail, "q"},
	"Bing":       seP{Bing, "q"},
	"Rambler":    seP{Rambler, "query"},
	"Myprom":     seP{Myprom, "query"},
	"Techserver": seP{Techserver, "query"},
	"Webalta":    seP{Webalta, "q"},
	"Nigma":      seP{Nigma, "s"},
}

func (s se) String() string {
	switch s {
	case Yandex:
		return "Yandex"
	case Google:
		return "Google"
	case Mail:
		return "Mail"
	case Bing:
		return "Bing"
	case Rambler:
		return "Rambler"
	case Myprom:
		return "Myprom"
	case Techserver:
		return "Techserver"
	case Webalta:
		return "Webalta"
	case Nigma:
		return "Nigma"
	}
	return ""
}

type SeQuery struct {
	Query    string
	EngineId se
}

func NewSeQuery(rawQuery string) (*SeQuery, error) {

	u, err := url.Parse(rawQuery)
	if err != nil {
		return nil, err
	}
	host, err := parseHost(u.Host)
	if err != nil {
		return nil, err
	}
	vals := u.Query()
	val, ok := seMap[host]
	if !ok {
		return nil, ErrUnknownSe
	}

	if err != nil {
		return nil, err
	}
	return &SeQuery{Query: vals.Get(val.param), EngineId: val.se}, nil

}

func (sq *SeQuery) Exist() bool {

	if sq.Query != "" && sq.EngineId > 0 {
		return true
	}
	return false
}

func (sq *SeQuery) SeName() (string, error) {
	if name := sq.EngineId.String(); name != "" {
		return name, nil
	}
	return "", ErrUnknownSe
}

func parseHost(fullHost string) (string, error) {

	parts := strings.Split(fullHost, ".")

	length := len(parts)

	if length > 1 && length <= 3 {
		return strings.Title(parts[length-2]), nil

	}
	if length > 3 {

		return strings.Title(parts[1]), nil
	}
	return "", ErrWrongHost
}
