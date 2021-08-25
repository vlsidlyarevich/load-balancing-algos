red=$(tput setaf 1)
green=$(tput setaf 2)
blue=$(tput setaf 4)
reset=$(tput sgr0)

ECHO_SERVER_DOCKERFILE=./echo-server/Dockerfile
LOAD_BALANCER_DOCKERFILE=./load-balancer/Dockerfile

if [[ ! -f "$ECHO_SERVER_DOCKERFILE" ]] || [[ ! -f "$LOAD_BALANCER_DOCKERFILE" ]]; then
  echo "{$red}Unable to build images: docker file doesn't exist!${reset}"
  exit
fi
echo "${blue}Running docker images script${reset}"
echo "Building ${green}echo-server${reset} docker image"
cd ./echo-server && docker build . -t echo-server
cd ..
echo "Building ${green}load-balancer${reset} docker image"
cd ./load-balancer && docker build . -t load-balancer

echo "${blue}Success!!!${reset}"
