#!/bin/bash

echo "Current PID: $$"

sig_term_handler() {
  echo "Caught SIGTERM, ignoring..."
  exit 0
}

sig_int_handler() {
  echo "Caught SIGINT, exiting..."
  exit 0
}

#trap sig_term_handler SIGTERM
#trap sig_int_handler SIGINT

while true
do
  sleep 1
done