import { Link } from "react-router-dom"
import { Bookmark } from "lucide-react"

type CourseCardProps = {
    courseName: string
    fieldName: string
    imageUrl: string
    courseId: number
}

const CourseCard: React.FC<CourseCardProps> = ({ courseName, fieldName, imageUrl, courseId }) => {
    return (
        <Link to={`/course/${courseId}`} className="block">
            <div className="border p-5 pb-2 w-[350px] h-[200px] pr-10 rounded-sm flex flex-col space-y-2">
                <div className="flex items-center space-x-2">
                <Bookmark className="stroke-2 w-5 h-5 text-gray-500" />
                    <p className="text-gray-500">{fieldName}</p>
                </div>
                <p className="font-medium pb-3 text-black break-words">{courseName}</p>
                <div className="flex-grow flex items-end justify-start">
                    <img src={imageUrl} className="w-16 h-16 object-contain" />
                </div>
            </div>
        </Link>
    )
}

export default CourseCard
