import { PayloadStepInfo } from "@/api/api"
import { server } from "@/configs/server"
import { useCallback, useEffect, useState } from "react"

export const useStepInfo = (stepId: number) => {
    const [stepInfo, setStepInfo] = useState<PayloadStepInfo | undefined>(undefined)
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState<Error | null>(null)

    const fetchStepInfo = useCallback(async () => {
        setIsLoading(true)
        setError(null)
        try {
            const response = await server.step.getStepInfo(stepId)
            setStepInfo(response.data)
        } catch (err) {
            setError(err as Error)
            setStepInfo(undefined)
        } finally {
            setIsLoading(false)
        }
    }, [stepId])

    useEffect(() => {
        fetchStepInfo()
    }, [fetchStepInfo])

    return { stepInfo, isLoading, error, refetch: fetchStepInfo }
}


