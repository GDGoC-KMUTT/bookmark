import { useEffect, useState } from 'react';
import { useAtom } from 'jotai';
import { courseInfoAtom } from '@/stores/course';
import { server } from '@/configs/server';
import courseBg from '../../assets/course_bg.png';
import { Bookmark } from 'lucide-react';

const Hero = () => {
	const [courseInfo, setCourseInfo] = useAtom(courseInfoAtom);  // Use the atom for state
	const [error, setError] = useState<string | null>(null);
	const courseId = 1; // Replace with dynamic courseId if needed

	useEffect(() => {
		const fetchCourseInfo = async () => {
			try {
				const response = await server.course.courseIdInfoList(courseId + "");
				if (response.success) {
					setCourseInfo(response.data);  // Update the Jotai atom with the fetched course data
					// console.log(response.data);
				} else {
					setError('Failed to fetch course info');
				}
			} catch (error) {
				setError('Error fetching course info');
			}
		};

		fetchCourseInfo();
	}, [courseId, setCourseInfo]);  // Dependency on courseId and setCourseInfo

	if (!courseInfo || error) {
		return (
			<div className="relative w-screen h-[480px] bg-cover bg-center" style={{ backgroundImage: `url(${courseBg})` }}>
				<div className="absolute inset-0 bg-black bg-opacity-30 flex flex-col justify-center pl-10">
					<div className="relative w-[500px] space-y-8 pl-10">
						<div className="relative flex flex-row">
							<h1 className="pl-4 text-2xl font-bold text-white font-light">Loading...</h1>
						</div>
						{/* {error && <div className=" pl-4 text-white">{error}</div>} */}
					</div>
				</div>
			</div>
		);
	}

	return (
		<div className="relative w-screen h-[480px] bg-cover bg-center" style={{ backgroundImage: `url(${courseBg})` }}>
			<div className="absolute inset-0 bg-black bg-opacity-30 flex flex-col justify-center pl-10">
				<div className="relative w-[30%] space-y-8 pl-10">
					<div className="relative flex flex-row">
						<Bookmark size={25} color="white" />
						<h1 className="pl-4 text-xl text-white font-light">COURSE / </h1>
						<h1 className="pl-2 text-xl text-white font-light">{courseInfo.field}</h1>
					</div>
					<h1 className="text-4xl text-white font-bold">{courseInfo.name}</h1>
					<button className="w-36 h-15 bg-primary text-fieldType-foreground rounded-lg text-xl">Enroll</button>
				</div>
			</div>
		</div>
	);
};

export default Hero;
