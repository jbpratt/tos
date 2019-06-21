FROM grafana/grafana:latest
USER root
RUN apt-get update && apt-get install -y curl gettext-base

WORKDIR /etc/grafana
COPY datasources ./datasources

WORKDIR /app
COPY entrypoint.sh ./
RUN chmod u+x entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]
