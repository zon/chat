#!/bin/bash

DATABASE=${1:-wurbs}

psql postgres -c "CREATE DATABASE \"$DATABASE\" ENCODING \"UTF8\" LC_COLLATE = \"en_US.UTF-8\" LC_CTYPE = \"en_US.UTF-8\";"
