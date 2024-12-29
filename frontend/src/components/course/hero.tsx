import courseBg from '../../assets/course_bg.png';
import { Bookmark } from 'lucide-react';
import { useAtom } from 'jotai';
import { courseInfoAtom } from '@/stores/course';
import { useCallback } from 'react';
import { server } from '@/configs/server';
import { toast } from "sonner"

const Hero = () => {
	const [courseInfo] = useAtom(courseInfoAtom);
	const courseName = courseInfo.name || '';
	const courseField = courseInfo.field || '';
	const courseId = courseInfo.course_id || 0; // Ensure `course_id` is present

	// Placeholder for authenticated user ID
	const userId = 1; // Replace with actual user ID from your authentication system

	const handleEnroll = useCallback(async () => {
		try {
			const response = await server.enroll.enrollCreate(userId, courseId);

			if (response) {
				toast.success('Enrolled successfully!');
			} else {
				toast.error('Enrollment failed. Please try again.');
			}
		} catch (error: any) {
			// Log detailed error for debugging
			console.error('Enrollment error:', error);

			// Handle specific error code for 409 Conflict
			if (error.response?.status === 409) {
				toast.error('User already enrolled.');
			} else {
				const errorMessage =
					error.response?.data?.message || 'An unexpected error occurred. Please try again.';
				toast.error(errorMessage);
			}
		}
	}, [userId, courseId]);


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
						className="w-36 h-15 bg-primary text-fieldType-foreground rounded-lg text-xl"
						onClick={handleEnroll}
					>
						Enroll
					</button>
				</div>
			</div>
		</div>
	);
};

export default Hero;
