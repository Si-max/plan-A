# plan-A-task1
Web application to return json response of

```json
{
  "timestamp": "<current date and time>",
  "hostname": "<Name of the host (IP also accepted)>",
  "engine": "<Name and/or version of the engine running>",
  "visitor ip": "<the IP address of the visitor>"
}
```
Application is dockerized with Alpine and used native golang packages to ensure small size

## How to get running

### Cloud
The deployment.yaml file creates a deployment and service of load balancer type

Clone the repo
```
git clone https://github.com/Si-max/Plan-A
```
and run

 ```
 kubectl apply -f deployment.yaml
 ```
 OR
 
 ```
 make deployment
 ```
 Get the external IP or dns from load balancer and you have the application running
 
 ### Local
 
 Running on local you need to install minikube which provides a cluster and needed tools to run Kubernetes, instructions to install [minikube](https://minikube.sigs.k8s.io/docs/start/)
 
 Run:
 ```
 make minikube
 ```
 and 
 
 ```
 make deployment
 ```
 then view services with
 ```
 kubectl get services
 ```
 Copy the external IP and run in browser to view the application.
 
 
 

