#!/bin/bash

send_command() {
    curl -X POST -d "$1" http://localhost:17000
}

random_float() {
    awk -v min=-0.25 -v max=0.25 'BEGIN{srand(); printf "%.2f\n", min+rand()*(max-min)}'
}

x=0.5
y=0.5

send_command "green"
send_command "figure $x $y"

while true; do
    delta_x=$(random_float)

    if (( $(awk "BEGIN {print ($x + $delta_x <= 0 || $x + $delta_x >= 1)}") )); then
        delta_x=$(awk "BEGIN {printf \"%.2f\", $delta_x * -1}")
    fi

    x=$(awk "BEGIN {printf \"%.2f\", $x + $delta_x}")

    send_command "move $x $y"
    send_command "update"
    sleep 0.01

    delta_y=$(random_float)

    if (( $(awk "BEGIN {print ($y + $delta_y <= 0 || $y + $delta_y >= 1)}") )); then
	delta_y=$(awk "BEGIN {printf \"%.2f\", $delta_y * -1}")
    fi

    y=$(awk "BEGIN {printf \"%.2f\", $y + $delta_y}")

    send_command "move $x $y"
    send_command "update"
    sleep 0.01
done
