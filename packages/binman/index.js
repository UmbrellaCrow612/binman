#!/usr/bin/env node
const { spawn } = require("child_process");
const path = require("path");
const os = require("os");

const platformMap = {
  win32: "windows",
  darwin: "darwin",
  linux: "linux",
};

const archMap = {
  x64: "amd64",
  arm64: "arm64",
  ia32: "386",
  arm: "arm",
};

const nodePlatform = os.platform();
const nodeArch = os.arch();

const platform = platformMap[nodePlatform];
const arch = archMap[nodeArch];

if (!platform) {
  console.error("Unsupported platform:", nodePlatform);
  process.exit(1);
}

if (!arch) {
  console.error("Unsupported architecture:", nodeArch);
  process.exit(1);
}

const ext = platform === "windows" ? ".exe" : "";
const binaryName = `binman-${platform}-${arch}${ext}`;
const binaryPath = path.resolve(__dirname, "bin", platform, binaryName);

const args = process.argv.slice(2);
const child = spawn(binaryPath, args, { stdio: "inherit" });

child.on("exit", (code) => process.exit(code));
