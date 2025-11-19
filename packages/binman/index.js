#!/usr/bin/env node
const { spawn } = require("child_process");
const path = require("path");
const os = require("os");

let binaryName;
switch (os.platform()) {
  case "win32":
    binaryName = "binman.exe";
    break;
  case "darwin":
    binaryName = "binman-darwin";
    break;
  case "linux":
    binaryName = "binman-linux";
    break;
  default:
    console.error("Unsupported platform:", os.platform());
    process.exit(1);
}

const binaryPath = path.resolve(__dirname, "bin", binaryName);
const args = process.argv.slice(2);

const child = spawn(binaryPath, args, { stdio: "inherit" });

child.on("exit", (code) => process.exit(code));
