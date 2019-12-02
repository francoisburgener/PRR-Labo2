
#!/bin/bash
echo "Bash version ${BASH_VERSION}..."

for i in {0..4}
  do 
     gnome-terminal -x bash -c "go run main.go -proc $i -N $1; exec bash"
     sleep 0.5s
 done


