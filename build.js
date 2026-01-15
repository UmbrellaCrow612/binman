#!/usr/bin/env node

/**
 * Cross-platform Go build script for binman
 */

const { execSync } = require('child_process');
const fs = require('fs');
const path = require('path');

console.log('\x1b[32mStarting cross-platform build for binman...\x1b[0m\n');

// Node -> Go mappings
const platformMap = {
  win32: 'windows',
  darwin: 'darwin',
  linux: 'linux',
  freebsd: 'freebsd',
  openbsd: 'openbsd',
  sunos: 'solaris',
  aix: 'aix'
};

const archMap = {
  x64: 'amd64',
  x32: '386',
  arm: 'arm',
  arm64: 'arm64',
  ia32: '386',
  mips: 'mips',
  mipsel: 'mipsle',
  ppc64: 'ppc64',
  s390x: 's390x'
};

// Define common targets
const targets = [
  { node_platform: 'win32', node_arch: 'x64' },
  { node_platform: 'win32', node_arch: 'x32' },
  { node_platform: 'win32', node_arch: 'arm64' },
  { node_platform: 'darwin', node_arch: 'x64' },
  { node_platform: 'darwin', node_arch: 'arm64' },
  { node_platform: 'linux', node_arch: 'x64' },
  { node_platform: 'linux', node_arch: 'x32' },
  { node_platform: 'linux', node_arch: 'arm' },
  { node_platform: 'linux', node_arch: 'arm64' },
  { node_platform: 'freebsd', node_arch: 'x64' },
  { node_platform: 'freebsd', node_arch: 'arm64' }
];

const sourceDir = path.resolve(__dirname, 'cli');
const outputBase = path.resolve(__dirname, 'packages', 'binman', 'bin');

// Check if source directory exists
if (!fs.existsSync(sourceDir)) {
  console.error(`\x1b[31mError: Source directory '${sourceDir}' not found!\x1b[0m`);
  process.exit(1);
}

// Build each target
for (const target of targets) {
  const goos = platformMap[target.node_platform];
  const goarch = archMap[target.node_arch];
  const outputDir = path.join(outputBase, target.node_platform, target.node_arch);
  const filename = goos === 'windows' ? 'binman.exe' : 'binman';
  const outputPath = path.join(outputDir, filename);

  // Ensure output directory exists
  fs.mkdirSync(outputDir, { recursive: true });

  console.log(`\x1b[36mBuilding for ${target.node_platform}/${target.node_arch} (GOOS=${goos} GOARCH=${goarch})...\x1b[0m`);

  try {
    // Run Go build
    execSync(`go build -o "${outputPath}"`, {
      stdio: 'inherit',
      cwd: sourceDir,
      env: { ...process.env, GOOS: goos, GOARCH: goarch, CGO_ENABLED: '0' }
    });

    console.log(`  \x1b[32m✓ Success: ${outputPath}\x1b[0m\n`);
  } catch (err) {
    console.error(`  \x1b[31m✗ Failed: ${err.message}\x1b[0m`);
    process.exit(1);
  }
}

console.log('\x1b[32mAll builds completed successfully! ✓\x1b[0m');
process.exit(0);
