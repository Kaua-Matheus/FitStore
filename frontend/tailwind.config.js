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
        'custom-primary': '#4C92FC',
        'custom-secondary': '#4C5BFC',
        'custom-tertiary': '#764CFC',

        'custom-extra-light': '#A3C8FF',
        'custom-extra-dark': '#7E9AC4',

        'custom-text': '#1f2937',
      }
    },
  },
  plugins: [],
}