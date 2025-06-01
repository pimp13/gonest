# GoNest CLI

A CLI tool for Go-Nest framework that helps you create modules, controllers, services, and more.

## Installation

```bash
go install github.com/pimp13/gonest/cmd/gonest@latest
```

## Usage

Generate a new module (includes controller and service):
```bash
gonest generate module users
# or short version
gonest g m users
```

Generate a controller:
```bash
gonest generate controller products
# or short version
gonest g c products
```

Generate a service:
```bash
gonest generate service orders
# or short version
gonest g s orders
```

## Available Commands

- `module` (alias: m): Generate a new module
- `controller` (alias: c): Generate a new controller
- `service` (alias: s): Generate a new service

## Help

For more information about the commands:
```bash
gonest --help
gonest generate --help
```

## License

MIT