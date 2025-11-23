# Stop on any errors
$ErrorActionPreference = "Stop"

# Root output directory
$BaseOutputDir = Join-Path $PWD "packages\binman\bin"

# Ensure base output directory exists
if (-not (Test-Path $BaseOutputDir)) {
    New-Item -ItemType Directory -Path $BaseOutputDir -Force | Out-Null
}

# Go into CLI folder
Push-Location "cli"

# Platforms and their common architectures
$BuildMatrix = @{
    "linux"   = @("amd64", "arm64", "386", "arm")
    "windows" = @("amd64", "arm64", "386")
    "darwin"  = @("amd64", "arm64")
}

foreach ($platform in $BuildMatrix.Keys) {
    foreach ($arch in $BuildMatrix[$platform]) {
        # Determine file extension
        $ext = if ($platform -eq "windows") { ".exe" } else { "" }

        # Create platform-specific folder
        $PlatformDir = Join-Path $BaseOutputDir $platform
        if (-not (Test-Path $PlatformDir)) {
            New-Item -ItemType Directory -Path $PlatformDir -Force | Out-Null
        }

        # Output binary path
        $OutputName = "binman-$platform-$arch$ext"
        $OutputPath = Join-Path $PlatformDir $OutputName

        Write-Host "Building for $platform/$arch -> $OutputPath"

        # Set Go cross-compilation environment
        $env:GOOS = $platform
        $env:GOARCH = $arch

        # Build
        go build -o $OutputPath

        # Clear env
        Remove-Item Env:GOOS
        Remove-Item Env:GOARCH
    }
}

# Return to original folder
Pop-Location

Write-Host "All builds complete!"
