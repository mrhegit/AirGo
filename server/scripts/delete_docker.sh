containerName=$( docker ps -a | awk 'NR==2{print $1}' )
imageName=$( docker image ls | awk 'NR==2{print $3}' )
docker stop $containerName
docker rm $containerName
docker rmi $imageName


docker run -ti \
  -v /$PWD/air/config.yaml:/air/config.yaml \
  -p 20080:80 \
  -p 20443:443 \
  --name airgo \
  --privileged=true \
  ppoiuty/airgo:latest


docker run -ti \
  -v /$PWD/air/config.yaml:/air/config.yaml \
  -p 20080:80 \
  -p 20443:443 \
  --name airgo \
  --privileged=true \
  demo:latest