from django.db import models
# from django.contrib.auth.models import User

from django.conf import settings

# adapted from https://github.com/sahat/satellizer


class UserAccounts(models.Model):
    user = models.OneToOneField(settings.AUTH_USER_MODEL, primary_key=True)

    facebook = models.CharField(max_length=512, blank=True, default='')
    github = models.CharField(max_length=512, blank=True, default='')
    google = models.CharField(max_length=512, blank=True, default='')
    linkedin = models.CharField(max_length=512, blank=True, default='')
    twitter = models.CharField(max_length=512, blank=True, default='')

    fackbook_id = models.PositiveIntegerField(default=0)
    facebook_code = models.CharField(max_length=120, blank=True, default='')
    fackbook_email = models.CharField(max_length=120, blank=True, default='')
    fackbook_name = models.CharField(max_length=120, blank=True, default='')


