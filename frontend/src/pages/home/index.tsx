import { useState, useEffect } from "react"
import { EnrollmentCard } from "@/components/home/EnrollmentCard"
import { RecentCard } from "@/components/home/RecentCard"
import { SuggestionCard } from "@/components/home/SuggestionCard"
import { StrengthAnalysis } from "@/components/home/StrengthAnalysis"
import { radarOptions } from "@/configs/chart"
import { server } from "@/configs/server"
import { PayloadEnrollmentListDTO, PayloadUserActivityResponse, PayloadCourseResponse } from "@/api/api"
import { RadarData } from "@/types/chart"
import { ScrollArea, ScrollBar } from "@/components/ui/scroll-area"
import { useNavigate } from "react-router-dom"
import { Loader2 } from "lucide-react"

const Portal = () => {
    // State
    const navigate = useNavigate()
    const [data, setData] = useState({
        enrollments: [] as PayloadEnrollmentListDTO[],
        recentActivity: [] as PayloadUserActivityResponse[],
        suggestions: [] as PayloadCourseResponse[],
        strengthData: null as RadarData | null,
    })
    const [gem, setGems] = useState<number[]>([])

    const [loading, setLoading] = useState(true)
    const [error, setError] = useState<Record<string, string | null>>({
        enrollment: null,
        recentActivity: null,
        strength: null,
        suggestion: null,
    })
    console.log(data.strengthData)

    // Fetch Functions
    const fetchData = async (fetchFn: () => Promise<any>, onSuccess: (data: any) => void, errorMessage: string) => {
        try {
            const response = await fetchFn()
            if (!response?.data) {
                throw new Error(`Invalid ${errorMessage} data structure`)
            }
            onSuccess(response.data)
        } catch (err) {
            return err instanceof Error ? err.message : "An unknown error occurred"
        }
    }

    useEffect(() => {
        const fetchAllData = async () => {
            setLoading(true)
            const [enrollmentError, recentActivityError, strengthError, suggestionError] = await Promise.all([
                fetchData(server.enroll.getUserEnrollments, (data) => setData((prev) => ({ ...prev, enrollments: data.enrollments })), "enrollment"),
                fetchData(
                    server.userActivity.getRecentActivity,
                    (data) => setData((prev) => ({ ...prev, recentActivity: data.activities })),
                    "recent activity"
                ),
                fetchData(
                    server.userStrength.getStrengthDataByUserId,
                    (data) => {
                        setGems((prev) => [...prev, ...data.data.map((item: any) => item.totalGems || 0)])

                        const radarData: RadarData = {
                            labels: data.data.map((item: any) => item.fieldName || "Unknown"),
                            datasets: [
                                {
                                    label: data.username,
                                    data: data.data.map((item: any) => item.totalGems || 0),
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
            setError({
                enrollment: enrollmentError ?? null,
                recentActivity: recentActivityError ?? null,
                strength: strengthError ?? null,
                suggestion: suggestionError ?? null,
            })
        }

        fetchAllData()
    }, [])

    // Render Section Component
    const renderSection = (title: string, content: JSX.Element, errorMessage: string | null, showScroll: boolean = true) => (
        <div className="flex flex-col space-y-4">
            <h2 className="text-xl font-semibold">{title}</h2>
            {errorMessage ? (
                <div className="text-red-500">Error: {errorMessage}</div>
            ) : showScroll ? (
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
            {data.enrollments?.map((enrollment, index) => (
                <div
                    className="cursor-pointer"
                    onClick={() => {
                        navigate(`/course/${enrollment.courseId}`)
                    }}
                >
                    <EnrollmentCard
                        key={enrollment.id || index}
                        course_name={enrollment.courseName || "Untitled Course"}
                        progress={enrollment.progress ?? 0}
                        id={enrollment.id}
                    />
                </div>
            ))}
        </div>
    )

    const strengthContent = data.strengthData ? (
        <StrengthAnalysis data={data.strengthData} options={radarOptions(gem)} />
    ) : (
        <div className="text-gray-500">No strength data</div>
    )

    const recentContent = (
        <ScrollArea className="w-full">
            <div className="flex space-x-4 pb-4">
                {data.recentActivity ? (
                    data.recentActivity.map((activity, index) => (
                        <div
                            key={activity.stepId || index}
                            className="w-80 flex-none cursor-pointer"
                            onClick={() => navigate(`/course/${activity.courseId}`)}
                        >
                            <RecentCard moduleTitle={activity.moduleTitle || "Untitled Module"} stepTitle={activity.stepTitle || "Untitled Step"} />
                        </div>
                    ))
                ) : (
                    <div className="text-gray-500">No recent activity found</div>
                )}
            </div>
            <ScrollBar orientation="horizontal" />
        </ScrollArea>
    )

    const suggestionContent = (
        <div className="flex w-full space-x-4 pb-4">
            {data.suggestions.length > 0 ? (
                data.suggestions.map((item, index) => (
                    <div className="flex-1 cursor-pointer" key={item.id || index} onClick={() => navigate(`/course/${item.id}`)}>
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
                    <div className="flex flex-col items-center justify-center h-full">
                        <Loader2 className="animate-spin" size={50} />
                        <p className="text-xl">Loading Home page</p>
                    </div>
                ) : (
                    <>
                        {renderSection("Enrollment", enrollmentContent, error.enrollment)}
                        {renderSection("Recent", recentContent, error.recentActivity)}
                        {renderSection("Strength Analysis", strengthContent, error.strength, false)}
                        {renderSection("Suggestion", suggestionContent, error.suggestion)}
                    </>
                )}
            </div>
        </div>
    )
}

export default Portal

