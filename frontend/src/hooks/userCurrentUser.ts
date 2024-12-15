import { PayloadProfile } from "@/api/api"
import { server } from "@/configs/server"
import { useEffect, useState } from "react"

export default function useCurrentUser() {
    const [currentUser, setCurrentUser] = useState<PayloadProfile | undefined>(undefined)
    const [isLoading, setIsLoading] = useState(true)
    const [error, setError] = useState<Error | null>(null)

    useEffect(() => {
        const fetchCurrentUser = async () => {
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
        }

        fetchCurrentUser()
    }, [])

    return { currentUser, isLoading, error }
}

