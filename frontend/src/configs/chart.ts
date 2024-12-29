import { Chart as ChartJS, RadialLinearScale, PointElement, LineElement, Filler, Tooltip, Legend } from "chart.js"

ChartJS.register(RadialLinearScale, PointElement, LineElement, Filler, Tooltip, Legend)

export const radarOptions = {
    scales: {
        r: {
            min: 0,
            max: 100,
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
                    size: 12,
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

