import { Link } from "react-router-dom"
type CourseCardProps = {
    courseName: string
    fieldName: string
    imageUrl: string
    courseId: number
}

const CourseCard: React.FC<CourseCardProps> = ({ courseName, fieldName, imageUrl, courseId }) => {
    return (
        <Link to={`/courses/${courseId}`} className="block">
            <div className="border p-5 pb-2 h-50 pr-10 rounded-sm flex flex-col space-y-2">
                <div className="flex items-center space-x-2">
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        strokeWidth={1.5}
                        stroke="currentColor"
                        className="w-4 h-4 text-gray-500"
                    >
                        <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            d="M17.593 3.322c1.1.128 1.907 1.077 1.907 2.185V21L12 17.25 4.5 21V5.507c0-1.108.806-2.057 1.907-2.185a48.507 48.507 0 0 1 11.186 0Z"
                        />
                    </svg>
                    <p className="text-gray-500">{fieldName}</p>
                </div>
                <p className="font-medium pb-5 text-black break-words">{courseName}</p>
                <div className="flex-grow flex items-end justify-start">
                    <img src={imageUrl} className="w-12 h-12 object-contain" />
                </div>
            </div>
        </Link>
    )
}

export default CourseCard
