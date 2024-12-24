import React from "react"
import { ArrowRight } from "lucide-react"
import { Chart as ChartJS, RadialLinearScale, PointElement, LineElement, Filler, Tooltip, Legend } from "chart.js"
import { Radar } from "react-chartjs-2"

ChartJS.register(RadialLinearScale, PointElement, LineElement, Filler, Tooltip, Legend)

const Portal = () => {
    const recentItems = [{ title: "Running MySQL Server with Docker" }, { title: "Test..." }, { title: "Test2.." }, { title: "Test3.." }]

    const suggestionItems = [
        { title: "Setup LED blink with ESP32 and Arduino IDE" },
        { title: "Setup LED blink with ESP32 and Arduino IDE" },
        { title: "Setup LED blink with ESP32 and Arduino IDE" },
        { title: "Setup LED blink with ESP32 and Arduino IDE" },
    ]

    const data = {
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
    const enrollmentItems = [
        { title: "Setup LED blink with ESP32 and Arduino IDE", progress: 33, image: "/src/assets/microcontroller.png" },
        { title: "Introduction to Database", progress: 45, image: "/src/assets/microcontroller.png" },
        { title: "Web Development Basics", progress: 60, image: "/src/assets/microcontroller.png" },
        { title: "IoT Fundamentals", progress: 25, image: "/src/assets/microcontroller.png" },
    ]

    const options = {
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

    return (
        <div className="flex flex-col w-screen bg-gray-50 ">
            <div className="flex justify-center mx-auto px-8 p-28">
                {/* Main Sections Container */}
                <div className="space-y-12">
                    {/* Enrollment Section */}
                    <div>
                        <h2 className="text-xl font-semibold mb-4">Enrollment</h2>
                        <div className="flex flex-wrap gap-4">
                            {enrollmentItems.map((item, index) => (
                                <div key={index} className="bg-white p-6 rounded-2xl shadow-md w-[300px] relative">
                                    <div className="flex justify-between items-start mb-2">
                                        <div>
                                            <div className="text-gray-500 font-medium text-sm mb-2">CONTINUE</div>
                                            <div className="text-gray-900 font-medium text-base mb-4">{item.title}</div>
                                        </div>
                                        <div className="flex gap-1.5 mt-1">
                                            <img src={item.image} alt="Course icon" className="w-12 h-12 object-contain opacity-80" />
                                        </div>
                                    </div>

                                    <div className="flex items-center gap-4">
                                        <div className="flex-grow">
                                            <div className="w-full h-1.5 bg-gray-100 rounded-full">
                                                <div className="h-full bg-green-500 rounded-full" style={{ width: `${item.progress}%` }}></div>
                                            </div>
                                        </div>
                                        <div className="flex items-center justify-center w-8 h-8 rounded-full bg-blue-500 text-white">
                                            <ArrowRight size={18} />
                                        </div>
                                    </div>
                                </div>
                            ))}
                        </div>
                    </div>

                    {/* Recent Section */}
                    <div>
                        <h2 className="text-xl font-semibold mb-4">Recent</h2>
                        <div className="flex flex-wrap gap-4">
                            {recentItems.map((item, index) => (
                                <div key={index} className="bg-white p-6 border border-gray rounded-2xl shadow-sm w-[calc(25%-12px)]">
                                    <div className="flex flex-col h-full">
                                        <div className="text-gray-500 text-sm font-medium mb-2 flex items-center gap-2">
                                            <svg className="w-4 h-4" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                                <path
                                                    d="M19 5v14H5V5h14m0-2H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2z"
                                                    fill="currentColor"
                                                />
                                            </svg>
                                            STEP
                                        </div>
                                        <div className="text-gray-900 font-medium mb-4">{item.title}</div>
                                        <div className="mt-auto">
                                            <img src="/src/assets/database.png" alt="Course icon" className="w-12 h-12 object-contain opacity-80" />
                                        </div>
                                    </div>
                                </div>
                            ))}
                        </div>
                    </div>

                    {/* Strength Analysis Section */}
                    <div>
                        <h2 className="text-xl font-semibold mb-4">Strength Analysis</h2>
                        <div className=" p-8 rounded-2xl shadow-sm">
                            <div className="h-[400px] max-w-2xl mx-auto">
                                <Radar data={data} options={options} />
                            </div>
                        </div>
                    </div>

                    {/* Suggestion Section */}
                    <div>
                        <h2 className="text-xl font-semibold mb-4">Suggestion</h2>
                        <div className="flex flex-wrap gap-4">
                            {suggestionItems.map((item, index) => (
                                <div key={index} className="bg-white p-6 border border-gray rounded-2xl shadow-sm w-[calc(25%-12px)]">
                                    <div className="flex flex-col h-full">
                                        <div className="text-gray-500 text-xs font-medium tracking-wide mb-2">COURSE / INFRASTRUCTURE</div>
                                        <div className="text-gray-900 font-medium mb-4">{item.title}</div>
                                        <div className="mt-auto">
                                            <img
                                                src="/src/assets/microcontroller.png"
                                                alt="Course icon"
                                                className="w-12 h-12 object-contain opacity-80"
                                            />
                                        </div>
                                    </div>
                                </div>
                            ))}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Portal

