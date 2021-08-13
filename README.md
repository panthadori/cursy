# cursy

Red team "troll" tool written in Go.

## Usage

```bash
DISPLAY=:0 ./cursy
```

The process ignores SIGINT and SIGTERM by default, so make sure to send SIGKILL or SIGSTOP instead (i.e., `pkill -9 cursy`).

## Example

![Example Run of Cursy](./cursyExample.gif)

## Installation

Download a release, or compile from source. Make sure to first install the dependencies for robotgo`.
