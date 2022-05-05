set +e
set -u
set -o pipefail

START_TIME=$SECONDS

RED='\033[0;31M'
GREEN='\033[0;32m'
NC='\033[0m'

cleanup () {
    docker-compose -p template down -t 30 || :
    docker-compose -p template kill || :
    docker-compose -p template rm -f || :
    docker rmi $(docker images --filter=reference="*template-server*" -q) || :
    docker volume rm -f $(docker volume ls -f "dangling=true") || :
    rm -r bin || :
    exit 0
}

if [ "$(uname)" == "Darwin"]
then
    trap 'cleanup ; printf "${RED}Mac local template server terminated.${NC}\n ;"'\
        SIGHUP SIGINT SIGQUIT SIGABRT SIGKILL SIGPIPE SIGTERM
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]
then
    trap 'cleanup ; printf "${RED}Linux local template server terminated.${NC}\n ;"'\
        HUP INT QUIT PIPE TERM
else
    trap 'cleanup ; printf "${RED}Local template server terminated.${NC}\n ;"'\
        HUP INT QUIT PIPE TERM
fi

docker-compose -p template down
# build first
docker-compose -p template up -d --build template-server

if [ $? -ne 0 ] ; then
    printf "${RED}Docker compose failed${NC}\n"
    exit -1
fi 

ELAPSED_TIME=$(($SECONDS - $START_TIME))
printf "${GREEN}Start server task takes ${ELAPSED_TIME} seconds.${NC}\n"

docker logs $(docker container ls -f name=="*template-server*" -q) -f

cleanup


