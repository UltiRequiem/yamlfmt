# YAMLFMT

[![Go Report Card](https://goreportcard.com/badge/github.com/UltiRequiem/yamlfmt)](https://goreportcard.com/report/github.com/UltiRequiem/yamlfmt)
[![Go Reference](https://pkg.go.dev/badge/github.com/UltiRequiem/yamlfmt/pkg.svg)](https://pkg.go.dev/github.com/UltiRequiem/yamlfmt/pkg)

A simple and extensible [yaml](https://yaml.org) formatter.

Code coverage: https://ulti.js.org/yamlfmt

## Installation

```bash
go install github.com/UltiRequiem/yamlfmt@latest
```

You can also use the binaries from
[releases](https://github.com/UltiRequiem/yamlfmt/releases).

## Usage

- To format one or more files and write to stdout:

  ```bash
  yamlfmt a.yaml b.yaml c.yaml
  ```

- To format one or more files in the replace mode:

  ```bash
  yamlfmt -w a.yaml b.yaml c.yaml
  ```

  If you want to log the files that have been formatted, you can use the `-l`
  flag also.

- To format stdin and write to stdout:

  ```bash
  cat a.yaml | yamlfmt
  ```

- To format stdin and write to a file:

  ```bash
  cat a.yaml | yamlfmt > b.yaml
  ```

- To format every file in your current directory and subdirectories:

  - Using `find`:

    ```bash
    yamlfmt -w $(find -name "*.yaml")
    ```

  - Using `fd`:

    ```bash
    yamlfmt -w $(fd -H -e yaml)
    ```

## Editor Integration

### Neovim / Vim

```viml
au FileType yaml let &l:formatprg= "yamlfmt /dev/stdin"
```

It can Probably integrate with others editors easily but I only use Neovim. If
you know how to integrate it with some other editor, please open a pull requests
or issue with the information.

## Authors

[Eliaz Bobadilla](https://ultirequiem.xyz) - Creator and Maintainer 💪

Big thanks to @Antoineio for helping with the tests and CI! 🎉

See also the full list of
[contributors](https://github.com/UltiRequiem/yamlfmt/contributors) who
participated in this project ✨

## Licence

Licensed under the MIT License 📄
