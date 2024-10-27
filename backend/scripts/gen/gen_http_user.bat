@echo off
cd ../../app/http
cwgo server --type HTTP --idl ../../scripts/idl/user.thrift --server_name http --module backend/app/http

