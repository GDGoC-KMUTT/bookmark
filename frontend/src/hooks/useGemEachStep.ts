import { PayloadGetGemsResponse } from "@/api/api"
import { server } from "@/configs/server"
import { useCallback, useEffect, useState } from "react"

export const useGemEachStep = (stepId: number) => {
    const [gemEachStep, setGemEachStep] = useState<PayloadGetGemsResponse | undefined>(undefined)
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState<Error | null>(null)

    const fetchGemEachStep = useCallback(async () => {
        setIsLoading(true)
        setError(null)
        try {
            const response = await server.step.getGemEachStep(stepId)
            setGemEachStep(response.data)
        } catch (err) {
            setError(err as Error)
            setGemEachStep(undefined)
        } finally {
            setIsLoading(false)
        }
    }, [stepId])

    useEffect(() => {
        fetchGemEachStep()
    }, [fetchGemEachStep])

    return { gemEachStep, isLoading, error, fetchGemEachStep }
}

