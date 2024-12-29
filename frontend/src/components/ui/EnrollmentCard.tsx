import { Card, CardContent } from "@/components/ui/card"
import { ArrowRight } from "lucide-react"

export interface EnrollmentCardProps {
    id?: string | number
    course_name: string
    progress?: number
}

export const EnrollmentCard = ({ course_name, progress = 0 }: EnrollmentCardProps) => (
    <Card className="w-[300px]">
        <CardContent className="p-6">
            <div className="flex justify-between items-start mb-2">
                <div>
                    <div className="text-gray-500 font-medium text-sm mb-2">CONTINUE</div>
                    <div className="text-gray-900 font-medium text-base mb-4" style={{ minHeight: "3rem" }}>
                        {course_name}
                    </div>
                </div>
                <div className="flex gap-1.5 mt-1">
                    <img
                        src="https://static.bookmark.scnd.app/asset/fieldicon/microcontroller.png"
                        alt="Course icon"
                        className="w-12 h-12 object-contain opacity-80"
                    />
                </div>
            </div>

            <div className="flex items-center gap-4">
                <div className="flex-grow">
                    <div className="w-full h-1.5 bg-gray-100 rounded-full">
                        <div
                            className="h-full bg-green-500 rounded-full transition-all duration-300"
                            style={{ width: `${Math.min(Math.max(progress, 0), 100)}%` }}
                        />
                    </div>
                </div>
                <div className="flex items-center justify-center w-8 h-8 rounded-full bg-blue-500 text-white">
                    <ArrowRight size={18} />
                </div>
            </div>
        </CardContent>
    </Card>
)

