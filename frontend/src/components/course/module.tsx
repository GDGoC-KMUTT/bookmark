import { Blocks, Check, Loader, SquareSquare } from 'lucide-react';
import { PayloadModuleStepResponse } from "../../api/api";
// import StepCard from "@/components/step/index";
import { useState } from "react";


type ModuleProps = {
	moduleTitle: string;
	moduleDescription: string;
	moduleImageUrl: string;
	steps: PayloadModuleStepResponse[];
};

const Module: React.FC<ModuleProps> = ({
	moduleTitle,
	moduleDescription,
	moduleImageUrl,
	steps,
}) => {
	const [activeStep, setActiveStep] = useState<{
		stepId: number;
		title: string;
		index: number;
	} | null>(null);

	return (
		<div className="relative sm:w-[70%] xl:w-[700px] bg-white border-2 border-gray-300 shadow-sm rounded-lg p-10 flex flex-col space-y-3 mb-10">
			{/* Green Circle */}
			<div className="absolute top-20 -left-2 mt-[-10px] ml-[-10px] w-10 h-10 bg-green-500 rounded-full"></div>

			{/* Module Info and Image */}
			<div className="flex flex-col sm:flex-row sm:space-x-6">
				{/* Module Info */}
				<div className="flex flex-col space-y-4 w-full sm:w-2/3">
					<div className="flex flex-row">
						<Blocks color="gray" />
						<p className="pl-2 text-gray-600 font-normal">MODULE</p>
					</div>
					<div>
						<h2 className="text-xl font-medium mb-3">{moduleTitle}</h2>
						<p>{moduleDescription}</p>
					</div>
				</div>

				{/* Image content */}
				{moduleImageUrl && (
					<div className="w-full sm:w-1/3 sm:flex sm:justify-center">
						<img
							src={moduleImageUrl}
							alt="module image"
							className="w-[200px] h-[120px] rounded-sm object-cover"
						/>
					</div>
				)}
			</div>

			{/* Steps Section */}
			<div className="flex flex-row">
				<SquareSquare color="gray" className="" />
				<p className="pl-2 text-gray-600 font-normal">STEPS</p>
			</div>
			<ul className="space-y-3">
				{steps.length > 0 ? (
					steps.map((step, index) => (
						<li
							key={step.id}
							className="flex items-center space-x-2 cursor-pointer"
							onClick={() => {
								if (step.id !== undefined) {
									setActiveStep({
										stepId: step.id,
										title: step.title || "No title available",
										index: index + 1,
									});
								} else {
									console.error("Step ID is undefined. Cannot activate this step.");
								}
							}}

						>
							<div
								className={`flex items-center justify-center w-6 h-6 rounded-full ${step.check === "true" ? "bg-green-500" : "bg-gray-500"
									}`}
							>
								{step.check === "true" ? (
									<Check className="text-white w-4 h-4" />
								) : (
									<Loader className="text-white w-4 h-4" />
								)}
							</div>
							<p className="pl-3 text-lg font-normal">
								{step.title || "No title available"}
							</p>
						</li>
					))
				) : (
					<p className="text-gray-500">No steps for this module.</p>
				)}
			</ul>

			{/* Render StepCard Conditionally */}
			{activeStep && (
				<div
					className="fixed inset-0 z-50 flex"
					onClick={() => setActiveStep(null)} // Close StepCard when clicking outside
				>
					{/* Remaining space outside StepCard */}
					<div className="flex-1 bg-black bg-opacity-50" />

					{/* StepCard container */}
					<div
						className="w-[85%] h-full bg-white shadow-lg"
						onClick={(e) => e.stopPropagation()} // Prevent click inside StepCard from triggering the outer onClick
					>
						{/* <StepCard
							stepId={activeStep.stepId} // Used to fetch step-specific data or perform actions
							title={activeStep.title}  // Step title
							index={activeStep.index}  // UI representation of step's order
						/> */}
					</div>
				</div>
			)}

		</div>
	);
};

export default Module;
