import logo1 from "@/assets/logo1.png"
import logo2 from "@/assets/logo2.png"
import { SERVER_HOST } from "@/configs/server"
import { ChevronRight } from "lucide-react"
import { toast } from "sonner"

const index = () => {
    const handleRedirect = async () => {
        await toast.promise(
            new Promise((resolve) => {
                setTimeout(() => {
                    resolve("Redirecting now...")
                    window.location.assign(`${SERVER_HOST}/login/redirect`)
                }, 1000) // 3 seconds delay before redirect
            }),
            {
                loading: "Preparing to redirect to the login page...",
                success: () => {
                    return "Successfully prepared for the login page redirect."
                },
            }
        )
    }

    return (
        <div className="flex flex-col w-screen">
            <div className={`flex items-center pl-20 h-[75vh] w-full bg-center bg-cover bg-[url('@/assets/banner3.jpg')] `}>
                <div className="flex flex-col gap-y-3 text-white">
                    <div className="flex flex-row items-center">
                        <img src={logo1} width={"96px"} height={"96px"} />
                        <h1 className="text-6xl ml-4">Bookmark</h1>
                    </div>
                    <div className="relative w-[400px] h-[70px] cursor-pointer" onClick={handleRedirect}>
                        <div className="absolute inset-0 bg-white opacity-20 rounded-[36px]"></div>
                        <div className="flex flex-row justify-center items-center h-full">
                            <h1 className="text-3xl text-center">Login</h1>
                            <ChevronRight className="absolute right-1" size={"50px"} />
                        </div>
                    </div>
                </div>
            </div>
            <div className="flex flex-col items-center justify-center px-20 pt-20">
                <img width={128} height={128} src={logo2} />
                <h1 className="text-5xl mt-6">
                    <span className="text-primary">Bookmark</span> what you learn
                </h1>
                <h2 className="text-5xl">
                    <span className="underline decoration-1 underline-offset-[5px]">not</span> what syllabus declare you to learn
                </h2>
                <div className="flex flex-col gap-y-10 mt-20 w-full px-20 py-14 bg-incompleteStep rounded-md text-center text-white">
                    <div className="flex flex-col gap-y-5">
                        <h3 className="text-4xl font-light">RETHINK WHAT EDUCATION LOOKS LIKE</h3>
                        <p className="text-xl font-medium">
                            ทุกคนน่าจะคุ้นชินกับระบบการศึกษาที่ตีกรอบมาว่าใครควรเรียนอะไรและประเมินผลแบบไหนเป็นอย่างดี
                            แต่คำถามคือทำไมการเรียนรู้ของทุกคนต้องถูกกำหนดตายตัวทั้งๆที่ทุกคนมีความถนัดเป็นของตัวเอง และในหลายครั้งเราต้องแลกเวลา
                            ความรู้สึก แรงบันดาลใจ ไปกับความพยายามเพื่อสอบให้ได้ซึ่งสิ่งที่เรียกว่าเกรด
                            ทั้งๆที่ยังเป็นปริศนาว่าในโลกของการนำความรู้ไปใช้งานจริง การเรียนในรูปแบบนี้มีประสิทธิภาพมากน้อยเพียงใด?
                        </p>
                    </div>
                    <div className="flex flex-col gap-y-5">
                        <h3 className="text-4xl font-light">TIME CHANGES, THINGS CHANGED</h3>
                        <p className="text-xl font-medium">
                            เมื่อเทคโนโลยีในปัจจุบันเปิดโอกาสในทุกคนสามารถค้นหาข้อมูลได้ด้วยตนเองจะเป็นจุดเริ่มต้นของหลักสูตรและวิธีการ
                            สอนแบบใหม่ที่สามารถเปิดโอกาสให้ทุกคนเรียนรู้ตามเรื่องที่ตัวเองสนใจและประเมินผลด้วยตัวเองได้
                            ไม่เหมือนกับการสอนในอดีตที่แหล่งความรู้มีจำกัด นั่นจะทำให้อนาคตของการเรียนและสอบตามหลักสูตรจะหมดไป
                        </p>
                    </div>
                    <div className="flex flex-col gap-y-5">
                        <h3 className="text-4xl font-light">THE REAL PURPOSE OF EDUCATION</h3>
                        <p className="text-xl font-medium">
                            อย่างหนึ่งที่เป็นหลักสำคัญในการเรียนรู้คือการได้ลองผิดลองถูก เพราะประสบการณ์และความเข้าใจนั้นมาจากการได้ทดลองด้วยตนเอง
                            หลายครั้งรูปแบบการสอนและสอบกำหนดว่าต้องทำอะไรให้แบบไหนในเวลาที่จำกัด แต่เราอาจลืมมองในมุมที่ว่าหลายครั้งสิ่งที่ผู้เรียน{" "}
                            ต้องการอาจเป็นเพียงแค่ไกด์สำหรับการทดลองด้วยตัวเองเท่านั้น แต่การตีกรอบเวลากลายเป็นแรงกดดันที่ทำให้การเรียนรู้หายไป
                        </p>
                    </div>
                </div>
                <div className="text-center mt-24 mb-32">
                    <h2 className="text-5xl font-normal">
                        We believe that <span className="font-bold"> ability to fail is ability to learn</span> <br />
                        <span className="font-light"> that’s how</span>
                        <span className="text-primary"> Bookmark </span>
                        <span className="font-light">works</span>
                    </h2>
                </div>
            </div>
        </div>
    )
}

export default index

