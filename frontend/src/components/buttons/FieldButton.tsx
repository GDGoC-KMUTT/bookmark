import { PayloadFieldType } from "@/api/api"
import React from "react"

interface IFieldButton {
    fieldType: PayloadFieldType
    isActive: boolean
    onClick: (fieldId: number) => void
}

const FieldButton: React.FC<IFieldButton> = ({ fieldType, isActive, onClick }) => {
    return (
        <button
            className={`rounded-lg px-4 py-2 my-2 h-[50px] font-medium hover:border-primary hover:border-2 ${
                isActive ? "bg-explore text-explore-foreground " : "text-foreground bg-transparent "
            } `}
            onClick={() => onClick(fieldType.id || 0)}
        >
            {fieldType.name}
        </button>
    )
}

export default FieldButton

