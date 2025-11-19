const os = require('os');
const path = require('path');
const { execSync } = require('child_process');

const platform = os.platform();

if (platform !== 'win32') {
    const binPath = path.resolve(__dirname, '..', 'bin');
    
    try {
        execSync(`chmod +x ${binPath}/*`, { stdio: 'inherit' });
        console.log('[umbr-binman] Successfully set executable permissions.');
    } catch (e) {
        // If chmod fails for some reason, log the error but don't stop installation
        console.error(`[umbr-binman] WARNING: Failed to set executable permissions using chmod.`, e.message);
        // Do not process.exit(1) here as we don't want to fail the npm installation
    }
} else {
    console.log(`[umbr-binman] Running on Windows, skipping executable permission assignment.`);
}