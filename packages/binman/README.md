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
```

---

## ▶️ Running Binman

### 🧰 Default (All Platforms)

To download and extract binaries for all defined platforms:

```bash
npx binman .
```

Binman will:

1. Parse your `binman.yml` configuration.
2. Download binaries for **all supported platforms**.
3. Verify their integrity via SHA256.
4. Extract archives into the `bin/` directory.
5. Clean up temporary downloads automatically.

---

### 🧩 Specific Platform Builds

You can also download binaries for **only one specific platform**, such as `linux`, `darwin`, or `windows`.

```bash
npx binman . linux
```

In this mode, Binman will:

1. Parse your configuration as usual.
2. Check if each binary defines a URL for the `linux` platform.
3. ❌ Exit with an error if any binary does **not** define the requested platform.
4. ✅ Download, verify, and extract only the `linux` variant for each binary.

Example Output:

```bash
npx binman . linux
Resolved path: /home/user/dev/binman
Found config file: /home/user/dev/binman/binman.yml
YAML file parsed successfully
All URLs resolved successfully
Processing platform: linux
Fetching linux -> https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-unknown-linux-musl.tar.gz
Downloaded linux -> /home/user/dev/binman/downloads/ripgrep/linux/ripgrep-15.1.0-x86_64-unknown-linux-musl.tar.gz
SHA256 verified for linux binary
Extracting linux binary...
Successfully removed downloads folder.
```

If the requested platform doesn’t exist for any binary, Binman will exit with a clear message:

```bash
❌ Binary 'ripgrep' does not define a URL for platform 'linux'
```

---

## 🧠 Why Use Umbr Binman?

* ✅ **Cross-platform binary management** (Linux, macOS, Windows)
* ✅ **Automatic checksum verification** (SHA256)
* ✅ **Simplifies setup for CI/CD environments**
* ✅ **Reproducible and versioned**
* ✅ **No manual download or extraction steps**

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
When a specific platform is specified, only that platform’s folder will be created.