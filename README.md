# Openlamp-cli
`openlamp-cli` is a CLI tool for controlling a BLE RGB lamp on Linux. <br>
Project was made by reverse-engineering the original mobile app and implementing my own BLE client in Go.

**Note:** This project is currently in the **early prototype stage** so things like device MAC address are hardcoded for development purposes.

---

## Features
- Turn lamp ON / OFF
- Set brightness
- Set RGB color
- Direct BLE communication
- Simple CLI usage

---

## Requirements
- Linux
- Bluetooth adapter
- BlueZ
- Go (>= 1.25)
- Lamp compatible with the `55 AA` BLE protocol

---

## Installation
```shell
git clone https://github.com/f1rq/openlamp-cli
cd openlamp-cli
go build -o openlamp ./cmd/openlamp-cli
```

---

## Usage
```shell
./openlamp turnon
./openlamp turnoff
./openlamp brightness <0-255>
./openlamp color <RRGGBB>
```

---

## Configuration
### Currently, the lamp MAC address is hardcoded in the source code.
Planned:
- config file in `~/.config`
- configurable MAC address

---

## Protocol notes
The lamp communicates over BLE using a custom binary protocol.<br>
General frame structure:
```
55 AA [CMD] [LEN] [DATA...] [CHECKSUM]
```
- Checksum is calculated as:<br>
    ```
    sum(bytes) & 0xFF
    ```
Examples:

| Action     | Hex Payload          |
|------------|----------------------|
| Turn ON    | `55AA01080501F1`     |
| Brightness | `55AA010801FF??`     |
| Color      | `55AA030802RRGGBB??` |
Note: ?? represents the checksum byte, calculated as the least significant byte of the sum of all preceding bytes.

---

## Disclaimer
This project is **not affiliated** with the lamp manufacturer.<br>
Protocol was discovered via reverse engineering for educational purposes.<br><br>
Use at **your own risk.**

---

## How it was made
I used an Android phone with the original lamp application for sniffing bytes through `adb logcat`.

---

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.