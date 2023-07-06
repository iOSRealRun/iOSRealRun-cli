#!/usr/bin/env sh
cd $(dirname $BASH_SOURCE)
if [ -f "./main.py"]; then
    python3 main.py
else
    chmod +x ./main.bin
    ./main.bin
fi
