import { useEffect, useState } from "react";
import { useAtom } from "jotai";
import { courseInfoAtom, courseContentAtom} from '@/stores/course';
import { moduleAtom } from '@/stores/module';
import { moduleStepsAtom } from '@/stores/moduleStep';
import Hero from "../../components/course/hero";
import Text from "../../components/course/text";
import Module from "../../components/course/module";
import { server } from '@/configs/server';

const Course = () => {
	const [courseInfo, setCourseInfo] = useAtom(courseInfoAtom); // Global state for course data
	const [courseContent, setCourseContent] = useAtom(courseContentAtom); // Global state for course content
	const [modules, setModules] = useAtom(moduleAtom); // Global state for module data
	const [moduleSteps, setModuleSteps] = useAtom(moduleStepsAtom);
	const [error, setError] = useState<string | null>(null);

	//! don't forget to fix this later
	const courseId = 1; // Replace with dynamic courseId if needed

	// Fetch course information
	useEffect(() => {
		const fetchCourseInfo = async () => {
			try {
				const response = await server.coursePage.getCoursePageInfo(courseId.toString());

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

	 // Fetch course content
	useEffect(() => {
		const fetchCourseContent = async () => {
			try {
				const response = await server.coursePage.getCoursePageContent(courseId.toString());
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

	// Fetch module info
	useEffect(() => {
		const fetchModuleInfo = async (moduleId: number) => {
			try {
			  if (modules.some((mod) => mod.module_id === moduleId)) {
				console.log(`Module ID: ${moduleId} already fetched.`);
				return;
			  }

			  const response = await server.module.getModuleInfo(moduleId.toString());
				console.log(response.data);

			  if (response.data) {
				const transformedModule = {
				  module_id: response.data.id ?? 0,
				  title: response.data.title ?? '',
				  description: response.data.description ?? '',
				  image_url: response.data.image_url ?? '',
				};

				setModules((prev) => [...prev, transformedModule]);
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


	// Fetch steps for each module
	useEffect(() => {
		const fetchModuleSteps = async (moduleId: number) => {
			try {
			  if (moduleSteps[moduleId]) {
				console.log(`Steps for Module ID: ${moduleId} already fetched.`);
				return;
			  }

			  const response = await server.moduleStep.getModuleSteps(moduleId.toString());
			  if (response.data && Array.isArray(response.data)) {
				const transformedSteps = response.data.map((step: any) => ({
				  step_id: step.id ?? 0,
				  title: step.title ?? '',
				  check: step.check ?? '',
				}));

				setModuleSteps((prev) => ({ ...prev, [moduleId]: transformedSteps }));
			  } else {
				console.log(`No step for Module ID: ${moduleId}`, response);
			  }
			} catch (err) {
			  console.error(`Error fetching steps for Module ID: ${moduleId}`, err);
			}
		  };



		// Fetch steps for all modules in the course content
		if (courseContent && Array.isArray(courseContent)) {
		  courseContent
			.filter((item) => item.type === "module")
			.forEach((module) => {
			  if (!moduleSteps[module.module_id]) {
				fetchModuleSteps(module.module_id);
			  }
			});
		}
	  }, [courseContent, moduleSteps, setModuleSteps]);


	  if (error) {
		return <div className="error">{error}</div>;
	  }

	return (
		<div className="relative w-full mt-[50px] mb-[150px] flex flex-col overflow-y-auto">
			<Hero key={courseId}/>
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
										moduleId={moduleData.module_id || 0}
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
