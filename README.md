
# Hyperping Exporter

[![tag](https://img.shields.io/github/tag/samber/hyperping-exporter.svg)](https://github.com/samber/hyperping-exporter/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18.0-%23007d9c)
[![GoDoc](https://godoc.org/github.com/samber/hyperping-exporter?status.svg)](https://pkg.go.dev/github.com/samber/hyperping-exporter)
![Build Status](https://github.com/samber/hyperping-exporter/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/samber/hyperping-exporter)](https://goreportcard.com/report/github.com/samber/hyperping-exporter)
[![Coverage](https://img.shields.io/codecov/c/github/samber/hyperping-exporter)](https://codecov.io/gh/samber/hyperping-exporter)
[![Contributors](https://img.shields.io/github/contributors/samber/hyperping-exporter)](https://github.com/samber/hyperping-exporter/graphs/contributors)
[![License](https://img.shields.io/github/license/samber/hyperping-exporter)](./LICENSE)

> A Prometheus Exporter for [Hyperping.io](https://hyperping.io)

## 🚀 Run

Using Docker:

```sh
docker run --rm -it -p 9312:9312 -e HYPERPING_TOKEN=xxxx samber/hyperping-exporter:v0.1.1
```

Or using a binary:

```sh
wget -O hyperping_exporter https://github.com/samber/hyperping-exporter/releases/download/v0.1.1/hyperping_exporter_0.1.1_linux_amd64
chmod +x hyperping_exporter
./hyperping_exporter --hyperping.token xxxx
```

## 💡 Usage

```sh
./hyperping_exporter
usage: hyperping_exporter --hyperping.token=HYPERPING.TOKEN [<flags>]

Flags:
  -h, --help                           Show context-sensitive help (also try --help-long and --help-man).
      --hyperping.token                Hyperping token ($HYPERPING_TOKEN)
      --namespace="hyperping"          Namespace for metrics ($HYPERPING_EXPORTER_NAMESPACE)
      --web.listen-address=":9312"     Address to listen on for web interface and telemetry. ($HYPERPING_EXPORTER_WEB_LISTEN_ADDRESS)
      --web.telemetry-path="/metrics"  Path under which to expose metrics. ($HYPERPING_EXPORTER_WEB_TELEMETRY_PATH)
      --log.format="txt"               Log format, valid options are txt and json ($HYPERPING_EXPORTER_LOG_FORMAT)
      --version                        Show application version.
```

## 📐 Metrics

### Probe status

Up vs down.

```
# HELP hyperping_monitor_status Probe status (0==down, 1==up)
# TYPE hyperping_monitor_status gauge
hyperping_monitor_status{monitor_id="mon_oungee2XIewoht",name="Landing page",project_id="proj_iiY5oo2Hfiepee",protocol="http",url="https://screeb.app"} 0
hyperping_monitor_status{monitor_id="mon_op3eeNgutahcha",name="API",project_id="proj_qui9looPieT0Ku",protocol="http",url="https://api.screeb.app"} 1
```

### Probe active

Paused vs active.

```
# HELP hyperping_monitor_active Probe active (0==paused, 1==active)
# TYPE hyperping_monitor_active gauge
hyperping_monitor_active{monitor_id="mon_oungee2XIewoht",name="Landing page",project_id="proj_iiY5oo2Hfiepee",protocol="http",url="https://screeb.app"} 0
hyperping_monitor_active{monitor_id="mon_op3eeNgutahcha",name="API",project_id="proj_qui9looPieT0Ku",protocol="http",url="https://api.screeb.app"} 1
```

## 🤝 Contributing

- Ping me on Twitter [@samuelberthe](https://twitter.com/samuelberthe) (DMs, mentions, whatever :))
- Fork the [project](https://github.com/samber/hyperping-exporter)
- Fix [open issues](https://github.com/samber/hyperping-exporter/issues) or request new features

Don't hesitate ;)

```bash
# Install some dev dependencies
make tools

# Run tests
make test
# or
make watch-test
```

## 👤 Contributors

![Contributors](https://contrib.rocks/image?repo=samber/hyperping-exporter)

## 💫 Show your support

Give a ⭐️ if this project helped you!

[![GitHub Sponsors](https://img.shields.io/github/sponsors/samber?style=for-the-badge)](https://github.com/sponsors/samber)

## 📝 License

Copyright © 2024 [Samuel Berthe](https://github.com/samber).

This project is [MIT](./LICENSE) licensed.
