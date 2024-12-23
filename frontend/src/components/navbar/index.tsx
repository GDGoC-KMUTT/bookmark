import { Gem, BookMarked } from "lucide-react";
import { useEffect, useState } from "react";
import { useAtom } from "jotai";
import { userProfileAtom, totalGemsAtom, currentCourse, progressAtom } from "@/stores/navbar";
import { server } from "@/configs/server";
import { Link, useNavigate } from "react-router-dom";

const Navbar = () => {
    const [userProfile, setUserProfile] = useAtom(userProfileAtom);
    const [totalGems, setTotalGems] = useAtom(totalGemsAtom);
    const [course, setCurrentCourse] = useAtom(currentCourse);
    const [progress, setProgress] = useAtom(progressAtom);
    const [error, setError] = useState<string | null>(null)
    const navigate = useNavigate();

    const handleClick = () => {
      navigate('/profile');
    };

    useEffect(() => {
        const fetchData = async () => {
            try {
                const profile = await server.profile.profileUserInfo();
                setUserProfile(profile.data);
                // console.log("Profile Data:", profile.data)

                const gemsResponse = await server.gems.getUserGems();
                if (gemsResponse.data) {
                    // console.log("Gems Data:", gemsResponse.data);
                    setTotalGems(gemsResponse.data.total as number);  // Set total gems atom
                } else {
                    setTotalGems(0);
                    setError("Gems data is unavailable.");
                }

                const currentCourse = await server.courses.getCurrentCourse();
                if (currentCourse.data && currentCourse.data.name) {
                    // console.log("Current course", currentCourse.data);
                    setCurrentCourse(currentCourse.data.name); 

                    const progressResponse = await server.progress.getCompletionPercentage(currentCourse.data.id as number);
                    setProgress(progressResponse.completion_percentage);
                    // console.log("Progress Data:", progressResponse.completion_percentage);
                } else {
                    setCurrentCourse('No active course');
                    setError("Course data is unavailable.");
                }
            } catch (error) {
                setError("Failed to fetch profile data or gem count.")
            }
        };

        fetchData();
    }, [setUserProfile, setTotalGems]);

    return (
        <div className="w-full bg-white h-[3rem] fixed top-0 shadow-md flex items-center px-6 py-3 justify-between">
            <div className="flex items-center space-x-8">
                <img src="src/assets/logo2.png" alt="Logo" className="w-8 h-8" />
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
                        <div></div>
                        <BookMarked className="text-foreground" size={20} />
                    </div>
                    <div className="items-center justify-center space-y-1">
                        <div className="font text-sm">
                            {course.length > 11 ? `${course.slice(0, 11)}...` : course}
                        </div>
                        <div className="relative w-24 h-1 bg-border rounded-full">
                            <div
                                className="absolute h-1 bg-progressBar rounded-full"
                                style={{ width: `${progress}%` }} // Replace with dynamic width later
                            ></div>
                        </div>
                    </div>
                </div>

                <div className="flex items-center space-x-1 text-foreground">
                    <Gem size={20} />
                    <span className="font-medium">{totalGems}</span>
                </div>

                {/* user profile image */}
                <div className="w-8 h-8 bg-border rounded-full" onClick={handleClick}>
                    {userProfile?.photoUrl && (
                        <img
                            src={userProfile.photoUrl}
                            alt="User Profile"
                            className="w-full h-full object-cover rounded-full"
                        />
                    )}
                </div>
            </div>
        </div>
    );
};

export default Navbar;
