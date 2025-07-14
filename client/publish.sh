#!/bin/bash

set -e

bun run build

b2 sync --delete --replace-newer dist b2://haralovich-wurbs/

# https://f005.backblazeb2.com/file/haralovich-wurbs/index.html
