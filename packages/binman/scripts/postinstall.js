#!/usr/bin/env node

const fs = require("fs");
const path = require("path");
const os = require("os");

const platform = os.platform();
const binDir = path.resolve(__dirname, "..", "bin");

const binaries = {
  linux: "binman-linux",
  darwin: "binman-darwin",
  win32: "binman.exe"
};

const binaryName = binaries[platform];
if (!binaryName) {
  console.log(`Unsupported platform: ${platform}`);
  process.exit(0);
}

const binaryPath = path.join(binDir, binaryName);

if (platform === "win32") {
  console.log("Windows detected — skipping chmod.");
  process.exit(0);
}

try {
  fs.chmodSync(binaryPath, 0o755);
  console.log(`✓ Executable permission set on ${binaryName}`);
} catch (err) {
  console.warn(`⚠ Could not set permissions: ${err.message}`);
}
