# Moment [![Actions Status](https://github.com/FantasticMao/moment/workflows/action/badge.svg)](https://github.com/FantasticMao/moment/actions) [![image](https://img.shields.io/badge/release-download-blue.svg)](https://github.com/FantasticMao/moment/releases) [![image](https://img.shields.io/badge/license-MIT-green.svg)](https://github.com/FantasticMao/moment/blob/master/LICENSE)

Moment 是一个简单（可能没用）的小工具，它可以把多张图片（一般是视频截图）的底部字幕，拼接到第一张图片底下。

## 在线演示

Moment 可以用来保存视频的精彩瞬间：

![image](doc/usage.gif)

## 下载和使用

Moment 下载地址：[https://github.com/FantasticMao/moment/release](https://github.com/FantasticMao/moment/release)。

Moment-1.0 (**已过时**) 以 Java 语言编写，以 .jar 文件发布，使用时需预先 [安装 Java 运行环境](https://www.baidu.com/s?wd=安装%20JRE)。

Moment-2.0 以 Go 语言编写，无需安装其它依赖，可以直接使用。Moment-2.0 提供了 macOS、Linux、Windows 三个平台的可执行文件。

Moment 使用手册：

macOS / Linux

```text
~$ ./moment --help
Usage:
  moment [OPTIONS]

Application Options:
  -p, --path=   The file path of image that need to be stitched
  -h, --height= The height of bottom subtitle which should be in unit px (default: 120)
      --out=    The output directory which is used to save the joined image (default: .)
      --help    Display this help message
```

Windows

```text
C:\Users\MaoMao\Desktop>moment.exe --help
Usage:
  moment.exe [OPTIONS]

Application Options:
  /p, /path:    The file path of image that need to be stitched
  /h, /height:  The height of bottom subtitle which should be in unit px (default: 120)
      /out:     The output directory which is used to save the joined image (default: .)
      /help     Display this help message
```

## 注意事项

使用 Moment 时需要注意以下几点：

- Moment 仅支持拼接 `jpeg`、`png` 格式的图片，且所有图片格式和宽度需要保持一致
- Moment 最终生成的图片名称为第一个 `-p` 参数指定的图片名称加上 `_final` 后缀
