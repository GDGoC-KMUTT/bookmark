// import { course } from "@/types/course";

import { server } from "@/configs/server"
import { toast } from "sonner"

const index = () => {
    const handleRedirect = async () => {
        try {
            await toast.promise(server.login.loginRedirect(), {
                loading: "Redirecting to login page...",
                success: () => {
                    // Handle success message or action
                    return "Successfully load login page"
                },
                error: (err) => {
                    // Handle error message or action
                    console.error("Failed to load login page", err)
                    return "Failed to load login page"
                },
            })
        } catch (error) {
            console.error("Failed to load login page:", error)
            // Handle error message or action
        }
    }
    return (
        <div className="flex justify-center items-center w-screen">
            <div className="text-center">
                <h1 className="text-4xl font-semibold text-orange-600">Welcome to Bookmark</h1>
                <div className="mt-6">
                    <button
                        className="bg-primary text-white rounded-md"
                        onClick={handleRedirect} // Placeholder functionality
                    >
                        Get Start!
                    </button>
                </div>
            </div>
        </div>
    )
}

export default index

