#### algorithm
```go 
----algorithm\
    |----algorithm-1.go
    |----binarytreevisit.go
    |----linklist.go
    |----maxlength_word.go
    |----trietree.go
1.  // MinChange  找硬币这个问题直接使用一维动态规划进行求解， 和背包问题有些类似
2.  // heapsorttest 百度一面 手写堆排序
3. // DeferFunc 观察defer的输出顺序
4. //ReplaceCompare 将所有的ab 替换成bba 的次数
5.//countMaxActivity  不同时间区间进行合并
6.// startgorountine 启动goroutine 的顺序测试
```
#### cache
```go 
----cache\
    |----byteview.go
    |----cache.go
    |----cache_test.go
    |----consistenthash\
    |    |----consistenthash.go
    |    |----consistent_test.go
    |----geecache.go
    |----geecachepb\
    |    |----geecachepb.proto
    |    |----protos\
    |    |    |----geecachepb.pb.go
    |----http.go
    |----http_test.go
    |----lru\
    |    |----lru.go
    |    |----lru_test.go
    |----peers.go
    |----peer_test.go
    |----run.sh
    |----singleflight\
    |    |----singleflight.go

```
#### chatroom
```go 
----chatroom\
    |----chat_ws.pack
    |----config\
    |    |----chatroom.yaml
    |----global\
    |    |----config.go
    |    |----init.go
    |----logic\
    |    |----broadcast.go
    |    |----message.go
    |    |----offline.go
    |    |----sensitive.go
    |    |----user.go
    |----main.go
    |----server\
    |    |----home.go
    |    |----websocket.go
    |----template\
    |    |----home.html
1.  实现对应的goroutine管理
2.  前端的websocket js
3.  单例广播模式
```
#### crontab-task
```go 
----crontab-task\
    |----crontab.go
    |----task.json
1.  实现进程启动， 进程重启， 进程定时重启，三种功能， 
2.  调用的库    gopsutil  crontab v3
3.task.json  配置文件   crontab.go 任务调度
```
#### design-pattern
```go 
----design-pattern\
    |----decorator.go
    |----interpret.go
    |----iteror.go
    |----mediaitor.go
    |----observer.go
    |----singleleton.go
1.观察者模式
2.单例模式
3.装饰器模式
4.迭代器模式
```
#### file-downloader
```go 
----file-downloader\
    |----download.mp4
    |----gao.mp4
    |----index.html
    |----prasevideo.go
    |----request.go

```
#### file-watch
```go 
----file-watch\
    |----config.yaml
    |----filewatch.go
    |----task.json

```
#### flash-sale
```go 
----flash-sale\
    |----gin.log
    |----main.go
    |----models.go
    |----router.go
    |----service.go
    |----sqls.go

```
#### ginoc
```go 
----ginoc\
    |    |----httpRequests\
    |    |    |----http-client.cookies
    |    |    |----http-requests-log.http
    |----1631783255556186200_ecb45559bda357aea824fde70cdde995.jpg
    |----main.go
    |----main2.go
    |----monitor.go

```
#### go-mod-init
```go 
----go-mod-init\
    |----main.go
    |----router\
    |    |----getRouter\
    |    |    |----getRouter.go
    |    |----router.go
    |----runRouter\
    |    |----runRouter.go
    |----test\
    |    |----test.go

```
#### go-protobuf
```go 
----go-protobuf\
    |----newdump
    |----serialproto.go
    |----student.pb.go
    |----student.proto
    |----student_pb2.py

```
#### go-webassemblely
```go 
----go-webassemblely\
    |----index.html
    |----lib.wasm
    |----main.go
    |----server.go
    |----wasm_exec.js

```
#### golang-features
```go 
----golang-features\
    |----a.json
    |----address.json
    |----alg.go
    |----cgotest.go
    |----client.go
    |----cmp_test.go
    |----comlielist.go
    |----computedis.go
    |----data_interface.pb.go
    |----deamon.go
    |----editdistance.go
    |----generics_comparable.go
    |----generics_iterable.go
    |----generics_restraint.go
    |----goid.go
    |----pprof.go
    |----service.go
    |----signal.go
    |----反射基础.go
1. 并发使用 map 
2. goroutine 的泄漏
3. 负载均衡算法的简单实现
4. 通过反射调用函数
```
#### goroutine-pool
```go 
----goroutine-pool\
    |----ants.go
    |----ants_test.go
    |----pool.go
    |----worker.go

```
#### grpcserver
```go 
----grpcserver\
    |----client\
    |    |----grpcclient.go
    |----pb\
    |    |----simple.pb.go
    |    |----simple.proto
    |    |----simple_grpc.pb.go
    |----server\
    |    |----grpcserver.go

```
#### httpserver
```go 
----httpserver\
    |----client.go
    |----server.go

```
#### kafka
```go 
----kafka\
    |----kafk.go
    |----kafkademo
    |----my.log
    |----test.ini

```
#### lock-free
```go 
----lock-free\
    |----esQueue.go
    |----lockfree.go

```
#### min-component
```go 
----min-component\
    |----channelchain.go
    |----concurrentmap.go
    |----flagprase.go
    |----getbalancer.go
    |----getemial.go
    |----getfuncbyname.go
    |----goroutineleak.go
    |----ioreader.go
    |----lrucache.go
    |----mmap.go
    |----pictureserver.go
    |----propressfile.go
    |----receive_signal.go
    |----rename.go
    |----scrapy.go
    |----slicegrow.go
    |----stopgoroutine.go
    |----stringbuilder.go
    |----swapalpha.go
    |----sync-cond.go
    |----tokenlimit.go
    |----xrate.go

```
#### nets
```go 
----nets\
    |----dial.go
    |----nets_test.go
    |----socket.go

```
#### orm
```go 
----orm\
    |    |----dataSources\
    |----dialect\
    |    |----dialect.go
    |    |----sqlite3.go
    |----gee.db
    |----log\
    |    |----log.go
    |----orm.go
    |----orm_test.go
    |----record_test.go
    |----schema\
    |    |----generateor.go
    |    |----generate_test.go
    |    |----schema.go
    |    |----schema_test.go
    |----session\
    |    |----raw.go
    |    |----record.go
    |    |----table.go
    |----table_test.go

```
#### painkiller
```go 
----painkiller\
    |----painkiller.go
    |----pill_string.go

```
#### prase-config
```go 
----prase-config\
    |----companies.json
    |----companies.yaml
    |----config.json
    |----readfile.go
    |----structure.go

```
#### redis
```go 
----redis\
    |----compent\
    |    |----redis.go
    |    |----redis_test.go
    |----httpget.go
    |----index.html
    |----main.go

```
#### refelect-struct
```go 
----refelect-struct\
    |----reflectstruct.go
    |----reflectstruct_test.go

```
#### rpcserver
```go 
----rpcserver\
    |----client.go
    |----server.go

```
#### shm
```go 
----shm\
    |----linux\
    |    |----shm_linux.go
    |----windows\
    |    |----shm_windows.go

```
#### sqlcon
```go 
----sqlcon\
    |----jwt.go
    |----mysqlconn.go
    |----readcommand.go
    |----sqlorm.go

```
#### tcpserver
```go 
----tcpserver\
    |----client.go
    |----server.go
    |----utils.go

```
#### tree.txt
```go 

```
#### udpserver
```go 
----udpserver\
    |----client.go
    |----pyau.ipynb
    |----server.go

```
#### web
```go 
----web\
    |----context.go
    |----gee.go
    |----handler.go
    |----logger.go
    |----recovy.go
    |----router.go
    |----trie.go
    |----web2_test.go
    |----web_test.go

```
#### websocketserver
```go 
----websocketserver\
    |----grilloa_client.go
    |----grilloa_server.go
    |----new.html
    |----nhooyr_client.go
    |----nhooyr_server.go

```
