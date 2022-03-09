# url-shortener

## Service for shortening links.

## Using
1. DB - `MySql`
1. Routing - `gorilla/mux`

## Routes
`GET: /` - get home page \
`POST: /create` - create new route \
`GET: /{short_url}` - redirect on corresponding long url

## Start
1. git clone github.com/EgorSkurihin/url-shortener && cd url-shortener 
1. docker-compose build
1. docker-compose up
1. follow the link http://localhost:8181