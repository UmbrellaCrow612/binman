const path = require("path");
const fs = require("fs").promises;
const fssync = require("fs");

/**
 * Automatically resolve an external binary's executable path using
 * Binman's download directory convention.
 *
 * Binman directory structure:
 *   basePath/packageName/os/architecture/<files...>
 *
 * Example:
 *   bin/
 *    └─ ripgrep/
 *        └─ linux/
 *            └─ x86_64/
 *                └─ rg   <-- executable
 *
 * This function determines the OS and architecture of the current machine,
 * navigates into the correct folder, looks for the executable, adds file
 * extension on Windows, and returns the resolved path.
 *
 * @param {string} packageName - The name of the binary package. This must match
 *                               the folder name Binman created.
 *
 * @param {string} exeName - The executable name WITHOUT extension.
 *                           e.g. "rg" for ripgrep.
 *
 * @param {string} basePath - Path to the folder where Binman places downloaded binaries.
 *                            Must follow binman's OS/arch directory structure.
 *
 * @returns {Promise<string | undefined>} The full path to the executable,
 *                                        or undefined if it cannot be resolved.
 */
async function binmanResolve(packageName, exeName, basePath) {
  if (!packageName || packageName.trim() === "") {
    throw new Error("Package name cannot be empty");
  }

  if (!exeName || exeName.trim() === "") {
    throw new Error("Exe name cannot be empty");
  }

  if (!basePath || !fssync.existsSync(basePath)) {
    return undefined;
  }

  let os = "";
  switch (process.platform) {
    case "win32":
      os = "windows";
      break;
    default:
      os = process.platform;
      break;
  }

  let arch = "";
  switch (process.arch) {
    case "x64":
      arch = "x86_64";
      break;
    case "arm64":
      arch = "aarch64";
      break;
    case "ia32":
      arch = "i386";
      break;
    default:
      arch = process.arch;
      break;
  }

  const exeFileName = os === "windows" ? `${exeName}.exe` : exeName;

  const targetDir = path.join(basePath, packageName, os, arch);

  if (!fssync.existsSync(targetDir)) {
    return undefined;
  }

  const files = await fs.readdir(targetDir);

  const match = files.find((f) => f === exeFileName);
  if (!match) return undefined;

  return path.join(targetDir, match);
}

module.exports = { binmanResolve };
