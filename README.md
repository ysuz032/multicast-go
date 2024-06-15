# multicast-go

## 構成

## 利用方法
### AWS EC2
#### 準備
AWSマネジメントコンソールでネットワークインターフェースを２つアタッチします

#### インストール
EC2にログインして以下の手順を実行します

goアプリをビルドします
```
make build
```
/usr/local/binにインストールします
```
sudo make install
```

#### 起動
EC2にログインして以下の手順を実行します

環境変数を設定
```
export MULTICAST_ADDR=224.0.0.1
export MULTICAST_PORT=28000
```
> **設定値について:**
>
> アドレスはマルチキャストアドレス（224.0.0.0 ~ 239.255.255.255）の範囲で指定してください。
> 
> ポートはwell-known-portを除いて、他のアプリケーションと被らない範囲のものを指定してください。
> 
> 例では `28000` としています。

receiver01を起動
```
nohup receiver --iface enX0 > /tmp/receiver01.log 2>&1 &
```

receiver02を起動
```
nohup receiver --iface enX1 > /tmp/receiver02.log 2>&1 &
```

senderを起動
```
sender
```

#### 送受信確認
senderがフォアグラウンドで動いているので、停止するまで待ちます。1秒ごとに１つマルチキャストパケットを送信するので約10秒待機してください。
senderアプリが停止したら、次のコマンドを実行して、受信ログが表示されているか確認します。
- receiver01
```
less /tmp/receiver01.log
```
以下のような受信ログが出ていればマルチキャストパケットが受信できています。ログ内のIPアドレス部分はネットワークインターフェースのプライベートIPアドレスが表示されるので、環境によって変わります。
```
nohup: ignoring input
use interface: enX0
received: message 0 from 10.0.0.12:44057
received: message 1 from 10.0.0.12:44057
received: message 2 from 10.0.0.12:44057
received: message 3 from 10.0.0.12:44057
received: message 4 from 10.0.0.12:44057
received: message 5 from 10.0.0.12:44057
received: message 6 from 10.0.0.12:44057
received: message 7 from 10.0.0.12:44057
received: message 8 from 10.0.0.12:44057
received: message 9 from 10.0.0.12:44057
```

- receiver02
```
less /tmp/receiver02.log
```
以下のような受信ログが出ていればマルチキャストパケットが受信できています。ログ内のIPアドレス部分はネットワークインターフェースのプライベートIPアドレスが表示されるので、環境によって変わります。
```
nohup: ignoring input
use interface: enX1
received: message 0 from 10.0.0.12:44057
received: message 1 from 10.0.0.12:44057
received: message 2 from 10.0.0.12:44057
received: message 3 from 10.0.0.12:44057
received: message 4 from 10.0.0.12:44057
received: message 5 from 10.0.0.12:44057
received: message 6 from 10.0.0.12:44057
received: message 7 from 10.0.0.12:44057
received: message 8 from 10.0.0.12:44057
received: message 9 from 10.0.0.12:44057
```
