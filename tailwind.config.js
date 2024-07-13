/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['internal/**/*.templ'],
  darkMode: 'class',
  theme: {
    extend: {
      fontFamily: {
        mono: ['GeistMono', 'monospace'],
      }
    },
  },
  plugins: [],
}

