# starterKit
Web app starterKit


DB setup
---------

when setting up the DB for the first time, follow these steps:

1. `createdb starterkit`
1. comment all installed apps below `'app'` in settings.py
1. `python manage.py migrate`
1. restore all installed apps below `'app'` in settings.py
1. `python manage.py migrate`
1. `python manage.py createsuperuser`



any time there is a change to a model run:

1. `python manage.py makemigrations`
1. `python manage.py migrate`

