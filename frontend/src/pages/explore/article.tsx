import React, { ChangeEvent, useState } from "react"
import { Search } from "lucide-react"
import { useNavigate } from "react-router-dom"
import useArticles from "@/hooks/useArticles"
import ArticleCard from "@/components/card/ArticleCard"

const ExploreArticle = () => {
    const [keyword, setKeyword] = useState<string>("")
    const navigate = useNavigate()
    const { articles, isLoading: isArticlesLoading } = useArticles()

    if (isArticlesLoading) {
        return <div>loading</div>
    }

    const handleKeywordChange = (event: ChangeEvent<HTMLInputElement>) => {
        setKeyword(event.target.value)
    }
    const filteredArticles = articles?.filter((article) => article.title?.toLowerCase().includes(keyword.toLowerCase()))

    return (
        <div className="min-h-screen flex flex-col items-center w-[100vw] pb-20">
            <div className="flex flex-col justify-center items-center  p-20 mt-12 bg-droplet1 bg-center bg-no-repeat">
                <h1 className="mb-[50px] text-explore-foreground">Explore something new!</h1>
                <div className="flex justify-center space-x-4 mb-28">
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
                        className="rounded-lg w-[120px] h-[50px] text-foreground bg-transparent hover:border-primary hover:border-2"
                        onClick={() => navigate("/explore/course")}
                    >
                        Course
                    </button>
                    <button
                        className="bg-explore rounded-lg w-[120px] h-[50px] bg-explore text-explore-foreground hover:border-primary hover:border-2 "
                        onClick={() => navigate("/explore/article")}
                        disabled={true}
                    >
                        Article
                    </button>
                </div>

                {/* <div className="flex justify-center flex-wrap space-x-3 max-w-[70%] my-[20px]">
                    {fieldTypes?.map((fieldType) => (
                        <CategoryButton
                            key={fieldType.id}
                            fieldType={fieldType}
                            isActive={activeField === fieldType.id}
                            onClick={() => handleFieldChange(fieldType.id || 0)}
                        />
                    ))}{" "}
                </div> */}
            </div>
            <div className="flex flex-wrap  items-center ">
                {filteredArticles?.map((article) => {
                    return <ArticleCard key={article.id} article={article} />
                })}
            </div>
        </div>
    )
}

export default ExploreArticle

