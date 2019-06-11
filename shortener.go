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

// TinyUrl это экспортируемая структура, которая реализует методы интерфейса shortener.
type TinyUrl struct {
	host    string
	hashLen int
	storage map[string]string
}

// NewTinyUrl создает новый инстанс TinyUrl.
func NewTinyUrl(host string, hashLen int) *TinyUrl {
	return &TinyUrl{host, hashLen, map[string]string{}}
}

// hash возвращает хеш url'а.
func (t *TinyUrl) hash(url string) string {
	return uniuri.NewLen(t.hashLen)
}

// Shorten сохраняет связь между исходным и сокращенным url'ами
// в inMemory-хранилище и возвращает сокращенный url.
func (t *TinyUrl) Shorten(url string) string {
	shorted := fmt.Sprintf("%s/%s", t.host, t.hash(url))

	t.storage[shorted] = url

	return shorted
}

// Resolve возвращает исходный url по его сокращенной версии.
func (t *TinyUrl) Resolve(url string) string {
	if value, ok := t.storage[url]; ok {
		return value
	}
	return ""
}
