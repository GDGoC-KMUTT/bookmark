import { ModelsUserEvaluate } from "@/api/api"
import { server } from "@/configs/server"
import { useCallback, useState } from "react"

export const useCheckStepEvalStatus = () => {
    const [userEvalStatus, setUserEvalStatus] = useState<ModelsUserEvaluate | undefined>(undefined)
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState<string | null>(null)

    const fetchStepEvalStatus = useCallback(async (userEvalId: number) => {
        setIsLoading(true)
        setError(null)
        try {
            const response = await server.step.checkStepEvalStatus({ userEvalId: userEvalId })
            setUserEvalStatus(response.data)
        } catch (err) {
            setError(`failed to fetch stepInfo: ${err}`)
            setUserEvalStatus(undefined)
        } finally {
            setIsLoading(false)
        }
    }, [])

    return { userEvalStatus, isLoading, error, fetchStepEvalStatus }
}

