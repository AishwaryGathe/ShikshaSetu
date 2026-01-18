
# ShikshaSetu – Cloud-Native Go Application on AWS EKS

ShikshaSetu is a cloud-native Go web application that demonstrates an end-to-end **CI/CD and GitOps workflow** using Docker, Kubernetes, Helm, GitHub Actions, and ArgoCD on AWS EKS.

This README provides **essential commands only**, assuming all configuration files already exist in the repository.

---

## Architecture Overview

**CI Flow**
- Code pushed to GitHub
- GitHub Actions builds and pushes Docker images

**CD Flow**
- ArgoCD watches the GitHub repository
- Automatically syncs changes to AWS EKS using Helm

**Runtime Flow**
- Pods → Service → Ingress → User

---

## Containerization

Run application locally:
```bash
go run main.go
```

Build Docker image:
```bash
docker build -t <docker-username>/go-web-app:v1 .
```

Run container:
```bash
docker run -p 8080:8080 <docker-username>/go-web-app:v1
```

Push image:
```bash
docker push <docker-username>/go-web-app:v1
```

---

## AWS EKS Configuration

Configure AWS CLI:
```bash
aws configure
```

Connect kubectl to EKS:
```bash
aws eks update-kubeconfig --region <region> --name <cluster-name>
```

---

## Kubernetes Deployment (Raw Manifests)

Apply manifests:
```bash
kubectl apply -f k8s/manifests/
```

Verify resources:
```bash
kubectl get all
```

---

## Ingress Controller

Install NGINX Ingress Controller:
```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.11.1/deploy/static/provider/aws/deploy.yaml
```

Check ingress:
```bash
kubectl get ing
```

Map ingress IP to local domain:
```text
<INGRESS_IP> shikshasetu.local
```

---

## Helm Deployment

Remove raw manifests:
```bash
kubectl delete -f k8s/manifests/
```

Install using Helm:
```bash
helm install go-web-app helm/go-web-app-chart
```

Uninstall:
```bash
helm uninstall go-web-app
```

---

## CI – GitHub Actions

Store the following GitHub secrets:
- `DOCKER_USERNAME`
- `DOCKER_PASSWORD`

GitHub Actions handles image build and push automatically.

---

## CD – ArgoCD

Install ArgoCD:
```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

Expose ArgoCD UI:
```bash
kubectl patch svc argocd-server -n argocd -p '{"spec":{"type":"LoadBalancer"}}'
```

Get ArgoCD URL:
```bash
kubectl get svc argocd-server -n argocd
```

Retrieve admin password:
```bash
kubectl edit secret argocd-initial-admin-secret -n argocd
```

Decode the password using Base64.

---

## Continuous Delivery

- Any push to GitHub triggers CI
- ArgoCD automatically syncs changes
- Updates are reflected live on:
```text
http://shikshasetu.local
```

---

## Summary

This project demonstrates a real-world cloud-native deployment using AWS EKS with automated CI/CD and GitOps practices, closely aligned with production environments.
