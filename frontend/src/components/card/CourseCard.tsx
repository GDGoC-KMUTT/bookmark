import { PayloadCourseWithFieldType } from "@/api/api"
import { Bookmark } from "lucide-react"
import { FC } from "react"
import { useNavigate } from "react-router-dom"

interface ICourseCard {
    course: PayloadCourseWithFieldType
}

const CourseCard: FC<ICourseCard> = ({ course }) => {
    const navigate = useNavigate()
    return (
        <div
            onClick={() => {
                navigate(`/course/${course.id}`)
            }}
            className="flex flex-col justify-between p-4 w-[380px] h-[200px] bg-background rounded-sm border border-border break-words m-2 cursor-pointer"
        >
            <div className="text-wrap">
                <div className="flex items-center text-border">
                    <Bookmark className="stroke-2 w-5 h-5" />
                    <p className="text-md font-normal">COURSE / {course.fieldName}</p>
                </div>

                <h2 className="text-foreground text-xl">{course.name}</h2>
            </div>
            <div className="w-[80px] ">
                <img src={course.fieldImageUrl} />
            </div>
        </div>
    )
}

export default CourseCard

