/**
 * set-env.js
 *
 * Reads environment variables from a .env file and generates
 * the corresponding TypeScript environment file for the Angular build.
 *
 * Usage: node set-env.js <mode> <envFilePath> <outputFilePath>
 *   mode        - localhost | remotedev | production
 *   envFilePath - path to the .env file (resolved relative to cwd)
 *   outputFilePath - path to write the generated TypeScript file
 */

const fs = require('fs');
const path = require('path');

const mode = process.argv[2];
const envFilePath = process.argv[3];
const outputFilePath = process.argv[4];

if (!mode || !envFilePath || !outputFilePath) {
  console.error('Usage: node set-env.js <mode> <envFilePath> <outputFilePath>');
  process.exit(1);
}

const resolvedEnvFile = path.resolve(process.cwd(), envFilePath);

function parseEnvFile(filePath) {
  if (!fs.existsSync(filePath)) {
    console.warn(`Warning: ${filePath} not found. Falling back to empty values.`);
    return {};
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

const vars = parseEnvFile(resolvedEnvFile);
const isProduction = mode === 'production';
const baseUrl = vars["BASE_URL"];
const baseUrlApi = vars["BASE_URL_API"];
const baseUrlOldSite = vars["BASE_URL_OLD_SITE"];


function escapeForTs(value) {
  return value
    .replace(/\\/g, '\\\\')
    .replace(/'/g, "\\'");
}

const outputContent = `export const environment = {
  production: ${isProduction},
  BASE_URL: '${baseUrl}',
  BASE_URL_API: '${baseUrlApi}',
  BASE_URL_OLD_SITE: '${baseUrlOldSite}'
};
`;

const resolvedOutput = path.resolve(process.cwd(), outputFilePath);
fs.writeFileSync(resolvedOutput, outputContent);
console.log(`Generated ${resolvedOutput} from ${resolvedEnvFile}`);
