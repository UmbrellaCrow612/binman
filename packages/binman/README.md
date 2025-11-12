# Umbr Binman

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
```

---

## ▶️ Running Binman

To download and extract binaries, simply run:

```bash
npx binman .
```

Binman will:

1. Parse your `binman.yml` configuration.
2. Download binaries for supported platforms.
3. Verify their integrity via SHA256.
4. Extract archives into the `bin/` directory.
5. Clean up temporary download files automatically.

---

## 🧩 Example Output

```powershell
PS C:\dev\binman\cli> npx binman .
Resolved path: C:\dev\binman\cli
Found config file: C:\dev\binman\cli\binman.yml
YAML file parsed successfully
All URLs resolved successfully
Successfully removed folder: C:\dev\binman\cli\bin
Successfully removed folder: C:\dev\binman\cli\downloads
Fetching darwin -> https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-apple-darwin.tar.gz
Downloaded darwin -> C:\dev\binman\cli\downloads\ripgrep\darwin\ripgrep-15.1.0-x86_64-apple-darwin.tar.gz
SHA256 verified for darwin binary
Fetching windows -> https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-pc-windows-msvc.zip
Downloaded windows -> C:\dev\binman\cli\downloads\ripgrep\windows\ripgrep-15.1.0-x86_64-pc-windows-msvc.zip
SHA256 verified for windows binary
Processing and extracting all platforms...
Successfully removed downloads folder.
```

---

## 🧠 Why Use Umbr Binman?

* ✅ **Cross-platform binary management** (Linux, macOS, Windows)
* ✅ **Automatic checksum verification** (SHA256)
* ✅ **Simplifies setup for CI/CD environments**
* ✅ **Keeps binaries versioned and reproducible**
* ✅ **Removes manual download/extract hassle**

---

## 📂 Output Structure

After running Binman, your project will contain:

```
bin/
  ripgrep/
    linux/
    darwin/
    windows/
```

Each folder contains the extracted binary files for the respective platform.

---