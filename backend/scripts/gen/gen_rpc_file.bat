@echo off
cd ../../app/rpc/file
cwgo server --type RPC --idl ../../../scripts/idl/file.proto --server_name filerpc --module backend/app/rpc/file -I ../../../scripts/idl
