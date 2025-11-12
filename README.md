# binman

A tool used to download external binary's and place them in the bin folder through a central config file


Convention 

```yml
binaries:
  - name: fos
    url:
      linux: https://github.com/user-attachments/files/23336838/fos_linux_amd64.zip
      darwin: https://github.com/user-attachments/files/23336836/fos_darwin_amd64.zip
      windows: https://github.com/user-attachments/files/23336840/fos_windows_amd64.zip
  - name: ripgrep
    url:
      linux: https://github.com/BurntSushi/ripgrep/releases/download/13.0.0/ripgrep-13.0.0-x86_64-unknown-linux-musl.tar.gz
      darwin: https://github.com/BurntSushi/ripgrep/releases/download/13.0.0/ripgrep-13.0.0-x86_64-apple-darwin.tar.gz
      windows: https://github.com/BurntSushi/ripgrep/releases/download/13.0.0/ripgrep-13.0.0-x86_64-pc-windows-msvc.zip
```

output example

```
bin
├─ fos
│  ├─ darwin
│  │  └─ [extracted files from fos_darwin_amd64.zip]
│  ├─ linux
│  │  └─ [extracted files from fos_linux_amd64.zip]
│  └─ windows
│     └─ [extracted files from fos_windows_amd64.zip]
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
Fetching darwin -> https://github.com/user-attachments/files/23336836/fos_darwin_amd64.zip
Downloaded darwin -> C:\dev\binman\cli\downloads\fos\darwin\fos_darwin_amd64.zip
Fetching windows -> https://github.com/user-attachments/files/23336840/fos_windows_amd64.zip
Downloaded windows -> C:\dev\binman\cli\downloads\fos\windows\fos_windows_amd64.zip
Fetching linux -> https://github.com/user-attachments/files/23336838/fos_linux_amd64.zip
Downloaded linux -> C:\dev\binman\cli\downloads\fos\linux\fos_linux_amd64.zip
Fetching windows -> https://github.com/BurntSushi/ripgrep/releases/download/13.0.0/ripgrep-13.0.0-x86_64-pc-windows-msvc.zip
Downloaded windows -> C:\dev\binman\cli\downloads\ripgrep\windows\ripgrep-13.0.0-x86_64-pc-windows-msvc.zip
Fetching linux -> https://github.com/BurntSushi/ripgrep/releases/download/13.0.0/ripgrep-13.0.0-x86_64-unknown-linux-musl.tar.gz
Downloaded linux -> C:\dev\binman\cli\downloads\ripgrep\linux\ripgrep-13.0.0-x86_64-unknown-linux-musl.tar.gz
Fetching darwin -> https://github.com/BurntSushi/ripgrep/releases/download/13.0.0/ripgrep-13.0.0-x86_64-apple-darwin.tar.gz
Downloaded darwin -> C:\dev\binman\cli\downloads\ripgrep\darwin\ripgrep-13.0.0-x86_64-apple-darwin.tar.gz
Processing folder: fos
Processing platform: darwin
Extracting ZIP: C:\dev\binman\cli\downloads\fos\darwin\fos_darwin_amd64.zip -> C:\dev\binman\cli\bin\fos\darwin
Processing platform: linux
Extracting ZIP: C:\dev\binman\cli\downloads\fos\linux\fos_linux_amd64.zip -> C:\dev\binman\cli\bin\fos\linux
Processing platform: windows
Extracting ZIP: C:\dev\binman\cli\downloads\fos\windows\fos_windows_amd64.zip -> C:\dev\binman\cli\bin\fos\windows
Processing folder: ripgrep
Processing platform: darwin
Extracting TAR.GZ: C:\dev\binman\cli\downloads\ripgrep\darwin\ripgrep-13.0.0-x86_64-apple-darwin.tar.gz -> C:\dev\binman\cli\bin\ripgrep\darwin
Processing platform: linux
Extracting TAR.GZ: C:\dev\binman\cli\downloads\ripgrep\linux\ripgrep-13.0.0-x86_64-unknown-linux-musl.tar.gz -> C:\dev\binman\cli\bin\ripgrep\linux
Processing platform: windows
Extracting ZIP: C:\dev\binman\cli\downloads\ripgrep\windows\ripgrep-13.0.0-x86_64-pc-windows-msvc.zip -> C:\dev\binman\cli\bin\ripgrep\windows
Successfully removed downloads folder: C:\dev\binman\cli\downloads
```