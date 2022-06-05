# docker-openweather
Création d'un conteneur docker qui récupère les données de l'API de openweather

docker login –u juliengan and enter my password # I login

### Build image in docker hub
sudo docker build . -t docker101:0.0.1

docker tag docker101:0.0.1 juliengan/docker101:0.0.1 # I tag it

docker push juliengan/docker101:0.0.1 # I publish my image to dockerhub

### Run the code

sudo docker run -it docker101:0.0.1 sh # I run the image

go run main.go # I run the code

docker run --env LAT="5.902785" --env LONG="102.754175" --env API_KEY="62bd02468799bb9568074245d9b8631e"
docker101:0.0.1/api:1.0.0 #API returning the weather using the image with location as input :

#### Sauvegarde variable d'environnement OpenWeatherAPI
export OPENWEATHERMAP_API_KEY="62bd02468799bb9568074245d9b8631e"

go get github.com/briandowns/openweathermap # retrieve openweathermap
