import { PayloadCourseWithFieldType } from "@/api/api"
import { server } from "@/configs/server"
import { useEffect, useState } from "react"

export default function useCourseWithFieldId(fieldId: number | undefined) {
    const [courses, setCourses] = useState<PayloadCourseWithFieldType[] | undefined>(undefined)
    const [isLoading, setIsLoading] = useState(true)
    const [error, setError] = useState<Error | null>(null)

    useEffect(() => {
        if (fieldId === undefined) return

        const fetchCourseWithFieldId = async () => {
            setIsLoading(true)
            setError(null)
            try {
                const response = await server.course.getCoursesByFieldId(Number(fieldId))
                setCourses(response.data)
            } catch (err) {
                setError(err as Error)
                setCourses(undefined)
            } finally {
                setIsLoading(false)
            }
        }

        fetchCourseWithFieldId()
    }, [fieldId])

    return { courses, isLoading, error }
}

