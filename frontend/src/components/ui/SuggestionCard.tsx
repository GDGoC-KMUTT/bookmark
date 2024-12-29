// Update SuggestionCard component
import { Card, CardContent } from "@/components/ui/card"
import { BookOpen } from "lucide-react"

export interface SuggestionCardProps {
    name: string
    field?: {
        name?: string
    }
}

export const SuggestionCard = ({ name = "Untitled Course", field = { name: "INFRASTRUCTURE" } }: SuggestionCardProps) => (
    <Card className="w-full h-full">
        <CardContent className="p-6">
            <div className="flex flex-col">
                <div className="flex items-center gap-2 mb-3">
                    <span className="text-gray-500 text-sm font-medium inline-flex items-center">
                        <BookOpen className="w-4 h-4 mr-1" />
                        COURSE / {field.name}
                    </span>
                </div>

                <div className="text-gray-900 font-medium text-base mb-4" style={{ minHeight: "3rem" }}>
                    {name}
                </div>

                <div className="mt-auto">
                    <img src="/src/assets/microcontroller.png" alt="Course icon" className="w-8 h-8 object-contain opacity-80" />
                </div>
            </div>
        </CardContent>
    </Card>
)
