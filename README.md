# ai  translate

[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg?style=flat-square)](https://github.com/996icu/996.ICU/blob/master/LICENSE)

> 获取粘贴版内容,判断是否为英语,是否翻译并通知翻译内容

> 仅在`Manjaro`上测试,其他平台未知,有问题或建议请提交`Pull Requests`或`Issues`

![](screenshots/1.png)

## 特色
 - 自动翻译 `clipboard` 为英语的内容
 - 支持 `Google Translate` 和 `YouDao Translate`
 - 开源,不用担心 `clipboard` 内容泄漏问题

## 使用

```
./aitranslate -auto // 然后就可以选中复制需要翻译的英文,就会自动翻译 显示翻译内容
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
  -auto
        auto translate for clipboard english content.
  -c string
        translate channel:  Google or YouDao. (default "Google")
  -v    show version and exit.
```

## screenshots
![](screenshots/2.png)