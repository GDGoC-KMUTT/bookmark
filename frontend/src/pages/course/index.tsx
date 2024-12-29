import { useEffect, useState } from "react";
import Hero from "../../components/course/hero";
import Text from "../../components/course/text";
import Module from "../../components/course/module";
import { server } from "@/configs/server";
import CourseCard from "../../components/coursecard/index";
import { PayloadModuleResponse, PayloadModuleStepResponse, PayloadCoursePage, PayloadCoursePageContent, PayloadSuggestCourse } from "../../api/api";

const Course = () => {
	const [courseInfo, setCourseInfo] = useState<PayloadCoursePage | null>(null);
	const [courseContent, setCourseContent] = useState<PayloadCoursePageContent[] | null>(null);
	const [modules, setModules] = useState<PayloadModuleResponse[]>([]);
	const [moduleSteps, setModuleSteps] = useState<Record<number, PayloadModuleStepResponse[]>>({});
	const [suggestCourses, setSuggestCourses] = useState<PayloadSuggestCourse[] | undefined>(undefined)
	const [error, setError] = useState<string | null>(null);

	//! Don't forget to replace with dynamic courseId if needed
	const courseId = 1;

	useEffect(() => {
		const fetchData = async () => {
			try {
				//* Fetch course information
				const courseInfoResponse = await server.coursePage.getCoursePageInfo(courseId.toString());
				if (courseInfoResponse.data) {
					setCourseInfo(courseInfoResponse.data);
				} else {
					setError("Failed to fetch course info.");
					return;
				}

				//* Fetch course content
				const courseContentResponse = await server.coursePage.getCoursePageContent(courseId.toString());
				if (courseContentResponse.data) {
					// console.log("Fetched course content:", courseContentResponse.data);

					// Transform the data to ensure it matches the interface
					const contentData = courseContentResponse.data.map((item: PayloadCoursePageContent) => ({
						coursePageId: item.coursePageId,
						moduleId: item.moduleId,
						order: item.order,
						text: item.text || '', // Default empty text if undefined
						type: item.type,
					}));
					setCourseContent(contentData);

					//* Fetch modules and their steps
					const modulesToFetch = contentData.filter((item) => item.type === "module");

					const modulePromises = modulesToFetch.map(async (module) => {
						const moduleResponse = await server.module.getModuleInfo(module.moduleId?.toString() || "");
						if (moduleResponse?.data) {
							return {
								module: moduleResponse.data,
								moduleId: module.moduleId,
							};
						}
						return null;
					});

					const fetchedModules = await Promise.all(modulePromises);
					const validModules = fetchedModules.filter((m): m is { module: PayloadModuleResponse; moduleId: number } => !!m);
					// console.log("Valid modules:", validModules);
					setModules(validModules.map((m) => m.module));

					// Fetch steps for each module
					const stepsPromises = validModules.map(async ({ moduleId }) => {
						const stepsResponse = await server.moduleStep.getModuleSteps(moduleId.toString());
						return { moduleId, steps: stepsResponse.data || [] };
					});

					const fetchedSteps = await Promise.all(stepsPromises);
					// console.log("fetchedSteps:", fetchedSteps);
					const stepsMap = fetchedSteps.reduce(
						(acc, { moduleId, steps }) => ({ ...acc, [moduleId]: steps }),
						{} as Record<number, PayloadModuleStepResponse[]>
					);
					// console.log("stepsMap:", stepsMap);
					setModuleSteps(stepsMap);
				} else {
					setError("Failed to fetch course content.");
				}

				//* Fetch suggest courses
				if (courseInfoResponse.data?.fieldId) {
					const suggestResponse = await server.coursePage.getSuggestCoursesByFieldId(courseInfoResponse.data.fieldId.toString());
					// console.log("Suggest courses:", suggestResponse);
					if (suggestResponse.data) {
						setSuggestCourses(suggestResponse.data);
					} else {
						setError("Failed to fetch suggest courses. Please try again.");
					}
				}
			} catch (err) {
				console.error("Error fetching data:", err);
				setError("An error occurred while fetching data. Please try again.");
			}
		};

		fetchData();
	}, [courseId]);

	if (error) {
		return <div className="error">{error}</div>;
	}

	return (
		<div className="relative w-full mt-[50px] mb-[150px] flex flex-col overflow-x-hidden overflow-y-auto">
			<Hero
				key={courseId}
				courseName={courseInfo?.name ?? ""}
				courseField={courseInfo?.field ?? ""}
				courseId={courseInfo?.id ?? 0}
			/>
			<div className="w-full flex flex-col items-center justify-center space-y-10 mt-20">
				{courseContent &&
					courseContent.map((item, index) => {
						if (item.type === "text") {
							return <Text key={index} content={item.text || ''} backgroundIndex={index} />;
						} else if (item.type === "module") {
							const moduleData = modules.find((mod) => mod.id === item.moduleId);

							if (moduleData) {
								return (
									<Module
										key={index}
										moduleTitle={moduleData.title || "Untitled Module"}
										moduleDescription={moduleData.description || "No description available."}
										moduleImageUrl={moduleData.imageUrl || "/default-image.png"}
										steps={moduleSteps[item.moduleId || 0] || []}
									/>
								);
							}
						}
						return null;
					})}

			</div>

			{Array.isArray(suggestCourses) && suggestCourses.length > 0 && (
				<>
					<div className="mt-20"></div>
					<p className="mt-20 text-2xl font-medium flex justify-center">What's next?</p>
					<div className="w-full flex justify-center mt-20">
						<div className="w-full max-w-[95%] flex justify-center ">
							<div className="overflow-x-auto overflow-y-hidden scrollbar-hide">
								<div className="flex justify-start">
									{suggestCourses.map((course, index) => (
										<div key={index} className="mr-10 last:-mr-20">
											<CourseCard
												courseName={course.name || "Untitled Course"}
												fieldName={course.fieldName || "Unknown Field"}
												imageUrl={course.fieldImageUrl || "/default-image.png"}
												courseId={course.id || 0}
											/>
										</div>
									))}
								</div>
							</div>
						</div>
					</div>
				</>
			)}

		</div>
	);
};

export default Course;
