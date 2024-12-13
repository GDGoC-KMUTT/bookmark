// import { course } from "@/types/course";

const index = () => {
    return (
        <div className="flex justify-center items-center w-screen">
            <div className="text-center">
                <h1 className="text-4xl font-semibold text-orange-600">Welcome to Bookmark</h1>
                <div className="mt-6">
                    <button
                        className="bg-orange-500 text-white py-2 px-4 rounded-lg hover:bg-orange-600"
                        onClick={() => alert("Going to Bookmarks Page!")} // Placeholder functionality
                    >
                        Get Started
                    </button>
                </div>
            </div>
        </div>
    )
}

export default index

