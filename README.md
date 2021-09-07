# YAMLFMT

[![GitMoji](https://img.shields.io/badge/Gitmoji-%F0%9F%8E%A8%20-FFDD67.svg)](https://gitmoji.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
![Lines Of Code](https://img.shields.io/tokei/lines/github.com/UltiRequiem/yamlfmt?color=blue&label=Total%20Lines)
![CodeQL](https://github.com/UltiRequiem/yamlfmt/workflows/CodeQL/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/UltiRequiem/yamlfmt)](https://goreportcard.com/report/github.com/UltiRequiem/yamlfmt)


### Usage

- To format one or more files and write to stdout:

  ```bash
  yamlfmt a.yaml b.yaml c.yaml
  ```

- To format one or more files in the replace mode:

  ```bash
  yamlfmt -w a.yaml b.yaml c.yaml
  ```

- To format stdin and write to stdout:

  ```bash
  cat a.yaml | yamlfmt
  ```

- To format stdin and write to a file:

  ```bash
  cat a.yaml | yamlfmt > b.yaml
  ```

### Editor Integration

- Neovim / Vim

Put this somewhere in your configuration:

```viml
au FileType yaml let &l:formatprg= "yamlfmt /dev/stdin"
```
