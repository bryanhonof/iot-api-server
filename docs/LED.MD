# LED 

This function stores `R` `G` `B` `Brightness` values associated with each API key. You can modify these values using an HTTPS PUT request, and retrieve the values for your application through an HTTPS GET request.

## put 

```bash
curl -X PUT -H "X-Api-Key: (your API KEY)" -d @LED.json https://iot.pxl.bjth.xyz/api/v1/LED
```

This command allows you to configure the `R`, `G`, `B`, and `Brightness` values by sending the data contained in the `LED.json` file. 
>Please note that you can also send HTTPS requests directly from your code. 

The JSON data to be sent is as follows:

```json
{"R": 0, "G": 330, "B": 0, "Brightness": 255}
```

## get 

```bash
curl -X GET -H "X-Api-Key: (your API KEY)" https://iot.pxl.bjth.xyz/api/v1/LED
```
Use this command to fetch the `LED` data stored on the CRUD endpoint. 
>Keep in mind that you can also incorporate HTTPS requests directly into your code

You will receive a JSON data packet in the following format:
```json
{"R": 0, "G": 330, "B": 0, "Brightness": 255}
```
## example code 

you find the example code for this library back under `iot-api-server/examples/LED_exmaple_code`
