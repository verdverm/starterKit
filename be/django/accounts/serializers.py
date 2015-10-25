from rest_framework import serializers

from accounts.models import UserAccounts

class UserAccountsSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = UserAccounts
        fields = (
        	'facebook', 
        	'google', 
        	'yahoo', 
        	'windows', 
        	'github', 
        	'twitter',
        	'soundcloud',
        	'dropbox',
        	)

