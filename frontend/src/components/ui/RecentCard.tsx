import React from "react"
import { Card, CardContent } from "@/components/ui/card"

export interface RecentCardProps {
    title: string
}

export const RecentCard = ({ title }: RecentCardProps) => (
    <Card className="w-[calc(25%-12px)]">
        <CardContent className="p-6">
            <div className="flex flex-col h-full">
                <div className="text-gray-500 text-sm font-medium mb-2 flex items-center gap-2">
                    <svg className="w-4 h-4" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M19 5v14H5V5h14m0-2H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2z" fill="currentColor" />
                    </svg>
                    STEP
                </div>
                <div className="text-gray-900 font-medium mb-4">{title}</div>
                <div className="mt-auto">
                    <img src="/src/assets/database.png" alt="Course icon" className="w-12 h-12 object-contain opacity-80" />
                </div>
            </div>
        </CardContent>
    </Card>
)
