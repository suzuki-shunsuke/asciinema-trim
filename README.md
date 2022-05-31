# asciinema-trim

Trim and change the playback speed of [asciinema](https://asciinema.org/)'s session

## Motivation

[asciinema](https://github.com/asciinema/asciinema) is an awesome project to record and share your terminal sessions.

After recoding the session, sometimes you would like to trim the session, or change the playback speed.
asciinema-trim is a CLI to trim the record file and change the playback speed.

## Install

* Homebrew
* [aqua](https://aquaproj.github.io/)
* [Download from GitHub Rleases](https://github.com/suzuki-shunsuke/asciinema-trim/releases)

```console
$ brew install suzuki-shunsuke/asciinema-trim/asciinema-trim
```

## Usage

```console
$ asciinema-trim <input file>
```

The trimmed record file is outputted to the standard output.

e.g.

```console
$ asciinema-trim input.cast > output.cast
```

## Example

### Trimming

input.cast

```
[1.000000, "o", "h"]
[2.000000, "o", "e"]
[4.500000, "o", "r"]
[7.000000, "o", "o"]
```

Insert `2` to trim 2 seconds.

```
[1.000000, "o", "h"]
[2.000000, "o", "e"]
2
[4.500000, "o", "r"]
[7.000000, "o", "o"]
```

Run asciiname-trim.

```
$ asciinema-trim input.cast > output.cast
```

Result

```
[1.000000, "o", "h"]
[2.000000, "o", "e"]
[2.500000, "o", "r"]
[5.000000, "o", "o"]
```

Float64 is also available.

```
[2.000000, "o", "e"]
1.5
[4.500000, "o", "r"]
```

### Change the speed

input.cast

```
[1.000000, "o", "a"]
[2.000000, "o", "b"]
[6.000000, "o", "c"]
[8.000000, "o", "d"]
[10.000000, "o", "e"]
```

Insert `*2` to set the playback speed to 2x.

```
[1.000000, "o", "a"]
[2.000000, "o", "b"]
*2
[6.000000, "o", "c"]
[8.000000, "o", "d"]
[10.000000, "o", "e"]
```

Run asciiname-trim.

```
$ asciinema-trim input.cast > output.cast
```

Result

```
[1.000000, "o", "a"]
[2.000000, "o", "b"]
[4.000000, "o", "c"]
[5.000000, "o", "d"]
[6.000000, "o", "e"]
```

Float64 is also available.

```
[2.000000, "o", "b"]
*1.5
[6.000000, "o", "c"]
```

When the playback speed is set multiple times, the playback speed is reset.
For example, when `*4` is set after `*2`, the playback speed is not 8x but 4x.

```
[1.000000, "o", "a"]
[2.000000, "o", "b"]
*2
[6.000000, "o", "c"]
[8.000000, "o", "d"]
*4
[10.000000, "o", "e"]
```

Result

```
[1.000000, "o", "a"]
[2.000000, "o", "b"]
[4.000000, "o", "c"]
[5.000000, "o", "d"]
[5.500000, "o", "e"]
```

## Release Note

[GitHub Releases](https://github.com/suzuki-shunsuke/asciinema-trim/releases)

## License

[MIT](LICENSE)
