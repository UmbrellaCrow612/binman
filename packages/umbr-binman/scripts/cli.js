import path from "path";
import { fileURLToPath } from "url";
import { spawn } from "child_process";
import fs from "fs";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const arch = process.arch;

let binName;
if (platform === "win32") {
  binName = `binman-${process.platform}-${arch}.exe`;
} else {
  binName = `binman-${process.platform}-${arch}`;
}

const binDir = path.resolve(__dirname, "../bin");
const binPath = path.join(binDir, binName);

// Check that the binary exists
if (!fs.existsSync(binPath)) {
  console.error(`Error: binary not found for ${platform}-${arch}`);
  process.exit(1);
}

// Get all args passed to npx binman
const args = process.argv.slice(2);

// Spawn the binary with the same arguments
const child = spawn(binPath, args, { stdio: "inherit" });

child.on("close", (code) => {
  process.exit(code);
});

child.on("error", (err) => {
  console.error("Spawn failed", err.message, err.cause, err.name, err.stack);
});