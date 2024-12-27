import { PayloadFieldType } from "@/api/api"
import { server } from "@/configs/server"
import { useEffect, useState } from "react"

export default function useFieldTypes() {
    const [fieldTypes, setFieldTypes] = useState<PayloadFieldType[] | undefined>(undefined)
    const [isLoading, setIsLoading] = useState(true)
    const [error, setError] = useState<Error | null>(null)

    useEffect(() => {
        const fetchFieldTypes = async () => {
            setIsLoading(true)
            setError(null)
            try {
                const response = await server.course.getAllFieldTypes()
                const sortedFieldTypes = response?.data?.sort((a, b) => (a.id || 0) - (b.id || 0))
                setFieldTypes(sortedFieldTypes)
            } catch (err) {
                setError(err as Error)
                setFieldTypes(undefined)
            } finally {
                setIsLoading(false)
            }
        }

        fetchFieldTypes()
    }, [])

    return { fieldTypes, isLoading, error }
}

