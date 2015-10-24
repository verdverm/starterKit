from django.conf.urls import url
# from rest_framework.urlpatterns import format_suffix_patterns

from accounts import views

urlpatterns = [
	url(r'^auth/facebook/', views.facebook),
	url(r'^auth/github/', views.github),
	url(r'^auth/google/', views.google),
	url(r'^auth/linkedin/', views.linkedin),
	url(r'^auth/twitter/', views.twitter),

	url(r'^auth/unlink/', views.unlink),
]

# urlpatterns = format_suffix_patterns(urlpatterns)

