你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
ARPG
====

A simple ARP (Address Resolution Protocol) support tool written in Go programming language.

![screenshot](screenshot.png)

## Key features

- Resolve MAC address from IP address
- Resolve IP address from MAC address
- Works on Windows, macOS and Linux

## Download

See [releases](https://github.com/mikan/arpg/releases) page.

## How it works

### IP to MAC

1. Check available network adapter information (e.g. eth0)
2. Send ICMP ping to the target using native `ping` command
3. Lookup ARP table using native `arp` (`ip`) command

### MAC to IP

1. Check available network adapter information (e.g. eth0)
2. Send ICMP ping to the *broadcast* address using native `ping` command
3. Lookup ARP table using native `arp` (`ip`) command

## Limitations

- Works only within a local area network
- IPv6 is not supported

## License

ARPG licensed under the [BSD 3-Clause License](LICENSE).

## Author

- [mikan](https://github.com/mikan)
