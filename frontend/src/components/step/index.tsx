import { BadgeCheck, Blocks, CircleSlash, Gem, MessageSquare, ShieldQuestion } from "lucide-react"
import { AspectRatio } from "../ui/aspect-ratio"
import { Button } from "../ui/button"
import { Label } from "../ui/label"
import { ScrollArea } from "../ui/scroll-area"
import { Sheet, SheetContent, SheetHeader, SheetTrigger } from "../ui/sheet"
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar"
import { Badge } from "../ui/badge"
import { Separator } from "../ui/separator"

const peoplePassed = [
    {
        avatar: "https://picsum.photos/1920/1080?random",
        name: "test1 gg",
    },
    {
        avatar: "https://github.com/shadcn.png",
        name: "test2 dd",
    },
    {
        avatar: "https://picsum.photos/1920/1080?random",
        name: "test3 hh",
    },
    {
        avatar: "https://github.com/shadcn.png",
        name: "test4 ll",
    },
    {
        avatar: "https://picsum.photos/1920/1080?random",
        name: "test5 pp",
    },
    {
        avatar: "https://github.com/shadcn.png",
        name: "test6 qq",
    },
    {
        avatar: "https://picsum.photos/1920/1080?random",
        name: "test7 rr",
    },
    {
        avatar: "https://picsum.photos/1920/1080?random",
        name: "test7 rr",
    },
]
export function Step() {
    const author = "hello world"
    const getFallBackName = (text: string) => {
        const test = text.split(" ")
        const firsTwo = test.slice(0, 2)
        return firsTwo.map((str) => str.substring(0, 1).toUpperCase())
    }

    return (
        <Sheet>
            <SheetTrigger asChild>
                <Button variant="outline">Open</Button>
            </SheetTrigger>
            <SheetContent className="w-[95%]">
                <ScrollArea className="h-full">
                    <SheetHeader>
                        {/* <SheetTitle>Edit profile</SheetTitle>
                        <SheetDescription>Make changes to your profile here. Click save when you're done.</SheetDescription> */}
                        <div className="h-60 overflow-hidden -z-100">
                            <AspectRatio ratio={16 / 9}>
                                <img src="https://proxy.bsthun.com/raspi/bookmark/Python.jpg" alt="Photo by Drew Beamer" />
                            </AspectRatio>
                        </div>
                    </SheetHeader>
                    <div className="p-6">
                        <div className="flex justify-between">
                            <div className="flex items-center gap-2">
                                <Blocks color="grey" size={"1rem"} />
                                <Label className="uppercase text-stone-500">Step</Label>
                            </div>
                            <div className="flex items-center gap-2">
                                <Gem color="grey" size={"1rem"} />
                                <Label className="text-stone-500">1/2</Label>
                            </div>
                        </div>
                        <div className="my-4">
                            <h2 className="text-2xl">ต่อ ESP32 board เข้ากับคอมพิวเตอร์</h2>
                            <p>
                                คือการจะคุม Microcontroller ได้เนี่ย เราต้องเขียนโปรแกรม ไส่มันผ่าน Arduino IDE ในโมดูลนี้เราจะมาลองตั้งค่าให้ Arduino
                                สามารถแฟลชโปรแกรมลงไปในบอร์ดให้ได้กัน
                            </p>
                        </div>
                        <div>
                            <Label>Author(s)</Label>
                            <div className="flex flex-col py-3 gap-2">
                                <div className="flex flex-row items-center gap-4">
                                    <Avatar>
                                        <AvatarImage src="https://github.com/shadcn.png" />
                                        <AvatarFallback>{getFallBackName(author)}</AvatarFallback>
                                    </Avatar>
                                    <p>{author}</p>
                                </div>
                                <div className="flex flex-row items-center gap-4">
                                    <Avatar>
                                        <AvatarImage src="https://github.com/shadcn.png" />
                                        <AvatarFallback>{getFallBackName(author)}</AvatarFallback>
                                    </Avatar>
                                    <p>{author}</p>
                                </div>
                            </div>
                        </div>
                        <div>
                            <Label>People Passed</Label>
                            <div className="flex relative py-3">
                                {peoplePassed.map((person, index) => {
                                    if (index < 5) {
                                        return (
                                            <div className="relative -ml-3 first:ml-0">
                                                <Avatar>
                                                    <AvatarImage src={person.avatar} alt={person.name} />
                                                    <AvatarFallback>{getFallBackName(person.name)}</AvatarFallback>
                                                </Avatar>
                                            </div>
                                        )
                                    }
                                })}
                                {peoplePassed.length > 5 && (
                                    <div className="relative -ml-3">
                                        <Avatar>
                                            <AvatarImage />
                                            <AvatarFallback>+{peoplePassed.length - 5}</AvatarFallback>
                                        </Avatar>
                                    </div>
                                )}
                            </div>
                        </div>
                        <div>
                            <div className="h-2/4 w-full bg-gray-100 rounded-sm flex flex-col p-3">
                                <Label>content</Label>
                                <Label>content</Label>
                                <Label>content</Label>
                                <Label>content</Label>
                                <Label>content</Label>
                                <Label>content</Label>
                                <Label>content</Label>
                                <Label>content</Label>
                                <Label>content</Label>
                            </div>
                        </div>
                        <div className="my-4">
                            <Badge className="bg-badge-outcome text-white gap-1 py-1">
                                <Gem size={"1rem"} />
                                <p className="uppercase">outcome</p>
                            </Badge>
                        </div>
                        <div className="my-4">
                            <Badge className="bg-badge-check text-white gap-1 py-1">
                                <ShieldQuestion size={"1rem"} />
                                <p className="uppercase">check</p>
                            </Badge>
                        </div>
                        <div className="my-4">
                            <Badge className="bg-badge-error text-white gap-1 py-1">
                                <CircleSlash size={"1rem"} />
                                <p className="uppercase">error</p>
                            </Badge>
                        </div>
                        <div className="my-4">
                            <Badge className="bg-badge-comment text-white gap-1 py-1">
                                <MessageSquare size={"1rem"} />
                                <p className="uppercase">comment</p>
                            </Badge>
                        </div>
                        <Separator />
                        <div className="my-4">
                            <Badge className="bg-badge-evaluate text-white gap-1 py-1">
                                <BadgeCheck size={"1rem"} />
                                <p className="uppercase">evaluate</p>
                            </Badge>
                        </div>
                    </div>
                </ScrollArea>
            </SheetContent>
        </Sheet>
    )
}

