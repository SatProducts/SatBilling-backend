# PodBilling backend

 Backend of Billing system for internet providers backend

# Build & Run
1) Install Go (project use v1.17.7)
2) Clone repo
```cmd
git clone https://github.com/PodProducts/PodBilling-backend.git
cd PodBilling-backend
```
3) Install libs
```cmd
go mod download
```
4) Build and run project\
Linux:
```cmd
go build -o server cmd/main.go
./server
```
Windows:
```cmd
go build -o server.exe cmd/main.go
server
```