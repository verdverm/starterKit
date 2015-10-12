# Zaha!

Simplify your application development.

```
// create and run apps locally with docker
zaha create myapp  // myapp is now the context
zaha db init
zaha start
zaha stop

zaha set app myOtherApp
zaha status

// create and control a mesosphere cluster 
zaha cloud create mycluster aws 4
zaha cloud status
zaha cloud resize mycluster 20
zaha cloud destroy mycluster

// push your app to the cloud
zaha push myapp mycluster
zaha resize flask 3
```

**...video...**


Zaha! is built on several important technologies:
- [Phusion/baseimage-docker](http://phusion.github.io/baseimage-docker/)
- [Mesosphere](http://mesosphere.com/)
- [Polymer](https://www.polymer-project.org/)

When you use Zaha! you'll get all this right out of the box.

- Security
- Scalability
- Fault-tolerance
- Mobile-first design
- A menagerie of app components
- Plug-in services
- Asset management
- User and admin accounts
- Live WYSIWYG content editing with drafts and history




With Zaha! you can focus on what you do best, building awesome apps!


### Contents

> Caution, the Zaha! toolchain is young and under rapid development. Expect change.


- [Install](#install)
- [Apps](#apps)
- [Clouds](#clouds)
- [Components](#components)
- [Features](#features)
- [Configuration](#configuration")
- [About](#about)
- [Contribute](#contribute)
















### Install

##### Dependencies

- [Docker](http://docker.com)
- [Golang](http://golang.org)
- Bower (if you want to use polymer)
- Ruby & Node for asset management

##### Go Get It!

``` Bash
go get -U github.com/zaha-io/zaha
zaha setup
```

Now run any command from any directory... Zaha!

> Note: Zaha! will create a '~/.zaha' directory 
> for storing configuration, templates,
> and information about your apps and clouds.













### Apps

Zaha! can create new web applications with a single command
and customize its components to fit your needs.

``` Bash
zaha app create <appName>
```

This creates a new Zaha! app called `<name>` in the current directory.
It also sets your current context to point to the new application.


```
sample output...
```

##### App Directory Layout

- clouds (your cloud configuration files)
- config (your app configuration files)
- dockers (your component configuration files)
- static (static application files)
- storage (persistent storage, don't touch!)
- webapp (your application code)

##### App Workflow

Start the app locally and you're ready to go

```
zaha start app
```

Open *chrome* to `localhost`

Zaha! also has live code in `dev` mode.



##### Working with multiple apps

Zaha! apps are completely contained within a single directory.
This means you can work with multiple apps by switching between them.

```
zaha set app <appName>
```

Worry not about where your code is, Zaha! keeps track for you.
You can even move directories around, but you will have to
change the copy of your config file in the '~/.zaha/apps' directory.























### Clouds

Zaha! makes application deployment a breeze.

##### Create a cloud
```
zaha cloud create <cloudName> <cloudProvider> <workerCount>
```

This will create a Mesosphere cluster 
on the cloud provider of your choice.
As of now only Digital Ocean is functional
with plans to expand to AWS and GCE.
If you have a Mesosphere master and slave
running on `localhost`, you may use that as well.


##### Resize your cloud
```
zaha cloud resize <cloudName> <newWorkerCount>
```


##### Deploying your app

Deploying your code requires two commands, a push and a restart.
Rolling updates can be enabled by setting the appropriate values
in the Mesosphere configuration for the component.

```
// first push your code to the cloud
zaha push <appName> <cloudName>

// then restart components on a remote cloud
zaha restart <component> <cloudName>
```




















### Components

Your Zaha! app will have several components 
spanning the technology stack which are
divided into functional groups.

**Frontend**
- Nginx
- Redis
- Flask

**Databases**

- Postgresql
- CouchDB
- Neo4j

**Backend**

- Ipython
- RabbitMQ

**Services**

You will find a few docker images 
for quickly creating micro-services 
which can talk to each other directly or through the messaging backend.

- Go based server
- Python based server
- Python based server with ML goodies.

> More components coming soon

##### Managing component dependencies

Each component has a yaml configuration file. 
Within that app you can add or remove any dependencies












### Features






### Configuration

##### Environment Variables

```
ZAHA_RUNMODE="dev"
ZAHA_APPDOMAIN="zaha.io"

ZAHA_SITE_OWNER_NAME="tony"
ZAHA_SITE_OWNER_PASS="secret"

ZAHA_CSRF_KEY="secret"
ZAHA_FLASK_KEY="secret"
ZAHA_SALT_LICK="secret"

ZAHA_GMAIL_SENDER="support@zaha.io"
ZAHA_GMAIL_USER="tony@zaha.io"
ZAHA_GMAIL_PASS="secret"

ZAHA_SQL_ENGINE="postgresql"
ZAHA_SQL_USER="super"
ZAHA_SQL_PASS="secret"
ZAHA_SQL_DB="somename"

These depend upon your situation

DIGITAL_OCEAN_TOKEN="..."
MAILCHIMP_API_KEY="..."
```



### About

Zaha! is named after a [great architect](http://www.zaha-hadid.com/)



### Contribute

Zaha! is on a journey to simplify and streamline application development.
We would be happy to have you join us on!

Here is an incomplete list of contributor ready functionality:

- GCE & AWS integration
- VPN integration
- Deal with docker and the iptables shenanigans 
- Adding your favorite components
- Application level RBAC
- Spruce up the live editting
- Interactive app and cloud creating
- Dockers & CLI commands for asset management
- Merge configuration directories into one `config` file
- Web UIs for the everything
- Dockerize `zaha setup` and `zaha create` and even the tool so the only dependency is Docker.
- ...

