import { PayloadProfile } from "@/api/api"
import { server } from "@/configs/server"
import { useEffect, useState, useCallback } from "react"

export default function useCurrentUser() {
    const [currentUser, setCurrentUser] = useState<PayloadProfile | undefined>(undefined)
    const [isLoading, setIsLoading] = useState(true)
    const [error, setError] = useState<Error | null>(null)

    const fetchCurrentUser = useCallback(async () => {
        setIsLoading(true)
        setError(null)
        try {
            const response = await server.profile.profileUserInfo()
            setCurrentUser(response.data)
        } catch (err) {
            setError(err as Error)
            setCurrentUser(undefined)
        } finally {
            setIsLoading(false)
        }
    }, [])

    useEffect(() => {
        fetchCurrentUser()
    }, [fetchCurrentUser])

    return { currentUser, isLoading, error, refetch: fetchCurrentUser }
}

