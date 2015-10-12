from django.conf.urls import url, include

from rest_framework.urlpatterns import format_suffix_patterns

import views

group_list = views.GroupViewSet.as_view({
    'get': 'list'
})
group_detail = views.GroupViewSet.as_view({
    'get': 'retrieve'
})

urlpatterns = [
	url(r'^users/$', views.UserList.as_view(), name='user-list'),
	url(r'^users/(?P<pk>[0-9]+)/$', views.UserDetail.as_view(), name='user-detail'),

	url(r'^groups/$', group_list, name='group-list'),
	url(r'^groups/(?P<pk>[0-9]+)/$', group_detail, name='group-detail'),
]

urlpatterns = format_suffix_patterns(urlpatterns)
