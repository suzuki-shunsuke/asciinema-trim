# asciinema-trim

Trim [asciinema](https://asciinema.org/)'s record file

## Motivation

[asciinema](https://github.com/asciinema/asciinema) is an awesome project to record and share your terminal sessions.

After recoding the session, sometimes you would like to trim the session.
asciinema-trim is a CLI to trim the record file.

## Install

* Homebrew
* [aqua](https://aquaproj.github.io/)
* [Download from GitHub Rleases](https://github.com/suzuki-shunsuke/asciinema-trim/releases)

```console
$ brew install suzuki-shunsuke/asciinema-trim/asciinema-trim
```

## How to use

1. Prepare for the input record file
1. Insert times you want to trim in the file
1. Run `asciinema-trim` then trimmed record file is generated

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

Please see [examples](examples).
There are three files.

* raw.cast: Generated by `asciinema rec raw.cast`
* input.cast: Update `raw.cast` to trim the record file
* output.cast: Generated by `asciinema-trim input.cast > output.cast`

```diff
$ diff raw.cast input.cast
```

## Release Note

[GitHub Releases](https://github.com/suzuki-shunsuke/asciinema-trim/releases)

## License

[MIT](LICENSE)
