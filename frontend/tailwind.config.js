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
      },

      keyframes: {
        transformRotate: {
          '0%': { transform: 'rotate(0deg)' },
          '100%': { transform: 'rotate(360deg)' },
        },

        slideInRight: {
          '0%': { 
            transform: 'translateX(100%)',
            opacity: '0'
          },
          '100%': { 
            transform: 'translateX(0)',
            opacity: '1'
          }
        },

        slideOutRight: {
          '0%': { 
            transform: 'translateX(0)',
            opacity: '1'
          },
          '100%': { 
            transform: 'translateX(100%)',
            opacity: '0'
          }
        },

        slideInBottom: {
          '0%': { 
            transform: 'translateY(200%)',
            opacity: '0'
          },
          '100%': { 
            transform: 'translateY(0%)',
            opacity: '1'
          }
        },
      },

      animation: {
        'rotate': 'transformRotate 20s linear infinite',

        'slide-in-right': 'slideInRight 0.8s ease-out',
        'slide-out-right': 'slideOutRight 0.8s ease-out',

        'slide-in-bottom': 'slideInBottom 0.8s ease-out',
      }
    },

  },
  plugins: [],
}