# mtool - modbus tool

This is a command line tool for experimenting with modbus.  **Please note that this is a work in progress**.

## Installing

If you have Go installed on your system you can fetch the latest version of this utility using:

```shell
go install github.com/lab5e/mtool/cmd/mtool@latest
```

## Usage

The `mtool` utility has several subcommands. You can list these, and the global command line options, by runnig `mtool -h`.

```shell
$ mtool -h
Usage:
  mtool [OPTIONS] <command>

Application Options:
      --device=          serial device [$MTOOL_DEVICE]
      --baud=            baud rate (default: 9600) [$MTOOL_BAUD]
      --databits=        data bits (default: 8) [$MTOOL_DATABITS]
      --parity=[N|E|O]   parity (default: N) [$MTOOL_PARITY]
      --stop=[1|2]       stop bits (default: 1) [$MTOOL_STOPBITS]
      --id=              device id (default: 1)
      --base=[2|8|10|16] output base (default: 10) [$MTOOL_OUTPUT_BASE]

Help Options:
  -h, --help             Show this help message

Available commands:
  rc  read coils
  rd  read discrete input
  rh  read holding register
  ri  read input register
  wh  write holding register
```

For the global (application) options we have provided environment variables which can make using `mtool` a bit more convenient.  For instance it may be preferable to at least set the `MTOOL_DEVICE` to the serial device you are speaking to so you don't have to clutter the command line with this.

### Device id

The device ID (`-id`) is the (hopefully) unique device ID on the modbus loop.  Before hooking up your devices, please make sure that each device is given a unique ID and that you have noted these somewhere.

### Output base

Sometimes you may want the values returned expressed in some other base than base 10.  To do this you can specify the `--base` option with `2` (binary), `8` (octal), `10` (decimal) or `16` (hexadecimal) as the base.

## Subcommands

Each subcommand has its own options. We have tried to keep things as consistent as possible, so many of them will be the same.  To list command specific options, just append `-h` to your command line for a specific option.

Example:

```shell
$ mtool rh -h
Usage:
  mtool [OPTIONS] rh [rh-OPTIONS]

Application Options:
      --device=          serial device [$MTOOL_DEVICE]
      --baud=            baud rate (default: 9600) [$MTOOL_BAUD]
      --databits=        data bits (default: 8) [$MTOOL_DATABITS]
      --parity=[N|E|O]   parity (default: N) [$MTOOL_PARITY]
      --stop=[1|2]       stop bits (default: 1) [$MTOOL_STOPBITS]
      --id=              device id (default: 1)
      --base=[2|8|10|16] output base (default: 10) [$MTOOL_OUTPUT_BASE]

Help Options:
  -h, --help             Show this help message

[rh command options]
          --addr=        address
          --count=       count (default: 1)
          --repeat=      repeat interval, if zero no repeat (default: 0)
          --json         format data as JSON
```

This still lists the global (application) options, but now you will also see the options for the `rh` (read holding register) command too.

- `--addr` refers to the short version of the holding register, eg `10`
- `--count` is the number of registers you want to read
- `--repeat` is the repetition interval. This is specified with unit, so 1 second is 1s, 10 milliseconds is 10ms etc. (Valid units are ns,ms,s,m,h)
- `--json` produces JSON output which can be useful if you want to parse the output later.


## Examples

### Repeatedly read holding register

This example reads holding register 6 from device id 1 ever 2 seconds and outputs the result as base 16 (hexadecimal). 

```shell
$ mtool rh --id 1 --addr 6 --base 16 --json --repeat=2s
{"time":"2023-01-30T18:02:33.170332Z","valueType":"holding_register","deviceID":1,"addr":6,"data":"00, 05"}
{"time":"2023-01-30T18:02:35.249169Z","valueType":"holding_register","deviceID":1,"addr":6,"data":"00, 05"}
{"time":"2023-01-30T18:02:37.328873Z","valueType":"holding_register","deviceID":1,"addr":6,"data":"00, 05"}
```

### Write holding register

This example writes holding register 6 on device 1, and sets it to the value 1 (decimal).  After successfully setting holding register 1, we output the response we get from the device.

```shell
$ bin/mtool wh --id 1 --addr 6 --value 1
deviceID=1 addr=6 res={0, 1}
```