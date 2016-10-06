# goon でスキーママイグレーションができるか？

この記事の真偽を確かめます。

http://qiita.com/vvakame/items/e017e7d955f82ddd8af1

以下の２つのシナリオで試験します。

* アップグレード - 新しいフィールドが struct に追加された場合で動くか
* ダウングレード - いままで存在していたフィールドが struct から消えた場合で動くか

# 結果

## キャッシュしていない場合

アップグレードでも、ダウングレードでも期待通り動く。

```
$ goapp test -v -run WithoutCache                                                                                                                                                                                                                                                                     [~/repos/GoonMigration]
2016/10/06 19:04:45 appengine: not running under devappserver2; using some default configuration
=== RUN   TestUpgradeEntityWithoutCache
INFO     2016-10-06 10:04:45,599 devappserver2.py:769] Skipping SDK update check.
WARNING  2016-10-06 10:04:45,599 devappserver2.py:785] DEFAULT_VERSION_HOSTNAME will not be set correctly with --port=0
WARNING  2016-10-06 10:04:45,632 simple_search_stub.py:1146] Could not read search indexes from /var/folders/h0/c79kbs8x6jj8scsd6j22c5hh01hx69/T/appengine.testapp.yuichi.murata/search_indexes
INFO     2016-10-06 10:04:45,636 api_server.py:205] Starting API server at: http://localhost:56655
INFO     2016-10-06 10:04:45,640 dispatcher.py:197] Starting module "default" running at: http://localhost:56656
INFO     2016-10-06 10:04:45,641 admin_server.py:116] Starting admin server at: http://localhost:56657
--- PASS: TestUpgradeEntityWithoutCache (2.70s)
        main_test.go:56: result: {_kind: ID:1 Name:}
=== RUN   TestDowngradeEntityWithoutCache
INFO     2016-10-06 10:04:48,301 devappserver2.py:769] Skipping SDK update check.
WARNING  2016-10-06 10:04:48,301 devappserver2.py:785] DEFAULT_VERSION_HOSTNAME will not be set correctly with --port=0
WARNING  2016-10-06 10:04:48,334 simple_search_stub.py:1146] Could not read search indexes from /var/folders/h0/c79kbs8x6jj8scsd6j22c5hh01hx69/T/appengine.testapp.yuichi.murata/search_indexes
INFO     2016-10-06 10:04:48,338 api_server.py:205] Starting API server at: http://localhost:56665
INFO     2016-10-06 10:04:48,341 dispatcher.py:197] Starting module "default" running at: http://localhost:56666
INFO     2016-10-06 10:04:48,342 admin_server.py:116] Starting admin server at: http://localhost:56667
--- PASS: TestDowngradeEntityWithoutCache (2.73s)
        main_test.go:104: result: {_kind: ID:1}
PASS
ok      _/Users/yuichi.murata/repos/GoonMigration       5.437s
```

## キャッシュしている場合

アップグレードでもダウングレードでも上手く動かない。
どちらも reflect.Set で失敗して panic する。

```
$ goapp test -v -run UpgradeWithCache                                                                                                                                                                                                                                                                 [~/repos/GoonMigration]
2016/10/06 19:07:53 appengine: not running under devappserver2; using some default configuration
=== RUN   TestUpgradeWithCache
INFO     2016-10-06 10:07:53,496 devappserver2.py:769] Skipping SDK update check.
WARNING  2016-10-06 10:07:53,496 devappserver2.py:785] DEFAULT_VERSION_HOSTNAME will not be set correctly with --port=0
WARNING  2016-10-06 10:07:53,529 simple_search_stub.py:1146] Could not read search indexes from /var/folders/h0/c79kbs8x6jj8scsd6j22c5hh01hx69/T/appengine.testapp.yuichi.murata/search_indexes
INFO     2016-10-06 10:07:53,532 api_server.py:205] Starting API server at: http://localhost:56704
INFO     2016-10-06 10:07:53,534 dispatcher.py:197] Starting module "default" running at: http://localhost:56705
INFO     2016-10-06 10:07:53,536 admin_server.py:116] Starting admin server at: http://localhost:56706
--- FAIL: TestUpgradeWithCache (2.63s)
panic: reflect.Set: value of type migrationtest.EntityV1 is not assignable to type migrationtest.EntityV2 [recovered]
        panic: reflect.Set: value of type migrationtest.EntityV1 is not assignable to type migrationtest.EntityV2

goroutine 5 [running]:
panic(0x406dc0, 0xc82024e1a0)
        /usr/local/Cellar/app-engine-go-64/1.9.38/share/app-engine-go-64/goroot/src/runtime/panic.go:481 +0x3e6
testing.tRunner.func1(0xc820086750)
        /usr/local/Cellar/app-engine-go-64/1.9.38/share/app-engine-go-64/goroot/src/testing/testing.go:467 +0x192
panic(0x406dc0, 0xc82024e1a0)
        /usr/local/Cellar/app-engine-go-64/1.9.38/share/app-engine-go-64/goroot/src/runtime/panic.go:443 +0x4e9
reflect.Value.assignTo(0x4e9140, 0xc8201aa200, 0x199, 0x5b0ae0, 0xb, 0x4fa320, 0x0, 0x0, 0x0, 0x0)
        /usr/local/Cellar/app-engine-go-64/1.9.38/share/app-engine-go-64/goroot/src/reflect/value.go:2164 +0x3be
reflect.Value.Set(0x4fa320, 0xc820250210, 0x199, 0x4e9140, 0xc8201aa200, 0x199)
        /usr/local/Cellar/app-engine-go-64/1.9.38/share/app-engine-go-64/goroot/src/reflect/value.go:1334 +0x95
github.com/mjibson/goon.(*Goon).GetMulti(0xc8201a2120, 0x3f1a40, 0xc8202541e0, 0x0, 0x0)
        /Users/yuichi.murata/go/src/github.com/mjibson/goon/goon.go:430 +0x602
github.com/mjibson/goon.(*Goon).Get(0xc8201a2120, 0x4e40c0, 0xc820250210, 0x0, 0x0)
        /Users/yuichi.murata/go/src/github.com/mjibson/goon/goon.go:378 +0x2b2
_/Users/yuichi.murata/repos/GoonMigration.TestUpgradeWithCache(0xc820086750)
        /Users/yuichi.murata/repos/GoonMigration/main_test.go:27 +0x4bd
testing.tRunner(0xc820086750, 0x85d800)
        /usr/local/Cellar/app-engine-go-64/1.9.38/share/app-engine-go-64/goroot/src/testing/testing.go:473 +0x98
created by testing.RunTests
        /usr/local/Cellar/app-engine-go-64/1.9.38/share/app-engine-go-64/goroot/src/testing/testing.go:582 +0x892
exit status 2
FAIL    _/Users/yuichi.murata/repos/GoonMigration       2.647s


$ goapp test -v -run DowngradeWithCache                                                                                                                                                                                                                                                               [~/repos/GoonMigration]
2016/10/06 19:08:07 appengine: not running under devappserver2; using some default configuration
=== RUN   TestDowngradeWithCache
INFO     2016-10-06 10:08:08,357 devappserver2.py:769] Skipping SDK update check.
WARNING  2016-10-06 10:08:08,357 devappserver2.py:785] DEFAULT_VERSION_HOSTNAME will not be set correctly with --port=0
WARNING  2016-10-06 10:08:08,387 simple_search_stub.py:1146] Could not read search indexes from /var/folders/h0/c79kbs8x6jj8scsd6j22c5hh01hx69/T/appengine.testapp.yuichi.murata/search_indexes
INFO     2016-10-06 10:08:08,391 api_server.py:205] Starting API server at: http://localhost:56715
INFO     2016-10-06 10:08:08,394 dispatcher.py:197] Starting module "default" running at: http://localhost:56716
INFO     2016-10-06 10:08:08,395 admin_server.py:116] Starting admin server at: http://localhost:56717
--- FAIL: TestDowngradeWithCache (2.61s)
panic: reflect.Set: value of type migrationtest.EntityV2 is not assignable to type migrationtest.EntityV1 [recovered]
        panic: reflect.Set: value of type migrationtest.EntityV2 is not assignable to type migrationtest.EntityV1

goroutine 5 [running]:
panic(0x406dc0, 0xc820173e10)
        /usr/local/Cellar/app-engine-go-64/1.9.38/share/app-engine-go-64/goroot/src/runtime/panic.go:481 +0x3e6
testing.tRunner.func1(0xc8200925a0)
        /usr/local/Cellar/app-engine-go-64/1.9.38/share/app-engine-go-64/goroot/src/testing/testing.go:467 +0x192
panic(0x406dc0, 0xc820173e10)
        /usr/local/Cellar/app-engine-go-64/1.9.38/share/app-engine-go-64/goroot/src/runtime/panic.go:443 +0x4e9
reflect.Value.assignTo(0x4fa320, 0xc8201c80c0, 0x199, 0x5b0ae0, 0xb, 0x4e9140, 0x0, 0x0, 0x0, 0x0)
        /usr/local/Cellar/app-engine-go-64/1.9.38/share/app-engine-go-64/goroot/src/reflect/value.go:2164 +0x3be
reflect.Value.Set(0x4e9140, 0xc820181e60, 0x199, 0x4fa320, 0xc8201c80c0, 0x199)
        /usr/local/Cellar/app-engine-go-64/1.9.38/share/app-engine-go-64/goroot/src/reflect/value.go:1334 +0x95
github.com/mjibson/goon.(*Goon).GetMulti(0xc8201a0120, 0x3f1a40, 0xc820181e80, 0x0, 0x0)
        /Users/yuichi.murata/go/src/github.com/mjibson/goon/goon.go:430 +0x602
github.com/mjibson/goon.(*Goon).Get(0xc8201a0120, 0x4e3fe0, 0xc820181e60, 0x0, 0x0)
        /Users/yuichi.murata/go/src/github.com/mjibson/goon/goon.go:378 +0x2b2
_/Users/yuichi.murata/repos/GoonMigration.TestDowngradeWithCache(0xc8200925a0)
        /Users/yuichi.murata/repos/GoonMigration/main_test.go:75 +0x4d3
testing.tRunner(0xc8200925a0, 0x85d830)
        /usr/local/Cellar/app-engine-go-64/1.9.38/share/app-engine-go-64/goroot/src/testing/testing.go:473 +0x98
created by testing.RunTests
        /usr/local/Cellar/app-engine-go-64/1.9.38/share/app-engine-go-64/goroot/src/testing/testing.go:582 +0x892
exit status 2
FAIL    _/Users/yuichi.murata/repos/GoonMigration       2.626s
```

