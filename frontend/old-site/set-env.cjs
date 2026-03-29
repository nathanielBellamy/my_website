/**
 * set-env.js
 *
 * Reads environment variables from a root .env file and generates
 * the corresponding Vite .env file for the old-site Svelte build.
 *
 * Usage: node set-env.js <mode>
 *   mode - prod | remotedev
 *
 * Mode mapping:
 *   prod      -> reads ../../.env/.env.production
 *   remotedev -> reads ../../.env/.env.remotedev
 */

const fs = require('fs');
const path = require('path');

const mode = process.argv[2];

const envFileMap = {
  prod: '../../.env/.env.production',
  remotedev: '../../.env/.env.remotedev',
};

if (!mode || !envFileMap[mode]) {
  console.error(`Usage: node set-env.js <mode>`);
  console.error(`  Supported modes: ${Object.keys(envFileMap).join(', ')}`);
  process.exit(1);
}

const envFilePath = path.resolve(__dirname, envFileMap[mode]);

function parseEnvFile(filePath) {
  if (!fs.existsSync(filePath)) {
    console.error(`Error: ${filePath} not found.`);
    process.exit(1);
  }
  const content = fs.readFileSync(filePath, 'utf-8');
  const vars = {};
  content.split('\n').forEach((line) => {
    const trimmed = line.trim();
    if (!trimmed || trimmed.startsWith('#')) return;
    const eqIndex = trimmed.indexOf('=');
    if (eqIndex > 0) {
      const key = trimmed.substring(0, eqIndex).trim();
      const value = trimmed.substring(eqIndex + 1).trim().replace(/^['"]|['"]$/g, '');
      vars[key] = value;
    }
  });
  return vars;
}

const vars = parseEnvFile(envFilePath);

const baseUrlOldSite = vars['BASE_URL_OLD_SITE'] || '';
// Strip protocol (https:// or http://) to get just the host
const viteBaseUrl = baseUrlOldSite.replace(/^https?:\/\//, '');

const isProduction = mode === 'prod';
const viteMode = isProduction ? 'production' : mode;
const recaptchaSiteKey = vars['RECAPTCHA_SITE_KEY'] || '';

const outputContent = [
  `VITE_BASE_URL=${viteBaseUrl}`,
  `VITE_ENV=${viteMode}`,
  `VITE_IS_PROD=${isProduction}`,
  `VITE_MODE=${viteMode}`,
  `VITE_RECAPTCHA_SITE_KEY=${recaptchaSiteKey}`,
  '',
].join('\n');

const outputPath = path.resolve(__dirname, `.env.${mode}`);
fs.writeFileSync(outputPath, outputContent);
