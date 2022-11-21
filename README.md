## How to reproduce

```shell
$ cd ./e2e/ && docker-compose up
$ go run main.go
```

**Results**

```shell
INFO[11-21|11:02:04] tx committed                             txhash=0x986b51917ed880545885c0d87e2edd78ed42841cb0c4415e3263f2a8f12636ce
INFO[11-21|11:02:04] deployed token contract to polygon edge  contractAddr=0xFEC99F22a85a2af6837c069b4C5C7c5922E9d7F2
INFO[11-21|11:02:04] [before] balanceOf alice                 balance=1000000
EROR[11-21|11:02:04] failed to execute 'transfer' method      err="failed to get header from block hash or block number"
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x5d0f38]

goroutine 1 [running]:
github.com/ethereum/go-ethereum/core/types.(*Transaction).Hash(0xc00075de50?)
        /home/scalalang/go/pkg/mod/github.com/ethereum/go-ethereum@v1.10.26/core/types/transaction.go:361 +0x38
main.main()
        /mnt/c/Users/cailhuiris/Projects/polygon-edge-sample/main.go:46 +0x507
exit status 2
```