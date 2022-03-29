deployment:
	kubectl apply -f deployment.yaml && kubectl get services
minikube:
	minikube start && minikube tunnel
