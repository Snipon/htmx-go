/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
      "./webroot/**/*.html",
      "./base.css",
  ],
  theme: {
    container: {
      center: true,
      padding: "3rem",
      screens: {
        "2xl": "1440px",
      },
    },
    extend: {
      colors: {
        foreground: "#333",
        background: "#f8f8f8",
        dimmed: "#ccc",
      },
      fontFamily: {
        inter: "'Inter', sans-serif",
      },
      fontSize: {
        xs: ".75rem",
        "2xs": ".5rem",
      },
      spacing: {
        xs: "1rem", 
        sm: "1.5rem",
        md: "2.25rem",
        lg: "5rem",
      },
      maxWidth: {
        fullScreen: "1440px",
        content: "1113px",
        contentSmall: "900px",
        contentSmallest: "659px",
      },
      animation: {
        "fade-in": "fade-in 500ms ease-in-out",
        "fade-out": "fade-out 500ms ease-in-out",
      },
      keyframes: {
        "fade-in": {
          "0%": { opacity: 0 },
          "100%": { opacity: 1 },
        },
        "fade-out": {
          "0%": { opacity: 1 },
          "100%": { opacity: 0 },
        },
      }
    },
  },
  plugins: [],
}
