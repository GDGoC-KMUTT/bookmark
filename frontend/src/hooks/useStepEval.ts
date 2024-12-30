import { PayloadStepEvalInfo } from "@/api/api"
import { server } from "@/configs/server"
import { useCallback, useEffect, useState } from "react"

export const useStepEval = (stepId: number) => {
    const [stepEval, setStepEval] = useState<PayloadStepEvalInfo[] | undefined>(undefined)
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState<string | null>(null)

    const fetchStepEval = useCallback(async () => {
        setIsLoading(true)
        setError(null)
        try {
            const response = await server.step.getStepEvaluate(stepId)
            setStepEval(response.data)
        } catch (err) {
            setError(`failed to fetch step evaluation: ${err}`)
            setStepEval(undefined)
        } finally {
            setIsLoading(false)
        }
    }, [stepId])

    useEffect(() => {
        fetchStepEval()
    }, [fetchStepEval])

    return { stepEval, isLoading, error, fetchStepEval }
}

