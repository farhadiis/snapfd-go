#!/bin/bash


cd order
docker build -t order:1.0.0 .
cd ..

cd order-process
docker build -t order-process:1.0.0 .
cd ..
