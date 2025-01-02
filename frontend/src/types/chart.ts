export interface RadarData {
    labels: string[]
    datasets: {
        label: string
        data: number[]
        backgroundColor: string
        borderColor: string
        borderWidth: number
        pointBackgroundColor: string
        pointBorderColor: string
        pointHoverBackgroundColor: string
        pointHoverBorderColor: string
    }[]
}

export interface RadarOptions {
    scales: {
        r: {
            min: number
            max: number
            beginAtZero: boolean
            angleLines: {
                display: boolean
                color: string
            }
            grid: {
                color: string
            }
            ticks: {
                stepSize: number
                display: boolean
                backdropColor: string
            }
            pointLabels: {
                font: {
                    size: number
                }
                color: string
            }
        }
    }
    plugins: {
        legend: {
            display: boolean
        }
    }
    maintainAspectRatio: boolean
}

