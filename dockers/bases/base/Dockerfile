# Use phusion/baseimage as base image. 
# See https://github.com/phusion/baseimage-docker
FROM phusion/baseimage:0.9.15
MAINTAINER Tony Worm docker@zaha.io


# Set correct environment variables.
ENV HOME /root

# Regenerate SSH host keys. baseimage-docker does not contain any
RUN /etc/my_init.d/00_regen_ssh_host_keys.sh

# Use baseimage-docker's init system.
CMD ["/sbin/my_init"]


### .........Custom Stuff............... ###

ADD dotfiles/profile /root/.myprofile
RUN echo "source /root/.myprofile" >> /root/.bashrc


# Update
RUN sudo apt-get update && sudo apt-get upgrade -y

# Install software
RUN sudo apt-get install -y \
	build-essential

### ........End Custom Stuff............ ###


# Clean up APT when done.
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
