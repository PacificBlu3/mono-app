Run with command `./build.sh`

Error logs in WSL: `docker logs -f nginx 1>/dev/null` 

Access logs in WSL: `docker logs -f nginx 2>/dev/null` 

Docker stop all: `docker stop $(docker ps -a -q)`

Docker delete all: `docker rm $(docker ps -a -q)`
