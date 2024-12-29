import { Blocks, Check, Loader, SquareSquare } from 'lucide-react';
import { PayloadModuleStepResponse } from "../../api/api";


type ModuleProps = {
    moduleTitle: string;
    moduleDescription: string;
    moduleImageUrl: string;
    steps: PayloadModuleStepResponse[];  // Accept an array of steps
};

const Module: React.FC<ModuleProps> = ({ moduleTitle, moduleDescription, moduleImageUrl, steps }) => {
	return (
		<div className="relative sm:w-[70%] xl:w-[800px] bg-white border-2 border-gray-300 shadow-sm rounded-lg p-10 flex flex-col sm:flex-row sm:space-x-6">
		  {/* Green Circle */}
		  <div className="absolute top-20 -left-2 mt-[-10px] ml-[-10px] w-10 h-10 bg-green-500 rounded-full"></div>

		  {/* Module Info */}
		  <div className={`flex flex-col space-y-4 ${moduleImageUrl ? 'sm:w-[60%]' : 'w-full'}`}>
			<div className="flex flex-row">
			  <Blocks color="gray" />
			  <p className="pl-2 text-gray-600 font-normal">MODULE</p>
			</div>
			<div>
			  <h2 className="text-xl font-medium mb-3">{moduleTitle}</h2>
			  <p>{moduleDescription}</p>
			</div>

			{/* Steps Section */}
			<div className="flex flex-row">
			  <SquareSquare color="gray" className="mt-3" />
			  <p className="pl-2 text-gray-600 mt-3 font-normal">STEPS</p>
			</div>
			<ul className="space-y-3 pt-3">
			  {steps.length > 0 ? (
				steps.map((step) => (
				  <li key={step.id} className="flex items-center space-x-2">
					<div className={`flex items-center justify-center w-6 h-6 rounded-full ${step.check === 'true' ? 'bg-green-500' : 'bg-gray-500'}`}>
					  {step.check === 'true' ? <Check className="text-white w-4 h-4" /> : <Loader className="text-white w-4 h-4" />}
					</div>
					<p className="pl-3 text-lg font-normal">{step.title || 'No title available'}</p>
				  </li>
				))
			  ) : (
				<p className="text-gray-500">No steps for this module.</p>
			  )}
			</ul>
		  </div>

		  {/* Image content */}
		  {moduleImageUrl && (
			<div className="mt-4 sm:mt-0">
			  <img src={moduleImageUrl} alt="module image" className="w-[350px] h-[210px] rounded-sm object-cover" />
			</div>
		  )}
		</div>
	);
};

export default Module;
