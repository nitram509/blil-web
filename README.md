# BliL-Web - Blinking Light Web

A web server/application, which makes a blinkstick and compatible devices available via REST API (HTTP + JSON)


#### License

The MIT License (MIT)


## Usage


### get all available devices

GET http://localhost:8080/

```json
{
    "version": "0.0.1",
    "name": "BliL - Blinking Light",
    "_embedded": {
        "leds": [
            {
                "number": 0,
                "type": "BlinkStick",
                "path": "USB_20a0_41e5_14100000",
                "_links": [
                    {
                        "self": {
                            "href": "http://localhost:8080/led/0",
                            "title": "Set or get color on this LED"
                        }
                    }
                ]
            }
        ]
    }
}
```

### set a color

POST http://localhost:8080/led/0/green

```json
{
    "number": 0,
    "color": "008000"
}
```

## Supported devices

* [blink(1)](http://blink1.thingm.com/)
* [LinkM / BlinkM](http://thingm.com/products/linkm/)
* [BlinkStick](http://www.blinkstick.com/)
* [Blync](http://www.blynclight.com/)
* [Busylight UC](http://www.busylight.com/busylight-uc.html)
* [Busylight Lync](http://www.busylight.com/busylight-lync.html)
* [DreamCheeky USBMailNotifier](http://www.dreamcheeky.com/webmail-notifier)

_powered by_ https://github.com/boombuler/led

