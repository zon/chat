#!/bin/bash

set -e

templ generate -watch -proxy="http://localhost:8080" -proxybind=0.0.0.0 -cmd="go run ./server" -open-browser=false