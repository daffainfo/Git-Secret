# Git-Secret
## Go scripts for finding an API key / some keywords in repository
![](https://img.shields.io/github/license/daffainfo/Git-Secret)
![](https://img.shields.io/github/issues/daffainfo/Git-Secret)
![](https://img.shields.io/github/forks/daffainfo/Git-Secret)
![](https://img.shields.io/github/stars/daffainfo/Git-Secret)
![](https://img.shields.io/github/last-commit/daffainfo/Git-Secret)

## Update V1.0.2 🚀
- Removing some checkers
- Adding example file contains github dorks
- Add support to `go install` command ([issue #5](https://github.com/daffainfo/Git-Secret/issues/5))

## Screenshoot 📷

![image](https://user-images.githubusercontent.com/36522826/128018595-990a9054-3d8a-4b1b-8c70-afc901f093eb.png)

## How to Install

To install the latest version:

```
go install github.com/daffainfo/Git-Secret@latest
```

Alternatively, you can clone this repo and build the project by running
`go build` inside the cloned repo.

## How to Use

```
./Git-Secret
```

* For path contain dorks, you can fill it with some keywords, for example
> keyword.txt
```
password
username
keys
access_keys
```

### Reference 📚

- https://github.com/odomojuli/RegExAPI
