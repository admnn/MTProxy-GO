MTProxy 代理
mtproxy代理是telegram专属轻量级代理工具，使用十分方便，telegram官方内置，无需另外安装软件。

注意：即使是mtg，目前也很容易被墙了，长期使用推荐 vless、trojan 等协议翻墙。如果一定要搭建 MTProxy 代理，建议搭配防火墙白名单使用。

手动搭建mtg代理教程请参阅：VPS 手动安装最新 MTProxy（ mtg 2.0 ）代理教程

 

已经有单独的代理软件的存在了，为什么要用一款内置的mtproxy代理工具呢？

对于轻量级翻墙用户而言，他们不需要翻墙上其他网站，只是需要使用telegram和朋友聊天，这个时候再安装一个代理软件就显得没有必要。

在向朋友交流梯子等敏感话题的时候，使用微信其实及其不安全，这个时候分享tg代理然给朋友然后转到telegram聊天，相对更加安全，这个可以作为翻墙第一步存在吧。

MTG
MTG即mtproxy-go版，是mtproxy的分支版本。原版mtproxy在使用一段时间后就会被墙，go版的因为有fake tls的加入，使用代理时的流量看起来像是访问正常网站所以不容易被墙。2020年原版mtproxy在国内已经完全不能使用，建议改用MTG。

一键安装
脚本说明: Mtproto Proxy Go版 一键编译安装管理脚本
系统支持: CentOS6+ / Debian7+ / Ubuntu14+
已修复脚本问题，支持mtg2.0
 

wget -N --no-check-certificate https://github.com/whunt1/onekeymakemtg/raw/master/mtproxy_go.sh && chmod +x mtproxy_go.sh && bash mtproxy_go.sh
配置文件在 /usr/local/mtproxy-go/mtproxy.conf ，可以手动修改，配置项详细介绍参见 mtg 文档

 

  MTProxy-Go 一键管理脚本 [v2.0.0]
  
  0. 升级脚本
————————————
  1. 安装 MTProxy
  2. 更新 MTProxy
  3. 卸载 MTProxy
————————————
  4. 启动 MTProxy
  5. 停止 MTProxy
  6. 重启 MTProxy
————————————
  7. 设置 账号配置
  8. 查看 账号信息
  9. 查看 日志信息
 10. 查看 链接信息
————————————

 当前状态: 已安装 并 已启动

 请输入数字 [0-10]:
输入1运行安装脚本

[信息] MTProxy服务 管理脚本下载完成 !
[信息] 开始设置 用户配置...
请输入 MTProxy 端口 [1-65535]
(默认: 443):443

========================
        端口 :  443
========================

请输入 MTProxy 密匙（普通密钥必须为32位，[0-9][a-z][A-Z]，建议留空随机生成）
(若需要开启TLS伪装建议直接回车):
端口号建议默认443，默认回车键即可，下一步也是直接回车，开启Fake TLS。
等待片刻即可安装完成，整个过程只需要几分钟时间。你也可以手动编译安装，下方有详细参考链接。

如果报错信息如下，请手动编译安装。

[信息] 开始检查编译环境！
[信息] 开始安装编译环境！
/dl/go1.14.4.linux-amd64.tar.gz: Scheme missing.
tar: go*linux-amd64.tar.gz: Cannot open: No such file or directory
tar: Error is not recoverable: exiting now
mv: cannot stat 'go': No such file or directory
[错误] go 安装失败 !
如果Unbuntu系统下报错信息如下：

mtproxy_go.sh: line 136: make: command not found
那么输入：

apt-get update
apt-get upgrade -y
apt-get install make
设置推广信息
在弹出需要设置Tag时使用MTProxybot注册。

关注 @MTProxybot 机器人。
发送/newproxy指令，bot返回添加方式。
发送 host:port，host即你的 vps 外网 ip 地址，port 就是端口号。
发送连接密码，即之前生成的 32 位随机字符串。
接下来 bot 会返回生成的分享链接和代理tag：
使用一键脚本生成的密钥并非上面需要的32位随机字符串，

ee271082e5da56c2877f215c225eb93ffe7777772e676f6f676c652e636f6d
#表示："ee"+"271082e5da56c2877f215c225eb93ffe"+"www.google.com".encode().hex()
ee后面的 271082e5da56c2877f215c225eb93ffe 才是需要的32位随机字符，在@MTProxybot 机器人注册后即可获得Tag，在终端中输入即可。

注意：mtg 2.0 不支持推广信息，如果需要放置推广频道，建议使用1.0版本或原版mtproxy。
