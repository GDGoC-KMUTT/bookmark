import { PayloadStepCommentInfo } from "@/api/api"
import { server } from "@/configs/server"
import { useCallback, useEffect, useState } from "react"

export const useStepComment = (stepId: number) => {
    const [stepComments, setStepComments] = useState<PayloadStepCommentInfo[] | undefined>(undefined)
    const [isLoading, setIsLoading] = useState(true)
    const [error, setError] = useState<Error | null>(null)

    const fetchStepComment = useCallback(async () => {
        setIsLoading(true)
        setError(null)
        try {
            const response = await server.step.getStepComment(stepId)
            setStepComments(response.data)
        } catch (err) {
            setError(err as Error)
            setStepComments(undefined)
        } finally {
            setIsLoading(false)
        }
    }, [])

    useEffect(() => {
        fetchStepComment()
    }, [fetchStepComment])

    return { stepComments, isLoading, error, fetchStepComment }
}

