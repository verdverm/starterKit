from rest_framework import status
from rest_framework.decorators import api_view, authentication_classes, permission_classes, throttle_classes
from rest_framework.response import Response
from rest_framework import permissions
from app.permissions import IsOwnerOrReadOnly, IsOwnerOnly


from django.contrib.auth.models import User
from accounts.models import UserAccounts
from accounts.serializers import UserAccountsSerializer

import app.local_settings as settings

import requests
from requests_oauthlib import OAuth1
import json
from urlparse import parse_qs, parse_qsl
from urllib import urlencode

from pprint import pprint

from django.core.exceptions import ObjectDoesNotExist

from rest_framework_jwt.settings import api_settings


def gen_jwt_token(user):
    jwt_payload_handler = api_settings.JWT_PAYLOAD_HANDLER
    jwt_encode_handler = api_settings.JWT_ENCODE_HANDLER

    payload = jwt_payload_handler(user)
    token = jwt_encode_handler(payload)
    return token



@api_view(['POST'])
# @authentication_classes([])  # use default
@permission_classes([permissions.IsAuthenticated, IsOwnerOnly])
@throttle_classes([])
def unlink(request):

    print "request.user:", request.user
    print "request.auth:", request.auth
    print "request.jwtuser:", request.jwtuser
    print "request.jwtauth:", request.jwtauth

    provider = request.data['provider']

    # hack!
    if provider == "live":
        provider = "windows"
    print "UNLINKING: ", provider

    current_user = request.jwtuser

    user = User.objects.get(pk=current_user.id)
    if user is None:
        print "NO USER!!"
        ### MAKE A NEW USER!!!

    try:
        if hasattr( user, 'useraccounts' ):
            print "HAS ACCOUNTS"
            accounts = user.useraccounts
            
            if getattr(accounts, provider) == False:
                return Response("Error: provider not linked " + provider)
            
            setattr(accounts, provider, False)
            accounts.save()

        else:
            print "NO ACCOUNTS"
            return Response("Error: no accounts")

        return Response("OK")
    except Exception, e:
        print "EXCEPTION: ", e
        return Response("EXCEPTION: " + str(e), status=500)


@api_view(['GET'])
# @authentication_classes([])
@permission_classes([permissions.IsAuthenticated, IsOwnerOnly])
@throttle_classes([])
def accounts(request):

    print "GETTING ACCOUNTS"

    print "request.user:", request.user
    print "request.auth:", request.auth
    print "request.jwtuser:", request.jwtuser
    print "request.jwtauth:", request.jwtauth

    current_user = request.jwtuser

    user = User.objects.get(pk=current_user.id)
    if user is None:
        print "NO USER!!"
        ### MAKE A NEW USER!!!

    try:
        if hasattr( user, 'useraccounts' ):
            print "HAS ACCOUNTS"
            accounts = user.useraccounts
            serializer = UserAccountsSerializer(accounts, context={'request': request})
            print serializer.data
            return Response(serializer.data)
        else:
            print "NO ACCOUNTS"
            return Response("Error: no accounts")

        return Response("success")
    except Exception, e:
        print "EXCEPTION: ", e
        return Response(str(e), status=500)



@api_view(['POST'])
@authentication_classes([])
@permission_classes([])
@throttle_classes([])
def facebook(request):
    access_token_url = 'https://graph.facebook.com/v2.3/oauth/access_token'
    graph_api_url = 'https://graph.facebook.com/v2.3/me'

    # print request.data

    params = {
        'client_id': request.data['clientId'],
        'redirect_uri': request.data['redirectUri'],
        'client_secret': settings.FACEBOOK_SECRET,
        'code': request.data['code']
    }

    # Step 1. Exchange authorization code for access token.
    r = requests.get(access_token_url, params=params)
    access_token = json.loads(r.text)

    # Step 2. Retrieve information about the current user.
    r = requests.get(graph_api_url, params=access_token)
    profile = json.loads(r.text)


    profile["token"] = request.jwtauth
    current_user = request.jwtuser

    user = User.objects.get(pk=current_user.id)
    if user is None:
        print "NO USER!!"
        ### MAKE A NEW USER!!!

    # print len(params["code"])
    # print len(access_token)

    if hasattr( user, 'useraccounts' ):
        print "HAS ACCOUNTS"
        accounts = user.useraccounts
        accounts.facebook = True
        accounts.facebook_code = request.data['code']
        accounts.facebook_token = access_token['access_token']
        accounts.facebook_id = profile["id"]
        accounts.facebook_name = profile["name"]
        accounts.save()
    else:
        print "NO ACCOUNTS"
        accounts = UserAccounts()
        accounts.user = user
        accounts.facebook = True
        accounts.facebook_code = request.data['code']
        accounts.facebook_token = access_token['access_token']
        accounts.facebook_id = profile["id"]
        accounts.facebook_name = profile["name"]
        accounts.save()
        user.useraccounts = accounts
        user.save()
        print user.useraccounts

    print profile
    return Response(profile)




@api_view(['POST'])
@authentication_classes([])
@permission_classes([])
@throttle_classes([])
def google(request):
    access_token_url = 'https://accounts.google.com/o/oauth2/token'
    people_api_url = 'https://www.googleapis.com/plus/v1/people/me/openIdConnect'

    payload = {
        'client_id': request.data['clientId'],
        'redirect_uri': request.data['redirectUri'],
        'client_secret': settings.GOOGLE_SECRET,
        'code': request.data['code'],
        'grant_type': 'authorization_code'
    }

   # Step 1. Exchange authorization code for access token.
    r = requests.post(access_token_url, data=payload)
    token = json.loads(r.text)
    headers = {'Authorization': 'Bearer {0}'.format(token['access_token'])}

    # Step 2. Retrieve information about the current user.
    r = requests.get(people_api_url, headers=headers)
    profile = json.loads(r.text)

    print profile

    profile["token"] = request.jwtauth
    current_user = request.jwtuser

    user = User.objects.get(pk=current_user.id)
    if user is None:
        print "NO USER!!"
        ### MAKE A NEW USER!!!

    # print len(params["code"])
    # print len(access_token)

    if hasattr( user, 'useraccounts' ):
        print "HAS ACCOUNTS"
        accounts = user.useraccounts
        accounts.google = True
        accounts.google_code = request.data['code']
        accounts.google_token = token['access_token']
        accounts.google_id = profile["sub"]
        accounts.google_name = profile["name"]
        accounts.google_email = profile["email"]
        accounts.save()
    else:
        print "NO ACCOUNTS"
        accounts = UserAccounts()
        accounts.user = user
        accounts.google = True
        accounts.google_code = request.data['code']
        accounts.google_token = token['access_token']
        accounts.google_id = profile["sub"]
        accounts.google_name = profile["name"]
        accounts.google_email = profile["email"]
        accounts.save()
        user.useraccounts = accounts
        user.save()
        print user.useraccounts

    print profile
    return Response(profile)


@api_view(['POST'])
@authentication_classes([])
@permission_classes([])
@throttle_classes([])
def yahoo(request):
    # access_token_url = 'https://graph.facebook.com/v2.3/oauth/access_token'
    # graph_api_url = 'https://graph.facebook.com/v2.3/me'

    print request.data

    # params = {
    #     'client_id': request.data['clientId'],
    #     'redirect_uri': request.data['redirectUri'],
    #     'client_secret': settings.FACEBOOK_SECRET,
    #     'code': request.data['code']
    # }

    # # Step 1. Exchange authorization code for access token.
    # r = requests.get(access_token_url, params=params)
    # access_token = json.loads(r.text)

    # # Step 2. Retrieve information about the current user.
    # r = requests.get(graph_api_url, params=access_token)
    # profile = json.loads(r.text)


    # profile["token"] = request.jwtauth
    current_user = request.jwtuser

    user = User.objects.get(pk=current_user.id)
    if user is None:
        print "NO USER!!"
        ### MAKE A NEW USER!!!

    return Response("dev roadblock")

    # print len(params["code"])
    # print len(access_token)

    # if hasattr( user, 'useraccounts' ):
    #     print "HAS ACCOUNTS"
    #     accounts = user.useraccounts
    #     accounts.facebook = True
    #     accounts.facebook_token = access_token['access_token']
    #     accounts.facebook_id = profile["id"]
    #     accounts.facebook_name = profile["name"]
    #     accounts.save()
    # else:
    #     print "NO ACCOUNTS"
    #     accounts = UserAccounts()
    #     accounts.user = user
    #     accounts.facebook = True
    #     accounts.facebook_token = access_token['access_token']
    #     accounts.facebook_id = profile["id"]
    #     accounts.facebook_name = profile["name"]
    #     accounts.save()
    #     user.useraccounts = accounts
    #     user.save()
    #     print user.useraccounts

    # print profile
    # return Response(profile)


@api_view(['POST'])
@authentication_classes([])
@permission_classes([])
@throttle_classes([])
def windows(request):
    access_token_url = 'https://login.live.com/oauth20_token.srf'
    profile_api_url = 'https://apis.live.net/v5.0/me'

    print request.data

    params = {
        'client_id': request.data['clientId'],
        'redirect_uri': request.data['redirectUri'],
        'client_secret': settings.WINDOWS_SECRET,
        'code': request.data['code'],
        'grant_type': 'authorization_code'
   }

    # Step 1. Exchange authorization code for access token.
    r = requests.get(access_token_url, params=params)
    access_token = json.loads(r.text)
    print "TOKEN: "
    pprint(access_token)
    print len(access_token['access_token'])


    # # Step 2. Retrieve information about the current user.
    r = requests.get(profile_api_url, params=access_token)
    profile = json.loads(r.text)
    print "PROFILE: "
    pprint(profile)


    current_user = request.jwtuser

    user = User.objects.get(pk=current_user.id)
    if user is None:
        print "NO USER!!"
        ### MAKE A NEW USER!!!

    def saver(accounts):
        if profile["name"] is None:
            profile["name"] = ""
        accounts.windows = True
        accounts.windows_code = request.data['code']
        accounts.windows_token = access_token['access_token']
        accounts.windows_id = profile["id"]
        accounts.windows_name = profile["name"]
        accounts.windows_email = profile["emails"]["account"]
        accounts.save()

    # return Response("dev roadblock")

    if hasattr( user, 'useraccounts' ):
        print "HAS ACCOUNTS"
        accounts = user.useraccounts
        saver(accounts)
    else:
        print "NO ACCOUNTS"
        accounts = UserAccounts()
        accounts.user = user
        saver(accounts)

        user.useraccounts = accounts
        user.save()
        print user.useraccounts

    profile["token"] = request.jwtauth
    print profile
    return Response(profile)



@api_view(['POST'])
@authentication_classes([])
@permission_classes([])
@throttle_classes([])
def github(request):
    access_token_url = 'https://github.com/login/oauth/access_token'
    users_api_url = 'https://api.github.com/user'

    print request.data

    params = {
        'client_id': request.data['clientId'],
        'redirect_uri': request.data['redirectUri'],
        'client_secret': settings.GITHUB_SECRET,
        'code': request.data['code']
    }

    # Step 1. Exchange authorization code for access token.
    r = requests.get(access_token_url, params=params)
    access_token = dict(parse_qsl(r.text))
    headers = {'User-Agent': 'starterKit'}

    # Step 2. Retrieve information about the current user.
    r = requests.get(users_api_url, params=access_token, headers=headers)
    profile = json.loads(r.text)


    profile["token"] = request.jwtauth
    current_user = request.jwtuser

    user = User.objects.get(pk=current_user.id)
    if user is None:
        print "NO USER!!"
        ### MAKE A NEW USER!!!

    if hasattr( user, 'useraccounts' ):
        print "HAS ACCOUNTS"
        accounts = user.useraccounts
        accounts.github = True
        accounts.github_code = request.data['code']
        accounts.github_token = access_token['access_token']
        accounts.github_id = profile["id"]
        accounts.github_name = profile["name"]
        accounts.github_login = profile["login"]
        accounts.github_email = profile["email"]
        accounts.save()
    else:
        print "NO ACCOUNTS"
        accounts = UserAccounts()
        accounts.user = user
        accounts.github = True
        accounts.github_code = request.data['code']
        accounts.github_token = access_token['access_token']
        accounts.github_id = profile["id"]
        accounts.github_name = profile["name"]
        accounts.github_login = profile["login"]
        accounts.github_email = profile["email"]
        accounts.save()
        user.useraccounts = accounts
        user.save()
        print user.useraccounts

    pprint(profile)
    return Response(profile)





@api_view(['POST'])
@authentication_classes([])
@permission_classes([])
@throttle_classes([])
def twitter(request):
    request_token_url = 'https://api.twitter.com/oauth/request_token'
    access_token_url = 'https://api.twitter.com/oauth/access_token'

    print request.data

    if request.data.get('oauth_token') and request.data.get('oauth_verifier'):
        auth = OAuth1(settings.TWITTER_CONSUMER_KEY,
                      client_secret=settings.TWITTER_CONSUMER_SECRET,
                      resource_owner_key=request.data.get('oauth_token'),
                      verifier=request.data.get('oauth_verifier'))
        r = requests.post(access_token_url, auth=auth)
        profile = dict(parse_qsl(r.text))
        print "PROFILE:"
        pprint(profile)

        current_user = request.jwtuser

        user = User.objects.get(pk=current_user.id)
        if user is None:
            print "NO USER!!"
            ### MAKE A NEW USER!!!

        def saver(accounts):
            accounts.twitter = True
            accounts.twitter_token = profile['oauth_token']
            accounts.twitter_secret = profile['oauth_token_secret']
            accounts.twitter_id = profile["user_id"]
            accounts.twitter_name = profile["screen_name"]
            # accounts.twitter_email = profile["email"]
            accounts.save()

        # return Response("dev roadblock")

        if hasattr( user, 'useraccounts' ):
            print "HAS ACCOUNTS"
            accounts = user.useraccounts
            pprint(accounts)
            saver(accounts)
        else:
            print "NO ACCOUNTS"
            accounts = UserAccounts()
            accounts.user = user
            saver(accounts)

            user.useraccounts = accounts
            user.save()
            print user.useraccounts

        profile["token"] = request.jwtauth
        pprint(profile);
        return Response(profile)
    else:
        oauth = OAuth1(settings.TWITTER_CONSUMER_KEY,
                       client_secret=settings.TWITTER_CONSUMER_SECRET,
                       callback_uri=settings.TWITTER_CALLBACK_URL)
        r = requests.post(request_token_url, auth=oauth)
        print "OAUTH_TOKEN: ", r.text
        oauth_token = dict(parse_qsl(r.text))
        oauth_token["token"] = request.jwtauth
        print "OAUTH_TOKENBB: ", oauth_token
        return Response(oauth_token)



@api_view(['POST'])
@authentication_classes([])
@permission_classes([])
@throttle_classes([])
def soundcloud(request):
    return Response("soundcloud!")




@api_view(['POST'])
@authentication_classes([])
@permission_classes([])
@throttle_classes([])
def dropbox(request):
    access_token_url = 'https://api.dropboxapi.com/1/oauth2/token'
    profile_api_url = 'https://api.dropboxapi.com/1/account/info'

    pprint(request.data)

    params = {
        'client_id': request.data['clientId'],
        'redirect_uri': request.data['redirectUri'],
        'client_secret': settings.DROPBOX_SECRET,
        'code': request.data['code'],
        'grant_type': 'authorization_code'
   }

    # Step 1. Exchange authorization code for access token.
    r = requests.post(access_token_url, params=params)
    print "TOKEN: ", r.text
    access_token = json.loads(r.text)
    pprint(access_token)
    print len(access_token['access_token'])

    headers = {'Authorization': 'Bearer {0}'.format(access_token['access_token'])}

    # # Step 2. Retrieve information about the current user.
    r = requests.get(profile_api_url, headers=headers)
    profile = json.loads(r.text)
    print "PROFILE: "
    pprint(profile)


    current_user = request.jwtuser

    user = User.objects.get(pk=current_user.id)
    if user is None:
        print "NO USER!!"
        ### MAKE A NEW USER!!!

    def saver(accounts):
        accounts.dropbox = True
        accounts.dropbox_code = request.data['code']
        accounts.dropbox_token = access_token['access_token']
        accounts.dropbox_id = profile["uid"]
        accounts.dropbox_name = profile["display_name"]
        accounts.dropbox_email = profile["email"]
        accounts.save()

    # return Response("dev roadblock")

    if hasattr( user, 'useraccounts' ):
        print "HAS ACCOUNTS"
        accounts = user.useraccounts
        pprint(accounts)
        saver(accounts)
    else:
        print "NO ACCOUNTS"
        accounts = UserAccounts()
        accounts.user = user
        saver(accounts)

        user.useraccounts = accounts
        user.save()
        print user.useraccounts

    profile["token"] = request.jwtauth
    print profile
    return Response(profile)
