package main

import (
    "sync"
    "net/http")

var mu sync.Mutex

func withLock(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        mu.Lock()
        defer mu.Unlock()
        handler(w, r)
    }
}

