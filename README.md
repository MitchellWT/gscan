
# Gscan

## About

Gscan is a Linux utility that allows for file system metadata capture and exporting of said collected meta data. Metadata collected includes:

- File path
- File size (in bits)

Exporting can be <strong>raw</strong> where all files (and sub files) in the target directory are captured in the export. Or exporting can be <strong>total</strong> where only the total size of the target dirtectory is exported, this total size is calculated by summing all files (and sub files) in the target directory.

Exporting can be in the following file formats:

- JSON
- HTML
- CSV

## WARNING: HTML Export

HTML exporting makes use of chart.js. Chart.js can <strong>NOT</strong> handle large datasets. This isnt an issue for <strong>total</strong> exporting but raw exports can have complications (the webpage exported will stall when loading in the JavaScript). This means / exporting is <strong>NOT</strong> recommeded (for raw exporting).

## Installation

To install from source run the following commands:

```
# This will produce a binary for your system
go build -o gscan cmd/gscan/main.go

# This will move the produced binary to your local bin's dir
sudo cp gscan /usr/local/bin
```

After the above steps run the following command for basic help:
```
gscan --help
```

## Usage

### Read Command

Example read:
```
gscan read ~/Desktop/
```

To get some basic help for file system metadata capture run the following:
```
gscan read --help
```

### Export Command

Example export to JSON for the last month (raw):
```
gscan export ~/Desktop/ --interval month
```

Example export to HTML (totaled):
```
gscan export ~/Desktop/ --format html --type total
```

To get some basic help for exporting collected metadata run the following:
```
gscan export --help
```

## License

[GPLv2](https://www.gnu.org/licenses/old-licenses/gpl-2.0.html)
