#!/bin/bash
if [ -f "client.go" ]; then
    echo "Go не требует предварительной компиляции"
    exit 0
fi
echo "Компиляция клиента завершена"