# ai translate

[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg?style=flat-square)](https://github.com/996icu/996.ICU/blob/master/LICENSE)

> 监听粘贴版新的内容并翻译成中文

> 仅在`Manjaro`上测试,其他平台未知,有问题或建议请提交`Pull Requests`或`Issues`

> `命令行版`请查看`cli`分支，`master`为`systray`系统托盘版本

|Linux| Windows 10|
|:---:|:---:|
|![](screenshots/linux.png)|![](screenshots/windows10.png)|

## 特色
 - 自动翻译 `clipboard` 的内容
 - 支持 `Google Translate` 和 `YouDao Translate`
 - 开源,不用担心 `clipboard` 内容泄漏问题

## TODO

 - [ ] GUI系统托盘执行,实现托盘切换翻译渠道及运行模式
 - [ ] 增加测试代码

## 使用
### 依赖
- Linux `xsel`或`xclip`
```
// Manjaro/Arch:
sudo pacman -S xsel xclip
// Debian/Ubuntu
sudo apt-get install xsel xclip
// Fedora
sudo dnf install xsel xclip
```
- Windows 10 `PowerShell`

### Linux or Mac
```
nohup ./aitranslate -c [Google|YouDao] & 
```
### Windows 
```
start /b ./aitranslate.exe -c [Google|YouDao] 
```

## Tips
> manjaro 可以开启`clipman`设置中的 `同步鼠标选择区`,实现选中自动翻译

## 依赖
### Linux `xsel`或`xclip`
```
// Manjaro/Arch:
sudo pacman -S xsel xclip
// Debian/Ubuntu
sudo apt-get install xsel xclip
// Fedora
sudo dnf install xsel xclip
```
### Windows 10 `PowerShell`

## screenshots
![](screenshots/2.png)