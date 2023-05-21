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
                    Name: "Theme one",
                    Primary: "#18181A",
                    Secondary: "#fff",
                    Highlight: "#9994b8",
                    HighlightTwo: "#fff",
                },],
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

