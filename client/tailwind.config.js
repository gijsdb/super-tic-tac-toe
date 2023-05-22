module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        bg_color: 'var(--bg-color)',
        main_color: 'var(--main-color )',
        caret_color: 'var(--caret-color )',
        sub_color: 'var(--sub-color )',
        sub_alt_color: 'var(--sub-alt-color )',
        text_color: 'var(--text-color )',
        error_color: 'var(--error-color )',
        error_extra_color: 'var(--error-extra-color )',
        colorful_error_color: 'var(--colorful-error-color )',
        colorful_error_extra_color: 'var(--colorful-error-extra-color )',
      },
    },
  },
  plugins: [],
}