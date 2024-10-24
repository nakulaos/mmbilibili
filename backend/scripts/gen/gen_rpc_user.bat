@echo off
cd ../../app/rpc/user
cwgo server --type RPC --idl ../../../scripts/idl/user.proto --server_name userrpc --module backend/app/rpc/user -I ../../../scripts/idl
