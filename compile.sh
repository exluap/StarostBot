#!/bin/zsh
mkdir builds
GOOS=windows GOARCH=386 go build -o StarostBot_32bit.exe && mv StarostBot_32bit.exe builds/
GOOS=windows GOARCH=386 go build -o StarostBot_64bit.exe && mv StarostBot_64bit.exe builds/
GOARCH=386 go build -o StarostBot_i386 && mv StarostBot_i386builds/
go build -o StarostBot_amd64 && mv StarostBot_amd64 builds/
GOARM=5 GOARCH=arm go build -o StarostBot_Arm5 && mv StarostBot_Arm5 builds/
GOARM=6 GOARCH=arm go build -o StarostBot_Arm6 && mv StarostBot_Arm6 builds/
GOARM=7 GOARCH=arm go build -o StarostBot_Arm7 && mv StarostBot_Arm7 builds/
GOARCH=arm64 go build -o StarostBot_Arm8 && mv StarostBot_Arm8 builds/
