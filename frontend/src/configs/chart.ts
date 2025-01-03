import { Chart as ChartJS, RadialLinearScale, PointElement, LineElement, Filler, Tooltip, Legend } from "chart.js"

ChartJS.register(RadialLinearScale, PointElement, LineElement, Filler, Tooltip, Legend)

export const radarOptions = (data: number[]) => {
    const maxDataValue = Math.max(...data.map((item) => item)) + 10 // Use data if max is undefined
    return {
        scales: {
            r: {
                min: 0,
                max: maxDataValue,
                beginAtZero: true,
                angleLines: {
                    display: true,
                    color: "rgba(0, 0, 0, 0.1)",
                },
                grid: {
                    color: "rgba(0, 0, 0, 0.05)",
                },
                ticks: {
                    stepSize: 20,
                    display: true,
                    backdropColor: "transparent",
                },
                pointLabels: {
                    font: {
                        size: 18,
                    },
                    color: "rgba(0, 0, 0, 0.7)",
                },
            },
        },
        plugins: {
            legend: {
                display: false,
            },
        },
        maintainAspectRatio: false,
    }
}

