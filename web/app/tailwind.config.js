import defaultTheme from 'tailwindcss/defaultTheme';
import forms from '@tailwindcss/forms';
import typography from '@tailwindcss/typography';

/** @type {import('tailwindcss').Config} */
export default {
    content: [
        './vendor/laravel/framework/src/Illuminate/Pagination/resources/views/*.blade.php',
        './vendor/laravel/jetstream/**/*.blade.php',
        './storage/framework/views/*.php',
        './resources/views/**/*.blade.php',
        './resources/js/**/*.vue',
    ],

    theme: {
        extend: {
            colors: {
                'brand-gray': '#F8F9FB',
                'dark-gray': '#E4E7EB',
                'font-gray': '#121825',
                'light-gray': '#9DA3AE',
            },
            screens: {
                '2xl': {'max': '1535px'},
                // => @media (max-width: 1535px) { ... }

                'xl': {'max': '1279px'},
                // => @media (max-width: 1279px) { ... }

                'lg': {'max': '860px'},
                // => @media (max-width: 1023px) { ... }

                'md': {'max': '767px'},
                // => @media (max-width: 767px) { ... }

                'sm': {'max': '639px'},
                // => @media (max-width: 639px) { ... }
            },
            fontFamily: {
                sans: ['Figtree', ...defaultTheme.fontFamily.sans],
            },
        },
    },

    plugins: [forms, typography],
};
