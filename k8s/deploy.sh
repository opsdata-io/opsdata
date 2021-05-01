#!/bin/bash

cd 	/drone/src/k8s/
echo "customizing Deployment files..."
mkdir -p ./deployment-ready/
cd ./deployment
for file in `ls *.yaml`
do
  echo "Working on $file"
  cat $file | sed "s/BUILD_NUMBER/${CI_BUILD_NUMBER}/g" > ../deployment-ready/"$file"
done


cd 	/drone/src/
echo "customizing helm values..."
mkdir -p ./helm-config-ready/mariadb-galera/
cat ./helm-config/mariadb-galera/values.yaml | sed "s/BUILD_NUMBER/${CI_BUILD_NUMBER}/g" > ./helm-config-ready/mariadb-galera/values.yaml

cd /drone/src/deployment-ready/
echo "Creating secret..."
kubectl create secret generic jwt-secret-key --from-literal=jwt-secret-key=${jwt-secret-key} -o yaml --dry-run > ./secret.yaml
