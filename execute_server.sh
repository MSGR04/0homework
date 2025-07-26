#!/bin/bash

if [ -f "server.go" ]; then
    echo "Запуск Go сервера..."
    go run server.go
fi