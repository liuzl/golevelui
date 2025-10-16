#!/usr/bin/env bash

bash -c "cd website && npm install && npm run build"

go run github.com/GeertJohan/go.rice/rice embed-go