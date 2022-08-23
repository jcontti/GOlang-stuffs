# This project use Golang + Docker + Minikube(kubernetes) + Helm
  
- Before running the app:  
 -- Install VirtualBox  
 -- Install Minikube  
 -- Install Docker  
 -- Start your kubernetes cluster using "minikube start --driver=docker"  

 
  
- Running the app:  
 -- Build the image:  
 ```
  docker build -t go-app-normal:latest .
 ```  
 -- Start the container:  
 ```
  docker run -d -p 80:80 --name web go-app-normal:latest
 ```  
 -- Go to http://localhost:80  