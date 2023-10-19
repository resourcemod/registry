import defaultTheme from 'tailwindcss/defaultTheme';
import forms from '@tailwindcss/forms';
import typography from '@tailwindcss/typography';

/** @type {import('tailwindcss').Config} */
export default {
    content: [
        './src/*.html',
        './src/**/*.vue',
    ],

    theme: {
        extend: {
            colors: {
                'brand-gray': '#F8F9FB',
                'dark-gray': '#E4E7EB',
                'font-gray': '#121825',
                'light-gray': '#9DA3AE',
            },
            fontFamily: {
                sans: ['Inter var', ...defaultTheme.fontFamily.sans],
            },
        },
    },

    plugins: [forms, typography],
};
