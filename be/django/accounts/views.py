from rest_framework import status
from rest_framework.decorators import api_view, authentication_classes, permission_classes, throttle_classes
from rest_framework.response import Response
from rest_framework import permissions

from django.contrib.auth.models import User
from accounts.models import UserAccounts
from accounts.serializers import UserAccountsSerializer

import app.local_settings as settings

import requests
import json
from urlparse import parse_qs, parse_qsl

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
@authentication_classes([])
@permission_classes([])
@throttle_classes([])
def unlink(request):
    print request.data

    return Response("success")





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

    # print "request.user:", request.user
    # print "request.auth:", request.auth
    # print "request.jwtuser:", request.jwtuser
    # print "request.jwtauth:", request.jwtauth
    profile["token"] = request.jwtauth
    current_user = request.jwtuser

    user = User.objects.get(pk=current_user.id)
    if user is None:
        print "NO USER!!"
        ### MAKE A NEW USER!!!

    print len(params["code"])
    print len(access_token)

    if hasattr( user, 'useraccounts' ):
        print "HAS ACCOUNTS"
        accounts = user.useraccounts
        accounts.facebook = access_token
        accounts.facebook_id = profile["id"]
        accounts.facebook_name = profile["name"]
        accounts.save()
    else:
        print "NO ACCOUNTS"
        accounts = UserAccounts()
        accounts.user = user
        accounts.facebook = params["code"]
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
def github(request):
    access_token_url = 'https://github.com/login/oauth/access_token'
    users_api_url = 'https://api.github.com/user'

    params = {
        'client_id': request.json['clientId'],
        'redirect_uri': request.json['redirectUri'],
        'client_secret': app.config['GITHUB_SECRET'],
        'code': request.json['code']
    }

    # Step 1. Exchange authorization code for access token.
    r = requests.get(access_token_url, params=params)
    access_token = dict(parse_qsl(r.text))
    headers = {'User-Agent': 'Satellizer'}

    # Step 2. Retrieve information about the current user.
    r = requests.get(users_api_url, params=access_token, headers=headers)
    profile = json.loads(r.text)

    print profile

    # # Step 3. (optional) Link accounts.
    # if request.headers.get('Authorization'):
    #     user = User.query.filter_by(github=profile['id']).first()
    #     if user:
    #         response = jsonify(message='There is already a GitHub account that belongs to you')
    #         response.status_code = 409
    #         return response

    #     payload = parse_token(request)

    #     user = User.query.filter_by(id=payload['sub']).first()
    #     if not user:
    #         response = jsonify(message='User not found')
    #         response.status_code = 400
    #         return response

    #     u = User(github=profile['id'], display_name=profile['name'])
    #     db.session.add(u)
    #     db.session.commit()
    #     token = create_token(u)
    #     return jsonify(token=token)

    # # Step 4. Create a new account or return an existing one.
    # user = User.query.filter_by(github=profile['id']).first()
    # if user:
    #     token = create_token(user)
    #     return jsonify(token=token)

    # u = User(github=profile['id'], display_name=profile['name'])
    # db.session.add(u)
    # db.session.commit()
    # token = create_token(u)
    # return jsonify(token=token)



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

    # print "request.user:", request.user
    # print "request.auth:", request.auth
    # print "request.jwtuser:", request.jwtuser
    # print "request.jwtauth:", request.jwtauth
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
        # accounts = user.useraccounts
        # accounts.facebook = access_token
        # accounts.facebook_id = profile["id"]
        # accounts.facebook_name = profile["name"]
        # accounts.save()
    else:
        print "NO ACCOUNTS"
        # accounts = UserAccounts()
        # accounts.user = user
        # accounts.facebook = params["code"]
        # accounts.facebook_id = profile["id"]
        # accounts.facebook_name = profile["name"]
        # accounts.save()
        # user.useraccounts = accounts
        # user.save()
        # print user.useraccounts

    print profile
    return Response(profile)

    # user = User.query.filter_by(google=profile['sub']).first()
    # if user:
    #     token = create_token(user)
    #     return jsonify(token=token)
    # u = User(google=profile['sub'],
    #          display_name=profile['name'])
    # db.session.add(u)
    # db.session.commit()
    # token = create_token(u)
    # return jsonify(token=token)


@api_view(['POST'])
@authentication_classes([])
@permission_classes([])
@throttle_classes([])
def linkedin(request):
    access_token_url = 'https://www.linkedin.com/uas/oauth2/accessToken'
    people_api_url = 'https://api.linkedin.com/v1/people/~:(id,first-name,last-name,email-address)'

    payload = dict(client_id=request.json['clientId'],
                   redirect_uri=request.json['redirectUri'],
                   client_secret=app.config['LINKEDIN_SECRET'],
                   code=request.json['code'],
                   grant_type='authorization_code')

    # Step 1. Exchange authorization code for access token.
    r = requests.post(access_token_url, data=payload)
    access_token = json.loads(r.text)
    params = dict(oauth2_access_token=access_token['access_token'],
                  format='json')

    # Step 2. Retrieve information about the current user.
    r = requests.get(people_api_url, params=params)
    profile = json.loads(r.text)

    print profile

    # user = User.query.filter_by(linkedin=profile['id']).first()
    # if user:
    #     token = create_token(user)
    #     return jsonify(token=token)
    # u = User(linkedin=profile['id'],
    #          display_name=profile['firstName'] + ' ' + profile['lastName'])
    # db.session.add(u)
    # db.session.commit()
    # token = create_token(u)
    # return jsonify(token=token)


@api_view(['POST'])
@authentication_classes([])
@permission_classes([])
@throttle_classes([])
def twitter(request):
    request_token_url = 'https://api.twitter.com/oauth/request_token'
    access_token_url = 'https://api.twitter.com/oauth/access_token'
    authenticate_url = 'https://api.twitter.com/oauth/authenticate'

    if request.args.get('oauth_token') and request.args.get('oauth_verifier'):
        auth = OAuth1(app.config['TWITTER_CONSUMER_KEY'],
                      client_secret=app.config['TWITTER_CONSUMER_SECRET'],
                      resource_owner_key=request.args.get('oauth_token'),
                      verifier=request.args.get('oauth_verifier'))
        r = requests.post(access_token_url, auth=auth)
        profile = dict(parse_qsl(r.text))

        print profile

        # user = User.query.filter_by(twitter=profile['user_id']).first()
        # if user:
        #     token = create_token(user)
        #     return jsonify(token=token)
        # u = User(twitter=profile['user_id'],
        #          display_name=profile['screen_name'])
        # db.session.add(u)
        # db.session.commit()
        # token = create_token(u)
        # return jsonify(token=token)
    else:
        oauth = OAuth1(app.config['TWITTER_CONSUMER_KEY'],
                       client_secret=app.config['TWITTER_CONSUMER_SECRET'],
                       callback_uri=app.config['TWITTER_CALLBACK_URL'])
        r = requests.post(request_token_url, auth=oauth)
        oauth_token = dict(parse_qsl(r.text))
        qs = urlencode(dict(oauth_token=oauth_token['oauth_token']))
        return redirect(authenticate_url + '?' + qs)
