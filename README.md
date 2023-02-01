# mtool - modbus tool

This is a command line tool for experimenting with modbus.  **Please note that this is a work in progress**.

## Installing

If you have Go installed on your system you can fetch the latest version of this utility using:

```shell
go install github.com/lab5e/mtool/cmd/mtool@latest
```

## Usage

The `mtool` utility has several subcommands. You can list these, and the global command line options, by runnig `mtool -h`.  

Note that some of the command line options have default values (which the help `-h` option will list).  The idea is to cut down on the number of command line options you have to specify by hopefully providing sensible defaults.  For instance the modbus device ID defaults to 1 since a lot of devices have this ID right out of the box.

### Environment variables

For the global (application) options we have provided environment variables which can make using `mtool` a bit more convenient.  For instance it may be preferable to at least set the `MTOOL_DEVICE` to the serial device you are speaking to so you don't have to clutter the command line by having to include this every time.  The `-h` help output will list the other command line parameters that have environment variables you can set.

### Device id

The device ID (`-id`) is the (hopefully) unique device ID on the modbus loop.  Before hooking up your devices, please make sure that each device is given a unique ID and that you have noted these somewhere.  

### Addresses

Note that modbus addresses (specified by the `--addr` flag) are one-based while the protocol itself is zero-based.  This means that the documentation will give an address as 10, while what is expected by the device is 9.  `mtool` subtracts 1 from the address you give it before putting it on the wire.

### Output base

Sometimes you may want the values returned expressed in some other base than base 10.  To do this you can specify the `--base` option with `2` (binary), `8` (octal), `10` (decimal) or `16` (hexadecimal) as the base.

### Repeatedly reading values

All of the reading commands support the `--repeat` flag.  When this flag is not specified the read command will be executed only once.  If it is set to a value it will repeat the read operation every time unit specified.

Time is expressed as a number followed by a unit.  Valid units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

This example reads holding register with address 1 every 500 milliseconds:

```shell
$ bin/mtool rh --addr 1 --repeat 500ms
2023-02-01T13:17:06Z [holding_register] deviceID=1 addr=1 count=1 data={230}
2023-02-01T13:17:07Z [holding_register] deviceID=1 addr=1 count=1 data={230}
2023-02-01T13:17:08Z [holding_register] deviceID=1 addr=1 count=1 data={230}
2023-02-01T13:17:08Z [holding_register] deviceID=1 addr=1 count=1 data={230}
```

### Debug output

If you want to see the messages that are sent and received on the bus, you can add the `--debug` flag.  This will print out messages prefixed with `modbus:` showing the bytes that are being transmitted and received. *The debug output is not affected by the `--base` option*.

## Subcommands

The mtool command itself does nothing.  To make it do things you have to specify *subcommands*. The following sections specify the subcommands.  In order to list the options available to each subcommand you should use the built-in help for each subcommand.  You can do this by issuing a command of the form

```shell
mtool <subcommand> -h
```

### `scan` -scan bus for devices

Scan the bus for devices by trying to read data from it.  Per default it will try to read holding register 1.  You can change what to read by setting the `--type` command line option to one of `holding`, `discrete`, `input` or `coil`.  You can change the address used for scanning by setting the `--addr` command line option.

Depending on the devices you may have to configure how long the scanning process should wait for response from each device.  This can be done with the `--scan-timeout` option.  The default timeout is 100ms.  If you don't find any devices during a scan, try to increase the timeout value.

### `rh` - read holding register

Read the holding register identified by `--addr`.

This example reads holding register 1

```shell
$ bin/mtool rh --addr 1
2023-02-01T13:05:28Z [holding_register] deviceID=1 addr=1 count=1 data={230}
```

### `wh` - write holding register

Write holding register identified by `--addr`.

This example writes the `int16` value `230` to holding register 1

```shell
$ bin/mtool wh --addr 1 --value 230
deviceID=1 addr=1 res={230}
```

You can check that the value has indeed been changed by reading the holding register back using the `rh` subcommand.

### `ri` - read input register

Read input register identified by `--addr`.

This example reads input register 4:

```shell
bin/mtool ri --addr 4
2023-02-01T13:03:50Z [input_register] deviceID=1 addr=4 count=1 data={2574}
```

### `rd` - read discrete input

Read discrete input identified by `--addr`.  Note that this is a bit field.

This example reads discrete input 4. Note that the output is in binary per default.  You can change the base with the `--base` command line option.

```shell
$ bin/mtool rd --addr 1 
2023-02-01T13:06:48Z [discrete_input] deviceID=1 addr=1 count=1 data={11111011, 0000000000011101}
```

### `rc` - read coils

Read coils identified by `-addr`.  Note that this is a bit field.  You can change the output base with the `--base` command line flag.

### `doc` - output this documentation

For convenience sake this README.md file is included in the binary.  You can spit it out as original markup or as HTML.
