#!/bin/bash

send_command() {
    curl -X POST -d "$1" http://localhost:17000
}

send_command "green"
send_command "bgrect 0.1 0.1 0.9 0.9"
send_command "figure 0.5 0.5"
send_command "update"