const { fontFamily } = require('tailwindcss/defaultTheme');

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{html,ts}",
  ],
  theme: {
    extend: {
      colors: {
        'background': '#0B1120',
        'primary-accent': '#22D3EE',
        'secondary-accent': '#D946EF',
        'body-text': '#CBD5E1',
        'heading-text': '#FFFFFF',
        'cmyk-c': '#00FFFF',
        'cmyk-m': '#FF00FF',
        'cmyk-y': '#FFFF00',
        'card': '#1E293B',
      },
      fontFamily: {
        'sans': ['"Work Sans"', ...fontFamily.sans],
        'heading': ['"Outfit"', ...fontFamily.sans],
      },
      animation: {
        'marquee': 'marquee 30s linear infinite',
      },
      keyframes: {
        marquee: {
          '0%': { transform: 'translateX(0%)' },
          '100%': { transform: 'translateX(-100%)' },
        }
      }
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
}
