module.exports = {
  extends: '../../../.eslintrc.js',
  overrides: [
    {
      files: ['src/**/*.ts'],
      rules: {
        'no-restricted-imports': [
          'error',
          'assert',
          'buffer',
          'child_process',
          'cluster',
          'crypto',
          'dgram',
          'dns',
          'domain',
          'events',
          'freelist',
          'fs',
          'http',
          'https',
          'module',
          'net',
          'os',
          'path',
          'punycode',
          'querystring',
          'readline',
          'repl',
          'smalloc',
          'stream',
          'string_decoder',
          'sys',
          'timers',
          'tls',
          'tracing',
          'tty',
          'url',
          'util',
          'vm',
          'zlib',
        ],
      },
    },
  ],
}
