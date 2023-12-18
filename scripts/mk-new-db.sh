#!/usr/bin/env bash

sqlite3 "${1}.db" <<EOF
    CREATE TABLE IF NOT EXISTS temperatures (
        id INTEGER NOT NULL PRIMARY KEY,
        value REAL, 
        scale TEXT
    );
EOF