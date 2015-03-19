## kangen v0.1.1

#### Requirements

  - [Golang](https://golang.org/) >= 1
  - [Redis](http://redis.io/) >= 1.3.10
  - [spf13/cobra](https://github.com/spf13/cobra)
  - [garyburd/redigo](https://github.com/garyburd/redigo)
  - [VividCortex/godaemon](https://github.com/VividCortex/godaemon)

#### Installation

    $ git clone git://github.com/hico-horiuchi/kangen.git
    $ cd kangen
    $ go get ./...
    $ sudo make install

#### Usage

    URL shortening tool by golang
    https://github.com/hico-horiuchi/kangen
    
    Usage:
      kangen [command]
    
    Available Commands:
      add         Add shorten URL
      remove      Remove shorten URL
      list        Show all pairs of shorten and URL
      server      Start kangen server (http daemon)
      version     Print kangen version
      help        Help about any command
    
    Use "kangen help [command]" for more information about a command.

#### License

kangen is released under the [MIT license](https://raw.githubusercontent.com/hico-horiuchi/ohgi/master/LICENSE).
