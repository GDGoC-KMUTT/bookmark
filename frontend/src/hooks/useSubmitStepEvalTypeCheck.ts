import { PayloadCreateUserEvalRes } from "@/api/api"
import { server } from "@/configs/server"
import { useCallback, useState } from "react"

export const useSubmitStepEvalTypeCheck = () => {
    const [createUserEvalIdResCheck, setCreateUserEvalIdResCheck] = useState<PayloadCreateUserEvalRes | undefined>(undefined)
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState<string | null>(null)

    const submitStepEvalTypeCheck = useCallback(async (stepEvalId: number) => {
        setIsLoading(true)
        setError(null)
        try {
            // Send only JSON if the type is not "image"
            const response = await server.step.submitStepEvalTypCheck({ stepEvalId: stepEvalId })
            setCreateUserEvalIdResCheck(response.data)
            console.log("Step evaluation submitted successfully:", response)
        } catch (error) {
            setError(`failed to submit step evaluation: ${error}`)
            console.error("Error submitting step evaluation:", error)
        }
    }, [])

    return { createUserEvalIdResCheck, isLoading, error, submitStepEvalTypeCheck }
}

