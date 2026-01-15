#!/usr/bin/env node

/**
 * postinstall script
 * Sets executable permissions for all binaries in ./bin recursively
 * Only runs chmod on non-Windows platforms
 */

const fs = require('fs');
const path = require('path');

const binDir = path.resolve(__dirname, '..', 'bin');

/**
 * Recursively walk a directory and apply a callback to each file
 */
function walkDir(dir, callback) {
  const entries = fs.readdirSync(dir);
  for (const entry of entries) {
    const fullPath = path.join(dir, entry);
    const stats = fs.statSync(fullPath);

    if (stats.isDirectory()) {
      walkDir(fullPath, callback);
    } else if (stats.isFile()) {
      callback(fullPath);
    }
  }
}

// Only run chmod on non-Windows systems
if (process.platform === 'win32') {
  console.log('Windows detected — skipping chmod.');
  process.exit(0);
}

try {
  if (!fs.existsSync(binDir)) {
    console.warn(`\x1b[33mWarning: bin directory not found: ${binDir}\x1b[0m`);
  } else {
    walkDir(binDir, (filePath) => {
      fs.chmodSync(filePath, 0o755);
      console.log(`\x1b[32mSet executable: ${filePath}\x1b[0m`);
    });
  }
} catch (err) {
  console.error(`\x1b[31mError setting executable permissions: ${err.message}\x1b[0m`);
  process.exit(1);
}
