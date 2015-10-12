#!/bin/bash

images=(

	# BASES
	base     
	java8    
	python   
	python-ml
	golang

	# SERVICES
	#  DATABASES
	psql      
	neo4j     
	couchdb   
	#  CACHING
	redis    
	#  MESSAGING
	rabbitmq 
	#  LANGUAGES
	goserv   
	pyserv   
	pyserv-ml

	# FRONTEND
	flask
	ipynb
	nginx

	# APPS
	# goapp 
	# pyapp
)

for img in ${images[@]}; do
	docker pull zaha/$img
done

