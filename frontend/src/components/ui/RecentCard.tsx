import { Card, CardContent } from "@/components/ui/card"
import { Blocks } from "lucide-react"

export interface RecentCardProps {
    moduleTitle: string
    stepTitle: string
}

export const RecentCard = ({ moduleTitle, stepTitle }: RecentCardProps) => (
    <Card className="w-full h-full">
        <CardContent className="p-6">
            <div className="flex flex-col">
                <div className="flex items-center gap-2 mb-3">
                    <span className="text-gray-500 text-sm font-medium inline-flex items-center">
                        <Blocks className="w-4 h-4 mr-1" />
                        STEP
                    </span>
                </div>

                <div className="text-gray-900 font-medium text-base mb-4" style={{ minHeight: "3rem" }}>
                    {moduleTitle}
                    <p className="text-sm text-gray-600">{stepTitle}</p>
                </div>

                <div className="mt-4">
                    <img src="/src/assets/database.png" alt="Step icon" className="w-12 h-12 object-contain opacity-80" />
                </div>
            </div>
        </CardContent>
    </Card>
)

