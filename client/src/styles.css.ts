import { style, globalStyle } from "@vanilla-extract/css"

export const theme = {
    bg: '#191825',
    pink: '#F86FBA',
}

globalStyle('html, body', {
    margin: 0,
    backgroundColor: theme.bg,
})

export const hello = style({
    color: theme.pink,
    fontSize: 28,
    margin: 20
})