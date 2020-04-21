# ai  translate

[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg?style=flat-square)](https://github.com/996icu/996.ICU/blob/master/LICENSE)

> 监听粘贴版新的英文内容并翻译成中文

> 仅在`Manjaro`上测试,其他平台未知,有问题或建议请提交`Pull Requests`或`Issues`

![](screenshots/1.png)

## 特色
 - 自动翻译 `clipboard` 为英语的内容
 - 支持 `Google Translate` 和 `YouDao Translate`
 - 开源,不用担心 `clipboard` 内容泄漏问题

## 使用

```
./aitranslate // 监听粘贴版新的英文内容并翻译
./aitranslate -once // 只运行翻译一次,完成后退出
```

## 依赖 `xsel`或`xclip`

```
// Manjaro/Arch:
sudo pacman -S xsel xclip
// Debian/Ubuntu
sudo apt-get install xsel xclip
// Fedora
sudo dnf install xsel xclip
```

## TODO
 - [ ] auto translate text for the clipboard and show on screen 
 - [ ] Support custom multi-language
 - [ ] add testing code

## 帮助
```
➜  ~ ./aitranslate -h
Usage of ./aitranslate:
  -c string
        translate channel:  Google or YouDao. (default "Google")
  -once
        run once, default auto translate for clipboard english content.
  -v    show version and exit.
```

## screenshots
![](screenshots/2.png)