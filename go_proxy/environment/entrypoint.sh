#!/bin/sh

cp /html/${FLAVOR}.html /usr/local/apache2/htdocs/index.html

exec "$@"