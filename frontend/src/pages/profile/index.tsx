import { useEffect, useState } from "react"
import { server } from "@/configs/server"
import { PayloadProfile } from "../../api/api"
import CourseCard from "../../components/coursecard"

const Profile = () => {
    const [profile, setProfile] = useState<PayloadProfile | null>(null)
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState<string | null>(null)

    const fetchProfileData = async () => {
        try {
            const profileData = await server.profile.profileUserInfo()
            if (profileData.data) {
                // console.log("Profile Data:", profileData.data)
                setProfile(profileData.data)
            } else {
                // console.log("Failed response:")
                setError("Failed to fetch profile data.")
                setProfile(null)
            }
        } catch (error) {
            console.error("Error fetching profile data:", error)
            setError("Error fetching profile data.")
            setProfile(null)
        } finally {
            setLoading(false)
        }
    }

    useEffect(() => {
        fetchProfileData()
    }, [])

    if (loading) {
        return <div>Loading...</div>
    }
    if (error) {
        return <div>{error}</div>
    }
    if (!profile) {
        return <div>Please Login first.</div>
    }
    return (
        <div className="flex flex-col w-full max-w-7xl mx-auto p-6 sm:p-10">
            <div className="flex flex-col items-start space-y-4">
                <h1 className="text-3xl mt-10 mb-4 font-medium">Profile</h1>
                <div className="flex flex-col sm:flex-row items-center space-y-4 sm:space-x-6 sm:space-y-0">
                    <div className="flex justify-center items-center">
                        <img
                            src={profile.photoUrl}
                            className="w-24 h-24 sm:w-32 sm:h-32 rounded-full object-cover shadow-md"
                            alt="Profile"
                        />
                    </div>
                    <div className="flex flex-col">
                        <h2 className="text-2xl sm:text-3xl font-medium">
                            {profile.firstname} {profile.lastname}
                        </h2>
                        {/* Change color */}
                        <p className="text-gray-500 text-sm sm:text-base">{profile.email}</p>
                    </div>
                </div>
            </div>

            <div className="mt-8 justify-center">
                <h1 className="text-2xl sm:text-3xl ml-5 font-medium">Enrolled Courses</h1>
                <div className="mt-5 mb-10  grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 justify-center sm:justify-start">
                    {/* <CourseCard courseName="Introduction to Computer Science" status="Completed" />
                    <CourseCard courseName="Web Development" status="In Progress" /> */}

                    <CourseCard/>
                    <CourseCard />
                    <CourseCard/>
                    <CourseCard />
                    <CourseCard/>
                    <CourseCard />
                    <CourseCard/>
                    <CourseCard />
                    {/* <div className="bg-slate-400 p-6 m-4 min-h-[200px] min-w-[250px] sm:min-w-[300px] flex items-center justify-center"> BOX1 </div>
                    <div className="bg-slate-500 p-6 m-4 min-h-[200px] min-w-[250px] sm:min-w-[300px] flex items-center justify-center"> BOX2 </div>
                    <div className="bg-slate-600 p-6 m-4 min-h-[200px] min-w-[250px] sm:min-w-[300px] flex items-center justify-center"> BOX3 </div>
                    <div className="bg-slate-700 p-6 m-4 min-h-[200px] min-w-[250px] sm:min-w-[300px] flex items-center justify-center"> BOX4 </div> */}
                </div>
            </div>
        </div>
    )
}

export default Profile
