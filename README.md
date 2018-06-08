# a-know

[![Latest Version](https://img.shields.io/github/release/a-know/a-know.svg)](https://github.com/a-know/a-know/releases)

`a-know` - CLI tool about me (a-know) written in Go.

# DESCRIPTION
`a-know` is CLI tool to show information about me (a-know).

I started developing this tool to learn golang and to learn how to release command line tools.

# INSTALLATION
## by `go get`

```
$ go get github.com/a-know/a-know
```
# USAGE

`$ a-know <command>`

## Available commands
### `twitter`

Display a-know's Twitter account name.

```
$ a-know twitter
@a_know
```

Use `--url` option, you can get twitter account page URL.

```
$ a-know twitter --url
https://twitter.com/a_know
```

### `blog`

Display a-know's blog URL.

```
$ a-know blog
http://blog.a-know.me
```

### `homepage`

Display a-know's homepage (portfolio page) URL.

```
$ a-know homepage
https://a-know.me/
```

### `github`

Display a-know's GitHub Account page.

```
$ a-know github
https://github.com/a-know
```

### `presentation`

Display a-know's slideshare URL.

```
$ a-know presentation
http://www.slideshare.net/aknow3373
```

### `photo`

Display a-know's photograph blog URL.

```
$ a-know photo
http://photos.a-know.me/
```

### `hatena`

Display a-know's Hatena profile page URL.

```
$ a-know hatena
http://profile.hatena.ne.jp/a-know/
```

### `ask`

Display a-know's ask.fm page URL.

```
$ a-know ask
http://ask.fm/a_know
```
