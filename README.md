# snp

[WIP] A VS Code code snippets generator

## Install

```
curl -sf https://gobinaries.com/djyde/snp | sh
```

## Usage

```
$ mkdir my-snippets
$ cd my-snippets

$ touch csl.snp
```

```js
// csl.snp

---
scope: javascript
description: console log
---

console.log($1)
```

```bash
$ snp -a
```

`snp -a` would write a code snippet file to your VS Code snippets folder. Then when you type `csl` in VS Code, the code snippet appear in IntelliSense.

> The program only find the files with `.snp` extension in current working directory.

## CLI Options

```
$ snp --help

Usage of snp:
  -p	Should print out the json text
  -u	Should automatically update snippet file in VS Code
  -v	Print version

```

### -p

Instead of updating the file on VS Code snippets folder, you could use `-p` to just print out the JSON and then use this output to do what you want, like output to a file:

```
snp -p > my-snippet.code-snippets
```

## Build

```
go install

go build
```
