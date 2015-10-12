FROM zaha/java8
MAINTAINER Tony Worm docker@zaha.io

## install neo4j according to http://www.neo4j.org/download/linux
# Import neo4j signing key
RUN wget -O - http://debian.neo4j.org/neotechnology.gpg.key | apt-key add - 
# Create an apt sources.list file
RUN echo 'deb http://debian.neo4j.org/repo stable/' > /etc/apt/sources.list.d/neo4j.list
# Find out about the files in neo4j repo ; install neo4j community edition
RUN apt-get update ; apt-get install neo4j -y


# Clean up APT when done.
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

## turn on indexing: http://chrislarson.me/blog/install-neo4j-graph-database-ubuntu
## enable neo4j indexing, and set indexable keys to name,age
RUN sed -i "s|#node_auto_indexing|node_auto_indexing|g" /var/lib/neo4j/conf/neo4j.properties
RUN sed -i "s|#node_keys_indexable|node_keys_indexable|g" /var/lib/neo4j/conf/neo4j.properties


## add launcher and set execute property
ADD scripts /scripts
RUN chmod +x /scripts/launch.sh

# Add daemon to be run by runit.
RUN mkdir /etc/service/neo4j
RUN ln -s /scripts/launch.sh /etc/service/neo4j/run

# Expose our port && data, log, and configuration directories.
EXPOSE 7474
VOLUME ["/var/lib/neo4j/data"]

# Use baseimage-docker's init system.
CMD ["/sbin/my_init"]
