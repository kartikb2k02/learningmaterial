#!/bin/sh
# Every gate, in the order that fails most informatively.
#
#   ./check.sh
#
# The build goes to a temporary directory on purpose. `hugo server` renders to
# disk in this version, so building into public/ while the server is running
# deletes and rewrites the very files it is serving.
set -e
go run ./tools/configcheck
go run ./tools/levelcheck
go run ./tools/stylecheck README.md CONTENT_GUIDE.md COURSE-MAP.md product-management
go run ./tools/diagramcheck README.md product-management
go run ./tools/gocheck -run product-management
go run ./tools/linkcheck -pending README.md product-management | tail -1
hugo --quiet --destination "$(mktemp -d)/site"
echo "build: ok"
