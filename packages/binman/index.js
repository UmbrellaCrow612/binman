const path = require("path");
const fs = require("fs").promises;
const fssync = require("fs");

/**
 * Automatically resolve an external binary's executable path using
 * Binman's download directory convention, trying multiple possible executable names.
 *
 * @param {string} packageName - The name of the binary package. Must match the folder name.
 * @param {string[]} exeNames - Array of possible executable names WITHOUT extension.
 *                              e.g. ["7zz_1", "7z"]
 * @param {string} basePath - Path to the folder where Binman places downloaded binaries.
 * @returns {Promise<string | undefined>} The full path to the first matching executable,
 *                                        or undefined if none are found.
 */
async function binmanResolve(packageName, exeNames, basePath) {
  if (!packageName || packageName.trim() === "") {
    throw new Error("Package name cannot be empty");
  }

  if (!Array.isArray(exeNames) || exeNames.length === 0) {
    throw new Error("exeNames must be a non-empty array of strings");
  }

  if (!basePath || !fssync.existsSync(basePath)) {
    return undefined;
  }

  // Determine OS (normalize Linux/Darwin/Windows)
  let os;
  switch (process.platform) {
    case "win32":
      os = "windows";
      break;
    case "darwin":
      os = "darwin";
      break;
    case "linux":
      os = "linux";
      break;
    default:
      os = process.platform;
      break;
  }

  let arch;
  switch (process.arch.toString()) {
    case "x64":
    case "amd64":
      arch = "x86_64";
      break;

    case "arm64": 
    case "aarch64": 
      arch = "arm64";
      break;

    case "ia32":
    case "x32":
    case "arm":
    case "armv7l":
    case "armv6l":
    case "ppc64":
    case "ppc64le":
    case "s390x":
    case "mips":
    case "mips64":
    case "riscv64":
    default:
      arch = process.arch;
      break;
  }

  const targetDir = path.join(basePath, packageName, os, arch);
  if (!fssync.existsSync(targetDir)) return undefined;

  const files = await fs.readdir(targetDir);

  for (const name of exeNames) {
    const exeFileName = os === "windows" ? `${name}.exe` : name;
    const match = files.find((f) => f === exeFileName);
    if (match) {
      return path.join(targetDir, match);
    }
  }

  return undefined;
}

module.exports = { binmanResolve };
