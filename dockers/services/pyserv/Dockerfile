FROM zaha/python
MAINTAINER Tony Worm docker@zaha.io

## add launcher and set execute property
ADD scripts /scripts
RUN chmod +x /scripts/launch.sh

# Add daemon to be run by runit.
RUN mkdir /etc/service/flask
RUN ln -s /scripts/launch.sh /etc/service/flask/run

RUN pip install -r /scripts/requirements.txt  

EXPOSE 5000

VOLUME ["/src"]
WORKDIR /src

# Use baseimage-docker's init system.
CMD ["/sbin/my_init"]
