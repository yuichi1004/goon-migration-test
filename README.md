# goon でスキーママイグレーションができるか？

この記事の真偽を確かめます。

http://qiita.com/vvakame/items/e017e7d955f82ddd8af1

以下の２つのシナリオで試験します。

* アップグレード - 新しいフィールドが struct に追加された場合で動くか
* ダウングレード - いままで存在していたフィールドが struct から消えた場合で動くか

# 結果

アップグレードでも、ダウングレードでも期待通り動く。

```
$ goapp test -v                                                                                                                                                                                                                                                                                       [~/repos/GoonMigration]
2016/10/06 19:18:41 appengine: not running under devappserver2; using some default configuration
=== RUN   TestUpgradeWithCache
INFO     2016-10-06 10:18:41,837 devappserver2.py:769] Skipping SDK update check.
WARNING  2016-10-06 10:18:41,837 devappserver2.py:785] DEFAULT_VERSION_HOSTNAME will not be set correctly with --port=0
WARNING  2016-10-06 10:18:41,869 simple_search_stub.py:1146] Could not read search indexes from /var/folders/h0/c79kbs8x6jj8scsd6j22c5hh01hx69/T/appengine.testapp.yuichi.murata/search_indexes
INFO     2016-10-06 10:18:41,872 api_server.py:205] Starting API server at: http://localhost:56897
INFO     2016-10-06 10:18:41,875 dispatcher.py:197] Starting module "default" running at: http://localhost:56898
INFO     2016-10-06 10:18:41,876 admin_server.py:116] Starting admin server at: http://localhost:56899
--- PASS: TestUpgradeWithCache (2.60s)
        main_test.go:32: result: %+v { 1 }
=== RUN   TestUpgradeEntityWithoutCache
INFO     2016-10-06 10:18:44,442 devappserver2.py:769] Skipping SDK update check.
WARNING  2016-10-06 10:18:44,442 devappserver2.py:785] DEFAULT_VERSION_HOSTNAME will not be set correctly with --port=0
WARNING  2016-10-06 10:18:44,476 simple_search_stub.py:1146] Could not read search indexes from /var/folders/h0/c79kbs8x6jj8scsd6j22c5hh01hx69/T/appengine.testapp.yuichi.murata/search_indexes
INFO     2016-10-06 10:18:44,479 api_server.py:205] Starting API server at: http://localhost:56906
INFO     2016-10-06 10:18:44,482 dispatcher.py:197] Starting module "default" running at: http://localhost:56907
INFO     2016-10-06 10:18:44,484 admin_server.py:116] Starting admin server at: http://localhost:56908
--- PASS: TestUpgradeEntityWithoutCache (2.60s)
        main_test.go:57: result: {_kind: ID:1 Name:}
=== RUN   TestDowngradeWithCache
INFO     2016-10-06 10:18:47,059 devappserver2.py:769] Skipping SDK update check.
WARNING  2016-10-06 10:18:47,059 devappserver2.py:785] DEFAULT_VERSION_HOSTNAME will not be set correctly with --port=0
WARNING  2016-10-06 10:18:47,090 simple_search_stub.py:1146] Could not read search indexes from /var/folders/h0/c79kbs8x6jj8scsd6j22c5hh01hx69/T/appengine.testapp.yuichi.murata/search_indexes
INFO     2016-10-06 10:18:47,094 api_server.py:205] Starting API server at: http://localhost:56916
INFO     2016-10-06 10:18:47,097 dispatcher.py:197] Starting module "default" running at: http://localhost:56917
INFO     2016-10-06 10:18:47,098 admin_server.py:116] Starting admin server at: http://localhost:56918
--- PASS: TestDowngradeWithCache (2.63s)
        main_test.go:81: result: {_kind: ID:1}
=== RUN   TestDowngradeEntityWithoutCache
INFO     2016-10-06 10:18:49,676 devappserver2.py:769] Skipping SDK update check.
WARNING  2016-10-06 10:18:49,676 devappserver2.py:785] DEFAULT_VERSION_HOSTNAME will not be set correctly with --port=0
WARNING  2016-10-06 10:18:49,706 simple_search_stub.py:1146] Could not read search indexes from /var/folders/h0/c79kbs8x6jj8scsd6j22c5hh01hx69/T/appengine.testapp.yuichi.murata/search_indexes
INFO     2016-10-06 10:18:49,709 api_server.py:205] Starting API server at: http://localhost:56925
INFO     2016-10-06 10:18:49,712 dispatcher.py:197] Starting module "default" running at: http://localhost:56926
INFO     2016-10-06 10:18:49,714 admin_server.py:116] Starting admin server at: http://localhost:56927
--- PASS: TestDowngradeEntityWithoutCache (2.59s)
        main_test.go:106: result: {_kind: ID:1}
PASS
ok      _/Users/yuichi.murata/repos/GoonMigration       10.431s
```

