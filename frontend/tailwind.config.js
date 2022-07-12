module.exports = {
  content: ['./src/**/*.{js,jsx,ts,tsx}'],
  theme: {
    extend: {
      spacing: {
        18: '4.5rem',
        50: '12.8rem',
        68: '16.5rem',
      },
      colors: {
        shipaDarkBlue: '#103c4c',
        shipaOrange: '#ff5444',
        shipaLime: '#e8fc7c',
        shipaGreen: '#a8dc8c',
        shipaGrey: '#504c4c',
      },
    },
  },
  plugins: [require('tailwind-scrollbar'), require('@tailwindcss/line-clamp')],
}
