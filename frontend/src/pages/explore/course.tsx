import { ChangeEvent, useEffect, useState } from "react"
import { Search } from "lucide-react"
import { useNavigate } from "react-router-dom"
import FieldButton from "@/components/buttons/FieldButton"
import CourseCard from "@/components/card/CourseCard"
import useCourseWithFieldId from "@/hooks/useCourseByFieldId"
import useFieldTypes from "@/hooks/useFieldTypes"
const ExploreCourse = () => {
    const [keyword, setKeyword] = useState<string>("")
    const [activeField, setActiveField] = useState<number | undefined>(undefined)
    const { fieldTypes, isLoading: isLoadingFieldTypes } = useFieldTypes()
    const { courses, isLoading: isLoadingCourseWithFiledId } = useCourseWithFieldId(activeField)

    const navigate = useNavigate()

    useEffect(() => {
        if (fieldTypes && fieldTypes.length > 0) {
            setActiveField(fieldTypes[0].id)
        }
    }, [fieldTypes])

    if (isLoadingFieldTypes || isLoadingCourseWithFiledId) {
        return <div>loading</div>
    }

    const handleKeywordChange = (event: ChangeEvent<HTMLInputElement>) => {
        setKeyword(event.target.value)
    }

    const handleFieldChange = (fieldId: number) => {
        setActiveField(fieldId)
        console.log(`Filtering results for field: ${fieldId}`)
    }

    return (
        <div className="min-h-screen flex flex-col items-center w-[100vw] ">
            <div className="flex flex-col justify-center items-center  p-20 mt-12 bg-droplet1 bg-center bg-no-repeat">
                <h1 className="mb-[50px] text-explore-foreground">Explore something new!</h1>
                <div className="flex justify-center space-x-4">
                    <div className="relative ">
                        <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 " />
                        <input
                            type="text"
                            className="pl-10 pr-4 py-2 bg-form rounded-full w-64  h-[50px]"
                            value={keyword}
                            onChange={handleKeywordChange}
                        />
                    </div>
                    <button
                        className="bg-explore rounded-lg w-[120px] h-[50px] bg-explore text-explore-foreground hover:border-primary hover:border-2 "
                        onClick={() => navigate("/explore/course")}
                        disabled={true}
                    >
                        Course
                    </button>
                    <button
                        className="rounded-lg w-[120px] h-[50px] text-foreground bg-transparent hover:border-primary hover:border-2"
                        onClick={() => navigate("/explore/article")}
                    >
                        Article
                    </button>
                </div>
                <div className="flex justify-center flex-wrap space-x-3 max-w-[70%] my-[20px]">
                    {fieldTypes?.map((fieldType) => (
                        <FieldButton
                            key={fieldType.id}
                            fieldType={fieldType}
                            isActive={activeField === fieldType.id}
                            onClick={() => handleFieldChange(fieldType.id || 0)}
                        />
                    ))}{" "}
                </div>
            </div>
            <div className="flex flex-wrap  items-center mb-20 ">
                {courses?.map((course) => {
                    return <CourseCard key={course.id} course={course} />
                })}
            </div>
        </div>
    )
}

export default ExploreCourse

