from rest_framework.request import Request
from django.utils.functional import SimpleLazyObject
from django.contrib.auth.middleware import get_user

from rest_framework_jwt.authentication import JSONWebTokenAuthentication


def get_user_jwt(request):
    user = get_user(request)
    if user.is_authenticated():
        return user, None
    try:
        user_jwt = JSONWebTokenAuthentication().authenticate(Request(request))
        print "user_jwt: ", user_jwt
        if user_jwt is not None:
            return user_jwt[0], user_jwt[1]
        else:
            return None, None
    except Exception, e:
        print "exception", e
        pass
    return user, None


class AuthenticationMiddlewareJWT(object):
    def process_request(self, request):
        assert hasattr(request, 'session'), "The Django authentication middleware requires session middleware to be installed. Edit your MIDDLEWARE_CLASSES setting to insert 'django.contrib.sessions.middleware.SessionMiddleware'."

        user,auth = get_user_jwt(request)
        print "GOT HERE!!!"
        print "user:", user
        print "auth:", auth
        print "-----"

        # request.user = SimpleLazyObject(lambda: user )
        request.jwtuser = SimpleLazyObject(lambda: user )
        request.jwtauth = auth
