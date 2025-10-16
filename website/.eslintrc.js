module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    'plugin:vue/recommended',
    'eslint:recommended',
    '@vue/typescript/recommended',
    'plugin:prettier/recommended', // Integrates Prettier
  ],
  parserOptions: {
    ecmaVersion: 2020,
  },
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    // Vue specific rules (simplified for Vue 2.7 compatibility)
    'vue/component-name-in-template-casing': 'off',
    'vue/custom-event-name-casing': 'off',
    'vue/no-unused-vars': 'warn',
    // TypeScript rules
    '@typescript-eslint/no-explicit-any': 'warn',
    '@typescript-eslint/no-unused-vars': 'warn',
    '@typescript-eslint/explicit-function-return-type': 'off',
    '@typescript-eslint/explicit-module-boundary-types': 'off',
    // General rules
    'prefer-const': 'warn',
    'no-var': 'error',
    'object-shorthand': 'warn',
  },
}
