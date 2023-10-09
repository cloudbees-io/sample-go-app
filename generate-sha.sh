#!/bin/bash

git config --global --add safe.directory /app
git rev-parse HEAD >sha.txt
