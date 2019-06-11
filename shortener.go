package shortener

import (
	"fmt"
	"github.com/dchest/uniuri"
)

// shortener это интерфейс, который группирует методы работы с "сокращателем" ссылок.
type shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

// tinyUrl это экспортируемая структура, которая реализует методы интерфейса shortener.
type tinyUrl struct {
	host    string
	hashLen int
	storage map[string]string
}

// NewTinyUrl создает новый инстанс tinyUrl.
func NewTinyUrl(host string, hashLen int) *tinyUrl {
	return &tinyUrl{host, hashLen, map[string]string{}}
}

// hash возвращает хеш url'а.
func (t *tinyUrl) hash(url string) string {
	return uniuri.NewLen(t.hashLen)
}

// Shorten сохраняет связь между исходным и сокращенным url'ами
// в inMemory-хранилище и возвращает сокращенный url.
func (t *tinyUrl) Shorten(url string) string {
	shorted := fmt.Sprintf("%s/%s", t.host, t.hash(url))

	t.storage[shorted] = url

	return shorted
}

// Resolve возвращает исходный url по его сокращенной версии.
func (t *tinyUrl) Resolve(url string) string {
	if value, ok := t.storage[url]; ok {
		return value
	}
	return ""
}
