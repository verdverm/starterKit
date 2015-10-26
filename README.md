# starterKit
Web app starterKit





## BE setup

Steps to setup the python environment
---------

Postgresql requirement: `apt-get install libpq-dev`


1. `cd be/django`
1. `virtualenv venv`
1. `source venv/bin/activate`
1. `pip install -r requirements.txt`


local_settings.py
-----------------

create the file `django/app/local_settings.py` with the following contents

```
import os

EMAIL_HOST_USER="..."
EMAIL_HOST_PASSWORD="..."

FACEBOOK_SECRET = os.environ.get('FACEBOOK_SECRET') or '...'
GOOGLE_SECRET = os.environ.get('GOOGLE_SECRET') or '...'
YAHOO_SECRET = os.environ.get('YAHOO_SECRET') or '...'
WINDOWS_SECRET = os.environ.get('WINDOWS_SECRET') or '...'
GITHUB_SECRET = os.environ.get('GITHUB_SECRET') or '...'
TWITTER_CONSUMER_KEY = os.environ.get('TWITTER_CONSUMER_KEY') or '...'
TWITTER_CONSUMER_SECRET = os.environ.get('TWITTER_CONSUMER_SECRET') or '...'
TWITTER_CALLBACK_URL = os.environ.get('TWITTER_CALLBACK_URL') or 'http://localhost:3000'
SOUNDCLOUD_SECRET = os.environ.get('SOUNDCLOUD_SECRET') or '...'
SPOTIFY_SECRET = os.environ.get('SPOTIFY_SECRET') or '...'
DROPBOX_SECRET = os.environ.get('DROPBOX_SECRET') or '...'
```

#### Setting up email

The email system is setup to use a gmail account by default.
You will need an app specific password

Mailchimp coming later...

#### Setting up OAuth

you will need to modify the `local_settings.py` file 
with the secrets.

A number of OAuth providers are available by default.
Each has their own interface and particulars.

See the [docs/OAUTH.md] file for details


DB setup
---------

Requires PostgreSQL

when setting up the DB for the first time, follow these steps:

1. `createdb starterkit`
1. comment all installed apps below `'app'` in django/app/settings.py
1. `python manage.py migrate`
1. restore all installed apps below `'app'` in django/app/settings.py
1. `python manage.py migrate`
1. `python manage.py createsuperuser`

any time there is a change to a model run:

1. `python manage.py makemigrations`
1. `python manage.py migrate`


Run the Backend
-----------------

From the django directory, run: `python manage.py runserver`



## Setting up the Frontend

Install dependencies
--------------------

Requires the following libraries:
`npm install -g gulp gulp-util karma karma-cli webpack`

install the local dependencies:

`npm install`

Run the Frontend
----------------

`gulp` or `gulp default`

Building for production
-----------------------

`gulp build` will create a bundle.min.js




## Contributing

We have lots of contributor ready 

Fork and clone this repository

Branching practices follow the methodology outlined at: 
[http://nvie.com/posts/a-successful-git-branching-model/](http://nvie.com/posts/a-successful-git-branching-model/)