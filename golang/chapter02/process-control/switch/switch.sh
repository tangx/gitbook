#!/bin/bash

case $1 in
"n1")
    echo "n1"
    ;;
"n2" | "n3")
    echo "n2 or n3"
    ;;
*)
    echo "default"
    ;;
esac
