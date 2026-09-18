# Moment

[![Actions Status](https://github.com/fantasticmao/moment/workflows/ci/badge.svg)](https://github.com/fantasticmao/moment/actions)
![Go Version](https://img.shields.io/github/go-mod/go-version/fantasticmao/moment)
[![Go Report Card](https://goreportcard.com/badge/github.com/fantasticmao/moment)](https://goreportcard.com/report/github.com/fantasticmao/moment)
[![Release](https://img.shields.io/github/v/release/fantasticmao/moment)](https://github.com/fantasticmao/moment/releases)
[![License](https://img.shields.io/github/license/fantasticmao/moment)](https://github.com/fantasticmao/moment/blob/main/LICENSE)

README [English](README.md) | [中文](README_ZH.md)

## What is it

Moment is a simple (and possibly useless) command line tool that can stitch the bottom subtitles of multiple pictures (
usually video screenshots) to the bottom of the first picture.

Moment can be used to record some beautiful video moments:

![usage](doc/usage.gif)

## Download and usage

Download Moment from [GitHub Releases](https://github.com/fantasticmao/moment/releases).

Moment-1.x (deprecated) is written in Java and released as a `.jar` file, which requires the pre-installed
Java-Runtime-Environment.

Moment-2.x is written in Go and can be used directly without other dependencies. Moment-2.x provides executable files
for macOS, Linux, and Windows.

When using Moment, you need to know the following:

- specify the height of the bottom subtitle with the `-h` option.
- only jpeg and png image formats are supported, and the format and width of all images need to be the same.
- the naming rule of the spliced images is: the name of the first image plus the current timestamp,
  i.e. `${firstImage}-${unixMilli}`.
