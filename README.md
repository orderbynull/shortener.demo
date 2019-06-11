**Пример использования:**
```
package main

import (
	"fmt"
	"github.com/orderbynull/shortener.demo"
)

func main() {
	s := shortener.NewTinyUrl("https://tiny.ru", 10)

	fmt.Println(s.Shorten("https://news.mail.ru/politics/37595116/?frommail=1"))
}
```

**Вывод:**
```
https://tiny.ru/0JFXoSpsfY
```

**Playground:** https://play.golang.org/p/7zyMZXjE9Uh