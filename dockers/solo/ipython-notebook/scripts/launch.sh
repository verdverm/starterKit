#!/bin/bash

ipython profile create

printf "\n\n\n"

mkdir -p /ipython/notebooks
cd /ipython/notebooks

# start nginx daemon
nginx

# start ipython in foreground
ipython notebook --matplotlib=inline
