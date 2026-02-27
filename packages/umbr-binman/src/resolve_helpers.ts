import fs from "node:fs/promises";
import path from "node:path";

/**
 * Resolve at runtime the path to the exe that binman downloaded in the bin dir for the given environment
 * @param binPath The base path to where the bin folder is located
 * @param packageName The name of the package to resolve
 * @param exeNames A list of possible executable names (without .exe extension)
 * @returns Path to the exe if found, otherwise undefined
 */
export async function binmanResolve(
  binPath: string,
  packageName: string,
  exeNames: string[],
): Promise<string | undefined> {
  await fs.access(binPath);

  const baseDir = path.join(
    binPath,
    packageName,
    process.platform,
    process.arch,
  );

  for (const name of exeNames) {
    const exeName = process.platform === "win32" ? `${name}.exe` : name;
    const finalExePath = path.join(baseDir, exeName);

    try {
      await fs.access(finalExePath);
      return finalExePath; 
    } catch {
    }
  }

  return undefined; 
}