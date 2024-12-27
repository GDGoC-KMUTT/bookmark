import { Radar } from "react-chartjs-2"
import { Card, CardContent } from "@/components/ui/card"
import { RadarData, RadarOptions } from "@/types/chart"

interface StrengthAnalysisProps {
    data: RadarData
    options: RadarOptions
}

export const StrengthAnalysis = ({ data, options }: StrengthAnalysisProps) => (
    <Card className="border-none shadow-none">
        <CardContent className="p-8">
            <div className="h-[400px] max-w-2xl mx-auto">
                <Radar data={data} options={options} />
            </div>
        </CardContent>
    </Card>
)

