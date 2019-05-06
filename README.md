![Aero Go Logo](docs/images/aero.go.png)

[![Godoc][godoc-image]][godoc-url]
[![Report][report-image]][report-url]
[![Tests][tests-image]][tests-url]
[![Coverage][coverage-image]][coverage-url]
[![Patreon][patreon-image]][patreon-url]

Aero is a high-performance web server with a clean API for web development.

## Installation

```shell
go get -u github.com/blitzprog/home/...
```

## Usage

![Aero usage](docs/usage.gif)

Run this in an empty directory:

```bash
aero -newapp
```

Now you can build your app with `go build` or use the more advanced [run](https://github.com/aerogo/run) tool.

## Features

- Makes it easy to reach top scores in [Lighthouse](https://developers.google.com/web/tools/lighthouse/), [PageSpeed](https://developers.google.com/speed/pagespeed/insights/) and [Mozilla Observatory](https://observatory.mozilla.org/)
- Optimized for high latency environments (mobile networks)
- Has a strict content security policy
- Calculates E-Tags out of the box
- Finishes ongoing requests on a server shutdown
- Supports HTTP/2, IPv6 and Web Manifest
- Automatic HTTP/2 push of configured resources
- Supports session data with custom stores
- Allows pushing live data to the client via SSE
- Provides http and https listener
- Shows response time and size for your routes
- Can run standalone without `nginx` babysitting it

## Optional

- [layout](https://github.com/aerogo/layout) as a layout system
- [pack](https://github.com/aerogo/pack) to compile Pixy, Scarlet and JS assets in record time
- [run](https://github.com/aerogo/run) which automatically restarts your server on code/template/style changes
- [pixy](https://github.com/aerogo/pixy) as a high-performance template engine similar to Jade/Pug
- [scarlet](https://github.com/aerogo/scarlet) as an aggressively compressing stylesheet preprocessor
- [nano](https://github.com/aerogo/nano) as a fast, decentralized and git-trackable database
- [api](https://github.com/aerogo/api) to automatically implement your REST API routes
- [markdown](https://github.com/aerogo/markdown) as an overly simplified markdown wrapper
- [http](https://github.com/aerogo/http) as an HTTP client with a simple and clean API
- [log](https://github.com/aerogo/log) for simple & performant logging

## Documentation

- [API](docs/API.md)
- [Configuration](docs/Configuration.md)
- [Benchmarks](docs/Benchmarks.md)

## Others

- [Slides for OWDDM talk](https://docs.google.com/presentation/d/166I69goLEVuvuFeeRfUu8c5lwl2_HAeSi2SZyzIuEKg/edit) (Osaka, May 2018)
- [Discord community](https://discord.gg/V3y4xTY)
- [Twitter account](https://twitter.com/aeroframework)
- [Facebook page](https://www.facebook.com/aeroframework/)

## Coding style

Please take a look at the [style guidelines](https://github.com/akyoto/quality/blob/master/STYLE.md) if you'd like to make a pull request.

## Patrons

| [![Scott Rayapoullé](https://avatars3.githubusercontent.com/u/11772084?s=70&v=4)](https://github.com/soulcramer) |
|---|
| [Scott Rayapoullé](https://github.com/soulcramer) |

Want to see [your own name here](https://www.patreon.com/eduardurbach)?

## Author

| [![Eduard Urbach on Twitter](https://gravatar.com/avatar/16ed4d41a5f244d1b10de1b791657989?s=70)](https://twitter.com/eduardurbach "Follow @eduardurbach on Twitter") |
|---|
| [Eduard Urbach](https://eduardurbach.com) |

[godoc-image]: https://godoc.org/github.com/blitzprog/home?status.svg
[godoc-url]: https://godoc.org/github.com/blitzprog/home
[report-image]: https://goreportcard.com/badge/github.com/blitzprog/home
[report-url]: https://goreportcard.com/report/github.com/blitzprog/home
[tests-image]: https://cloud.drone.io/api/badges/blitzprog/home/status.svg
[tests-url]: https://cloud.drone.io/blitzprog/home
[coverage-image]: https://codecov.io/gh/blitzprog/home/graph/badge.svg
[coverage-url]: https://codecov.io/gh/blitzprog/home
[patreon-image]: https://img.shields.io/badge/patreon-donate-green.svg
[patreon-url]: https://www.patreon.com/eduardurbach
