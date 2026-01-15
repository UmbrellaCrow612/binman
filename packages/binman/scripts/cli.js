#!/usr/bin/env node

/**
 * binman CLI launcher
 * Detects platform/arch and runs the correct shipped binary
 */

const { spawn } = require('child_process');
const path = require('path');
const fs = require('fs');

// Map Node.js process.platform/arch to folder names and Go executable
const platformMap = {
  win32: 'win32',
  darwin: 'darwin',
  linux: 'linux',
  freebsd: 'freebsd',
  openbsd: 'openbsd',
  sunos: 'sunos',
  aix: 'aix'
};

const archMap = {
  x64: 'x64',
  x32: 'x32',
  ia32: 'x32',
  arm: 'arm',
  arm64: 'arm64',
  mips: 'mips',
  mipsel: 'mipsel',
  ppc64: 'ppc64',
  s390x: 's390x'
};

// Detect current platform and architecture
const nodePlatform = process.platform;
const nodeArch = process.arch;

const osFolder = platformMap[nodePlatform];
const archFolder = archMap[nodeArch];

if (!osFolder || !archFolder) {
  console.error(`Unsupported platform or architecture: ${nodePlatform}/${nodeArch}`);
  process.exit(1);
}

// Build the path to the binary
const binName = osFolder === 'win32' ? 'binman.exe' : 'binman';
const binPath = path.resolve(__dirname, '..', 'bin', osFolder, archFolder, binName);

// Check if binary exists
if (!fs.existsSync(binPath)) {
  console.error(`Binary not found for this platform/arch: ${binPath}`);
  process.exit(1);
}

// Grab arguments passed to the Node CLI and forward them to the binary
const args = process.argv.slice(2);

// Spawn the binary
const child = spawn(binPath, args, { stdio: 'inherit' });

child.on('close', (code) => {
  process.exit(code);
});
