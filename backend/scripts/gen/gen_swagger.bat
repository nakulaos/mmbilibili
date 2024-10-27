@echo off
cd  ../idl
goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api user.api -dir ../swagger
goctl api plugin -plugin goctl-swagger="swagger -filename file.json" -api file.api -dir ../swagger
