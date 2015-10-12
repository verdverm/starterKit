from django.conf.urls import url
from rest_framework.urlpatterns import format_suffix_patterns

from todos import views

urlpatterns = [
    url(r'^todos/$', views.TodoList.as_view(), name='todo-list'),
    url(r'^todos/(?P<pk>[0-9]+)/$', views.TodoDetail.as_view(), name='todo-detail'),
    url(r'^todos/(?P<pk>[0-9]+)/pomodoros/$', views.TodoPomodoroList.as_view(), name='todo-pomodoro-list'),

    url(r'^pomodoros/$', views.PomodoroList.as_view(), name='pomodoro-list'),
    url(r'^pomodoros/(?P<pk>[0-9]+)/$', views.PomodoroDetail.as_view(), name='pomodoro-detail'),
]

urlpatterns = format_suffix_patterns(urlpatterns)
