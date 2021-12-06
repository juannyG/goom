# goom

### Getting started

```
git clone git@github.com:juannyG/goom.git
cd goom
docker-compose build
docker-compose run --rm goom go test ./... -v -cover
```

### Running locally

Make sure you're running the whiskey database and traefik images. Then `docker-compose up` and you can navigate to a browser:

http://goom.ordergroove.localhost/product?merchant=67525f8ca4772569c35f326c274cad70&product=1

and should get back a response of

```
{
  "autoship_enabled": true,
  "live":true,
  "merchant":"67525f8ca4772569c35f326c274cad70",
  "product":"1"
}
```
