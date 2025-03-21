#!/bin/bash

psql postgres -c 'CREATE DATABASE "wurbs" ENCODING "UTF8" LC_COLLATE = "en_US.UTF-8" LC_CTYPE = "en_US.UTF-8";'