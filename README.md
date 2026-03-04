# 🚀 In-Memory Cache (Redis-Lite)

### 🎯 Goal

Build a concurrency-safe in-memory key-value store with expiration support.

This project will teach:

- sync.Mutex or sync.RWMutex
- TTL handling
- Goroutines
- Background cleanup
- Data structure design

---

## 📝 Requirements

Create:

    type Cache struct

Functions:

- Set(key string, value string, ttl time.Duration)
- Get(key string) (string, bool)
- Delete(key string)
- Size() int

---

## 🔎 Features

1. Keys should expire automatically.
2. Expired keys should not be returned.
3. Concurrency safe.
4. No race conditions.

---

## 🔎 Bonus

- Background cleanup goroutine.
- Configurable cleanup interval.
- Avoid memory leak.

---

## 🔎 Think About

- Should you use RWMutex?
- Where do you store expiration time?
- What happens under heavy concurrency?

---
