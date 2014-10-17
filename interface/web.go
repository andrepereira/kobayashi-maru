package web

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Kobayashi Maru InfoSec Training System"
	})
	m.Run()
}
