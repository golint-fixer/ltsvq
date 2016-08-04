ltsvq [![Build Status](https://travis-ci.org/hnakamur/ltsvq.png)](https://travis-ci.org/hnakamur/ltsvq) [![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hyperium/hyper/master/LICENSE)
=====

ltsvq reads a LTSV file and prints subset fields of specified labels.

## Install

```
go get github.com/hnakamur/ltsvq
```

## Usage

```
$ ./ltsvq -h
Usage of ./ltsvq:
  -f string
        filename (default "-")
  -l string
        labels (ex. time,url,status)
```

## Example

```
$ cat a.log
time:2016-05-30T02:21:28.135713584Z	level:Debug	msg:This is a debug message	key:key1	intValue:234
time:2016-05-30T02:21:28.135744631Z	level:Info	msg:hello, world	key:key1	value:value1
time:2016-05-30T02:21:28.135772957Z	level:Error	err:demo error	stack: [main.b() /home/hnakamur/gocode/src/github.com/hnakamur/ltsvlog/cmd/example/main.go:28 +0x1ba],[main.a() /home/hnakamur/gocode/src/github.com/hnakamur/ltsvlog/cmd/example/main.go:22 +0x14],[main.main() /home/hnakamur/gocode/src/github.com/hnakamur/ltsvlog/cmd/example/main.go:16 +0x56c]
time:2016-05-30T02:21:28.135804911Z	level:Info	msg:goodbye, world	foo:bar	nilValue:<nil>	bytes:0x612f62
$ ./ltsvq -l time,msg < a.log
time:2016-05-30T02:21:28.135713584Z     msg:This is a debug message
time:2016-05-30T02:21:28.135744631Z     msg:hello, world
time:2016-05-30T02:21:28.135772957Z
time:2016-05-30T02:21:28.135804911Z     msg:goodbye, world
$ ./ltsvq -l time,msg -f a.log
time:2016-05-30T02:21:28.135713584Z     msg:This is a debug message
time:2016-05-30T02:21:28.135744631Z     msg:hello, world
time:2016-05-30T02:21:28.135772957Z
time:2016-05-30T02:21:28.135804911Z     msg:goodbye, world
```

## License
MIT
