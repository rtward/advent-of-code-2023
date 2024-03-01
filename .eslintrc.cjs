module.exports = {
  env: {
    browser: true,
    node: true,
    es2017: true,
  },
  extends: [
    'eslint:recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:prettier/recommended',
    'prettier',
  ],
  overrides: [
    {
      files: ['**/*.tsx'],
      rules: {
        '@typescript-eslint/naming-convention': [
          'error',
          {
            selector: 'variable',
            format: ['strictCamelCase', 'PascalCase', 'UPPER_CASE'],
            filter: {
              regex: 'XY',
              match: false,
            },
            leadingUnderscore: 'allow',
            trailingUnderscore: 'allow',
          },
        ],
      },
    },
  ],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    project: './tsconfig.json',
    sourceType: 'module',
  },
  ignorePatterns: ['.eslintrc.cjs'],
  plugins: ['@typescript-eslint', 'prettier'],
  root: true,
  rules: {
    eqeqeq: ['error', 'always'],
    '@typescript-eslint/array-type': ['error', { default: 'array-simple' }],
    '@typescript-eslint/explicit-module-boundary-types': 'off',
    '@typescript-eslint/member-delimiter-style': [
      'error',
      {
        multiline: {
          delimiter: 'none',
          requireLast: true,
        },
        singleline: {
          delimiter: 'semi',
          requireLast: false,
        },
      },
    ],
    '@typescript-eslint/no-unused-vars': [
      'error',
      { argsIgnorePattern: '^_', ignoreRestSiblings: true },
    ],
    '@typescript-eslint/prefer-nullish-coalescing': 'error',
    '@typescript-eslint/prefer-optional-chain': 'error',

    'no-fallthrough': 'off',
  },
}
