docker build -t hello-go-image:latest ./Apps/HelloGo

kubectl apply -R -f ./Deployment