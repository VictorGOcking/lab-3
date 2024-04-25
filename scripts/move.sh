#!/bin/bash

send_command() {
    curl -X POST -d "$1" http://localhost:17000
}

x=0
y=0
step=0

send_command "green"
send_command "figure $x $y"

while true; do
    send_command "move $x $y"
    x=$(awk "BEGIN {printf \"%.2f\", $x + $step}")
    y=$(awk "BEGIN {printf \"%.2f\", $y + $step}")
    
    if (( $(awk "BEGIN {print ($x <= 0 && $y <= 0)}") )); then
        step=0.05
    elif (( $(awk "BEGIN {print ($x >= 1 && $y >= 1)}") )); then
        step=-0.05
    fi
    
    send_command "update"
    sleep 0.01
done
