import { ArrowBigUp, BadgeCheck, Blocks, CircleSlash, Gem, Loader2, MessageSquare, ShieldQuestion } from "lucide-react"
import { AspectRatio } from "@/components/ui/aspect-ratio"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { ScrollArea } from "@/components/ui/scroll-area"
import { Sheet, SheetContent, SheetHeader, SheetTrigger } from "@/components/ui/sheet"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Badge } from "@/components/ui/badge"
import { Separator } from "@/components/ui/separator"
import { TooltipProvider, Tooltip, TooltipTrigger, TooltipContent } from "@radix-ui/react-tooltip"
import { getFallbackName } from "@/utils/getFallbackName"
import { useStepInfo } from "@/hooks/useStepInfo"
import { useGemEachStep } from "@/hooks/useGemEachStep"
import { useStepComment } from "@/hooks/useStepComment"
import { Input } from "@/components/ui/input"
import { useState } from "react"
import { useComment } from "@/hooks/useComment"
import { PayloadStepCommentInfo } from "@/api/api"
import useCurrentUser from "@/hooks/userCurrentUser"
import { useUpVote } from "@/hooks/useUpVote"
import { useStepEval } from "@/hooks/useStepEval"
import EvalTypeCard from "./eval-type"
import MarkdownRenderer from "../ui/markdown-renderer"

type StepProps = {
    stepId: number
}

const StepCard: React.FC<StepProps> = ({ stepId }) => {
    const [userComment, setUserComment] = useState<string>("")

    const { stepInfo, error: errorStepInfo, isLoading: isLoadingStepInfo } = useStepInfo(stepId)
    const {
        gemEachStep,
        error: errorGetGemEachStep,
        isLoading: isLoadingGetGemEachStep,
        fetchGemEachStep: refetchGemEachStep,
    } = useGemEachStep(stepId)
    const { stepComments, setStepComments, error: errorGetStepComment, fetchStepComment: refetchStepComment } = useStepComment(stepId)
    const { stepEval, error: errorGetStepEval } = useStepEval(stepId)

    const { isLoading: isLoadingCommentOnStep, commentOnStep } = useComment()
    const { currentUser } = useCurrentUser()
    const { upvoteComment } = useUpVote()

    const submitComment = async () => {
        if (userComment !== "") {
            try {
                await commentOnStep(stepId, userComment)
                const newComment: PayloadStepCommentInfo = {
                    userInfo: {
                        firstName: currentUser?.firstname,
                        lastname: currentUser?.lastname,
                        photoUrl: currentUser?.photoUrl,
                    },
                    comment: userComment,
                    upVote: 0,
                }

                // Optimistically update the comments list in the UI
                stepComments?.unshift(newComment)
                refetchStepComment()
                setUserComment("")
                console.log("Comment submitted successfully!")
            } catch (err) {
                console.error("Error submitting comment:", err)
            }
        }
    }

    const upVote = async (stepCommentId: number) => {
        try {
            await upvoteComment(stepCommentId)
            // Optimistically toggle the upVote count
            setStepComments((prevComments) =>
                prevComments?.map((comment) =>
                    comment.stepCommentId === stepCommentId
                        ? {
                              ...comment,
                              upVote: comment.hasUpVoted
                                  ? (comment.upVote ?? 0) - 1 // Decrease if already upvoted
                                  : (comment.upVote ?? 0) + 1, // Increase if not upvoted
                          }
                        : comment
                )
            )

            // Optionally refetch comments if necessary
            // await refetchStepComment();
        } catch (err) {
            console.log("Error toggling upvote: ", err)

            // Optionally revert optimistic update in case of failure
            setStepComments((prevComments) =>
                prevComments?.map((comment) =>
                    comment.stepCommentId === stepCommentId
                        ? {
                              ...comment,
                              upVote: comment.hasUpVoted
                                  ? (comment.upVote ?? 0) + 1 // Revert decrease
                                  : (comment.upVote ?? 0) - 1, // Revert increase
                          }
                        : comment
                )
            )
        }
    }
    console.log(errorStepInfo === null)

    return (
        <Sheet>
            <SheetTrigger asChild>
                <Button variant="outline">Open</Button>
            </SheetTrigger>
            <SheetContent className="w-[95%]">
                <ScrollArea className="h-full">
                    {errorStepInfo !== null ? (
                        <div className="text-2xl w-full h-screen flex justify-center items-center">
                            <div className="flex flex-col items-center">
                                <p className="font-bold">{errorStepInfo}</p>
                                <p>Please try again</p>
                            </div>
                        </div>
                    ) : (
                        <>
                            {isLoadingStepInfo ? (
                                <div className="text-2xl w-full h-screen flex justify-center items-center">
                                    <div className="flex flex-col items-center">
                                        <Loader2 className="animate-spin" />
                                        <p>Loading Step Information</p>
                                    </div>
                                </div>
                            ) : (
                                <>
                                    <SheetHeader>
                                        <div className="relative h-60 overflow-hidden">
                                            <AspectRatio>
                                                <img
                                                    src={stepInfo?.step?.banner}
                                                    alt={stepInfo?.step?.title}
                                                    className="absolute inset-0 w-full h-full object-cover"
                                                />
                                            </AspectRatio>
                                        </div>
                                    </SheetHeader>
                                    {stepInfo && (
                                        <div className="p-20">
                                            <div className="flex justify-between">
                                                <div className="flex items-center gap-2">
                                                    <Blocks color="grey" size={"1rem"} />
                                                    <Label className="uppercase text-stone-500">Step</Label>
                                                </div>
                                                <div className="flex items-center gap-2">
                                                    <Gem color="grey" size={"1rem"} />
                                                    {errorGetGemEachStep !== null ? (
                                                        <Label className="text-red-500">{errorGetGemEachStep}</Label>
                                                    ) : (
                                                        <>
                                                            {isLoadingGetGemEachStep ? (
                                                                <Loader2 className="animate-spin" />
                                                            ) : (
                                                                <Label className="text-stone-500">{`${gemEachStep?.currentGems}/${gemEachStep?.totalGems}`}</Label>
                                                            )}
                                                        </>
                                                    )}
                                                </div>
                                            </div>
                                            <div className="my-4">
                                                <h2 className="text-2xl">{stepInfo?.step?.title}</h2>
                                                <p>{stepInfo?.step?.description}</p>
                                            </div>
                                            <div>
                                                <p className="text-base font-bold">Author(s)</p>
                                                {stepInfo?.authors && (
                                                    <div className="flex flex-col py-3 gap-2">
                                                        {stepInfo?.authors?.map((author) => {
                                                            return (
                                                                <div className="flex flex-row items-center gap-4" key={author.userId}>
                                                                    <Avatar>
                                                                        <AvatarImage
                                                                            src={author.photoUrl}
                                                                            alt={`${author.firstName} ${author.lastName}`}
                                                                        />
                                                                        <AvatarFallback>
                                                                            {getFallbackName(`${author.firstName} ${author.lastName}`)}
                                                                        </AvatarFallback>
                                                                    </Avatar>
                                                                    <p>{`${author.firstName} ${author.lastName}`}</p>
                                                                </div>
                                                            )
                                                        })}
                                                    </div>
                                                )}
                                            </div>
                                            <div>
                                                {stepInfo.userPassed && (
                                                    <>
                                                        <p className="text-base font-bold">People Passed</p>
                                                        <div className="flex relative py-3">
                                                            {stepInfo.userPassed?.map((person, index) => {
                                                                if (index < 5) {
                                                                    return (
                                                                        <div className="relative -ml-3 first:ml-0">
                                                                            <TooltipProvider>
                                                                                <Tooltip>
                                                                                    <TooltipTrigger asChild>
                                                                                        <Avatar>
                                                                                            <AvatarImage
                                                                                                src={person.photoUrl}
                                                                                                alt={`${person.firstName} ${person.lastName}`}
                                                                                            />
                                                                                            <AvatarFallback>
                                                                                                {getFallbackName(
                                                                                                    `${person.firstName} ${person.lastName}`
                                                                                                )}
                                                                                            </AvatarFallback>
                                                                                        </Avatar>
                                                                                    </TooltipTrigger>
                                                                                    <TooltipContent>
                                                                                        <p>{`${person.firstName} ${person.lastName}`}</p>
                                                                                    </TooltipContent>
                                                                                </Tooltip>
                                                                            </TooltipProvider>
                                                                        </div>
                                                                    )
                                                                }
                                                            })}
                                                            {stepInfo.userPassed.length > 5 && (
                                                                <div className="relative -ml-3">
                                                                    <Avatar>
                                                                        <AvatarImage />
                                                                        <AvatarFallback>+{stepInfo.userPassed.length - 5}</AvatarFallback>
                                                                    </Avatar>
                                                                </div>
                                                            )}
                                                            {stepInfo.userPassed.length === 0 && <Label>No one has passed this step yet.</Label>}
                                                        </div>
                                                    </>
                                                )}
                                            </div>
                                            <div>
                                                <div className="h-2/4 w-full bg-white rounded-sm flex flex-col p-3 border-[1px] border-gray-400">
                                                    {stepInfo.step?.content && <MarkdownRenderer content={stepInfo.step?.content} />}
                                                </div>
                                            </div>
                                            <div className="my-4">
                                                <Badge className="bg-badge-outcome text-white gap-1 py-1 hover:bg-badge-outcome">
                                                    <Gem size={"1rem"} />
                                                    <p className="uppercase text-base">outcome</p>
                                                </Badge>
                                                <div className="h-2/4 w-full bg-white rounded-sm flex flex-col p-3 border-[1px] border-gray-400 mt-4">
                                                    {stepInfo.step?.outcome && <MarkdownRenderer content={stepInfo.step?.outcome} />}
                                                </div>
                                            </div>
                                            <div className="my-4">
                                                <Badge className="bg-badge-check text-white gap-1 py-1 hover:bg-badge-check">
                                                    <ShieldQuestion size={"1rem"} />
                                                    <p className="uppercase text-base">check</p>
                                                </Badge>
                                                <div className="h-2/4 w-full bg-white rounded-sm flex flex-col p-3 border-[1px] border-gray-400 mt-4">
                                                    {stepInfo.step?.check && <MarkdownRenderer content={stepInfo.step?.check} />}
                                                </div>
                                            </div>
                                            <div className="my-4">
                                                <Badge className="bg-badge-error text-white gap-1 py-1 hover:bg-badge-error">
                                                    <CircleSlash size={"1rem"} />
                                                    <p className="uppercase text-base">error</p>
                                                </Badge>
                                                <div className="h-2/4 w-full bg-white rounded-sm flex flex-col p-3 border-[1px] border-gray-400 mt-4">
                                                    {stepInfo.step?.error && <MarkdownRenderer content={stepInfo.step?.error} />}
                                                </div>
                                            </div>
                                            <div className="my-4">
                                                <Badge className="bg-badge-comment text-white gap-1 py-1 hover:bg-badge-comment">
                                                    <MessageSquare size={"1rem"} />
                                                    <p className="uppercase text-base">comment</p>
                                                </Badge>
                                                {errorGetStepComment !== null ? (
                                                    <div className="text-red-500">
                                                        <Label> {errorGetStepComment}</Label>
                                                    </div>
                                                ) : (
                                                    <>
                                                        <div className="pt-8">
                                                            {stepComments?.map((cm) => {
                                                                return (
                                                                    <div className="flex items-start pb-6 justify-between">
                                                                        <div className="flex">
                                                                            <Avatar>
                                                                                <AvatarImage
                                                                                    alt={`${cm.userInfo?.firstName} ${cm.userInfo?.lastname}`}
                                                                                    src={cm.userInfo?.photoUrl}
                                                                                />
                                                                                <AvatarFallback>
                                                                                    {getFallbackName(
                                                                                        `${cm.userInfo?.firstName} ${cm.userInfo?.lastname}`
                                                                                    )}
                                                                                </AvatarFallback>
                                                                            </Avatar>
                                                                            <div>
                                                                                <p className="ps-4 font-bold">{`${cm.userInfo?.firstName} ${cm.userInfo?.lastname}`}</p>
                                                                                <p className="ps-4">{cm.comment} </p>
                                                                            </div>
                                                                        </div>
                                                                        <div
                                                                            className="flex items-center cursor-pointer"
                                                                            onClick={() => upVote(cm.stepCommentId ?? 0)}
                                                                        >
                                                                            {/* TODO add upvote comment */}
                                                                            <ArrowBigUp className="text-explore-foreground" />
                                                                            <p className="ps-2 text-explore-foreground">{cm.upVote}</p>
                                                                        </div>
                                                                    </div>
                                                                )
                                                            })}
                                                        </div>

                                                        <div className="flex w-full items-center space-x-2">
                                                            <Input
                                                                type="text"
                                                                placeholder="Writing something..."
                                                                onChange={(e) => setUserComment(e.target.value)}
                                                                value={userComment}
                                                                className="rounded-sm"
                                                            />
                                                            <Button
                                                                className="bg-neutral-950 text-white hover:bg-neutral-800 hover:border-neutral-800 rounded-sm"
                                                                disabled={userComment === "" || isLoadingCommentOnStep}
                                                                onClick={submitComment}
                                                            >
                                                                {isLoadingCommentOnStep && <Loader2 className="animate-spin" />}
                                                                {isLoadingCommentOnStep ? "Please wait" : "Comment"}
                                                            </Button>
                                                        </div>
                                                    </>
                                                )}
                                            </div>
                                            <Separator />
                                            <div className="my-4">
                                                <Badge className="bg-badge-evaluate text-white gap-1 py-1 hover:bg-badge-evaluate">
                                                    <BadgeCheck size={"1rem"} />
                                                    <p className="uppercase text-base">evaluate</p>
                                                </Badge>
                                                {errorGetStepEval !== null ? (
                                                    <div className="text-red-500">
                                                        <Label> {errorGetStepEval}</Label>
                                                    </div>
                                                ) : (
                                                    <>
                                                        {stepEval && (
                                                            <>
                                                                {stepEval.map((el) => {
                                                                    return (
                                                                        <EvalTypeCard
                                                                            stepEvalId={el.stepEvalId}
                                                                            stepId={el.stepId}
                                                                            question={el.question}
                                                                            key={el.stepEvalId}
                                                                            type={el.type}
                                                                            userEval={el.userEval}
                                                                            refetchGetGem={refetchGemEachStep}
                                                                        />
                                                                    )
                                                                })}
                                                            </>
                                                        )}
                                                    </>
                                                )}
                                            </div>
                                        </div>
                                    )}
                                </>
                            )}
                        </>
                    )}
                </ScrollArea>
            </SheetContent>
        </Sheet>
    )
}

export default StepCard

