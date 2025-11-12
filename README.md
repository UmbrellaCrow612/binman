# binman

A tool used to download external binary's and place them in the bin folder through a central config file


Convention 

```yml
binaries:
  - name: ripgrep
    version: 15.1.0
    url:
      linux: https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-unknown-linux-musl.tar.gz
      darwin: https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-apple-darwin.tar.gz
      windows: https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-pc-windows-msvc.zip
    sha256:
      linux: 1c9297be4a084eea7ecaedf93eb03d058d6faae29bbc57ecdaf5063921491599
      darwin: 64811cb24e77cac3057d6c40b63ac9becf9082eedd54ca411b475b755d334882
      windows: 124510b94b6baa3380d051fdf4650eaa80a302c876d611e9dba0b2e18d87493a
```

output example

```
bin
└─ ripgrep
   ├─ darwin
   │  └─ [extracted files from ripgrep-13.0.0-x86_64-apple-darwin.tar.gz]
   ├─ linux
   │  └─ [extracted files from ripgrep-13.0.0-x86_64-unknown-linux-musl.tar.gz]
   └─ windows
      └─ [extracted files from ripgrep-13.0.0-x86_64-pc-windows-msvc.zip]
```

exmpale 

```powershell
PS C:\dev\binman\cli> .\cli.exe .
Resolved path: C:\dev\binman\cli
Found config file: C:\dev\binman\cli\binman.yml
YAML file parsed successfully
Successfully removed folder: C:\dev\binman\cli\bin
Successfully removed folder: C:\dev\binman\cli\downloads
Fetching windows -> https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-pc-windows-msvc.zip
Downloaded windows -> C:\dev\binman\cli\downloads\ripgrep\windows\ripgrep-15.1.0-x86_64-pc-windows-msvc.zip
SHA256 verified for C:\dev\binman\cli\downloads\ripgrep\windows\ripgrep-15.1.0-x86_64-pc-windows-msvc.zip
Fetching linux -> https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-unknown-linux-musl.tar.gz
Downloaded linux -> C:\dev\binman\cli\downloads\ripgrep\linux\ripgrep-15.1.0-x86_64-unknown-linux-musl.tar.gz
SHA256 verified for C:\dev\binman\cli\downloads\ripgrep\linux\ripgrep-15.1.0-x86_64-unknown-linux-musl.tar.gz
Fetching darwin -> https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-apple-darwin.tar.gz
Downloaded darwin -> C:\dev\binman\cli\downloads\ripgrep\darwin\ripgrep-15.1.0-x86_64-apple-darwin.tar.gz
SHA256 verified for C:\dev\binman\cli\downloads\ripgrep\darwin\ripgrep-15.1.0-x86_64-apple-darwin.tar.gz
Processing folder: ripgrep
Processing platform: darwin
Extracting TAR.GZ: C:\dev\binman\cli\downloads\ripgrep\darwin\ripgrep-15.1.0-x86_64-apple-darwin.tar.gz -> C:\dev\binman\cli\bin\ripgrep\darwin
Processing platform: linux
Extracting TAR.GZ: C:\dev\binman\cli\downloads\ripgrep\linux\ripgrep-15.1.0-x86_64-unknown-linux-musl.tar.gz -> C:\dev\binman\cli\bin\ripgrep\linux
Processing platform: windows
Extracting ZIP: C:\dev\binman\cli\downloads\ripgrep\windows\ripgrep-15.1.0-x86_64-pc-windows-msvc.zip -> C:\dev\binman\cli\bin\ripgrep\windows
Successfully removed downloads folder: C:\dev\binman\cli\downloads
PS C:\dev\binman\cli> 
```