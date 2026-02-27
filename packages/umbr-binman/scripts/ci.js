import path from "path";
import { fileURLToPath } from "url";
import { spawn } from "child_process";
import fs from "fs";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// Detect current platform and arch
const platform = process.platform;
const arch = process.arch;

let binName;
if (platform === "win32") {
  binName = `binman-windows-${arch}.exe`;
} else if (platform === "darwin") {
  binName = `binman-darwin-${arch}`;
} else {
  binName = `binman-linux-${arch}`;
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