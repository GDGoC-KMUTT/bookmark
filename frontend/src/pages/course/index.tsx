import { useEffect, useState } from "react";
import { useAtom } from "jotai";
import { courseInfoAtom, courseContentAtom } from '@/stores/course';
import { moduleAtom } from '@/stores/module';
import Hero from "../../components/course/hero";
import Text from "../../components/course/text";
import Module from "../../components/course/module";
import { server } from '@/configs/server';

const Course = () => {
	const [courseInfo, setCourseInfo] = useAtom(courseInfoAtom); // Global state for course data
	const [courseContent, setCourseContent] = useAtom(courseContentAtom); // Global state for course content
	const [modules, setModules] = useAtom(moduleAtom); // Global state for module data
	const [error, setError] = useState<string | null>(null);

	const courseId = 1; // Replace with dynamic courseId if needed

	useEffect(() => {
		const fetchCourseInfo = async () => {
			try {
				const response = await server.course.getCourseInfo(courseId.toString());
				// console.log(response);

				if (response.data) {
					const courseData = response.data;
					// console.log(courseData.field);

					// Transform courseData to match the required shape
					const transformedCourse = {
						course_id: courseData.id ?? 0,
						name: courseData.name ?? '',
						field: courseData.field ?? '',
					};

					setCourseInfo(transformedCourse); // Store course info in the global atom
					// console.log(courseInfo);
				} else {
					setError("Failed to fetch course info");
				}
			} catch (err) {
				setError("Error fetching course info");
			}
		};

		fetchCourseInfo();
	}, [courseId, setCourseInfo]);

	useEffect(() => {
		const fetchCourseContent = async () => {
			try {
				const response = await server.course.getCourseContent(courseId.toString());
				if (response.data) {
					const courseContentData = response.data;

					// Transform courseContentData to match the required atom structure
					const transformedCourseContent = courseContentData.map((content: any) => ({
						course_id: content.course_id ?? 0,
						order: content.order ?? 0,
						type: content.type ?? '',
						text: content.text ?? '',
						module_id: content.module_id ?? 0,
					}));

					setCourseContent(transformedCourseContent); // Store course content in the global atom
				} else {
					setError("Failed to fetch course content");
				}
			} catch (err) {
				setError("Error fetching course content");
			}
		};

		fetchCourseContent();
	}, [courseId, setCourseContent]);

	useEffect(() => {
		const fetchModuleInfo = async (moduleId: number) => {
			try {
				const response = await server.module.getModuleInfo(moduleId.toString());
				if (response.data) {
					const moduleData = response.data;

					// Ensure the data matches the atom structure
					const transformedModule = {
						module_id: moduleData.id ?? 0, // Assuming `id` is the correct key in the response
						title: moduleData.title ?? '',
						description: moduleData.description ?? '',
						image_url: moduleData.image_url ?? '',
					};

					setModules((prev) => [
						...prev.filter((module) => module.module_id !== moduleId), // Avoid duplicates
						transformedModule, // Add the new module data
					]);
				} else {
					console.error(`Failed to fetch module info for ID: ${moduleId}`);
				}
			} catch (err) {
				console.error(`Error fetching module info for ID: ${moduleId}`, err);
			}
		};

		// Fetch module data for all modules in the course content
		if (courseContent && Array.isArray(courseContent)) {
			courseContent
				.filter((item) => item.type === "module")
				.forEach((module) => {
					if (!modules.some((mod) => mod.module_id === module.module_id)) {
						fetchModuleInfo(module.module_id);
					}
				});
		}
	}, [courseContent, modules, setModules]);


	if (error) {
		return <div className="error">{error}</div>;
	}

	return (
		<div className="relative w-full mt-[50px] mb-[150px] flex flex-col overflow-y-auto">
			<Hero key={courseId} courseField={courseInfo.field} courseName={courseInfo.name} />
			<div className="w-full flex flex-col items-center justify-center space-y-10 mt-20">
				{courseContent &&
					courseContent.map((item, index) => {
						if (item.type === "text") {
							return <Text key={index} content={item.text} />;
						} else if (item.type === "module") {
							const moduleData = modules.find((mod) => mod.module_id === item.module_id);
							if (moduleData) {
								return (
									<Module
										key={index}
										title={moduleData.title || ''}
										description={moduleData.description || ''}
										imageUrl={moduleData.image_url || ''}
									/>
								);
							}
						}
						return null;
					})}
				</div>
		</div>
	);
};

export default Course;
