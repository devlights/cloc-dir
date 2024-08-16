# Overview

[cloc](https://github.com/AlDanial/cloc)をディレクトリ単位で実行し、結果を出力するラッパーです。

[cloc](https://github.com/AlDanial/cloc)は、ディレクトリを指定すると、その配下のディレクトリ全てを集計するので

１階層分の結果を出力するためには、都度ディレクトリを指定して実行する必要があります。

本ツールは、内部で [cloc](https://github.com/AlDanial/cloc) を実行して、ディレクトリ単位の結果を出力するものです。

# Requirements

```cloc``` が ```$PATH``` 内に存在するか、本実行ファイルと同じディレクトリにあること。

# Install

```sh
go install github.com/devlights/cloc-dir/cmd/cloc-dir@latest
```

# Usage

```sh
$ cloc-dir.exe -help
Usage of cloc-dir.exe:
  -dir string
        directory (default ".")
  -lang string
        Language (default "C#")
  -sep string
        separator (default ",")
```

```sh
# デフォルトでは区切り文字にカンマを付与して出力します
$ cloc-dir.exe -dir /path/to/dir -lang 'C#'

# sepオプションでタブを指定することも出来ます
$ cloc-dir.exe -dir /path/to/dir -sep "\t" -lang 'Java'
```
