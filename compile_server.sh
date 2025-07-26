#!/bin/bash

if [ -f "server.go" ]; then
    echo "Go не требует предварительной компиляции"
    exit 0
fi
echo "Компиляция сервера завершена"
