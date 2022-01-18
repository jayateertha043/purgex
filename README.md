<h1 align="center">purgex</h1>

>Multi-Threaded PURGE Request Method Check Tool


## REQUIREMENTS AND INSTALLATION

Build purgex:
```
git clone https://github.com/jayateertha043/purgex.git
cd purgex
go build purgex.go
```

or

Install using Go get:

```
go install github.com/jayateertha043/purgex@latest
```

Run purgex:

```
.\purgex -h
```


>Note:Ensure you have git version>1.8

## USAGE:

```
type urls.txt|purgex
```

```
cat urls.txt | purgex
```

```
echo "https://jayateerthag.in/" | purgex -nobanner
```

Flags:

```
Usage of purgex:
  -headers string
        To use Custom Headers headers.json file
  -maxrequest int
        Maximum requests/urls to try (default 1000)
  -nobanner
        Disable Banner
  -nostatus
        Outputs only urls with status code between 100-300
  -proxy string
        Use custom proxy [http://ip:port or https://ip:port]
  -t int
        Enter amount of threads (default 8)
  -timeout int
        Enter request timeout in seconds (default 3)
```

## Author

👤 **Jayateertha G**

* Twitter: [@jayateerthaG](https://twitter.com/jayateerthaG)
* Github: [@jayateertha043](https://github.com/jayateertha043)

