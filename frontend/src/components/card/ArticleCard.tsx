import { PayloadArticle } from "@/api/api"
import { Text } from "lucide-react"
import { FC } from "react"

interface IArticleCard {
    article: PayloadArticle
}

const ArticleCard: FC<IArticleCard> = ({ article }) => {
    const openArticle = () => {
        if (article?.href) {
            window.open(article.href, "_blank", "noopener,noreferrer")
        }
    }
    return (
        <div
            className="flex flex-col p-4 w-[380px] h-[200px] bg-background rounded-sm border border-border break-words m-2 cursor-pointer"
            onClick={openArticle}
        >
            <div className="text-wrap">
                <div className="flex items-center text-border">
                    <Text className="stroke-2 w-5 h-5" />
                    <p className="text-md font-normal">Article</p>
                </div>

                <h2 className="text-foreground text-xl">{article.title}</h2>
            </div>
        </div>
    )
}

export default ArticleCard

