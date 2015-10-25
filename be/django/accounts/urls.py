from django.conf.urls import url
# from rest_framework.urlpatterns import format_suffix_patterns

from accounts import views

urlpatterns = [
	url(r'^auth/accounts/', views.accounts),

	url(r'^auth/facebook/', views.facebook),
	url(r'^auth/google/', views.google),
	url(r'^auth/yahoo/', views.yahoo),
	url(r'^auth/live/', views.windows),
	url(r'^auth/github/', views.github),
	url(r'^auth/twitter/', views.twitter),
	url(r'^auth/soundcloud/', views.soundcloud),
	url(r'^auth/dropbox/', views.dropbox),

	url(r'^auth/unlink/', views.unlink),
]

# urlpatterns = format_suffix_patterns(urlpatterns)

