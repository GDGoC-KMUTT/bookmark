/** @type {import('tailwindcss').Config} */
export default {
    content: ["./index.html", "./src/**/*.{ts,tsx,js,jsx}"],
    theme: {
        extend: {
            borderRadius: {
                DEFAULT: "val(--radius)",
                lg: "var(--radius-large)",
                md: "var(--radius-medium)",
                sm: "var(--radius-small)",
            },
            colors: {
                background: "hsl(var(--background))",
                foreground: "hsl(var(--foreground))",
                primary: "hsl(var(--chart))",
                secondary: "hsl(var(--secondary))",

                border: "hsl(var(--border))",

                chart: {
                    bg: "hsl(var(--chart),0.5)",
                    border: "hsl(var(--chart-border))",
                    pointBg: "hsl(var(--chart-point))",
                    pointBorder: "hsl(var(--chart-point-border))",
                    pointHover: "hsl(var(--chart-point-hover))",
                    pointHoverBorder: "hsl(var(--chart-point-hover-border))",
                },
                progressBar: "hsl(var(--progress-bar))",
                continueCourse: "hsl(var(--continue-course))",
                fieldType: {
                    DEFAULT: "hsl(var(--field-type)",
                    foreground: "hsl(var(--field-type-foreground))",
                },
                footer: "hsl(var(--footer)",
                codeTag: "hsl(var(--code-tag)",
                badge: {
                    outcome: "hsl(var(--outcome-badge)",
                    check: "hsl(var(--check-badge)",
                    error: "hsl(var(--error-badge)",
                    comment: "hsl(var(--comment-badge)",
                    evaluate: "hsl(var(--evaluate-badge)",
                },
                evalPassed: {
                    DEFAULT: "hsl(var--check-passed)",
                    foreground: "hsl(var(--check-passed-foreground))",
                },
                evalFailed: {
                    DEFAULT: "hsl(var(--check-failed)",
                    foreground: "hsl(var(--check-failed-foreground))",
                },
                form: "hsl(var(--form)",
                completeStep: "hsl(var(--complete-step)",
                incompleteStep: "hsl(var(--incomplete-step)",
                explore: {
                    DEFAULT: "hsl(var(--explore)",
                    foreground: "hsl(var(--explore-foreground))",
                },
            },
        },
        fontFamily: {
            sans: ['"Ubuntu"', '"Noto Sans Thai"', "sans-serif", "Source Code Pro"],
        },
    },
    plugins: [require("tailwindcss-animate")],
}

