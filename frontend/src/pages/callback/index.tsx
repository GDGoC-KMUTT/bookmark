import { server } from "@/configs/server"
import { useEffect } from "react"
import { useNavigate, useSearchParams } from "react-router-dom"
import { toast } from "sonner"

const Callback = () => {
    const [searchParams] = useSearchParams()
    const navigate = useNavigate()
    const code = searchParams.get("code")

    const handleUserCallBack = async () => {
        await toast.promise(server.login.loginCallBack({ code: code ?? "" }), {
            loading: "callback...",
            success: () => {
                navigate("/home")
                return "Successfully callback"
            },
            error: () => {
                return "failed to callback"
            },
        })
    }

    useEffect(() => {
        handleUserCallBack()
    }, [])

    return (
        <div className="flex justify-center items-center w-screen">
            <div className="text-center">
                <h1 className="text-4xl font-semibold text-orange-600">Welcome to Bookmark</h1>
                <div className="mt-6">...Callback...</div>
            </div>
        </div>
    )
}

export default Callback

