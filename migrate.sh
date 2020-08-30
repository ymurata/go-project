#!/bin/sh

COMMAND=$*

cd `dirname $0`
migrate -database ${DATABASE_URL} -path ./migrations ${COMMAND}
