#!/bin/bash

go build -o disk/web index.go
go build -o disk/proxy proxy.go
go build -o disk/imageWorker imageWorker.go
