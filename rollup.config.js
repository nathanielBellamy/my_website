import { wasm } from '@rollup/plugin-wasm';

export default {
  input: 'src/main.ts',
  output: {
    dir: 'output',
    format: 'cjs'
  },
  plugins: [wasm()]
};
