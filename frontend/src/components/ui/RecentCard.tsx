import { Card, CardContent } from "@/components/ui/card"
import { Blocks } from "lucide-react" // Import the Blocks icon

export interface RecentCardProps {
    moduleTitle: string // Renamed for clarity
    stepTitle: string
}

export const RecentCard = ({ moduleTitle, stepTitle }: RecentCardProps) => (
    <Card className="w-[calc(25%-12px)]">
        <CardContent className="p-6">
            <div className="flex flex-col h-full">
                {/* Replace the SVG with the Blocks icon */}
                <div className="text-gray-500 text-sm font-medium mb-2 flex items-center gap-2">
                    <Blocks className="w-4 h-4" /> {/* Blocks Icon */}
                    STEP
                </div>
                <div className="text-gray-900 font-medium mb-2">{moduleTitle || "Untitled Module"}</div>
                <div className="text-gray-700 text-sm mb-4">{stepTitle || "Untitled Step"}</div>
                <div className="mt-auto">
                    <img src="/src/assets/database.png" alt="Course icon" className="w-12 h-12 object-contain opacity-80" />
                </div>
            </div>
        </CardContent>
    </Card>
)

