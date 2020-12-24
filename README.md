# snp

VS Code code snippets generator.

## Install

```
curl -sf https://gobinaries.com/djyde/snp | sh
```

## Usage

> See [my snippets](https://github.com/djyde/snippets) as example.

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

### The `.snp` file

`.snp` is just a file with [front matter](https://jekyllrb.com/docs/front-matter) and the snippet code. You should pass at least these attributes:

- [scope](https://code.visualstudio.com/docs/editor/userdefinedsnippets#_language-snippet-scope)
- **description** The description of this snippet

The file name would become the snippet prefix (the trigger text).

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
