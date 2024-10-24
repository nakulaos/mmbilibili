#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/userrpc"
exec "$CURDIR/bin/userrpc"
