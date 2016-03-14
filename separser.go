package separser

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

var (
	ErrUnknownSe = errors.New("Unknown search engine")
	ErrWrongHost = errors.New("Wrong Host")
)

var SE = map[string][]string{

	"Yandex":     {"1", "text"},
	"Google":     {"2", "q"},
	"Mail":       {"3", "q"},
	"Bing":       {"4", "q"},
	"Rambler":    {"5", "query"},
	"Myprom":     {"6", "query"},
	"Techserver": {"7", "query"},
	"Webalta":    {"8", "q"},
	"Nigma":      {"9", "s"},
}

type SeQuery struct {
	Query string
	SeId  int
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
	ptr, ok := SE[host]
	if !ok {
		return nil, ErrUnknownSe
	}

	id, err := strconv.Atoi(ptr[0])

	if err != nil {
		return nil, err
	}
	return &SeQuery{Query: vals.Get(ptr[1]), SeId: id}, nil

}

func (sq *SeQuery) Exist() bool {

	if sq.Query != "" && sq.SeId > 0 {
		return true
	}
	return false
}

func (sq *SeQuery) SeName() (string, error) {
	i := strconv.Itoa(sq.SeId)
	for k, v := range SE {
		if v[0] == i {
			return k, nil
		}
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
