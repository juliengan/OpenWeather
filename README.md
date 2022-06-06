# Docker : wrapper en go qui récupère les données CurrentWeatherData de OpenWeatherMap (OWM)

Utilisation de Alpine 

docker login –u username # login

### Build image in docker hub
sudo docker build . -t image

docker tag image username/image # I tag it

docker push username/image # I publish my image to dockerhub

### Run the code

sudo docker run -it docker101:0.0.1 go run main.go 

docker run --env LAT="5.902785" --env LONG="102.754175" --env API_KEY="62bd02468799bb9568074245d9b8631e"
docker101:0.0.1/api:1.0.0 #API returning the weather using the image with location as input :

#### Sauvegarde variable d'environnement OpenWeatherAPI
export OPENWEATHERMAP_API_KEY=***

go get github.com/briandowns/openweathermap # retrieve openweathermap
