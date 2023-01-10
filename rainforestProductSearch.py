import requests
import json

# set up the request parameters
params = {
'api_key': '63FFD605C840421D9F5FC4433C106F90',
  'amazon_domain': 'amazon.com',
  'asin': 'B073JYC4XM',
  'type': 'product'
}

# make the http GET request to Rainforest API
api_result = requests.get('https://api.rainforestapi.com/request', params)

# print the JSON response from Rainforest API
print(json.dumps(api_result.json()))