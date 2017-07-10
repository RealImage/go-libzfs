package main

import (
    "fmt"
    "os"
    "github.com/RealImage/go-libzfs"
)

func main () {
    poolName := os.Args[1]
    fmt.Println("Trying to open pool having name " + poolName)

    pool, err := zfs.PoolOpen(poolName)
    if err != nil {
        fmt.Println("Failed to open pool: " + err.Error())
        return
    }

    defer pool.Close()

    status, err := pool.Status()
    if err != nil {
        fmt.Println("Failed to get status: " + err.Error())
        return
    }

    if status != zfs.PoolStatusOk {
        fmt.Println("Zpool status failed with code: " + string(status))
        return
    }

    state, err := pool.State()
    if err != nil {
        fmt.Println("Failed to get Zpool state: " + err.Error())
        return
    }

    fmt.Println("Zpool state: " + zfs.PoolStateToName(state))
}
