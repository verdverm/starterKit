from django.db import models
import datetime

class Todo(models.Model):
    created = models.DateTimeField(auto_now_add=True)
    name = models.CharField(max_length=100, blank=True, default='')
    description = models.TextField()

    pomodoros_started = models.PositiveIntegerField(default=0)
    pomodoros_completed = models.PositiveIntegerField(default=0)

    owner = models.ForeignKey('auth.User', related_name='todos')


class Pomodoro(models.Model):
    started_at = models.DateTimeField(auto_now_add=True)
    ended_at = models.DateTimeField(default=datetime.date.min)
    completed = models.BooleanField(default=False)

    owner = models.ForeignKey('auth.User', related_name='pomodoros')
    todo = models.ForeignKey('todos.Todo', related_name='pomodoros')

