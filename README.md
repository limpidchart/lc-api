# lc-api

Limpidchart API service

# Observability

You can scrap [Prometheus](https://prometheus.io) `/metrics` endpoint on the `LC_METRICS_ADDRESS` (`0.0.0.0:54013` by default).  
Currently there is only a `request_duration_seconds` histogram with the default Prometheus buckets (`.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10`).  

You can use [PromQL](https://prometheus.io/docs/prometheus/latest/querying/basics/) to build some useful visualisations from it (queries based on [Weave Works](https://www.weave.works/blog/of-metrics-and-middleware/) article):

 * `sum(irate(request_duration_seconds_count{job="default/hello"}[1m])) by (status_code)` - instantaneous QPS of the service, by status code, aggregated across all replicas;
 * `sum(irate(request_duration_seconds_count{job="default/hello", status_code=~"5.."}[1m]))` - rate of requests returning 500s;
 * `sum(irate(request_duration_seconds_count{job="default/hello"}[1m])) by (instance)` - QPS per instance;
 * `sum(rate(request_duration_seconds_sum{job="default/hello",ws="false"}[5m])) / sum(rate(request_duration_seconds_count{job="default/hello",ws="false"}[5m]))` - 5-minute moving average latency of requests to the service;
 * `histogram_quantile(0.99, sum(rate(request_duration_seconds_bucket{job="default/hello",ws="false"}[5m])) by (le))` - 5-min moving 99th percentile request latency;

Every request is also logged to `dev/stderr`, duration field contains value in seconds:

 * gRPC example: `{"level":"info","time":"2021-08-15T12:41:28Z","protocol":"grpc","request_id":"26fab748-7e41-4901-882f-dbe6babb4a6f","ip":"192.168.0.1:55722","code":"OK","method":"CreateChart","duration":0.0091757}`
 * HTTP example: `{"level":"info","time":"2021-08-15T12:41:28Z","protocol":"http","request_id":"343177ef-078c-4396-a8ae-71a06f42de57","ip":"192.168.0.1","user_agent":"lc-acceptance-tests","referer":"","code":201,"method":"POST","path":"/v0/charts","resp_bytes_written":12863,"duration":0.0194194}`
