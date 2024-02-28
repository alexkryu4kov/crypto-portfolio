import json
import os

from django.shortcuts import render, get_object_or_404
from .models import Cryptocurrency


def crypto_list(request):
    cryptos = Cryptocurrency.objects.all()
    return render(request, 'crypto/crypto_list.html', {'cryptos': cryptos})


def crypto_detail(request, pk):
    crypto = get_object_or_404(Cryptocurrency, pk=pk)

    crypto_name = crypto.name

    try:
        with open(os.path.join('/code/data', 'crypto_prices.json'), 'r') as file:
            lines = file.readlines()
            prices = [json.loads(line) for line in lines if line.strip()]
    except IOError as e:
        prices = []
        print("Error loading price data {e")

    labels = [entry['timestamp'] for entry in prices]
    data = [entry["prices"][crypto_name]['usd'] for entry in prices if crypto_name in entry['prices']]

    context = {
        'crypto': {
            'name': crypto_name,
            'symbol': crypto.symbol,
        },
        'labels': labels,
        'data': data,
    }
    return render(request, 'crypto/crypto_detail.html', context)
