minikube start
eval $(minikube docker-env)
docker build -t gocache-server .
kubectl delete deployment gocache-server
kubectl delete service gocache-server
kubectl apply -f deployment.yml
kubectl apply -f service.yml
minikube service gocache-server