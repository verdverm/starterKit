from rest_framework import serializers
from todos.models import Todo, Pomodoro

import datetime
from django.utils import timezone



class TodoSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Todo
        fields = ('url', 'id', 'owner', 'name', 'description', 'pomodoros_started', 'pomodoros_completed', 'pomodoros')
        read_only_fields = ('owner', 'pomodoros_started', 'pomodoros_completed', 'pomodoros')


class PomodoroSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Pomodoro
        fields = ('url', 'id', 'started_at', 'ended_at', 'completed', 'owner', 'todo')
        read_only_fields = ('started_at', 'ended_at', 'completed', 'owner', 'todo')


    def update(self, instance, validated_data):
        print "GOT HERE  $$$$$$"
        now = timezone.now()
        start = instance.started_at
        ended_set = instance.ended_at == datetime.date.min
        if not ended_set:
            past = start + datetime.timedelta(minutes=1)
            diff = now-start
            instance.ended_at = now
            instance.completed = now > past
            # update the todo
            todo = Todo.objects.filter(id=instance.todo_id).first()
            todo.pomodoros_completed += 1
            todo.save()
        instance.save()
        return instance


