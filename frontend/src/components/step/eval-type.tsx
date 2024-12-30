import { FilePen, Loader2 } from "lucide-react"
import React, { FC, useEffect, useState } from "react"
import { Label } from "@/components/ui/label"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import ResultCard from "./result-card"
import { PayloadUserEvalResult } from "@/api/api"
import { useSubmitStepEval } from "@/hooks/useSubmitStepEval"
import { Dialog, DialogContent, DialogTrigger } from "@/components/ui/dialog"
import { useSubmitStepEvalTypeCheck } from "@/hooks/useSubmitStepEvalTypeCheck"
import { useCheckStepEvalStatus } from "@/hooks/useCheckStepEvalStatus"

type EvalTypeCardProps = {
    stepId: number | undefined
    stepEvalId: number | undefined
    type: string | undefined
    question: string | undefined
    userEval: PayloadUserEvalResult | undefined
    refetchGetGem: () => Promise<void>
}

const EvalTypeCard: FC<EvalTypeCardProps> = ({ stepId, stepEvalId, type, question, userEval, refetchGetGem }) => {
    const [answer, setAnswer] = useState<string>("")
    const [file, setFile] = useState<File | null>(null)
    const [isChecked, setIsChecked] = useState<boolean>(false)
    const [isSubmit, setIsSubmit] = useState<boolean>(false)

    const { createUserEvalRes, isLoading: isLoadingSubmitEval, error: errorSubmitEval, submitStepEval } = useSubmitStepEval()
    const { createUserEvalIdResCheck, error: errorSubmitTypeCheck, submitStepEvalTypeCheck } = useSubmitStepEvalTypeCheck()
    const { userEvalStatus, fetchStepEvalStatus } = useCheckStepEvalStatus()

    // Handle file input change
    const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (e.target.files && e.target.files[0]) {
            setFile(e.target.files[0]) // Store the selected file
        }
    }

    const handleCheckboxChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setIsChecked(event.target.checked) // Access the checked state
        if (stepEvalId) submitStepEvalTypeCheck(stepEvalId)
    }

    // Submit the file to the API
    const handleSubmit = async () => {
        if (type && stepId && stepEvalId) {
            if (answer !== "") {
                await submitStepEval(stepId, stepEvalId, type, undefined, answer)
                setIsSubmit(true)
            } else if (file) {
                await submitStepEval(stepId, stepEvalId, type, file)
                setIsSubmit(true)
            }
        }
    }

    const userEvalId = createUserEvalRes?.userEvalId || createUserEvalIdResCheck?.userEvalId || userEval?.userEvalId
    const userSubmission = createUserEvalRes?.userSubmission || userEval?.content
    const stepEvalPassed = isSubmit || userEval?.pass === null
    const passStatus = userEvalStatus?.pass ?? userEval?.pass
    const commentStatus = userEvalStatus?.comment ?? userEval?.comment
    const checkPassAndCommentStatus = passStatus !== null && passStatus !== undefined && commentStatus !== null && commentStatus !== undefined

    useEffect(() => {
        setFile(null)
        setAnswer("")
        setIsChecked(false)
        setIsSubmit(false)
    }, [])

    useEffect(() => {
        if (userEval?.userEvalId != null && userEval.content != null) {
            setAnswer(userEval.content)
        }
    }, [userEval?.content, userEval?.userEvalId])

    useEffect(() => {
        if (stepEvalPassed && userEvalId) {
            const intervalId = setInterval(() => {
                fetchStepEvalStatus(userEvalId)
            }, 5000) // Check status every 5 seconds

            // Stop checking once status is received
            if (userEvalStatus?.pass !== null && userEvalStatus?.pass !== undefined) {
                refetchGetGem()
                clearInterval(intervalId)
                setIsSubmit(false) // Reset submission state
            }

            return () => clearInterval(intervalId) // Cleanup on unmount
        }
    }, [isSubmit, userEvalStatus, stepEvalPassed, userEvalId, fetchStepEvalStatus, refetchGetGem])

    return (
        <>
            <div className="flex gap-2 pt-8">
                <FilePen />
                <p>{question}</p>
            </div>
            {type === "check" && (
                <div className="ps-9">
                    <div className="flex justify-start space-x-2 pt-4">
                        <input
                            id="checkbox"
                            type="checkbox"
                            className="cursor-pointer"
                            onChange={handleCheckboxChange}
                            checked={!!userEvalId || isChecked}
                            disabled={!!userEvalId || isChecked}
                        />
                        <Label className="cursor-pointer" htmlFor="checkbox">
                            Mark as completed
                        </Label>
                    </div>
                    {errorSubmitTypeCheck !== null ? (
                        <Label className="text-red-500">{errorSubmitTypeCheck}</Label>
                    ) : (
                        <>{userEvalId && <p className="text-xs pt-2">submission Id: {userEvalId}</p>}</>
                    )}
                </div>
            )}

            {type === "image" && (
                <div className="ps-9">
                    {!userEvalId && (
                        <div className="flex pe-7 py-3 items-end space-x-2">
                            <div className=" w-[25rem]">
                                <Label htmlFor="picture">Please upload image</Label>
                                <Input id="picture" type="file" className="rounded-sm" onChange={handleFileChange} />
                            </div>
                            <Button
                                type="submit"
                                className="bg-neutral-950 text-white hover:bg-neutral-800 hover:border-neutral-800 rounded-sm"
                                disabled={!file}
                                onClick={handleSubmit}
                            >
                                {isLoadingSubmitEval ? <Loader2 className="animate-spin" /> : <></>}
                                {isLoadingSubmitEval ? "Submitting..." : "Submit"}
                            </Button>
                        </div>
                    )}
                    {userEvalId && userSubmission && (
                        <Dialog>
                            <DialogTrigger asChild>
                                <div className="w-[450px] cursor-pointer pt-4 flex items-center justify-center overflow-hidden h-auto">
                                    <img src={userSubmission} alt={`submissionId:${userEvalId}`} className="max-w-full max-h-full object-contain" />
                                </div>
                            </DialogTrigger>
                            <DialogContent className="flex items-center justify-center overflow-hidden">
                                <img src={userSubmission} alt={`submissionId:${userEvalId}`} className="max-w-full max-h-[90vh] object-contain" />
                            </DialogContent>
                        </Dialog>
                    )}
                    {errorSubmitEval !== null ? (
                        <Label className="text-red-500">{errorSubmitEval}</Label>
                    ) : (
                        <>{userEvalId && <p className="text-xs pt-2">submission Id: {userEvalId}</p>}</>
                    )}
                    {stepEvalPassed && passStatus === null && commentStatus === null && (
                        <div className="flex space-x-2 pt-3">
                            <Loader2 className="animate-spin" />
                            <p>Please wait, this may take a few minutes as we check your answer</p>
                        </div>
                    )}
                </div>
            )}
            {type === "text" && (
                <div className="ps-9">
                    <div className="flex space-x-2">
                        <Input
                            type="text"
                            placeholder="write your answer..."
                            className="rounded-sm"
                            value={answer}
                            disabled={!!userEvalId}
                            onChange={(e) => setAnswer(e.target.value)}
                        />
                        <Button
                            type="submit"
                            className="bg-neutral-950 text-white hover:bg-neutral-800 hover:border-neutral-800 rounded-sm"
                            disabled={userEval?.content !== null || userEval?.userEvalId !== null}
                            onClick={handleSubmit}
                        >
                            Submit
                        </Button>
                    </div>
                    {errorSubmitEval !== null ? (
                        <Label className="text-red-500">{errorSubmitEval}</Label>
                    ) : (
                        <>{userEvalId && <p className="text-xs pt-2">submission Id: {userEvalId}</p>}</>
                    )}
                    {stepEvalPassed && passStatus === null && commentStatus === null && (
                        <div className="flex space-x-2 pt-3">
                            <Loader2 className="animate-spin" />
                            <p>Please wait, this may take a few minutes as we check your answer</p>
                        </div>
                    )}
                </div>
            )}

            {type !== "check" && checkPassAndCommentStatus && <ResultCard pass={passStatus} comment={commentStatus} />}
        </>
    )
}

export default EvalTypeCard

