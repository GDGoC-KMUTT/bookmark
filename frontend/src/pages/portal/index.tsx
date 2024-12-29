import { useState, useEffect } from "react"
import { EnrollmentCard } from "@/components/ui/EnrollmentCard"
import { RecentCard } from "@/components/ui/RecentCard"
import { SuggestionCard } from "@/components/ui/SuggestionCard"
import { StrengthAnalysis } from "@/components/ui/StrengthAnalysis"
import { radarOptions } from "@/configs/chart"
import { server } from "@/configs/server"
import { ResponseEnrollmentListDTO, ResponseUserActivityResponse, ResponseCourseResponse } from "@/api/api"
import { RadarData } from "@/types/chart"
import { ScrollArea, ScrollBar } from "@/components/ui/scroll-area"

const Portal = () => {
    // State
    const [data, setData] = useState({
        enrollments: [] as ResponseEnrollmentListDTO[],
        recentActivity: [] as ResponseUserActivityResponse[],
        suggestions: [] as ResponseCourseResponse[],
        strengthData: null as RadarData | null,
    })

    const [loading, setLoading] = useState(true)
    const [error, setError] = useState<string | null>(null)

    // Fetch Functions
    const fetchData = async (fetchFn: () => Promise<any>, onSuccess: (data: any) => void, errorMessage: string) => {
        try {
            const response = await fetchFn()
            if (!response?.data) {
                throw new Error(`Invalid ${errorMessage} data structure`)
            }
            onSuccess(response.data)
        } catch (err) {
            setError(err instanceof Error ? err.message : "An unknown error occurred")
        }
    }

    useEffect(() => {
        const fetchAllData = async () => {
            setLoading(true)
            await Promise.all([
                fetchData(server.enroll.getUserEnrollments, (data) => setData((prev) => ({ ...prev, enrollments: data.enrollments })), "enrollment"),
                fetchData(
                    server.userActivity.getRecentActivity,
                    (data) => setData((prev) => ({ ...prev, recentActivity: data.activities })),
                    "recent activity"
                ),
                fetchData(
                    server.userStrength.getStrengthDataByUserId,
                    (data) => {
                        const radarData: RadarData = {
                            labels: data.map((item: any) => item.field_name || "Unknown"),
                            datasets: [
                                {
                                    label: "User Strength",
                                    data: data.map((item: any) => item.total_gems || 0),
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
                        setData((prev) => ({ ...prev, strengthData: radarData }))
                    },
                    "strength"
                ),
                fetchData(server.userStrength.getSuggestionCourse, (data) => setData((prev) => ({ ...prev, suggestions: data })), "suggestion"),
            ])
            setLoading(false)
        }

        fetchAllData()
    }, [])

    // Render Section Component
    const renderSection = (title: string, content: JSX.Element, showScroll: boolean = true) => (
        <div className="flex flex-col space-y-4">
            <h2 className="text-xl font-semibold">{title}</h2>
            {showScroll ? (
                <ScrollArea className="w-full">
                    {content}
                    <ScrollBar orientation="horizontal" />
                </ScrollArea>
            ) : (
                content
            )}
        </div>
    )

    // Section Contents
    const enrollmentContent = (
        <div className="flex w-max space-x-4">
            {data.enrollments.map((enrollment, index) => (
                <EnrollmentCard
                    key={enrollment.id || index}
                    course_name={enrollment.course_name || "Untitled Course"}
                    progress={enrollment.progress ?? 0}
                    image="/src/assets/default-image.png"
                    id={enrollment.id}
                />
            ))}
        </div>
    )

    const strengthContent = data.strengthData ? (
        <StrengthAnalysis data={data.strengthData} options={radarOptions} />
    ) : (
        <div className="text-gray-500">Loading strength data...</div>
    )

    const recentContent = (
        <div className="flex w-full space-x-4 pb-4">
            {data.recentActivity.length > 0 ? (
                data.recentActivity.map((activity, index) => (
                    <div className="flex-1" key={activity.step_id || index}>
                        {" "}
                        {/* Key moved to parent div */}
                        <RecentCard moduleTitle={activity.module_title || "Untitled Module"} stepTitle={activity.step_title || "Untitled Step"} />
                    </div>
                ))
            ) : (
                <div className="text-gray-500">No recent activity found</div>
            )}
        </div>
    )

    const suggestionContent = (
        <div className="flex w-full space-x-4 pb-4">
            {data.suggestions.length > 0 ? (
                data.suggestions.map((item, index) => (
                    <div className="flex-1" key={item.id || index}>
                        {" "}
                        {/* Key moved to parent div */}
                        <SuggestionCard name={item.name || "Untitled Course"} field={item.field} />
                    </div>
                ))
            ) : (
                <div className="text-gray-500">No suggestions available</div>
            )}
        </div>
    )
    return (
        <div className="flex flex-col h-screen w-screen bg-gray-50 overflow-auto">
            <div className="flex flex-col flex-1 px-20 py-28 space-y-12">
                {loading ? (
                    <div className="text-gray-500">Loading...</div>
                ) : error ? (
                    <div className="text-red-500">Error: {error}</div>
                ) : (
                    <>
                        {renderSection("Enrollment", enrollmentContent)}
                        {renderSection("Recent", recentContent)}
                        {renderSection("Strength Analysis", strengthContent, false)}
                        {renderSection("Suggestion", suggestionContent)}
                    </>
                )}
            </div>
        </div>
    )
}

export default Portal

