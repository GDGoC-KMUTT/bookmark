// pages/portal/index.tsx
import { useState } from "react"
import { Button } from "@/components/ui/button"
import { EnrollmentCard } from "@/components/ui/EnrollmentCard"
import { RecentCard } from "@/components/ui/RecentCard"
import { SuggestionCard } from "@/components/ui/SuggestionCard"
import { StrengthAnalysis } from "@/components/ui/StrengthAnalysis"
import { radarData, radarOptions } from "@/configs/chart"

const Portal = () => {
    const [enrollmentItems, setEnrollmentItems] = useState([
        { title: "Setup LED blink with ESP32 and Arduino IDE", progress: 33, image: "/src/assets/microcontroller.png" },
        { title: "Introduction to Database", progress: 45, image: "/src/assets/microcontroller.png" },
        { title: "Web Development Basics", progress: 60, image: "/src/assets/microcontroller.png" },
        { title: "IoT Fundamentals", progress: 25, image: "/src/assets/microcontroller.png" },
    ])

    const [recentItems, setRecentItems] = useState([
        { title: "Running MySQL Server with Docker" },
        { title: "Test..." },
        { title: "Test2.." },
        { title: "Test3.." },
    ])

    const [suggestionItems, setSuggestionItems] = useState([
        { title: "Setup LED blink with ESP32 and Arduino IDE" },
        { title: "Setup LED blink with ESP32 and Arduino IDE" },
        { title: "Setup LED blink with ESP32 and Arduino IDE" },
        { title: "Setup LED blink with ESP32 and Arduino IDE" },
    ])

    const addEnrollmentItem = () => {
        const newItem = {
            title: `New Course ${enrollmentItems.length + 1}`,
            progress: 0,
            image: "/src/assets/microcontroller.png",
        }
        setEnrollmentItems([...enrollmentItems, newItem])
    }

    const addRecentItem = () => {
        const newItem = {
            title: `New Recent Item ${recentItems.length + 1}`,
        }
        setRecentItems([...recentItems, newItem])
    }

    const addSuggestionItem = () => {
        const newItem = {
            title: `New Suggestion ${suggestionItems.length + 1}`,
        }
        setSuggestionItems([...suggestionItems, newItem])
    }

    return (
        <div className="flex flex-col w-screen bg-gray-50">
            <div className="flex justify-center mx-auto px-8 p-28">
                <div className="space-y-12">
                    {/* Enrollment Section */}
                    <div>
                        <div className="flex justify-between items-center mb-4">
                            <h2 className="text-xl font-semibold">Enrollment</h2>
                            <Button onClick={addEnrollmentItem} variant="outline">
                                Add Enrollment
                            </Button>
                        </div>
                        <div className="flex flex-wrap gap-4">
                            {enrollmentItems.map((item, index) => (
                                <EnrollmentCard key={index} {...item} />
                            ))}
                        </div>
                    </div>

                    {/* Recent Section */}
                    <div>
                        <div className="flex justify-between items-center mb-4">
                            <h2 className="text-xl font-semibold">Recent</h2>
                            <Button onClick={addRecentItem} variant="outline">
                                Add Recent
                            </Button>
                        </div>
                        <div className="flex flex-wrap gap-4">
                            {recentItems.map((item, index) => (
                                <RecentCard key={index} {...item} />
                            ))}
                        </div>
                    </div>

                    {/* Strength Analysis Section */}
                    <div>
                        <h2 className="text-xl font-semibold mb-4">Strength Analysis</h2>
                        <StrengthAnalysis data={radarData} options={radarOptions} />
                    </div>

                    {/* Suggestion Section */}
                    <div>
                        <div className="flex justify-between items-center mb-4">
                            <h2 className="text-xl font-semibold">Suggestion</h2>
                            <Button onClick={addSuggestionItem} variant="outline">
                                Add Suggestion
                            </Button>
                        </div>
                        <div className="flex flex-wrap gap-4">
                            {suggestionItems.map((item, index) => (
                                <SuggestionCard key={index} {...item} />
                            ))}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Portal

