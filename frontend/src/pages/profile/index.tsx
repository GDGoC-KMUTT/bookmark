import { useEffect, useState } from "react"
import { server } from "@/configs/server"
import { PayloadProfile, PayloadEnrollwithCourse } from "../../api/api"
import CourseCard from "../../components/coursecard"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { getFallbackName } from "@/utils/getFallbackName"

const Profile = () => {
    const [profile, setProfile] = useState<PayloadProfile | undefined>(undefined)
    const [enrolledCourses, setEnrolledCourses] = useState<PayloadEnrollwithCourse[] | undefined>(undefined)
    const [error, setError] = useState<string | null>(null)

    useEffect(() => {
        const fetchData = async () => {
            try {
                // Fetch profile data
                const profileResponse = await server.profile.profileUserInfo()
                if (profileResponse.code === 200) {
                    setProfile(profileResponse.data)
                } else {
                    setError("Failed to fetch profile data. Please try again.")
                    return
                }

                // Fetch enrolled courses
                const coursesResponse = await server.courses.getEnrollCourseByUserId()
                if (coursesResponse.code === 200) {
                    setEnrolledCourses(coursesResponse.data)
                } else {
                    setError("Failed to fetch enrolled courses. Please try again.")
                }
            } catch (error) {
                console.error("Error during data fetching: ", error)
                setError("An error occurred while fetching data. Please try again.")
            }
        }

        fetchData()
    }, [])

    if (error) {
        return (
            <div className="flex items-center justify-center w-screen">
                <div className="text-center">
                    <h1 className="text-4xl font-bold mb-4">{error}</h1>
                    <p className="text-3xl text-gray-700">Please refresh the page and try again.</p>
                </div>
            </div>
        )
    }

    return (
        <div className="items-start sm:p-20 p-10 min-h-screen">
            <div className="flex flex-col items-start space-y-4">
                <h1 className="text-3xl mt-10 mb-4 font-medium">Profile</h1>
                <div className="flex flex-col sm:flex-row items-center  space-y-4 sm:space-x-6 sm:space-y-0">
                    {profile && (
                        <>
                            <Avatar className="w-28 h-28 sm:w-24 sm:h-24 rounded-full object-cover shadow-md bg-slate-200 text-continueCourse text-xl font-bold">
                                <AvatarImage src={profile.photoUrl} alt={`${profile.firstname} ${profile.lastname}`} />
                                <AvatarFallback>{getFallbackName(`${profile.firstname} ${profile.lastname}`)}</AvatarFallback>
                            </Avatar>
                            <div className="flex flex-col self-center">
                                <h2 className="text-2xl sm:text-3xl font-medium">
                                    {profile.firstname} {profile.lastname}
                                </h2>
                                <p className="text-gray-500 text-sm sm:text-base">{profile.email}</p>
                            </div>
                        </>
                    )}
                </div>
            </div>

            <div className="mt-10 sm:mt-20">
                <h1 className="text-3xl font-medium">Enrolled Courses</h1>
                <div className="mt-8 mb-10 gap-6 justify-center sm:justify-start">
                    {enrolledCourses?.length === 0 ? (
                        <p className="ml-10">No courses enrolled yet.</p>
                    ) : (
                        <div className="flex flex-wrap mt-8 mb-10 gap-6">
                            {enrolledCourses?.map((course) => (
                                <CourseCard
                                    key={course.id}
                                    courseName={course.courseName?.name ?? ""}
                                    fieldName={course.fieldName ?? ""}
                                    imageUrl={course.fieldImageUrl ?? ""}
                                    courseId={course.courseId ?? 0}
                                />
                            ))}
                        </div>
                    )}
                </div>
            </div>
        </div>
    )
}
export default Profile
