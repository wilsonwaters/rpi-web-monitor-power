# rpi-web-monitor-power
Control monitor power on a Raspberry Pi through http

## Prerequisites

`apt install golang`

## Compile

`cd rpi-web-monitor-power`
`go build`

## Running

You probably want to run this as a service. But for testing you can run it from the command line

`cd rpi-web-monitor-power`
`./rpi-web-monitor-power
Turn monitor off: http://localhost:9990/monitor-power/1/1?power=off
Turn monitor off: http://localhost:9990/monitor-power/1/1?power=on

## API
/monitor-power/{version]/{monitor-number}?power={on|off}
