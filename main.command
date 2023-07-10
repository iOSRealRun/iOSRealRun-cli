#!/usr/bin/env sh
cd "$(dirname "${BASH_SOURCE[0]}")"
if [ -f "./main.py" ]; then
    python3 main.py
else
    chmod +x ./main
    ./main
fi
