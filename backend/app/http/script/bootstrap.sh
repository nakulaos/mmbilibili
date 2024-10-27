#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=userapi
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}