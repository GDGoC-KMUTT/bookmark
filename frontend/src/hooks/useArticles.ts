import { PayloadArticle } from "@/api/api"
import { server } from "@/configs/server"
import { useEffect, useState } from "react"

const useArticles = () => {
    const [articles, setArticles] = useState<PayloadArticle[] | undefined>(undefined)
    const [isLoading, setIsLoading] = useState(true)
    const [error, setError] = useState<Error | null>(null)

    useEffect(() => {
        const fetchAllArticles = async () => {
            setIsLoading(true)
            setError(null)
            try {
                const response = await server.article.getAllArticles()
                setArticles(response.data)
            } catch (err) {
                setError(err as Error)
                setArticles(undefined)
            } finally {
                setIsLoading(false)
            }
        }

        fetchAllArticles()
    }, [])

    return { articles, isLoading, error }
}

export default useArticles

