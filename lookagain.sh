#!/bin/bash

find . -type f -name "*.sh" \
    | xargs -n1 basename \
    | cut -d "." -f 1 \
    | sort -r
