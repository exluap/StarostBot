#!/bin/zsh
GOOS=windows GOARCH=386 go build -o StarostBot_32bit.exe
GOOS=windows GOARCH=386 go build -o StarostBot_64bit.exe
GOARCH=386 go build -o StarostBot_i386
go build -o StarostBot_amd64
GOARM=5 GOARCH=arm go build -o StarostBot_Arm5
GOARM=6 GOARCH=arm go build -o StarostBot_Arm6
GOARM=7 GOARCH=arm go build -o StarostBot_Arm7
GOARCH=arm64 go build -o StarostBot_Arm8
