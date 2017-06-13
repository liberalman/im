#!/bin/bash

CWD=`pwd`
SERVER=("gateway" "access" "notify" "register" "manager" "logic")
SERVICE=("idgen")

# See how we were called.
case "$1" in
  build)
	[ "$EUID" != "0" ] && exit 4
    
    for svr in ${SERVER[*]}
    do
        if [ -d $CWD/server/$svr ]; then
            pushd $CWD/server/$svr/ >/dev/null 2>&1
            go build
            echo "build "$CWD/server/$svr/$svr
            popd >/dev/null 2>&1
        else
            echo "no "$CWD/server/$svr
        fi
    done
    
    for svr in ${SERVICE[*]}
    do
        if [ -d $CWD/service/$svr ]; then
            pushd $CWD/service/$svr/ >/dev/null 2>&1
            go build
            echo "build "$CWD/service/$svr/$svr
            popd >/dev/null 2>&1
        else
            echo "no "$CWD/service/$svr
        fi
    done
    
        ;;
  clean)
	[ "$EUID" != "0" ] && exit 4
    
    for svr in ${SERVER[*]}
    do
        rm $CWD/server/$svr/$svr
    done
    for svr in ${SERVICE[*]}
    do
        rm $CWD/service/$svr/$svr
    done
    ;;
  restart|reload|force-reload)
        cd "$CWD"
	$0 stop
	$0 start
	rc=$?
	;;
  *)
        echo $"Usage: $0 {build|clean}"
        exit 2
esac

exit $rc
