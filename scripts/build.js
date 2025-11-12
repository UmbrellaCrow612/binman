#!/usr/bin/env node
const { spawnSync } = require('child_process');
const fs = require('fs');
const path = require('path');

// Path to your CLI Go project
const goProjectDir = path.resolve(__dirname, '../cli');

// Output directory
const binDir = path.resolve(__dirname, '../packages/binman/bin');
if (!fs.existsSync(binDir)) fs.mkdirSync(binDir, { recursive: true });

const platforms = {
    linux: 'binman-linux',
    darwin: 'binman-darwin',
    windows: 'binman.exe'
};

console.log('Starting build for all platforms...\n');

for (const [platform, outputName] of Object.entries(platforms)) {
    const outputPath = path.join(binDir, outputName);
    console.log(`Building for ${platform} -> ${outputPath}`);

    // Environment variables
    const env = { ...process.env, GOOS: platform, GOARCH: 'amd64' };

    // Run 'go build' inside the CLI project folder
    const result = spawnSync('go', ['build', '-o', outputPath], {
        stdio: 'inherit',
        cwd: goProjectDir,
        env
    });

    if (result.status !== 0) {
        console.error(`Failed to build for ${platform}`);
        process.exit(1);
    }
}

console.log('\nAll binaries built successfully in ./bin');
