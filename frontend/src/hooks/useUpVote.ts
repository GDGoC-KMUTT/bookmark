import { server } from "@/configs/server"
import { useCallback, useState } from "react"

export const useUpVote = () => {
    const [success, setSuccess] = useState<boolean>(false)
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState<Error | null>(null)

    const upvoteComment = useCallback(async (stepCommentId: number) => {
        setIsLoading(true)
        setError(null)
        setSuccess(false)
        try {
            const response = await server.step.upVoteStepComment({ stepCommentId })
            if (response.code === 201) {
                setSuccess(true)
                setIsLoading(false)
            }
        } catch (err) {
            setError(err as Error)
        } finally {
            setIsLoading(false)
        }
    }, [])

    return { success, isLoading, error, upvoteComment }
}

