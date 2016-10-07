package migrationtest

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"

	"github.com/mjibson/goon"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

func init() {
	http.HandleFunc("/tests/upgrade", upgradeHandler)
	http.HandleFunc("/tests/downgrade", downgradeHandler)
}

func setCache(ctx context.Context, g *goon.Goon, ver string) error {
	var err error
	switch ver {
	case "v1":
		err = g.Get(&EntityV1{ID: 1})
	case "v2":
		err = g.Get(&EntityV2{ID: 1})
	case "no":
		memcache.Flush(ctx)
	default:
		err = fmt.Errorf("unknown version: %s\n", ver)
		memcache.Flush(ctx)
	}
	g.FlushLocalCache()
	return err
}

func upgradeHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if rv := recover(); rv != nil {
			fmt.Fprintf(w, "panic! %v\n", rv)
		}
	}()

	cv := r.URL.Query().Get("cached")

	g := goon.NewGoon(r)
	ctx := appengine.NewContext(r)
	memcache.Flush(ctx)

	v1 := EntityV1{ID: 1}
	g.Put(&v1)
	g.FlushLocalCache()

	if err := setCache(ctx, g, cv); err != nil {
		fmt.Fprintf(w, "err: %v\n", err)
		return
	}

	v2 := EntityV2{ID: 1}
	if err := g.Get(&v2); err != nil {
		fmt.Fprint(w, "err: %v\n", err)
		return
	}

	fmt.Fprintf(w, "Test Passed (cached: %v): %+v\n", cv, v2)
}

func downgradeHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if rv := recover(); rv != nil {
			fmt.Fprintf(w, "panic! %v\n", rv)
		}
	}()

	cv := r.URL.Query().Get("cached")

	g := goon.NewGoon(r)
	ctx := appengine.NewContext(r)
	memcache.Flush(ctx)

	v2 := EntityV2{ID: 1, Name: "Mike"}
	g.Put(&v2)
	g.FlushLocalCache()

	if err := setCache(ctx, g, cv); err != nil {
		fmt.Fprint(w, "err: %v\n", err)
		return
	}

	v1 := EntityV1{ID: 1}
	if err := g.Get(&v1); err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	fmt.Fprintf(w, "Test Passed (cached: %v): %+v\n", cv, v1)
}
