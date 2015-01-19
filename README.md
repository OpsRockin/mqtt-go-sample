# mqtt-go-sample
Simple pubsub program to demonstrate MQTT with Go

## Usage

### MQTT のセットアップ

Mac OS X で [mosquitto](http://mosquitto.org/) をインストールして、起動します。

```bash
$ brew install mosquitto
$ brew install czmq zeromq
$ echo export PATH='/usr/local/sbin:$PATH' >> ~/.bash_profile
$ source ~/.bash_profile
$ mosquitto -c /usr/local/etc/mosquitto/mosquitto.conf
1421637760: mosquitto version 1.3.5 (build date 2014-10-27 15:13:47+0000) starting
1421637760: Config loaded from /usr/local/etc/mosquitto/mosquitto.conf.
1421637760: Opening ipv4 listen socket on port 1883.
1421637760: Opening ipv6 listen socket on port 1883.

```

### Pub/Sub テスト

`mosquitto` に付属する `mosquitto_sub` と `mosquitto_pub` コマンドを別々のターミナルで実行し、簡単な動作をテストしてみます。

まず、subscriber を `my/topic` トピックに待機させます。 

```bash
$ mosquitto_sub -t my/topic
```

つぎに、`hello`　メッセージを `my/topic` トピックにパブリッシュします。

```bash
$ mosquitto_pub -t my/topic -m "hello"
```

すると、subscriber がメッセージを受け取り標準出力します。

```bash
$ mosquitto_sub -t my/topic
hello
```

### Go で Pub/Sub

go で書いた publisher と subscriber を実行します。

```bash
$ mkdir mqtt-client-go
$ echo 'export GOPATH=$(PWD)' > .envrc && direnv allow .
```

publisher をスタート。

```bash
$ go get git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git
$ go run pub.go
Connected to tcp://127.0.0.1:1883

```

別のターミナルで subscriber をスタート。

```bash
$ go run sub.go
Connected to tcp://127.0.0.1:1883
```

メッセージを pub/sub 。

#### PUBLISHER

```bash
$ go run pub.go
Connected to tcp://127.0.0.1:1883
hello!  
```

#### SUBSCRIBER

```bash
$ go run sub.go 
Connected to tcp://127.0.0.1:1883
Received message on topic: NPC158.local
Message: hello!
```

### おまけ

vim で go を記述するときは、`.vimrc` に以下を追加しておくとハードタブのインデントが有効になりはかどります。

```
au BufNewFile,BufRead *.go set noexpandtab tabstop=8 shiftwidth=8
```

### ソース

[package mqtt](https://godoc.org/git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git)

