deployment:
	kubectl apply -f deployment.yaml
minikube:
	minikube start && minikube tunnel
