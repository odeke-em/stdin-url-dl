# stdin-url-dl
download from stdin line by line

## Installation

```shell
$ go get github.com/odeke-em/stdin-url-dl
```

## Usage

```shell
$ cat urls.txt | stdin-url-dl

dfe3f252-52a1-4aa6-84d8-74eb8852611e
cb14e0ac-87a6-475d-b9b0-9efb1b5d45b5
5aedaee3-1c2c-4397-a6e3-abc3152e4849
fde1f47b-ca1d-47a6-a594-5c5e1b10e3b3
44f132ca-ba34-41a7-8327-75b1cc0895fd
```

The UUIDs will be the content of each url line by line downloaded and written to disk
