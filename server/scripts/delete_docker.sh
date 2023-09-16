containerName=$( docker ps -a | awk 'NR==2{print $1}' )
imageName=$( docker image ls | awk 'NR==2{print $3}' )
docker stop $containerName
docker rm $containerName
docker rmi $imageName
