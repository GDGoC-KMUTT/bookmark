import { useEffect, useState } from "react"
import { server } from "@/configs/server"
import { PayloadProfile, PayloadEnrollwithCourse } from "../../api/api"
import CourseCard from "../../components/coursecard"

const Profile = () => {
    const [profile, setProfile] = useState<PayloadProfile | null>(null)
    const [enrolledCourses, setEnrolledCourses] = useState<PayloadEnrollwithCourse[]>([])
    const [error, setError] = useState<string | null>(null)

    const fetchProfileData = async () => {
        try {
            const profileData = await server.profile.profileUserInfo()
            if (profileData.data) {
                console.log("Profile Data:", profileData.data)
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
        }
    }

    const fetchEnrolledCourses = async () => {
        try {
            const coursesData = await server.courses.getEnrollCourseByUserId()
            if (coursesData.data) {
                console.log("Course Data:", coursesData.data)

                setEnrolledCourses(coursesData.data)
            } else {
                setError("Failed to fetch enrolled courses.")
            }
        } catch (error) {
            console.error("Error fetching enrolled courses: hereeeee", error)
            setError("Error fetching enrolled courses.")
        }
    }

    useEffect(() => {
        const fetchData = async () => {
            await fetchProfileData()
            await fetchEnrolledCourses()
        }
        fetchData()
    }, [])

    if (error) {
        return <div>{error}</div>
    }
    if (!profile) {
        return <div>Please Login first.</div>
    }
    return (
        <div className="flex flex-col items-start w-full max-w-7xl mx-auto p-10">
            <div className="flex flex-col items-start space-y-4">
                <h1 className="text-3xl mt-10 mb-4 font-medium">Profile</h1>
                <div className="flex flex-col sm:flex-row items-start space-y-4 sm:space-x-6 sm:space-y-0">
                    <div className="flex justify-center items-center">
                        <img src={profile.photoUrl} className="w-20 h-20 rounded-full object-cover shadow-md"/>
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

            <div className="mt-16">
                <h1 className="text-2xl sm:text-3xl ml-5 font-medium">Enrolled Courses</h1>
                <div className="mt-8 mb-10 grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 justify-center sm:justify-start">
                    {enrolledCourses.map((course) => (
                        <CourseCard
                            key={course.id}
                            courseName={course.course_name?.name ?? ""}
                            fieldName={course.field_name ?? ""}
                            imageUrl={course.field_image_url?? ""}
                            courseId={course.course_id ?? 0}
                        />
                    ))}
                </div>
            </div>
        </div>
    )
}
export default Profile
