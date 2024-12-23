import { SERVER_HOST } from "@/configs/server"
import { toast } from "sonner"

const index = () => {
    const handleRedirect = async () => {
        await toast.promise(
            new Promise((resolve) => {
                setTimeout(() => {
                    resolve("Redirecting now...")
                    window.location.assign(`${SERVER_HOST}/login/redirect`)
                }, 1000) // 3 seconds delay before redirect
            }),
            {
                loading: "Preparing to redirect to the login page...",
                success: () => {
                    return "Successfully prepared for the login page redirect."
                },
            }
        )
    }

    return (
        <div className="flex justify-center items-center w-screen">
            <div className="text-center">
                <h1 className="text-4xl font-semibold text-orange-600">Welcome to Bookmark</h1>
                <div className="mt-6">
                    <button
                        className="bg-primary text-white rounded-md"
                        onClick={handleRedirect} // Trigger the redirect inside the handler
                    >
                        Get Started!
                    </button>
                </div>
            </div>
        </div>
    )
}

export default index

