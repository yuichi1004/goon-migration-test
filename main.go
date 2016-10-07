package migrationtest

import (
	"fmt"
	"net/http"

	"github.com/mjibson/goon"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

func init() {
	http.HandleFunc("/tests/upgrade", upgradeHandler)
	http.HandleFunc("/tests/downgrade", downgradeHandler)
}

func upgradeHandler(w http.ResponseWriter, r *http.Request) {
	withCache := "yes" == r.URL.Query().Get("cached")

	g := goon.NewGoon(r)
	ctx := appengine.NewContext(r)
	memcache.Flush(ctx)

	v1 := EntityV1{ID: 1}
	g.Put(&v1)

	if err := g.Get(&v1); err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}
	g.FlushLocalCache()
	if !withCache {
		memcache.Flush(ctx)
	}

	v2 := EntityV2{ID: 1}
	if err := g.Get(&v2); err != nil {
		fmt.Fprint(w, "err: %v", err)
		return
	}

	fmt.Fprintf(w, "Test Passed (cached: %v): %+v", withCache, v2)
}

func downgradeHandler(w http.ResponseWriter, r *http.Request) {
	withCache := "yes" == r.URL.Query().Get("cached")

	ctx := appengine.NewContext(r)
	v2 := EntityV2{ID: 1, Name: "Mike"}
	g := goon.FromContext(ctx)
	g.Put(&v2)

	if err := g.Get(&v2); err != nil {
		fmt.Fprint(w, "err: %v", err)
		return
	}
	g.FlushLocalCache()
	if !withCache {
		memcache.Flush(ctx)
	}

	v1 := EntityV1{ID: 1}
	if err := g.Get(&v1); err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	fmt.Fprintf(w, "Test Passed (cached: %v): %+v", withCache, v1)
}
