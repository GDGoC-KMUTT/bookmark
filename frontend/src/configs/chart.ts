import { Chart as ChartJS, RadialLinearScale, PointElement, LineElement, Filler, Tooltip, Legend } from "chart.js"

ChartJS.register(RadialLinearScale, PointElement, LineElement, Filler, Tooltip, Legend)

export const radarData = {
    labels: ["Microcontroller", "Database", "Internet of Things", "Web Development", "Operating System"],
    datasets: [
        {
            label: "Strength Quotient",
            data: [35, 65, 85, 45, 30],
            backgroundColor: "rgba(255, 178, 166, 0.3)",
            borderColor: "rgba(255, 178, 166, 1)",
            borderWidth: 1,
            pointBackgroundColor: "rgba(255, 178, 166, 1)",
            pointBorderColor: "#fff",
            pointHoverBackgroundColor: "#fff",
            pointHoverBorderColor: "rgba(255, 178, 166, 1)",
        },
    ],
}

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
