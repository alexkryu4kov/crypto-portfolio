from django.db import models

# Create your models here.


class Cryptocurrency(models.Model):
    name = models.CharField(max_length=100)
    symbol = models.CharField(max_length=100)

    def __str__(self):
        return f'{self.name} ({self.symbol})'
