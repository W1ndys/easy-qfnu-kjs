/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: '#884F22',
          light: '#A67C52',
          lighter: '#C4A882',
          dark: '#5D3615',
          50: '#FDF8F3',
          100: '#F5E6D3',
          200: '#E8CDB0',
          300: '#D4AD82',
          400: '#C4956A',
          500: '#884F22',
          600: '#76451C',
          700: '#5D3615',
          800: '#442710',
          900: '#2C190A',
        },
        clay: {
          canvas: '#FAF5EF',
          cardBg: 'rgba(255, 255, 255, 0.70)',
          foreground: '#332A24',
          muted: '#6B5D52',
          accent: '#884F22',
          'accent-alt': '#C4956A',
          success: '#10B981',
          warning: '#F59E0B',
          info: '#0EA5E9',
        },
      },
      borderRadius: {
        clay: '20px',
        'clay-md': '24px',
        'clay-lg': '32px',
        'clay-xl': '48px',
      },
      boxShadow: {
        clayCard: [
          '16px 16px 32px rgba(136, 79, 34, 0.08)',
          '-10px -10px 24px rgba(255, 255, 255, 0.9)',
          'inset 6px 6px 12px rgba(136, 79, 34, 0.03)',
          'inset -6px -6px 12px rgba(255, 255, 255, 1)',
        ].join(', '),
        clayCardHover: [
          '20px 20px 40px rgba(136, 79, 34, 0.12)',
          '-12px -12px 28px rgba(255, 255, 255, 0.95)',
          'inset 6px 6px 12px rgba(136, 79, 34, 0.03)',
          'inset -6px -6px 12px rgba(255, 255, 255, 1)',
        ].join(', '),
        clayButton: [
          '12px 12px 24px rgba(136, 79, 34, 0.25)',
          '-8px -8px 16px rgba(255, 255, 255, 0.4)',
          'inset 4px 4px 8px rgba(255, 255, 255, 0.4)',
          'inset -4px -4px 8px rgba(0, 0, 0, 0.08)',
        ].join(', '),
        clayButtonHover: [
          '14px 14px 28px rgba(136, 79, 34, 0.3)',
          '-10px -10px 20px rgba(255, 255, 255, 0.5)',
          'inset 4px 4px 8px rgba(255, 255, 255, 0.4)',
          'inset -4px -4px 8px rgba(0, 0, 0, 0.1)',
        ].join(', '),
        clayPressed: [
          'inset 10px 10px 20px rgba(136, 79, 34, 0.08)',
          'inset -10px -10px 20px #ffffff',
        ].join(', '),
        claySurface: [
          '30px 30px 60px rgba(136, 79, 34, 0.06)',
          '-30px -30px 60px #ffffff',
          'inset 10px 10px 20px rgba(136, 79, 34, 0.03)',
          'inset -10px -10px 20px rgba(255, 255, 255, 0.8)',
        ].join(', '),
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
