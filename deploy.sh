#!/bin/bash

echo "customizing Deployment files..."
mkdir ./deployment-ready/
cd ./deployment
for file in `ls *.yaml`
do
  echo "Working on $file"
  cat $file | sed "s/BUILD_NUMBER/${CI_BUILD_NUMBER}/g" > ../deployment-ready/"$file"
done
