import { useState, useEffect } from "react"
import { EnrollmentCard } from "@/components/ui/EnrollmentCard"
import { RecentCard } from "@/components/ui/RecentCard"
import { SuggestionCard } from "@/components/ui/SuggestionCard"
import { StrengthAnalysis } from "@/components/ui/StrengthAnalysis"
import { radarData, radarOptions } from "@/configs/chart"
import { server } from "@/configs/server"
import {
    ResponseEnrollmentListDTO,
    ResponseInfoResponseArrayResponseStrengthDataResponse,
    ResponseInfoResponseResponseEnrollmentListResponse,
    ResponseInfoResponseResponseUserActivitiesResponse,
    ResponseStrengthDataResponse,
    ResponseUserActivityResponse,
} from "@/api/api"
import { RadarData } from "@/types/chart"
import { ScrollArea, ScrollBar } from "@/components/ui/scroll-area"

const Portal = () => {
    const [enrollmentItems, setEnrollmentItems] = useState<ResponseEnrollmentListDTO[]>([])
    const [recentActivity, setRecentActivity] = useState<ResponseUserActivityResponse[]>([])
    const [strengthData, setStrengthData] = useState<RadarData | null>(null)

    const [loading, setLoading] = useState(true)
    const [error, setError] = useState<string | null>(null)
    const [debugInfo, setDebugInfo] = useState<string>("")

    const fetchEnrollments = async () => {
        try {
            const response: ResponseInfoResponseResponseEnrollmentListResponse = await server.enroll.getUserEnrollments()
            if (!response?.data?.enrollments || !Array.isArray(response.data.enrollments)) {
                throw new Error("Invalid enrollment data structure")
            }
            const enrollments = response.data.enrollments
            setEnrollmentItems(enrollments)
        } catch (err) {
            const errorMessage = err instanceof Error ? err.message : "An unknown error occurred"
            setError(errorMessage)
            setDebugInfo((prev) => `${prev}\nError: ${errorMessage}`)
        } finally {
            setLoading(false)
        }
    }

    const fetchRecent = async () => {
        try {
            const response: ResponseInfoResponseResponseUserActivitiesResponse = await server.userActivity.getRecentActivity()
            console.log("Fetched Recent Activity Response:", response) // Debug log

            // Validate the structure and check if activities exist
            if (!response?.data?.activities || !Array.isArray(response.data.activities)) {
                throw new Error("Invalid recent activity data structure")
            }

            const activities = response.data.activities

            setRecentActivity(activities)
        } catch (error) {
            const errorMessage = error instanceof Error ? error.message : "An unknown error occurred"
            setError(errorMessage)
            setDebugInfo((prev) => `${prev}\nError: ${errorMessage}`)
        } finally {
            setLoading(false)
        }
    }

    const fetchStrength = async () => {
        try {
            const response: ResponseInfoResponseArrayResponseStrengthDataResponse = await server.userStrength.getStrengthDataByUserId()
            console.log("Fetched Strength Data Response:", response)

            if (!response?.data || !Array.isArray(response.data)) {
                throw new Error("Invalid strength data structure")
            }

            // Transform strengthData into RadarData format
            const strengthData = response.data
            const radarData: RadarData = {
                labels: strengthData.map((item) => item.field_name || "Unknown"),
                datasets: [
                    {
                        label: "User Strength",
                        data: strengthData.map((item) => item.total_gems || 0),
                        backgroundColor: "rgba(54, 162, 235, 0.2)",
                        borderColor: "rgba(54, 162, 235, 1)",
                        borderWidth: 1,
                        pointBackgroundColor: "rgba(54, 162, 235, 1)",
                        pointBorderColor: "#fff",
                        pointHoverBackgroundColor: "#fff",
                        pointHoverBorderColor: "rgba(54, 162, 235, 1)",
                    },
                ],
            }

            setStrengthData(radarData)
        } catch (error) {
            const errorMessage = error instanceof Error ? error.message : "An unknown error occurred"
            setError(errorMessage)
            setDebugInfo((prev) => `${prev}\nError: ${errorMessage}`)
        } finally {
            setLoading(false)
        }
    }

    useEffect(() => {
        fetchEnrollments()
        fetchRecent()
        fetchStrength()
    }, [])

    var enrollmentItems2 = [
        { course_name: "Course 1", progress: 0, id: 1 },
        { course_name: "Course 2", progress: 0, id: 2 },
        { course_name: "Course 3", progress: 0, id: 3 },
        { course_name: "Course 4", progress: 0, id: 4 },
        { course_name: "Course 5", progress: 0, id: 5 },
        { course_name: "Course 6", progress: 0, id: 6 },
        { course_name: "Course 7", progress: 0, id: 7 },
        { course_name: "Course 8", progress: 0, id: 8 },
        { course_name: "Course 9", progress: 0, id: 9 },
        { course_name: "Course 10", progress: 0, id: 10 },
    ]

    var recentActivity2 = [
        { module_title: "Module 1", step_title: "Step 1" },
        { module_title: "Module 2", step_title: "Step 2" },
        { module_title: "Module 3", step_title: "Step 3" },
        { module_title: "Module 4", step_title: "Step 4" },
        { module_title: "Module 5", step_title: "Step 5" },
        { module_title: "Module 6", step_title: "Step 6" },
        { module_title: "Module 7", step_title: "Step 7" },
        { module_title: "Module 8", step_title: "Step 8" },
        { module_title: "Module 9", step_title: "Step 9" },
        { module_title: "Module 10", step_title: "Step 10" },
    ]

    const [suggestionItems, setSuggestionItems] = useState([
        { title: "Setup LED blink with ESP32 and Arduino IDE" },
        { title: "Setup LED blink with ESP32 and Arduino IDE" },
        { title: "Setup LED blink with ESP32 and Arduino IDE" },
        { title: "Setup LED blink with ESP32 and Arduino IDE" },
    ])

    return (
        <div className="flex flex-col h-screen w-screen bg-gray-50 overflow-auto">
            <div className="flex flex-col flex-1 px-8 py-28 space-y-12">
                {/* Enrollment Section */}
                <div className="flex flex-col space-y-4">
                    <h2 className="text-xl font-semibold">Enrollment</h2>
                    {loading ? (
                        <div className="text-gray-500">Loading enrollments...</div>
                    ) : error ? (
                        <div className="flex flex-col gap-2">
                            <div className="text-red-500">Error: {error}</div>
                            <div className="text-sm text-gray-600 whitespace-pre-wrap">{debugInfo}</div>
                        </div>
                    ) : (
                        <ScrollArea className="w-full">
                            <div className="flex w-max space-x-4">
                                {enrollmentItems.map((enrollment, index) => (
                                    <EnrollmentCard
                                        key={enrollment.id || index}
                                        course_name={enrollment.course_name || "Untitled Course"}
                                        progress={enrollment.progress ?? 0}
                                        image="/src/assets/default-image.png"
                                        id={enrollment.id}
                                    />
                                ))}
                            </div>
                            <ScrollBar orientation="horizontal" />
                        </ScrollArea>
                    )}
                </div>

                {/* Recent Section */}
                <div className="flex flex-col space-y-4">
                    <h2 className="text-xl font-semibold">Recent</h2>
                    <ScrollArea className="w-full">
                        {recentActivity2.length > 0 ? (
                            <div className="flex flex-wrap gap-4">
                                {recentActivity.map((activity, index) => (
                                    <RecentCard
                                        key={activity.step_id || index}
                                        moduleTitle={activity.module_title || "Untitled Module"}
                                        stepTitle={activity.step_title || "Untitled Step"}
                                    />
                                ))}
                            </div>
                        ) : (
                            <div className="text-gray-500">No recent activity found</div>
                        )}
                        <ScrollBar orientation="horizontal" />
                    </ScrollArea>
                </div>

                {/* Strength Analysis Section */}
                <div className="flex flex-col space-y-4">
                    <h2 className="text-xl font-semibold">Strength Analysis</h2>

                    {strengthData ? (
                        <StrengthAnalysis data={strengthData} options={radarOptions} />
                    ) : (
                        <div className="text-gray-500">Loading strength data...</div>
                    )}
                </div>

                {/* Suggestion Section */}
                <div className="flex flex-col space-y-4">
                    <h2 className="text-xl font-semibold">Suggestion</h2>
                    <ScrollArea className="w-full h-48">
                        <div className="flex flex-wrap gap-4">
                            {suggestionItems.map((item, index) => (
                                <SuggestionCard key={index} {...item} />
                            ))}
                        </div>
                        <ScrollBar orientation="horizontal" />
                    </ScrollArea>
                </div>
            </div>
        </div>
    )
}

export default Portal

