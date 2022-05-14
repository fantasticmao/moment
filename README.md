# Moment

[![Actions Status](https://github.com/fantasticmao/moment/workflows/ci/badge.svg)](https://github.com/fantasticmao/moment/actions)
[![image](https://img.shields.io/badge/release-download-blue.svg)](https://github.com/fantasticmao/moment/releases)
[![image](https://img.shields.io/badge/license-MIT-green.svg)](https://github.com/fantasticmao/moment/blob/master/LICENSE)

Moment 是一个简单（可能没用）的小工具，它可以把多张图片（一般是视频截图）的底部字幕，拼接到第一张图片底下。

## 在线演示

Moment 可以用来保存视频的精彩瞬间：

![usage](doc/usage.gif)

## 下载和使用

Moment 下载地址请见 [GitHub Releases](https://github.com/fantasticmao/moment/releases) 。

~~Moment-1.0 (已过时) 以 Java 语言编写，以 .jar 文件发布，使用时需预先 [安装 Java 运行环境](https://www.baidu.com/s?wd=安装%20JRE) 。~~

Moment-2.x 以 Go 语言编写，无需安装其它依赖，可以直接使用。Moment-2.x 提供了 macOS、Linux、Windows 三个平台的可执行文件。

## 注意事项

使用 Moment 时需要注意以下几点：

- 通过 `-h` 选项来显式指定底部字幕的高度
- 仅支持 jpeg、png 图片格式，所有图片的格式和宽度需要保持一致
- 拼接而成的图片名称生成规则为：第一张图片名称加上当前时间戳，即 `${firstImage}-${unixMilli}`
