// import { course } from "@/types/course";

const index = () => {
    return (
        <div className="flex justify-center items-center w-screen">
            <div className="text-center">
                <h1 className="text-4xl font-semibold text-orange-600">Welcome to Bookmark</h1>
                <div className="mt-6">
                    <button
                        className="bg-primary text-white rounded-md"
                        onClick={() => alert("Going to Bookmarks Page!")} // Placeholder functionality
                    >
                        Get Start!
                    </button>
                </div>
            </div>
        </div>
    )
}

export default index

