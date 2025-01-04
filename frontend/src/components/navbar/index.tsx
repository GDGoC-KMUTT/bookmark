import type { PayloadProfile } from "@/api/api"
import BookmarkLogo from "@/assets/logo2.png"
import { server } from "@/configs/server"
import { BookMarked, Gem } from "lucide-react"
import { useEffect, useState } from "react"
import { Link, useNavigate } from "react-router-dom"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { getFallbackName } from "@/utils/getFallbackName"
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "@/components/ui/tooltip"

const Navbar = () => {
    const [userProfile, setUserProfile] = useState<PayloadProfile | undefined>(undefined)
    const [totalGems, setTotalGems] = useState<number | null>(null)
    const [course, setCurrentCourse] = useState<string>("")
    const [progress, setProgress] = useState<number | null>(null)
    const navigate = useNavigate()
    // hot fix
    const handleClick = () => {
        navigate("/profile")
    }

    useEffect(() => {
        const fetchData = async () => {
            try {
                const profile = await server.profile.profileUserInfo()
                setUserProfile(profile.data)

                const gemsResponse = await server.gems.getUserGems()
                setTotalGems(gemsResponse.data?.total || 0)

                const currentCourse = await server.courses.getCurrentCourse()
                if (currentCourse.data?.name) {
                    setCurrentCourse(currentCourse.data.name)

                    const progressResponse = await server.progress.getCompletionPercentage(currentCourse.data.id as number)
                    setProgress(progressResponse.data || 0)
                } else {
                    setCurrentCourse("No active course")
                }
            } catch (error) {
                if (typeof error === "object" && error !== null && "response" in error && typeof (error as any).response?.status === "number") {
                    const responseError = error as { response: { status: number } }
                    if (responseError.response.status === 500) {
                        setCurrentCourse("No active course")
                    }
                } else {
                    setCurrentCourse("No active course")
                }
            }
        }

        fetchData()
    }, [])

    return (
        <div className="w-full bg-white h-[3rem] fixed top-0 shadow-md flex items-center px-6 py-3 justify-between z-[99]">
            <div className="flex items-center space-x-8">
                <Link to="/home">
                    <img src={BookmarkLogo} alt="bookmarkLogo" className="w-8 h-8" />
                </Link>
                <div className="flex space-x-8">
                    <Link to="/home" className="text-gray-500 font-medium hover:text-explore-foreground transition-colors">
                        Home
                    </Link>
                    <Link to="/explore" className="text-gray-500 font-medium hover:text-explore-foreground transition-colors">
                        Explore
                    </Link>
                </div>
            </div>

            <div className="flex items-center justify-center space-x-6">
                <div className="flex items-center space-x-2">
                    <div className="space-y-1">
                        <BookMarked className="text-foreground" size={20} />
                    </div>
                    <div className="items-center justify-center space-y-1">
                        <TooltipProvider>
                            <Tooltip>
                                <TooltipTrigger asChild>
                                    <div className="font text-sm">{course.length > 11 ? `${course.slice(0, 11)}...` : course}</div>
                                </TooltipTrigger>
                                <TooltipContent>
                                    <p>{course}</p>
                                </TooltipContent>
                            </Tooltip>
                        </TooltipProvider>
                        <div className="relative w-24 h-1 bg-border rounded-full">
                            <div className="absolute h-1 bg-progressBar rounded-full" style={{ width: `${progress}%` }}></div>
                        </div>
                    </div>
                </div>

                <div className="flex items-center space-x-1 text-foreground">
                    <Gem size={20} />
                    <span className="font-medium">{totalGems ?? 0}</span>
                </div>

                {/* User profile with fallback to avatar */}
                <div className="w-8 h-8 cursor-pointer" onClick={handleClick}>
                    <Avatar className="w-8 h-8 rounded-full bg-slate-200 text-center font-bold text-sm flex items-center justify-center">
                        <AvatarImage
                            src={userProfile?.photoUrl || undefined}
                            alt="User Profile"
                            className="w-full h-full object-cover rounded-full"
                        />
                        {!userProfile?.photoUrl && userProfile && (
                            <AvatarFallback>{getFallbackName(`${userProfile.firstname} ${userProfile.lastname}`)}</AvatarFallback>
                        )}
                    </Avatar>
                </div>
            </div>
        </div>
    )
}

export default Navbar

