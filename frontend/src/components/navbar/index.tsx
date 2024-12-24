import React from "react"

const Navbar = () => {
    return (
        <nav className="w-full bg-white h-16 fixed top-0 shadow-lg flex items-center px-10 z-50">
            <div className="flex items-center gap-8">
                <div className="flex items-center">
                    <img src="/src/assets/logo2.png" alt="Logo" className="w-8 h-8 object-contain" />
                </div>
                <div className="flex gap-6 text-sm">
                    <span className="text-gray-700 font-medium">Home</span>
                    <span className="text-gray-400">Explore</span>
                </div>
            </div>
            <div className="flex items-center gap-4 ml-auto">
                <div className="w-8 h-8 rounded-full bg-gray-200"></div>
            </div>
        </nav>
    )
}

export default Navbar

