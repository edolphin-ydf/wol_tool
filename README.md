openssl req -x509 -nodes -days 36500 -newkey rsa:2048 -keyout wol.key -out wol.crt

可以认证的WOL

要求:
可以运行golang的路由器
支持WOL的pc

原理:
在路由器运行此工具的server端, 监听连接, 当此工具的client端连接上后, 发送登陆和WOL包

当server端验证通过且接收到WOL包之后, 会广播MagicPacket, 当支持WOL的pc收到MagicPacket

即被唤醒