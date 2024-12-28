import { Card, CardContent } from "@/components/ui/card"

export interface SuggestionCardProps {
    title: string
}

export const SuggestionCard = ({ title }: SuggestionCardProps) => (
    <Card className="w-[calc(25%-12px)]">
        <CardContent className="p-6">
            <div className="flex flex-col h-full">
                <div className="text-gray-500 text-xs font-medium tracking-wide mb-2">COURSE / INFRASTRUCTURE</div>
                <div className="text-gray-900 font-medium mb-4">{title}</div>
                <div className="mt-auto">
                    <img src="/src/assets/microcontroller.png" alt="Course icon" className="w-12 h-12 object-contain opacity-80" />
                </div>
            </div>
        </CardContent>
    </Card>
)

