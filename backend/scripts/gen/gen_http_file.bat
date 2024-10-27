@echo off
cd ../../app/http
cwgo server --type HTTP --idl ../../scripts/idl/file.thrift --server_name http --module backend/app/http


