tag=$1
if [ tag == '' ]; then
    tag=qa
fi
if docker buildx build -f ./Dockerfile -t hfunc/user-service:"${tag}" . ;then
  docker push hfunc/user-service:"${tag}"
fi