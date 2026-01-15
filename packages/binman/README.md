# Binman CLI

**Binman** is a command-line tool for downloading, extracting, and organizing prebuilt binaries for multiple platforms and architectures. It uses a `binman.yml` configuration file to define packages, platforms, and architectures.

---

## Features

* Download packages for multiple platforms (`linux`, `darwin`, `win32`) and architectures (`x64`, `arm64`) etc.
* Verify downloads using SHA256 checksums.
* Extract and organize binaries into a consistent folder structure.
* Support for selective package, platform, and architecture builds.
* Verbose logging for detailed progress tracking.

---

## Usage

Run the CLI from the base folder containing your `binman.yml` file:

```powershell
.\cli.exe <base_path> [flags]
```

### Flags

| Flag             | Description                                                             |
| ---------------- | ----------------------------------------------------------------------- |
| `-platforms`     | Comma-separated list of platforms to build (e.g., `linux,win32,darwin`) |
| `-architectures` | Comma-separated list of architectures (e.g., `x64,arm64`)               |
| `-packages`      | Comma-separated list of packages to download (e.g., `ripgrep,fsearch`)  |
| `-verbose`       | Enable or disable detailed logging during execution (default: `true`)   |

---

### Example Commands

**Download a single package for a specific platform and architecture:**

```powershell
.\cli.exe . --platforms=win32 --architectures=x64 --packages=fsearch --verbose=true
```

**Download all packages defined in `binman.yml` for all platforms and architectures:**

```powershell
.\cli.exe .
```

# Example binman.yml config file 


```yaml
- name: ripgrep
  platforms:
    linux:
      x64:
        url: https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-unknown-linux-musl.tar.gz
        sha256: "1c9297be4a084eea7ecaedf93eb03d058d6faae29bbc57ecdaf5063921491599"
        pattern: "^rg$"
    darwin:
      x64:
        url: https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-apple-darwin.tar.gz
        sha256: "64811cb24e77cac3057d6c40b63ac9becf9082eedd54ca411b475b755d334882"
        pattern: "^rg$"
    win32:
      x64:
        url: https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-pc-windows-msvc.zip
        sha256: "124510b94b6baa3380d051fdf4650eaa80a302c876d611e9dba0b2e18d87493a"
        pattern: "^rg\\.exe$"

- name: fsearch
  platforms:
    linux:
      x64:
        url: https://github.com/UmbrellaCrow612/fsearch/releases/download/V0.01-6-g3ac8ef4-1-ge6242a7/fsearch_V0.01-6-g3ac8ef4-1-ge6242a7_linux_amd64.zip
        sha256: "787535a8a916864d7da280c385aaf4d6347c0c26fc4613c2dceb1a7b6261dce2"
        pattern: "(?i)^fsearch$"
    darwin:
      x64:
        url: https://github.com/UmbrellaCrow612/fsearch/releases/download/V0.01-6-g3ac8ef4-1-ge6242a7/fsearch_V0.01-6-g3ac8ef4-1-ge6242a7_darwin_amd64.zip
        sha256: "b2e7b014394a0eec5b820f802ed98fe70ad6a30d840775c7ed5fd68900c8f037"
        pattern: "(?i)^fsearch$"
    win32:
      x64:
        url: https://github.com/UmbrellaCrow612/fsearch/releases/download/V0.01-6-g3ac8ef4-1-ge6242a7/fsearch_V0.01-6-g3ac8ef4-1-ge6242a7_windows_amd64.zip
        sha256: "e18e6cfcba3ba5fb6626ab7dc65e3ef3ce7b3d996c806eee659c5457f64290a3"
        pattern: "(?i)^fsearch(.exe)?$"

- name: gopls
  platforms:
    linux:
      x64:
        url: https://github.com/UmbrellaCrow612/go-tools/releases/download/v0.0.3/gopls-linux-amd64.tar.gz
        sha256: "f7163011d877bd16b611836f729da7ff0f44ffaa372caf9aea96cda4b3f9f59b"
        pattern: "^gopls$"
    darwin:
      x64:
        url: https://github.com/UmbrellaCrow612/go-tools/releases/download/v0.0.3/gopls-darwin-arm64.tar.gz
        sha256: "5ea718a1b0ee6ca6c12da6c7db093b378469e173ef9288adba65d43f5e2c3786"
        pattern: "^gopls$"
    win32:
      x64:
        url: https://github.com/UmbrellaCrow612/go-tools/releases/download/v0.0.3/gopls-windows-amd64.tar.gz
        sha256: "3bca2855397d459b71997dd8cb22cf43efcbb950d22aeb14c3663f403179b0ec"
        pattern: "^gopls(.exe)?$"

```