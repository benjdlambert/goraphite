<h1 align="center">GoRaphite</h1>

<p align="center">
  <a href="https://travis-ci.org/benjdlambert/goraphite" target="_blank">
    <img src="https://img.shields.io/travis/benjdlambert/goraphite.svg?maxAge=30">
  </a>
</p>

<p align="center">
  <b>Graphite</b> client written in <b>Go</b>.
</p>

## Usage

### Quickstart
```go
host, port := "myhost.com", 123
client, err := NewGoraphite(host, port)
```

### FindMetrics
You can find metrics, by using the `FindMetrics` `struct` in the `query` package and the `FindMetrics` `func` on the `Client` like so:

```go
metrics, err = client.FindMetrics(
    query.FindMetrics{
        Query: "collectd.*"
    },
)
```

This will return you a list of `Metrics`.

### GetMetrics
You can get metrics for a given wildcard too, by using the `GetMetrics` `func` and the `GetMetrics` `struct` in the `query` package and specifying a `target`

```go
metrics, err = client.GetMetrics(
    query.GetMetrics{
        Target: "collectd.*"
    },
)
```

This will return you a list of `Targets` with the `Datapoints` array.


