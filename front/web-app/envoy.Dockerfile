FROM envoyproxy/envoy:latest

COPY ./envoy.yml /etc/envoy/envoy.yml

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yml