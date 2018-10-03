# transmission-peerport
Sets the Peer-Port of Transmission to a Port of Perfect Privacy's Default Port-Forwarding

## Installation 

Debian/Ubuntu:

```
$ apt install golang git
$ git clone https://github.com/Mave95/transmission-peerport.git
```

## Usage

Change the following line in ```transmission-peerport.go``` to fit your Transmission config:
```
transmission, err := transmissionrpc.New("127.0.0.1", "transmission", "transmission", nil)
```

```$ cd transmission-peerport
$ go run transmission-peerport.go
```

It is recommended to run this script every couple of minutes to keep the Peer-Port up to date. One way to do so is by using a Cronjob like this:
```
*/15 * * * * /usr/bin/go run /$PATH/transmission-peerport.go
```
