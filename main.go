package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
)

const tpl = `
<meta name="go-import" content="%s git ssh://git@git.vonechain.com/%s">
`

var addr = flag.String("listen", "127.0.0.1:9090", "listening address")

func main() {
	flag.Parse()
	http.HandleFunc("/", handleGoGet)
	log.Println("listen ", addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

func handleGoGet(w http.ResponseWriter, r *http.Request) {
	pkgPath := packagePath(r.URL.Path)
	host := r.Host
	root := fmt.Sprintf("%s/%s", host, pkgPath)
	_, _ = fmt.Fprintf(w, tpl, root, pkgPath)
}

func packagePath(p string) string {
	s := strings.Split(p, "/")
	items := make([]string, 0, 2)
	for _, n := range s {
		if n == "" {
			continue
		}
		items = append(items, n)
		if len(items) == 2 {
			break
		}
	}
	return path.Join(items...)
}
