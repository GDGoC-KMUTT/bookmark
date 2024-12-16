import { Gem, BookMarked } from "lucide-react"

const Navbar = () => {
    return (
        <div className="w-full bg-white h-[3rem] fixed top-0 shadow-md flex items-center px-6 py-3 justify-between">
            <div className="flex items-center space-x-8">
                <img src="src/assets/logo2.png" alt="Logo" className="w-8 h-8" />
                <div className="flex space-x-8">
                    <div className="text-gray-500 font-medium hover:text-explore-foreground transition-colors">
                        Home
                    </div>
                    <div className="text-gray-500 font-medium hover:text-explore-foreground transition-colors">
                        Explore
                    </div>
                </div>
            </div>

            <div className="flex items-center justify-center space-x-6">
                <div className="flex items-center space-x-2">
                    <div className="space-y-1">
                        <div></div>
                        <BookMarked className="text-foreground" size={20} />
                    </div>
                    <div className="items-center justify-center space-y-1">
                        <div className="font text-sm">Connect ESP w...</div>
                        <div className="relative w-24 h-1 bg-border rounded-full">
                            <div
                                className="absolute h-1 bg-progressBar rounded-full"
                                style={{ width: "60%" }} // Replace with dynamic width later
                            ></div>
                        </div>
                    </div>
                </div>

                <div className="flex items-center space-x-1 text-foreground">
                    <Gem size={20} />
                    <span className="font-medium">0</span>
                </div>

                <div className="w-8 h-8 bg-border rounded-full"></div>
            </div>
        </div>
    )
}

export default Navbar

