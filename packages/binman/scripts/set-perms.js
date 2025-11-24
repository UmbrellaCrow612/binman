const os = require("os");
const path = require("path");
const { execSync } = require("child_process");
const fs = require("fs");

const rootBin = path.resolve(__dirname, "..", "bin");

function chmodRecursive(/** @type {string}*/ dir) {
  for (const entry of fs.readdirSync(dir)) {
    const full = path.join(dir, entry);
    const stats = fs.statSync(full);

    if (stats.isDirectory()) {
      chmodRecursive(full);
    } else {
      try {
        execSync(`chmod +x "${full}"`);
      } catch (/** @type {any}*/ e) {
        console.error(`[umbr-binman] Failed chmod for`, full, e.message);
      }
    }
  }
}

if (os.platform() !== "win32") {
  chmodRecursive(rootBin);
  console.log("[umbr-binman] Executable permissions applied recursively.");
}
