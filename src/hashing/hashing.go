package hashing

import (
    "hash/fnv"
)

func HashString(key string) uint32 {
    hasher := fnv.New32()
    hasher.Write([]byte(key))
    return hasher.Sum32()
}
