/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,ts,jsx,tsx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        'custom-background': '#dfdfdfff',

        'custom-primary': '#E5007F',
        'custom-secondary': '#BB1F75',
        'custom-tertiary': '#66334F',

        'custom-extra-light': '#FFD7F2',
        'custom-extra-dark': '#332B2F',

        'custom-light': '#EDEEFF',
        'custom-gray': '#A6A7B3',
        'custom-dark': '#868791',
      }
    },
  },
  plugins: [],
}