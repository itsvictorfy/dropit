FROM grafana/grafana:7.5.17
RUN grafana-cli plugins install redis-datasource
