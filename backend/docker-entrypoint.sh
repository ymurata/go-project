#!/bin/sh
set -e

case $1 in
    run)
        sh ./migrate.sh up
        fresh
        ;;
    *)
        exec "$@"
        ;;
esac
