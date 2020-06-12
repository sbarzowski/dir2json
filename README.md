# dir2json

This is an exteremely simple tool which saves a contents of a directory as JSON.
A directory is represented as a JSON object with fields corresponding to name of the contents.
Regular file contents are represented as strings.

## Installation

```
go get https://github.com/sbarzowski/dir2json
```

## Example Usage with Jsonnet

```
dir2json my_dir > my_dir.json
jsonnet --tla-var-code my_dir=my_dir.json -e 'function(my_dir) my_dir["sub_dir"]["file"]'

```
