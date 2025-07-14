#!/bin/bash

set -e

VITE_REST_HOST=https://api.wurbs.chat bun run build

b2 sync --delete --replace-newer dist b2://haralovich-wurbs/
