#!/bin/bash

function abc() {
    n=$1
    if [ $n -gt 100 ]; then
        {
            echo "无效"
        }
    elif [ $n -gt 90 ]; then
        {
            echo "优秀"
        }
    else
        {
            echo "一般"
        }
    fi
}

abc 99
