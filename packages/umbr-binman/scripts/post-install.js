import fs from "fs";
import path from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// Determine current platform and architecture
const platform = process.platform;
const arch = process.arch;

let binName;
if (platform === "win32") {
  binName = `binman-${platform}-${arch}.exe`;
} else {
  binName = `binman-${platform}-${arch}`;
}

const binDir = path.resolve(__dirname, "../bin");
const binPath = path.join(binDir, binName);

// Delete all other binaries
const files = fs.readdirSync(binDir);
for (const file of files) {
  if (file !== binName) {
    fs.unlinkSync(path.join(binDir, file));
    console.log(`Deleted: ${file}`);
  }
}

// Make current binary executable (Linux/macOS)
if (platform !== "win32") {
  if (!fs.existsSync(binPath)) {
    console.error(`Binary does not exit at ${binPath} for binman`);
    throw new Error(`Binary does not exit at ${binPath} for binman`)
  }
  fs.chmodSync(binPath, 0o755);
}

console.log(`Setup complete: ${binPath} is ready to run`);