import { PayloadCourseWithFieldType } from "@/api/api"
import { Bookmark } from "lucide-react"
import { FC } from "react"

interface ICourseCard {
    course: PayloadCourseWithFieldType
}

const CourseCard: FC<ICourseCard> = ({ course }) => {
    return (
        <div className="flex flex-col justify-between p-4 w-[380px] h-[200px] bg-background rounded-sm border border-border break-words m-2">
            <div className="text-wrap">
                <div className="flex items-center text-border">
                    <Bookmark className="stroke-2 w-5 h-5" />
                    <p className="text-md font-normal">COURSE / {course.field_name}</p>
                </div>

                <h2 className="text-foreground text-xl">{course.name}</h2>
            </div>
            <div className="w-[60px] h-[40px] bg-slate-300">
                <img src={course.field_image} />
            </div>
        </div>
    )
}

export default CourseCard

