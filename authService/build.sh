tag=$1
if [ tag == '' ]; then
    tag=qa
fi
if docker buildx build -f ./Dockerfile -t hfunc/auth-service:"${tag}" . ;then
  docker push hfunc/auth-service:"${tag}"
fi