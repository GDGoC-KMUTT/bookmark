import StepCard from "@/components/step"
import { ModuleInfo, StepInfo } from "@/types/step"

const modules: ModuleInfo[] = [
    {
        moduleId: 1,
        name: " Setup board environment",
    },
    {
        moduleId: 2,
        name: "รู้จักกับ serial console",
    },
    {
        moduleId: 3,

        name: "ทำบอร์ดให้ควบคุม actuators ภายนอก",
    },
    {
        moduleId: 4,
        name: "ทำบอร์ดให้ควบคุมผ่าน network ได้",
    },
    {
        moduleId: 5,
        name: "Quest: ลองสร้าง LED switcher API กัน",
    },
]

const steps: StepInfo[] = [
    { stepId: 1, moduleId: 1, title: "เชื่อมต่อ board เข้ากับคอมพิวเตอร์" },
    { stepId: 2, moduleId: 1, title: "ติดตั้ง Arduino IDE" },
    { stepId: 3, moduleId: 1, title: "ติดตั้ง ESP32 Driver" },
    { stepId: 4, moduleId: 1, title: "ลอง upload sketch ขึ้นไปบน board" },
    { stepId: 5, moduleId: 1, title: "ลองเปลี่ยนตัวแปรดูซักหน่อย" },
    { stepId: 6, moduleId: 1, title: "ทำความเข้าใจ Serial Console" },
    { stepId: 7, moduleId: 2, title: "การตั้งค่า Serial Monitor" },
    { stepId: 8, moduleId: 2, title: "การส่งข้อมูลแบบซับซ้อน" },
    { stepId: 9, moduleId: 2, title: "ใช้ snprintf กับ Serial Console" },
    { stepId: 10, moduleId: 2, title: "ควบคุม On-board LED ผ่าน Serial Console" },
    { stepId: 11, moduleId: 3, title: "รู้จักกับ GPIO PIN" },
    { stepId: 12, moduleId: 3, title: "รู้จักกับ Breadboard" },
    { stepId: 13, moduleId: 3, title: "เชื่อมต่อ LED เข้ากับ GPIO" },
    { stepId: 14, moduleId: 3, title: "เขียนโค้ดควบคุม LED" },
    { stepId: 15, moduleId: 4, title: "เขียนโค๊ตให้เชื่อมต่อ Wi-Fi ตอนบูต" },
    { stepId: 16, moduleId: 4, title: "สร้าง HTTP Server พร้อม Ping Pong Endpoint" },
    { stepId: 17, moduleId: 4, title: "สร้าง HTTP Server ให้" },
    { stepId: 18, moduleId: 5, title: "ควบคุม on-board LED ด้วย HTTP Call" },
    { stepId: 19, moduleId: 5, title: "ลองสร้าง LED switcher" },
    { stepId: 20, moduleId: 5, title: "ลองสร้าง LED status API" },
]

const Home = () => {
    return (
        <div>
            <div className="ps-10">
                {modules.map((module) => (
                    <div key={module.moduleId} className="py-4">
                        <h2>
                            Module {module.moduleId}: {module.name}
                        </h2>
                        <div>
                            {steps
                                .filter((step) => step.moduleId === module.moduleId)
                                .map((step, index) => (
                                    <StepCard key={step.stepId} stepId={step.stepId} title={step.title} check={false} />
                                ))}
                        </div>
                    </div>
                ))}
            </div>
        </div>
    )
}

export default Home

