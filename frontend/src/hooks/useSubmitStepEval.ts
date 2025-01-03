import { PayloadCreateUserEvalRes } from "@/api/api"
import { server } from "@/configs/server"
import { SubmitStepEval } from "@/types/step"
import { useCallback, useState } from "react"

export const useSubmitStepEval = () => {
    const [createUserEvalRes, setCreateUserEvalRes] = useState<PayloadCreateUserEvalRes | undefined>(undefined)
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState<string | null>(null)

    const submitStepEval = useCallback(async (stepId: number, stepEvalId: number, type: string, file?: File, content?: string) => {
        setIsLoading(true)
        setError(null)
        const body: SubmitStepEval = {
            stepId: stepId,
            stepEvalId: stepEvalId,
        }
        if (type === "text") {
            body.content = content
        }

        if (type === "image" && file) {
            try {
                const response = await server.step.submitStepEval({ data: JSON.stringify(body), file: file })
                setCreateUserEvalRes(response.data)
                console.log("File uploaded successfully:", response)
                setIsLoading(false)
            } catch (error) {
                setIsLoading(false)
                setError(`failed to submit step evaluation: ${error}`)
                console.error("Error submitting step evaluation:", error)
            }
        } else {
            try {
                // Send only JSON if the type is not "image"
                const response = await server.step.submitStepEval({
                    data: JSON.stringify(body),
                })
                setCreateUserEvalRes(response.data)
                setIsLoading(false)
            } catch (error) {
                setIsLoading(false)
                setError(`failed to submit step evaluation: ${error}`)
                console.error("Error submitting step evaluation:", error)
            }
        }
    }, [])

    return { createUserEvalRes, isLoading, error, submitStepEval }
}

