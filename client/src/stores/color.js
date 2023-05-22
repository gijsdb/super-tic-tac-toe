import { defineStore } from 'pinia'

export const useColorStore = defineStore('color', {
    state: () => {
        return {
            Show: false,
            ActiveTheme: {
                name: 'default',
                bg_color: '#313131',
                main_color: '#f66e0d',
                caret_color: '#f66e0d',
                sub_color: '#616161',
                sub_alt_color: '#2b2b2b',
                text_color: '#f5e6c8',
                error_color: '#e72d2d',
                error_extra_color: '#7e2a33',
                colorful_error_color: '#e72d2d',
                colorful_error_extra_color: '#7e2a33',
            },
            Themes: [
                {
                    name: 'blueberry',
                    bg_color: '#212b42',
                    main_color: '#add7ff',
                    caret_color: '#962f7e',
                    sub_color: '#5c7da5',
                    sub_alt_color: '#1b2334',
                    text_color: '#91b4d5',
                    error_color: '#df4576',
                    error_extra_color: '#d996ac',
                    colorful_error_color: '#df4576',
                    colorful_error_extra_color: '#d996ac',
                },
                {
                    name: 'default',
                    bg_color: '#313131',
                    main_color: '#f66e0d',
                    caret_color: '#f66e0d',
                    sub_color: '#616161',
                    sub_alt_color: '#2b2b2b',
                    text_color: '#f5e6c8',
                    error_color: '#e72d2d',
                    error_extra_color: '#7e2a33',
                    colorful_error_color: '#e72d2d',
                    colorful_error_extra_color: '#7e2a33',
                },
                {
                    name: 'earthly',
                    bg_color: '#292521',
                    main_color: '#509452',
                    caret_color: '#1298ba',
                    sub_color: '#f5ae2d',
                    sub_alt_color: '#1d1b18',
                    text_color: '#e6c7a8',
                    error_color: '#7e2a33',
                    error_extra_color: '#ff645a',
                    colorful_error_color: '#7e2a33',
                    colorful_error_extra_color: '#ff645a',
                },
                {
                    name: 'luna',
                    bg_color: '#221c35',
                    main_color: '#f67599',
                    caret_color: '#f67599',
                    sub_color: '#5a3a7e',
                    sub_alt_color: '#2f2346',
                    text_color: '#ffe3eb',
                    error_color: '#efc050',
                    error_extra_color: '#c5972c',
                    colorful_error_color: '#efc050',
                    colorful_error_extra_color: '#c5972c',
                },
                {
                    name: 'roboto',
                    bg_color: '#0c0d11',
                    main_color: '#7ebab5',
                    caret_color: '#7ebab5',
                    sub_color: '#454864',
                    sub_alt_color: '#171a25',
                    text_color: '#f6f5f5',
                    error_color: '#ff4754',
                    error_extra_color: '#b02a33',
                    colorful_error_color: '#ff4754',
                    colorful_error_extra_color: '#b02a33',
                }
            ],
        }
    },
    actions: {
        toggleThemePicker() {
            this.Show = !this.Show
        },
        changeTheme(theme) {
            this.Themes.forEach((t) => {
                if (t.name === theme) {
                    this.ActiveTheme = t
                    document.documentElement.style.setProperty('--bg-color', t.bg_color);
                    document.documentElement.style.setProperty('--main-color', t.main_color);
                    document.documentElement.style.setProperty('--caret-color', t.caret_color);
                    document.documentElement.style.setProperty('--sub-color', t.sub_color);
                    document.documentElement.style.setProperty('--sub-alt-color', t.sub_alt_color);
                    document.documentElement.style.setProperty('--text-color', t.text_color);
                    document.documentElement.style.setProperty('--error-color', t.error_color);
                    document.documentElement.style.setProperty('--error-extra-color', t.error_extra_color);
                    document.documentElement.style.setProperty('--colorful-error-color', t.colorful_error_extra_color);
                    document.documentElement.style.setProperty('--colorful-error-extra-color', t.colorful_error_extra_color);
                }
            })
        },
    }
})

