import { runCommand, safeRun } from "node-github-actions";
import { Logger } from "node-logy";
import path, { dirname } from "node:path";
import fs from "fs/promises";
import { fileURLToPath } from "node:url";

const logger = new Logger({ showCallSite: true, saveToLogFiles: false });

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const platformMap = {
  linux: "linux",
  darwin: "darwin",
  windows: "windows",
};

const archMap = {
  arm64: "arm64",
  amd64: "amd64",
};

const platforms = Object.keys(platformMap);
const archs = Object.keys(archMap);

safeRun(
  async () => {
    const outputDir = path.join(__dirname, "./packages/umbr-binman/bin");
    logger.info("Output dir: ", outputDir);
    await fs.mkdir(outputDir, { recursive: true });

    for (const platform of platforms) {
      for (const arch of archs) {
        const exeName =
          platform === "windows"
            ? `binman-${platform}-${arch}.exe`
            : `binman-${platform}-${arch}`;
        const outputPath = path.join(outputDir, exeName);

        logger.info(`Building for ${platform}/${arch} -> ${exeName}`);

        // Build command
        await runCommand(
          "go",
          ["build", "-o", outputPath],
          {
            env: {
              ...process.env,
              GOOS: platformMap[platform],
              GOARCH: archMap[arch],
              CGO_ENABLED: "0",
            },
          },
          120,
        );

        logger.info(`Built ${exeName}`);
      }
    }
  },
  {
    exitFailCode: 1,
    exitOnFailed: true,
    timeoutMs: 10 * 60 * 1000,
    onFail: (err) => {
      logger.error("Failed tro build ", err);
    },
    onAfter: () => {
      logger.info("Finished build");
    },
  },
);
