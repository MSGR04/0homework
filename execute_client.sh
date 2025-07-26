#!/bin/bash

if [ -f "client.go" ]; then
    echo "Запуск Go клиента..."
    go run client.go
fi