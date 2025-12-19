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

        'custom-light': '#FFFFFF',
        'custom-light-gray': '#F4F4F4',
        'custom-gray': '#5E5E5E',
        'custom-dark': '#252525',
      },
      fontFamily: {
        "outfit": ['Outfit', 'sans-serif'],
        "fugaz": ['Fugaz One', 'cursive'],
        "sans": ["Outfit", "ui-sans-serif", "system-ui"],
      }
    },
  },
  plugins: [],
}