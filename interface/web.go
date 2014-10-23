package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/wishes", func(r render.Render) {
		r.HTML(200, "standard", nil)
	})

	bind := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	fmt.Printf("listening on %s...", bind)
	err := http.ListenAndServe(bind, m)
	if err != nil {
		panic(err)
	}

}
