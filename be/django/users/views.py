from rest_framework import generics, permissions, viewsets

from django.contrib.auth.models import User, Group
from serializers import UserSerializer, GroupSerializer

from app.permissions import IsOwnerUserToo



class UserList(generics.ListAPIView):
    permission_classes = (permissions.IsAdminUser,)
    queryset = User.objects.all()
    serializer_class = UserSerializer


class UserDetail(generics.RetrieveAPIView):
    permission_classes = (permissions.IsAuthenticatedOrReadOnly,IsOwnerUserToo)

    queryset = User.objects.all()
    serializer_class = UserSerializer

class GroupViewSet(viewsets.ModelViewSet):
    """
    API endpoint that allows groups to be viewed or edited.
    """
    permission_classes = (permissions.IsAdminUser,)
    queryset = Group.objects.all()
    serializer_class = GroupSerializer
    