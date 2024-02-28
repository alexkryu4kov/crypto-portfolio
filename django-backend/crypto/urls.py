from django.urls import path
from . import views

urlpatterns = [
    path('', views.crypto_list, name='crypto_list'),
    path('<int:pk>/', views.crypto_detail, name='crypto_detail'),
]
