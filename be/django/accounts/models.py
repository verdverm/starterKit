from django.db import models
# from django.contrib.auth.models import User

from django.conf import settings

# adapted from https://github.com/sahat/satellizer


class UserAccounts(models.Model):
    user = models.OneToOneField(settings.AUTH_USER_MODEL, primary_key=True)

    facebook   = models.BooleanField(default=False)
    google     = models.BooleanField(default=False)
    yahoo      = models.BooleanField(default=False)
    windows    = models.BooleanField(default=False)
    github     = models.BooleanField(default=False)
    twitter    = models.BooleanField(default=False)
    soundcloud = models.BooleanField(default=False)
    dropbox    = models.BooleanField(default=False)

    facebook_code  = models.CharField(max_length=512, blank=True, default='')
    facebook_token = models.CharField(max_length=512, blank=True, default='')
    facebook_id    = models.CharField(max_length=120, blank=True, default='')
    fackbook_email = models.CharField(max_length=120, blank=True, default='')
    fackbook_name  = models.CharField(max_length=120, blank=True, default='')

    google_code    = models.CharField(max_length=64, blank=True, default='')
    google_token   = models.CharField(max_length=512, blank=True, default='')
    google_id      = models.CharField(max_length=120, blank=True, default='')
    google_email   = models.CharField(max_length=120, blank=True, default='')
    google_name    = models.CharField(max_length=120, blank=True, default='')
 
    github_code    = models.CharField(max_length=64, blank=True, default='')
    github_token   = models.CharField(max_length=512, blank=True, default='')
    github_id      = models.CharField(max_length=120, blank=True, default='')
    github_name    = models.CharField(max_length=120, blank=True, default='')
    github_login   = models.CharField(max_length=120, blank=True, default='')
    github_email   = models.CharField(max_length=120, blank=True, default='')

    windows_code   = models.CharField(max_length=64, blank=True, default='')
    windows_token  = models.CharField(max_length=1024, blank=True, default='')
    windows_id     = models.CharField(max_length=120, blank=True, default='')
    windows_name   = models.CharField(max_length=120, blank=True, default='')
    windows_email  = models.CharField(max_length=120, blank=True, default='')

    yahoo_token = models.CharField(max_length=512, blank=True, default='')

    twitter_token  = models.CharField(max_length=120, blank=True, default='')
    twitter_secret = models.CharField(max_length=120, blank=True, default='')
    twitter_id     = models.CharField(max_length=120, blank=True, default='')
    twitter_name   = models.CharField(max_length=120, blank=True, default='')

    dropbox_code   = models.CharField(max_length=64, blank=True, default='')
    dropdox_token  = models.CharField(max_length=512, blank=True, default='')
    dropbox_id     = models.CharField(max_length=120, blank=True, default='')
    dropbox_name   = models.CharField(max_length=120, blank=True, default='')
    dropbox_email  = models.CharField(max_length=120, blank=True, default='')
