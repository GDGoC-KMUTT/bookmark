import { SquareCheck, SquareX } from "lucide-react"
import React, { FC } from "react"

type ResultCardProps = {
    pass: boolean
    comment?: string
}

const IconMap: Record<string, React.ReactNode> = {
    true: <SquareCheck />,
    false: <SquareX />,
}

const BgColorAndTextColorMap: Record<string, string> = {
    true: "bg-check-passed text-green-600 rounded-sm p-4",
    false: "bg-check-failed text-red-600 rounded-sm p-4",
}

const ResultCard: FC<ResultCardProps> = ({ pass, comment }) => {
    return (
        <div className={BgColorAndTextColorMap[pass.toString()]}>
            <div className="flex gap-3 ">
                {IconMap[pass.toString()]}
                <p>Check {pass ? "Passed" : "Failed"}</p>
            </div>
            {comment && <p>{comment}</p>}
        </div>
    )
}

export default ResultCard

