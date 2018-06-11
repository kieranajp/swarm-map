import resolve from 'rollup-plugin-node-resolve';
import babel from 'rollup-plugin-babel';

export default {
  input: 'client/assets/js/main.js',
  output: {
    file: 'client/assets/dist/bundle.js',
    format: 'cjs'
  },
  watch: {
    include: 'client/**'
  },
  plugins: [
    resolve(),
    babel({
      exclude: 'node_modules/**'
    })
  ]
};
