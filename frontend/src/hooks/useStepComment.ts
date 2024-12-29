import { PayloadStepCommentInfo } from "@/api/api"
import { server } from "@/configs/server"
import { useCallback, useEffect, useState } from "react"

export const useStepComment = (stepId: number) => {
    const [stepComments, setStepComments] = useState<PayloadStepCommentInfo[] | undefined>(undefined)
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState<Error | null>(null)

    const fetchStepComment = useCallback(async () => {
        console.log("fetchStepComment")

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
    }, [stepId])

    useEffect(() => {
        fetchStepComment()
    }, [fetchStepComment])

    return { stepComments, setStepComments, isLoading, error, fetchStepComment }
}

