@echo off
cd ../../app/http/user
cwgo server --type HTTP --idl ../../../scripts/idl/user.thrift --server_name userapi --module backend/app/http/user

