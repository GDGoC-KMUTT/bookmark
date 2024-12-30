import useCurrentUser from "@/hooks/userCurrentUser"
import { cn } from "@/utils/cn"
import { PropsWithChildren, useEffect, useState } from "react"
import { useLocation, useNavigate } from "react-router-dom"
import bookmarkLogo from "../../assets/logo2.png"

const AppLoading: React.FC<PropsWithChildren> = ({ children }) => {
    const navigate = useNavigate()
    const location = useLocation()

    const { currentUser, isLoading, refetch } = useCurrentUser()
    const [isLoaded, setIsLoaded] = useState(false)

    useEffect(() => {
        if (isLoading) {
            return
        }

        // If on `/callback`, refetch user info and prevent redirection
        if (location.pathname === "/callback") {
            refetch().then(() => {
                setIsLoaded(true)
            })
            return
        }

        // Redirect based on login state and root path
        if (currentUser) {
            if (location.pathname === "/welcome" || location.pathname === "/") {
                navigate("/home", { replace: true }) // Redirect to portal
            } else {
                setIsLoaded(true) // Allow navigation for other paths
            }
        } else {
            if (location.pathname !== "/welcome") {
                navigate("/welcome", { replace: true }) // Redirect to welcome
            } else {
                setIsLoaded(true)
            }
        }
    }, [isLoading, currentUser, navigate, location.pathname, refetch])

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
                <img src={bookmarkLogo} alt="Bookmark" className={"w-96 h-96 animate-spin-360"} />
            </div>
            {isLoaded && children}
        </>
    )
}

export default AppLoading

