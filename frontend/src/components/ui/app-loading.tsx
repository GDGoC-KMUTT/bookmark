import { PropsWithChildren, useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"
import bookmarkLogo from "../../assets/logo2.png"
import { cn } from "@/utils/cn"
import useCurrentUser from "@/hooks/userCurrentUser"

const AppLoading: React.FC<PropsWithChildren> = ({ children }) => {
    const navigate = useNavigate()

    const { currentUser, isLoading } = useCurrentUser()
    const isLoggedIn = currentUser != null
    const [isLoaded, setIsLoaded] = useState(false)

    const verify = currentUser?.id !== undefined

    useEffect(() => {
        if (isLoading) {
            return
        }

        if (!isLoggedIn && !isLoading) {
            if (verify) {
                setIsLoaded(true)
                return
            }

            navigate("/welcome")
            const timeout = setTimeout(() => {
                setIsLoaded(true)
            }, 500)

            return () => {
                clearTimeout(timeout)
            }
        }
        navigate("/portal")
        setIsLoaded(true)
    }, [isLoading, isLoggedIn, navigate, verify])

    return (
        <>
            <div
                id="app-loading"
                className={cn(
                    "flex z-[9999] transition-all duration-1000 items-center justify-center bg-incompleteStep fixed inset-0 pointer-events-none",
                    {
                        "opacity-0": isLoaded,
                        "opacity-100": !isLoaded,
                    }
                )}
            >
                <img src={bookmarkLogo} alt="Bookmark" className={cn(`w-96 h-96`)} />
            </div>
            {isLoaded ? children : null}
        </>
    )
}

export default AppLoading

