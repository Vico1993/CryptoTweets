#!/bin/bash
set -e

mongo <<-EOJS
    use cryptotweets
    db.createUser({
        user: $(_js_escape "$MONGO_USER"),
        pwd: $(_js_escape "$MONGO_USER_PWD"),
        roles: [ { role: 'readWrite', db: $(_js_escape "$MONGO_INITDB_DATABASE") } ]
    })
EOJS