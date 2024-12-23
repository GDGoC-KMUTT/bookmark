// type CourseCardProps = {
//     courseName: string;
//     status: string;
// };

// const CourseCard = ({ courseName, status }: CourseCardProps) => {
//     const CourseCard = () => {
//     return (
//         <div className="w-full bg-footer h-[3rem] fixed bottom-0 shadow-md flex justify-between items-center px-4">
//             <p className="text-black font-medium bg-red-500">Hello</p>
//             {/* <div className="text-black font-medium">{courseName}</div> */}
//             {/* <div className={`text-sm ${status === 'Completed' ? 'text-green-500' : 'text-yellow-500'}`}>{status}</div> */}
//         </div>
//     );
// }

// export default CourseCard;

const CourseCard = () => {
    return (
        <div className="border p-4 rounded-sm flex flex-col items-start space-y-2">
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="size-6">
                    <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M17.593 3.322c1.1.128 1.907 1.077 1.907 2.185V21L12 17.25 4.5 21V5.507c0-1.108.806-2.057 1.907-2.185a48.507 48.507 0 0 1 11.186 0Z"
                    />
                </svg>
            </div>
            <p className="">COURSE / INFRASTRUCTURE</p>

            <p className="pb-8">Setup LED blink with ESP32 and Arduino IDE</p>
            <img className="" src="https://static.bookmark.scnd.app/asset/fieldicon/microcontroller.png"></img>
        </div>
    )
}

export default CourseCard
