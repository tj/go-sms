# go-sms

Package sms provides a small wrapper around AWS SNS to make SMS usage less obscure,
conflating APIs FTW.

## Example

```` go
sms.Send("Hello World", "+19999999999")
````

## Running Tests

First export the `PHONE` environment variable to test against your number, and set your AWS credentials via the `AWS_*` environment variables.

## Badges

[![GoDoc](https://godoc.org/github.com/tj/go-sms?status.svg)](https://godoc.org/github.com/tj/go-sms)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)
[![](http://apex.sh/images/badge.svg)](https://apex.sh/ping/)

---

> [tjholowaychuk.com](http://tjholowaychuk.com) &nbsp;&middot;&nbsp;
> GitHub [@tj](https://github.com/tj) &nbsp;&middot;&nbsp;
> Twitter [@tjholowaychuk](https://twitter.com/tjholowaychuk)
