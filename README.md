# lc-api

Limpidchart API service

# Observability

You can scrap [Prometheus](https://prometheus.io) `/metrics` endpoint on the `LC_API_HTTP_ADDRESS` (`0.0.0.0:54012` by default).  
Currently there is only a `request_duration_seconds` histogram with the default Prometheus buckets (`.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10`).  

You can use [PromQL](https://prometheus.io/docs/prometheus/latest/querying/basics/) to build some useful visualisations from it (queries based on [Weave Works](https://www.weave.works/blog/of-metrics-and-middleware/) article):

 * `sum(irate(request_duration_seconds_count{job="default/hello"}[1m])) by (status_code)` - instantaneous QPS of the service, by status code, aggregated across all replicas;
 * `sum(irate(request_duration_seconds_count{job="default/hello", status_code=~"5.."}[1m]))` - rate of requests returning 500s;
 * `sum(irate(request_duration_seconds_count{job="default/hello"}[1m])) by (instance)` - QPS per instance;
 * `sum(rate(request_duration_seconds_sum{job="default/hello",ws="false"}[5m])) / sum(rate(request_duration_seconds_count{job="default/hello",ws="false"}[5m]))` - 5-minute moving average latency of requests to the service;
 * `histogram_quantile(0.99, sum(rate(request_duration_seconds_bucket{job="default/hello",ws="false"}[5m])) by (le))` - 5-min moving 99th percentile request latency;
