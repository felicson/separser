package separser

import (
	//	"fmt"
	"net/url"
	"testing"
)

const qs string = "http://images.yandex.ua/yandsearch?source=wiz&fp=4&uinfo=ww-1263-wh-850-fw-1038-fh-598-pd-1&tld=ua&p=4&text=%D0%BA%D0%B0%D1%80%D1%82%D0%B8%D0%BD%D0%BA%D0%B8%20%D0%BF%D0%BE%D0%BB%D0%B8%D1%8D%D1%82%D0%B8%D0%BB%D0%B5%D0%BD%20-%20%D0%B0%D0%BD%D1%82%D0%B8%D0%BA%D0%BE%D1%80%D1%80%D0%BE%D0%B7%D0%B8%D0%B9%D0%BD%D0%BE%D0%B5%20%D0%BF%D0%BE%D0%BA%D1%80%D1%8B%D1%82%D0%B8%D0%B5&noreask=1&pos=139&rpt=simage&lr=146&img_url=http%3A%2F%2Fcdn.stpulscen.ru%2Fsystem%2Fimages%2Fproduct%2F000%2F049%2F088_thumb.jpg"

func TestSequery(t *testing.T) {

	q, _ := NewSeQuery(qs)
	if q.Query == "" {
		t.Error("Wrong Query")
	}
	if q.SeId != 1 {
		t.Error("Wrong SeId")
	}

	if !q.Exist() {
		t.Error("Empty struct")
		t.Log(q)
	}
	if s, _ := q.SeName(); s != "Yandex" {
		t.Error("Wrong SE")
	}

}

func TestParseHost(t *testing.T) {
	index, _ := parseHost("www.google.com.ua")
	if index != "Google" {
		t.Error("Wrong Host")
	}
}

func TestSeName(t *testing.T) {
	q, _ := NewSeQuery(qs)
	index, _ := q.SeName()
	if index != "Yandex" {
		t.Error("Wrong Host")
	}
}

var sq *SeQuery

func BenchmarkParse(b *testing.B) {

	for n := 0; n < b.N; n++ {
		sq, _ = NewSeQuery(qs)
	}

}

var h string

func BenchmarkParseHost(b *testing.B) {
	for n := 0; n < b.N; n++ {
		h, _ = parseHost("www.yandex.ru")
	}
}

func BenchmarkParseUrl(b *testing.B) {
	for n := 0; n < b.N; n++ {
		u, _ := url.Parse(qs)
		_ = u.Host
	}
}

func BenchmarkSeName(b *testing.B) {

	b.StopTimer()
	sq, _ := NewSeQuery(qs)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_, _ = sq.SeName()

	}
}
