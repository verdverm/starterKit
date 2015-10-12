#!/bin/bash

ipython profile create nbserver

printf "\n\n\n"

cp /scripts/ipython_notebook_config.py /ipython/profile_nbserver/ipython_notebook_config.py

mkdir /ipython/notebooks
cd /ipython/notebooks

ipython notebook --profile=nbserver
