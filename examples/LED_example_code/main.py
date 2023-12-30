from led import RGBdata, Neopixel
import argparse
import requests
import json
import time
from time import sleep
import os

leds = Neopixel(1)
leds.fill(RGBdata(0,0,0,255))
print(leds.colors()) 

API_KEY = 1222
url = 'https://iot.pxl.bjth.xyz/api/v1/LED'
headers = {
    'X-Api-Key': str(API_KEY)  # Fix the header format
}

# main loop
data = {
    "R": 255,
    "G": 0,
    "B": 255,
    "Brightness": 0
}
RGBSend = requests.put(url, json=data, headers=headers)
if RGBSend.status_code != 200:
    print("error : "+ str(RGBSend.status_code))

RGBRecieve = requests.get(url, headers=headers)
if RGBRecieve.status_code != 200:
    print("error : "+ str(RGBRecieve.status_code))

LED = json.loads(RGBRecieve.json())
leds.fill(RGBdata(LED['R'],LED['G'],LED['B'],LED['Brightness']))
print(leds.colors()) 