#! /bin/bash
user=$(whoami)
kubectl create clusterrolebinding cluster-admin-binding --clusterrole=cluster-admin --user=$user
#applying tekton pipelines
kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml

##TODO: check status for kubectl get pods --namespace tekton-pipelines --watch
sleep 5
#applying tekton triggers
kubectl apply --filename https://storage.googleapis.com/tekton-releases/triggers/latest/release.yaml

#create namespace 
kubectl create namespace getting-started
kubectl apply -f /opt/dkube/admin-role.yaml
kubectl apply -f /opt/dkube/webhook-role.yaml
kubectl apply -f /opt/dkube/mypipeline.yaml
kubectl apply -f /opt/dkube/triggers.yaml
kubectl apply -f /opt/dkube/secrets.yaml

#dashboard
kubectl apply --filename https://storage.googleapis.com/tekton-releases/dashboard/latest/tekton-dashboard-release.yaml
sleep 5
kubectl get svc tekton-dashboard --namespace tekton-pipelines -oyaml > dashboard.yaml && sed -i "s/ClusterIP/NodePort/" dashboard.yaml && kubectl replace -f dashboard.yaml


#TODO: display dashboard port
kubectl get svc tekton-dashboard --namespace tekton-pipelines
#TODO: display event lsiterner port 
kubectl get svc -n getting-started el-getting-started-listener
