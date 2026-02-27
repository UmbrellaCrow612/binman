Here’s a clean **README-ready table** of all major **GOOS** and **GOARCH** values you can pass to `go build`.
This is formatted so you can paste it directly into a README.

---

# 🏗️ Go Cross-Compilation Matrix

## **Supported `GOOS` Values**

| GOOS        | Description                |
| ----------- | -------------------------- |
| `aix`       | IBM AIX                    |
| `android`   | Android systems            |
| `darwin`    | macOS                      |
| `dragonfly` | DragonFly BSD              |
| `freebsd`   | FreeBSD                    |
| `illumos`   | Illumos distributions      |
| `ios`       | Apple iOS                  |
| `js`        | WebAssembly via JS runtime |
| `linux`     | Linux distributions        |
| `netbsd`    | NetBSD                     |
| `openbsd`   | OpenBSD                    |
| `plan9`     | Plan 9                     |
| `solaris`   | Oracle Solaris             |
| `windows`   | Microsoft Windows          |

---

## **Supported `GOARCH` Values**

| GOARCH     | Description                 |
| ---------- | --------------------------- |
| `386`      | 32-bit x86                  |
| `amd64`    | 64-bit x86-64               |
| `amd64p32` | (rare) 32-bit x86-64 hybrid |
| `arm`      | ARM (32-bit)                |
| `arm64`    | ARM (64-bit)                |
| `loong64`  | LoongArch 64-bit            |
| `mips`     | MIPS (big-endian)           |
| `mipsle`   | MIPS (little-endian)        |
| `mips64`   | MIPS64 (big-endian)         |
| `mips64le` | MIPS64 (little-endian)      |
| `ppc64`    | PowerPC 64 big-endian       |
| `ppc64le`  | PowerPC 64 little-endian    |
| `riscv64`  | RISC-V 64-bit               |
| `s390x`    | IBM Z Systems               |
| `wasm`     | WebAssembly                 |

---

## **Additional Architecture Modifiers**

### **ARM (`GOARM`)**

| GOARM | Meaning                         |
| ----- | ------------------------------- |
| `5`   | ARMv5                           |
| `6`   | ARMv6                           |
| `7`   | ARMv7 (Raspberry Pi 2/3/Zero 2) |

### **AMD64 (`GOAMD64`)**

| GOAMD64 | Meaning          |
| ------- | ---------------- |
| `v1`    | Base x86-64      |
| `v2`    | Adds SSE3, SSSE3 |
| `v3`    | Adds AVX, AVX2   |
| `v4`    | Adds AVX-512     |

---

## **Example Cross-Compile Command**

```sh
GOOS=linux GOARCH=arm64 go build -o myapp-linux-arm64
```
