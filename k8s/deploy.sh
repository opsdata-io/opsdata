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
cat ./helm-config/mariadb-galera/master.yaml | sed "s/BUILD_NUMBER/${CI_BUILD_NUMBER}/g" | sed "s/REPO_USER/${REPO_USER}/g" > ./helm-config-ready/mariadb-galera/master.yaml
cat ./helm-config/mariadb-galera/dev.yaml | sed "s/BUILD_NUMBER/${CI_BUILD_NUMBER}/g" > ./helm-config-ready/mariadb-galera/dev.yaml
cat ./helm-config/mariadb-galera/staging.yaml | sed "s/BUILD_NUMBER/${CI_BUILD_NUMBER}/g" > ./helm-config-ready/mariadb-galera/staging.yaml
cat ./helm-config/mariadb-galera/prod.yaml | sed "s/BUILD_NUMBER/${CI_BUILD_NUMBER}/g" > ./helm-config-ready/mariadb-galera/prod.yaml

cd 	/drone/src/k8s/
echo "customizing master files..."
REPO_USER=`echo ${DRONE_REPO} | awk -F'/' '{print $1}'`
echo "Drone Repo: ${DRONE_REPO}"
echo "Repo User: ${REPO_USER}"
mkdir -p ./ingress-ready/
cat ./ingress/master.yaml | sed "s/REPO_USER/${REPO_USER}/g" > ./ingress-ready/master.yaml