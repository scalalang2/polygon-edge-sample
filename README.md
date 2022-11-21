## How to reproduce

```shell
$ cd ./e2e/ && docker-compose up
$ go run main.go
```

**Results**

```shell
INFO[11-21|11:05:43] tx committed                             txhash=0xbcbe7ab14d48f54ab3163d5fd2855c8462e77d42c3a0a87150cec4e580b25f33
INFO[11-21|11:05:43] deployed token contract to polygon edge  contractAddr=0x1E5dB98eF20902Fa3c0B95Ef1781Ae260BFF77C7
INFO[11-21|11:05:43] let's transfer 100 tokens to bob from alice
INFO[11-21|11:05:43] [before] balanceOf alice                 balance=0
INFO[11-21|11:05:43] [before] balanceOf bob                   balance=1000000
EROR[11-21|11:05:44] failed to execute 'transfer' method      err="failed to get header from block hash or block number"
exit status 1
```