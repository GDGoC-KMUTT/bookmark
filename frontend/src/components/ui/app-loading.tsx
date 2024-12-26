import { PropsWithChildren, useEffect, useState } from "react"
import { useLocation, useNavigate } from "react-router-dom"
import bookmarkLogo from "../../assets/logo2.png"
import { cn } from "@/utils/cn"
import useCurrentUser from "@/hooks/userCurrentUser"

const AppLoading: React.FC<PropsWithChildren> = ({ children }) => {
    const navigate = useNavigate()
    const location = useLocation()

    const { currentUser, isLoading, refetch } = useCurrentUser() // Destructure refetch
    const [isLoaded, setIsLoaded] = useState(false)

    const isLoggedIn = !!currentUser // Ensure it's a boolean

    useEffect(() => {
        if (isLoading) {
            return
        }

        if (location.pathname === "/callback") {
            refetch()
            setIsLoaded(true)
            return
        }

        if (isLoggedIn) {
            navigate("/portal", { replace: true })
        } else {
            navigate("/welcome", { replace: true })
        }

        setIsLoaded(true)
    }, [isLoading, isLoggedIn, navigate, location.pathname, refetch])
    return (
        <>
            <div
                id="app-loading"
                className={cn(
                    "flex z-[9999] transition-all duration-1000 items-center justify-center bg-gray-300 fixed inset-0 pointer-events-none",
                    {
                        "opacity-0": isLoaded,
                        "opacity-100": !isLoaded,
                    }
                )}
            >
                <img src={bookmarkLogo} alt="Bookmark" className="w-96 h-96" />
            </div>
            {isLoaded && children}
        </>
    )
}

export default AppLoading

