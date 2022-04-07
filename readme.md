# YAMLFMT

[![Go Report Card](https://goreportcard.com/badge/github.com/UltiRequiem/yamlfmt)](https://goreportcard.com/report/github.com/UltiRequiem/yamlfmt)

A simple and extensible [yaml](https://yaml.org) formatter.

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

[Eliaz Bobadilla](https://ultirequiem.com) - Creator and Maintainer 💪

See also the full list of
[contributors](https://github.com/UltiRequiem/yamlfmt/contributors) who
participated in this project ✨

## Support

Open an Issue, I will check it a soon as possible 👀

If you want to hurry me up a bit
[send me a tweet](https://twitter.com/UltiRequiem) 😆

Consider [supporting me on Patreon](https://patreon.com/UltiRequiem) if you like
my work 🙏

Don't forget to start the repo ⭐

## Licence

Licensed under the MIT License 📄
