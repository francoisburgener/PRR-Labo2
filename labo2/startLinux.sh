#!/bin/bash
echo ""
echo "Initialize all process"
echo ""
for ((i = 0; i <= $1-1; i++ )); do
  gnome-terminal -- bash -c "go run main.go -proc $i -N $1; exec bash"
  sleep 0.5s
done
