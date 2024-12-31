import { server } from "@/configs/server"
import { useCallback, useState } from "react"

export const useComment = () => {
    const [success, setSuccess] = useState<boolean>(false)
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState<string | null>(null)

    const commentOnStep = useCallback(async (stepId: number, comment: string) => {
        setIsLoading(true)
        setError(null)
        setSuccess(false)
        try {
            const response = await server.step.commentOnStep({ stepId: stepId, content: comment })
            if (response.code === 201) {
                setSuccess(true)
                setIsLoading(false)
            }
        } catch (err) {
            setError(`failed to comment: ${err}`)
        } finally {
            setIsLoading(false)
        }
    }, [])

    return { success, isLoading, error, commentOnStep }
}

