from rest_framework import mixins, generics
from rest_framework import permissions
from rest_framework import renderers
from rest_framework.response import Response

from todos.models import Todo, Pomodoro
from todos.serializers import TodoSerializer, PomodoroSerializer

from bluepomos.permissions import IsOwnerOrReadOnly, IsOwnerOnly

import datetime

class TodoList(generics.ListCreateAPIView):
    permission_classes = (permissions.IsAuthenticated, IsOwnerOrReadOnly,)
    serializer_class = TodoSerializer

    def get_queryset(self):
        user = self.request.user
        return Todo.objects.filter(owner=user).all()

    def perform_create(self, serializer):
        serializer.save(owner=self.request.user)


class TodoDetail(generics.RetrieveUpdateDestroyAPIView):
    permission_classes = (permissions.IsAuthenticated, IsOwnerOnly,)
    queryset = Todo.objects.all()
    serializer_class = TodoSerializer


class TodoPomodoroList(generics.ListCreateAPIView):
    permission_classes = (permissions.IsAuthenticated, IsOwnerOrReadOnly,)
    serializer_class = PomodoroSerializer

    def get_queryset(self):
        pk = self.kwargs.get('pk',None)
        user = self.request.user
        return Pomodoro.objects.filter(owner=user, todo=pk).all()

    def get(self, request, *args, **kwargs):
        print "GOT HERE ###"
        return self.list(request, *args, **kwargs)

    def perform_create(self, serializer):
        print "GOT HERE @@@"
        pk = self.kwargs.get('pk',None)
        user = self.request.user
        todo = Todo.objects.filter(id=pk).first()
        todo.pomodoros_started += 1
        todo.save()
        serializer.save(owner=user, todo=todo)



class PomodoroList(generics.ListCreateAPIView):
    permission_classes = (permissions.IsAuthenticated, IsOwnerOrReadOnly,)
    serializer_class = PomodoroSerializer

    def get_queryset(self):
        user = self.request.user
        return Pomodoro.objects.filter(owner=user).all()

    def perform_create(self, serializer):
        serializer.save(owner=self.request.user)


class PomodoroDetail(generics.RetrieveUpdateAPIView):
    permission_classes = (permissions.IsAuthenticated, IsOwnerOnly,)
    queryset = Pomodoro.objects.all()
    serializer_class = PomodoroSerializer

    def put_queryset(self):
        pid = self.kwargs.get('pk',None)
        user = self.request.user
        return Pomodoro.objects.filter(owner=user, id=pid, end=datetime.date.min).all()


