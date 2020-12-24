# snp

[WIP] A VS Code code snippets generator

## Install

```
curl -sf https://gobinaries.com/djyde/snp | sh
```

## Usage

```
mkdir my-snippets
cd my-snippets

touch csl.snp
```

```js
---
scope: javascript
description: console log
---

// csl.snp
console.log($1)
```

```bash
snp -a
```

`snp -a` would write a code snippet file to your VS Code snippets folder. Then when you type `csl` in VS Code, the code snippet will show up.

## Build

```
go install

go build
```
