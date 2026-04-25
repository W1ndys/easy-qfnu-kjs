/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: '#884F22',
          light: '#B87948',
          lighter: '#E7CFBA',
          dark: '#5F3517',
          50: '#FAF8F6',
          100: '#F3E5D8',
          200: '#E7CFBA',
          300: '#D1C7BE',
          400: '#B87948',
          500: '#884F22',
          600: '#76441E',
          700: '#5F3517',
          800: '#332D29',
          900: '#1F1B18',
        },
        clay: {
          canvas: '#F8F5F2',
          cardBg: '#FFFFFF',
          foreground: '#1F1B18',
          muted: '#4C433D',
          accent: '#884F22',
          'accent-alt': '#B87948',
          success: '#156B52',
          warning: '#9A5A00',
          info: '#1D4ED8',
        },
        page: '#F8F5F2',
        card: '#FFFFFF',
        subtle: '#E5DED7',
        secondary: '#4C433D',
      },
      borderRadius: {
        clay: '10px',
        'clay-md': '12px',
        'clay-lg': '16px',
        'clay-xl': '20px',
      },
      boxShadow: {
        clayCard: '0 1px 2px rgba(31, 27, 24, 0.06), 0 4px 12px rgba(31, 27, 24, 0.04)',
        clayCardHover: '0 6px 20px rgba(31, 27, 24, 0.08), 0 2px 6px rgba(31, 27, 24, 0.05)',
        clayButton: '0 6px 16px rgba(136, 79, 34, 0.18)',
        clayButtonHover: '0 8px 20px rgba(136, 79, 34, 0.22)',
        clayPressed: '0 0 0 3px rgba(184, 121, 72, 0.18)',
        claySurface: '0 10px 30px rgba(31, 27, 24, 0.08)',
      },
      animation: {
        'clay-float': 'clay-float 8s ease-in-out infinite',
        'clay-float-delayed': 'clay-float-delayed 10s ease-in-out infinite',
        'clay-float-slow': 'clay-float-slow 12s ease-in-out infinite',
        'clay-breathe': 'clay-breathe 6s ease-in-out infinite',
      },
      keyframes: {
        'clay-float': {
          '0%, 100%': { transform: 'translateY(0) rotate(0deg)' },
          '50%': { transform: 'translateY(-20px) rotate(2deg)' },
        },
        'clay-float-delayed': {
          '0%, 100%': { transform: 'translateY(0) rotate(0deg)' },
          '50%': { transform: 'translateY(-15px) rotate(-2deg)' },
        },
        'clay-float-slow': {
          '0%, 100%': { transform: 'translateY(0) rotate(0deg)' },
          '50%': { transform: 'translateY(-30px) rotate(5deg)' },
        },
        'clay-breathe': {
          '0%, 100%': { transform: 'scale(1)' },
          '50%': { transform: 'scale(1.02)' },
        },
      },
    },
  },
  plugins: [],
}
