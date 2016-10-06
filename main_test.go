package migrationtest

import (
	"testing"

	"github.com/mjibson/goon"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/memcache"
)

func TestUpgradeWithCache(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	v1 := EntityV1{ID: 1}
	g := goon.FromContext(ctx)
	g.Put(&v1)

	if err = g.Get(&v1); err != nil {
		t.Errorf("err: %v", err)
	}

	v2 := EntityV2{ID: 1}
	if err = g.Get(&v2); err != nil {
		t.Errorf("err: %v", err)
	}

	t.Log("result: %+v", v2)
}

func TestUpgradeEntityWithoutCache(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	v1 := EntityV1{ID: 1}
	g := goon.FromContext(ctx)
	g.Put(&v1)

	if err = g.Get(&v1); err != nil {
		t.Errorf("err: %v", err)
	}
	g.FlushLocalCache()
	memcache.Flush(ctx)

	v2 := EntityV2{ID: 1}
	if err = g.Get(&v2); err != nil {
		t.Errorf("err: %v", err)
	}

	t.Logf("result: %+v", v2)
}

func TestDowngradeWithCache(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	v2 := EntityV2{ID: 1, Name: "Mike"}
	g := goon.FromContext(ctx)
	g.Put(&v2)

	if err = g.Get(&v2); err != nil {
		t.Errorf("err: %v", err)
	}

	v1 := EntityV1{ID: 1}
	if err = g.Get(&v1); err != nil {
		t.Errorf("err: %v", err)
	}

	t.Logf("result: %+v", v1)
}

func TestDowngradeEntityWithoutCache(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	v2 := EntityV2{ID: 1}
	g := goon.FromContext(ctx)
	g.Put(&v2)

	if err = g.Get(&v2); err != nil {
		t.Errorf("err: %v", err)
	}
	g.FlushLocalCache()
	memcache.Flush(ctx)

	v1 := EntityV1{ID: 1}
	if err = g.Get(&v1); err != nil {
		t.Errorf("err: %v", err)
	}

	t.Logf("result: %+v", v1)
}
