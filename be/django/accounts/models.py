from django.db import models
from django.contrib.auth.models import User

# adapted from https://github.com/sahat/satellizer


class UserAccounts(models.Model):
    user = models.OneToOneField(User)

    facebook = models.CharField(max_length=120, blank=True, default='')
    github = models.CharField(max_length=120, blank=True, default='')
    google = models.CharField(max_length=120, blank=True, default='')
    linkedin = models.CharField(max_length=120, blank=True, default='')
    twitter = models.CharField(max_length=120, blank=True, default='')


