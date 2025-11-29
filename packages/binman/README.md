# umbr-binman

A cross-platform utility to download, verify, and extract external binaries into a central `bin` folder using a YAML configuration file.
This package wraps the `binman` CLI for seamless integration in Node.js projects.

---

## Installation

Install as a dev dependency:

```bash
npm install umbr-binman --save-dev
```

## Usage

After installation, you can run the CLI using:

```bash
npx binman <path-to-config> [flags]
```

Example:

```bash
npx binman .
```

### CLI Flags

* `--platforms=linux,windows` — comma-separated list of platforms to fetch
* `--architectures=x86_64` — comma-separated architectures to fetch
* `--no-clean` — skip cleaning the `bin` and `downloads` folders before fetching

### Example

```bash
npx binman . --platforms=linux,windows --architectures=x64
```

Output:

```
[2025-11-23 16:58:05] Resolved path: /project/cli
[2025-11-23 16:58:05] Found config file: /project/cli/binman.yml
[2025-11-23 16:58:05] Target platforms: linux, windows
[2025-11-23 16:58:05] Target architectures: x86_64
[2025-11-23 16:58:05] YAML file parsed successfully
[2025-11-23 16:58:05] Removed /project/cli/bin
[2025-11-23 16:58:05] Removed /project/cli/downloads
[2025-11-23 16:58:05] Fetching https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-pc-windows-gnu.zip
[2025-11-23 16:58:06] SHA256 verified
[2025-11-23 16:58:06] Fetching https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-unknown-linux-musl.tar.gz
[2025-11-23 16:58:06] SHA256 verified
```

---

## Folder Conventions

### Downloads folder

```
downloads/<package-name>/<platform>/<architecture>/<source-file>
```

### Bin folder

```
bin/<package-name>/<platform>/<architecture>/<binary>
```

Example:

```
bin/ripgrep/
├─ linux/
│  └─ x86_64/
└─ windows/
   └─ x86_64/
```

---

## Example `binman.yml`

Valid platforms are `windows`, `linux`, or `darwin` for now.

Valid architecture keys are those found in Node's `process.arch`. Only valid keys are permitted — alias names such as `x86_64` are not allowed. For example, use `x64` instead of the alias.

```yaml
binaries:
  - name: ripgrep
    urls:
      linux:
        x64: https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-unknown-linux-musl.tar.gz
      windows:
        x64: https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-pc-windows-gnu.zip

    sha256:
      linux:
        x64: 1c9297be4a084eea7ecaedf93eb03d058d6faae29bbc57ecdaf5063921491599
      windows:
        x64: 0bf217086ecb1392070020810b888bd405cb1dd5f088c16c45d9de1e5ea6b638

    patterns:
      linux:
        x64: "^rg$"
      windows:
        x64: "^rg\\.exe$"
```

---

# Binman Resolve

Binman resolve will try to resolve the binary path to a specific binary you defined above in the config using the package name key, the possible executable names, and the path to the bin directory.

```js
import { binmanResolve } from "umbr-binman";

let exePath = await binmanResolve("ripgrep", ["rg"], "path/to/bin");
```
