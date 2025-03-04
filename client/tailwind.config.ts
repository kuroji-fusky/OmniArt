import type { Config } from "tailwindcss"
import typography from "@tailwindcss/typography"
import forms from "@tailwindcss/forms"

const config: Config = {
  content: ["./src/**/*.{html,js,svelte,ts}"],
  // For flexible light/dark mode
  darkMode: ["variant", [
    "html.dark &",
    "@media (prefers-color-scheme: dark) { & }"
  ]],
  theme: {
    extend: {},
  },
  plugins: [typography, forms]
}

export default config
