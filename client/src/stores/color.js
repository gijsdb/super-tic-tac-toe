import { defineStore } from 'pinia'

export const useColorStore = defineStore('color', {
    state: () => {
        return {
            Show: false,
            ActiveTheme: {
                Name: "Default",
                Primary: "#292521",
                Secondary: "#e6c7a8",
                Highlight: "#509452",
                HighlightTwo: "#f5ae2d"
            },
            Themes: [
                {
                    Name: "Default",
                    Primary: "#292521",
                    Secondary: "#e6c7a8",
                    Highlight: "#509452",
                    HighlightTwo: "#f5ae2d"
                },
                {
                    Name: "8008",
                    Primary: "#333A45",
                    Secondary: "#8E98A8",
                    Highlight: "#E9ECF0",
                    HighlightTwo: "#F44C7F",
                },
                {
                    Name: "Carbon",
                    Primary: "#313131",
                    Secondary: "#5D5D5D",
                    Highlight: "#F5E6C8",
                    HighlightTwo: "#F66E0D",
                },
                {
                    Name: "Ever Brush",
                    Primary: "#141B1E",
                    Secondary: "#838887",
                    Highlight: "#DADADA",
                    HighlightTwo: "#8CCF7E",
                }],
        }
    },
    actions: {
        toggleThemePicker() {
            this.Show = !this.Show
        },
        changeTheme(theme) {
            this.Themes.forEach((t) => {
                if (t.Name === theme) {
                    this.ActiveTheme = t
                }
            })
        }
    }
})

