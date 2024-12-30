import courseBg from '../../assets/course_bg.png';
import { Bookmark } from 'lucide-react';
import { useCallback } from 'react';
import { server } from '@/configs/server';
import { toast } from "sonner"
import { PayloadEnrollwithCourse } from "../../api/api"
import { useState, useEffect } from "react";


type HeroProps = {
    courseName: string
    courseField: string
	courseId: number
}

const Hero: React.FC<HeroProps> = ({ courseName, courseField, courseId }) => {
	const [enrolledCourses, setEnrolledCourses] = useState<PayloadEnrollwithCourse[] | undefined>(undefined);
	const [isEnrolled, setIsEnrolled] = useState<boolean>(false);
	const [error, setError] = useState<string | null>(null);

	const fetchEnrolledCourses = useCallback(async () => {
		try {
			// Fetch enrolled courses
			const coursesResponse = await server.courses.getEnrollCourseByUserId();
			if (coursesResponse.code === 200 && coursesResponse.data) {
				setEnrolledCourses(coursesResponse.data);
				const enrolled = coursesResponse.data.some((course: PayloadEnrollwithCourse) => course.courseId === courseId);
				setIsEnrolled(enrolled);
			} else {
				setError("Failed to fetch enrolled courses. Please try again.");
			}
		} catch (error) {
			console.error("Error during data fetching: ", error);
			setError("An error occurred while fetching data. Please try again.");
		}
	}, [courseId]);


	useEffect(() => {
		fetchEnrolledCourses();
	}, [fetchEnrolledCourses]);

	const handleEnroll = useCallback(async () => {
		try {
			const response = await server.enroll.enrollCreate(courseId);
			console.log('Enrollment response:', response);
			if (response) {
				toast.success('Enrolled successfully!');
				// Re-fetch enrolled courses after successful enrollment
				fetchEnrolledCourses();
			} else {
				toast.error('Enrollment failed. Please try again.');
			}
		} catch (error: any) {
			console.error('Enrollment error:', error);

			if (error.response?.status === 409) {
				toast.error('User already enrolled.');
			} else {
				const errorMessage =
					error.response?.data?.message || 'An unexpected error occurred. Please try again.';
				toast.error(errorMessage);
			}
		}
	}, [courseId, fetchEnrolledCourses]);

	return (
		<div className="relative w-screen h-[480px] bg-cover bg-center" style={{ backgroundImage: `url(${courseBg})` }}>
			<div className="absolute inset-0 bg-black bg-opacity-30 flex flex-col justify-center pl-10">
				<div className="relative w-[60%] space-y-8 pl-10">
					<div className="relative flex flex-row">
						<Bookmark size={25} color="white" />
						<h1 className="pl-4 text-xl text-white font-light">COURSE / </h1>
						<h1 className="pl-2 text-xl text-white font-light">{courseField}</h1>
					</div>
					<h1 className="text-4xl text-white font-bold">{courseName}</h1>
					<button
						className={`w-36 h-15 bg-primary text-fieldType-foreground rounded-lg text-xl`}
						onClick={handleEnroll}
					>
						{isEnrolled ? "Enrolled" : "Enroll"}
					</button>
				</div>
			</div>
		</div>
	);
};

export default Hero;
