# binman

A tool used to download external binary's and place them in the bin folder through a central config file


# Download convention

```bash
path/downloads/package-name/operating-system/archecture/source-code
```

# Bin convention

```bash
path/bin/package-name/operating-system/archecture/source-code
```



# CLI API

```
Usage: binman <path> [..flags..]
```

Flags

- `--platforms=linux,windows etc`: comma seperated platforms to fetch 
- `--architectures=x86_64`: commoa seperated arch to fetch only
- `--no-clean`: passed to turn off pattern cleaning


# Example 

```bash
.\cli.exe . --platforms=linux,windows --architectures=x86_64
```

Gets windows and linux platform binarys and only x86_64


```bash
PS C:\dev\binman\cli> .\cli.exe . --platforms=linux,windows --architectures=x86_64
[2025-11-23 16:58:05] Resolved path: C:\dev\binman\cli
[2025-11-23 16:58:05] Found config file: C:\dev\binman\cli\binman.yml
[2025-11-23 16:58:05] Target platforms: linux, windows
[2025-11-23 16:58:05] Target architectures: x86_64
[2025-11-23 16:58:05] YAML file parsed successfully
[2025-11-23 16:58:05] Removed C:\dev\binman\cli\bin
[2025-11-23 16:58:05] Removed C:\dev\binman\cli\downloads
[2025-11-23 16:58:05] Fetching https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-pc-windows-gnu.zip
[2025-11-23 16:58:06] SHA256 verified for C:\dev\binman\cli\downloads\ripgrep\windows\x86_64\ripgrep-15.1.0-x86_64-pc-windows-gnu.zip
[2025-11-23 16:58:06] Fetching https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-unknown-linux-musl.tar.gz
[2025-11-23 16:58:06] SHA256 verified for C:\dev\binman\cli\downloads\ripgrep\linux\x86_64\ripgrep-15.1.0-x86_64-unknown-linux-musl.tar.gz
[2025-11-23 16:58:06] Skipping fetch aarch64
PS C:\dev\binman\cli> 
```

Examnple bin folder

```bash
PS C:\dev\binman\cli\bin> tree
Folder PATH listing for volume Windows
Volume serial number is 124E-B996
C:.
└───ripgrep
    ├───linux
    │   └───x86_64
    └───windows
        └───x86_64
PS C:\dev\binman\cli\bin> 
```