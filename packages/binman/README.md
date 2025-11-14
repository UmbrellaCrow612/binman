# **Umbr Binman**

**Umbr Binman** is an NPM wrapper around a CLI utility for downloading, verifying, and extracting external binaries using a simple YAML configuration file.
It’s designed to simplify dependency setup across **Linux**, **macOS**, and **Windows** environments — perfect for build pipelines and developer tooling.

---

## 🚀 Installation

Install locally as a development dependency:

```bash
npm i umbr-binman --save-dev
```

Or use directly via `npx`:

```bash
npx binman .
```

---

## ⚙️ Usage

Create a configuration file named **`binman.yml`** in your project root.

### Example Configuration

```yml
binaries:
  - name: ripgrep
    version: 15.1.0
    url:
      linux: https://github.com/BurntSushi/ripgrep/releases/download/{{version}}/ripgrep-{{version}}-x86_64-unknown-linux-musl.tar.gz
      darwin: https://github.com/BurntSushi/ripgrep/releases/download/{{version}}/ripgrep-{{version}}-x86_64-apple-darwin.tar.gz
      windows: https://github.com/BurntSushi/ripgrep/releases/download/{{version}}/ripgrep-{{version}}-x86_64-pc-windows-msvc.zip
    sha256:
      linux: 1c9297be4a084eea7ecaedf93eb03d058d6faae29bbc57ecdaf5063921491599
      darwin: 64811cb24e77cac3057d6c40b63ac9becf9082eedd54ca411b475b755d334882
      windows: 124510b94b6baa3380d051fdf4650eaa80a302c876d611e9dba0b2e18d87493a
    pattern:
      linux: "^rg$"
      darwin: "^rg$"
      windows: "^rg\\.exe$"
```

### Key-Value Placeholders

Each binary can define as many key-value pairs as needed, such as `version`, `arch`, or custom variables.
You can then use `{{key}}` placeholders in URLs, patterns, or other fields to dynamically replace them at runtime.

For example, if you add a `platform` key:

```yml
binaries:
  - name: mytool
    version: 2.0.1
    platform: x86_64
    url:
      linux: https://example.com/downloads/{{version}}/{{platform}}/mytool.tar.gz
```

When Umbr Binman processes this binary, it replaces `{{version}}` and `{{platform}}` with the values defined under the binary, producing the final URL:

```
https://example.com/downloads/2.0.1/x86_64/mytool.tar.gz
```

This makes it easy to manage multiple binaries, versions, or architectures without repeating URLs.

---

## 🛠️ CLI API

The CLI usage is as follows:

```bash
binman [path] [..flags..]
```

For example:

```bash
binman .
```

This resolves `.` to the current folder, finds the configuration file, downloads the binaries, and cleans them according to the specified `pattern`.

### Flags

* `--no-clean`
  Prevents automatic cleanup of downloaded binary folders. By default, files not matching the `pattern` regex for the platform are removed.
  Example: using this flag keeps all files instead of cleaning them.

* `--platform=<platform>`
  Downloads and processes binaries only for the specified platform (`linux`, `darwin`, or `windows`).
  Example: `--platform=linux` will only download and clean Linux binaries.


# Example


```yml
binaries:
  - name: ripgrep
    version: 15.1.0
    url:
      linux: https://github.com/BurntSushi/ripgrep/releases/download/{{version}}/ripgrep-{{version}}-x86_64-unknown-linux-musl.tar.gz
      darwin: https://github.com/BurntSushi/ripgrep/releases/download/{{version}}/ripgrep-{{version}}-x86_64-apple-darwin.tar.gz
      windows: https://github.com/BurntSushi/ripgrep/releases/download/{{version}}/ripgrep-{{version}}-x86_64-pc-windows-msvc.zip
    sha256:
      linux: 1c9297be4a084eea7ecaedf93eb03d058d6faae29bbc57ecdaf5063921491599
      darwin: 64811cb24e77cac3057d6c40b63ac9becf9082eedd54ca411b475b755d334882
      windows: 124510b94b6baa3380d051fdf4650eaa80a302c876d611e9dba0b2e18d87493a
    pattern:
      linux: "^rg$"
      darwin: "^rg$"
      windows: "^rg\\.exe$"
```

## Building for linux

```bash
binman . --no-clean --platform=linux
```

stdout

```powershell
Resolved path: C:\dev\binman\cli
Found config file: C:\dev\binman\cli\binman.yml
Skipping bin folder cleaning (--no-clean)
Target platform: linux
All URLs resolved successfully
YAML file parsed successfully
Successfully removed folder: C:\dev\binman\cli\bin
Successfully removed folder: C:\dev\binman\cli\downloads
Fetching linux -> https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-unknown-linux-musl.tar.gz
Downloaded linux -> C:\dev\binman\cli\downloads\ripgrep\linux\ripgrep-15.1.0-x86_64-unknown-linux-musl.tar.gz
SHA256 verified for C:\dev\binman\cli\downloads\ripgrep\linux\ripgrep-15.1.0-x86_64-unknown-linux-musl.tar.gz
Processing folder: ripgrep
Processing platform: linux
Extracting TAR.GZ: C:\dev\binman\cli\downloads\ripgrep\linux\ripgrep-15.1.0-x86_64-unknown-linux-musl.tar.gz -> C:\dev\binman\cli\bin\ripgrep\linux
Bin clean logic skipped
Successfully removed downloads folder: C:\dev\binman\cli\downloads
```


Bin folder

```powershell
Folder PATH listing for volume Windows
Volume serial number is 124E-B996
C:\DEV\BINMAN\CLI\BIN
\---ripgrep
    \---linux
        |   COPYING
        |   LICENSE-MIT
        |   README.md
        |   rg
        |   UNLICENSE
        |   
        +---complete
        |       rg.bash
        |       rg.fish
        |       _rg
        |       _rg.ps1
        |
        \---doc
                CHANGELOG.md
                FAQ.md
                GUIDE.md
                rg.1
```