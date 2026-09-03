#!/bin/sh

set -eu

docker run --rm \
  -v /etc/letsencrypt:/etc/letsencrypt \
  -v /var/lib/letsencrypt:/var/lib/letsencrypt \
  -v /opt/1panel/apps/openresty/openresty/root:/var/www/acme \
  certbot/certbot:latest renew --webroot -w /var/www/acme --quiet

cp /etc/letsencrypt/live/82.157.174.249/fullchain.pem \
  /opt/1panel/apps/openresty/openresty/conf/ssl/voice-room-fullchain.pem
cp /etc/letsencrypt/live/82.157.174.249/privkey.pem \
  /opt/1panel/apps/openresty/openresty/conf/ssl/voice-room-privkey.pem
chmod 644 /opt/1panel/apps/openresty/openresty/conf/ssl/voice-room-fullchain.pem
chmod 600 /opt/1panel/apps/openresty/openresty/conf/ssl/voice-room-privkey.pem

docker exec 1Panel-openresty-vMGH nginx -t
docker exec 1Panel-openresty-vMGH nginx -s reload
