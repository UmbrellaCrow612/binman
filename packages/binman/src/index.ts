import fs = require("fs/promises");
import path = require("path");

/**
 * Resolves at runtime a `binman.yml` `exe` binary path based on the download convention of `binman`.
 * Using Node's `process.arch` and `process.platform`, the final downloaded content follows this convention:
 * 
 * ```
 * basePath/packageName/platform/arch/...files...
 * ```
 * 
 * It then tries to get the required `exe` based on Node's `process.arch` and `process.platform`.
 * 
 * Example config:
 * 
 * ```yaml
 * - name: ripgrep
 *   platforms:
 *     linux:
 *       x64:
 *         url: https://github.com/BurntSushi/ripgrep/releases/download/15.1.0/ripgrep-15.1.0-x86_64-unknown-linux-musl.tar.gz
 *         sha256: "1c9297be4a084eea7ecaedf93eb03d058d6faae29bbc57ecdaf5063921491599"
 *         pattern: "^rg$"
 * ```
 * 
 * @param packageName - The name of the `binman.yml` package, e.g., "ripgrep"
 * @param exeNames - List of possible executable names, e.g., ["rg", "rpx"]. Do not include .exe extension
 * @param basePath - The base path where `binman` downloaded the binaries
 * @returns The full path to the executable, or undefined if not found
 */
const binmanResolve = async function (
  packageName: string,
  exeNames: string[],
  basePath: string
): Promise<string | undefined> {
  const platform = process.platform;
  const arch = process.arch;
  
  const binDir = path.join(basePath, packageName, platform, arch);
  
  try {
    await fs.access(binDir);
    
    const files = await fs.readdir(binDir, { withFileTypes: true });
    
    for (const exeName of exeNames) {
      for (const file of files) {
        if (!file.isFile()) continue;
        
        const fileName = file.name;
        
        if (platform === 'win32') {
          if (fileName === `${exeName}.exe`) {
            return path.join(binDir, fileName);
          }
        } else {
          if (fileName === exeName) {
            const filePath = path.join(binDir, fileName);
            
            try {
              await fs.access(filePath, fs.constants.X_OK);
              return filePath;
            } catch {
              continue;
            }
          }
        }
      }
    }
    
    return undefined;
  } catch (error) {
    return undefined;
  }
};

export = binmanResolve;