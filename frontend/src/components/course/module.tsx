import { Blocks, Check, Loader, SquareSquare } from 'lucide-react';
import { useAtom } from 'jotai';
import { moduleAtom } from '@/stores/module';
import { moduleStepsAtom } from '@/stores/moduleStep';

const Module = ({ moduleId }: { moduleId: number }) => {
	const [module] = useAtom(moduleAtom);
	const moduleData = module.find((mod) => mod.module_id === moduleId);
	const [moduleSteps] = useAtom(moduleStepsAtom);
	const steps = moduleSteps[moduleId] || [];

	if (!moduleData) {
		return <div>Loading module data...</div>;
	}

	return (
		<div className="relative sm:w-[70%] xl:w-[800px] bg-white border-2 border-gray-300 shadow-sm rounded-lg p-10 flex flex-col sm:flex-row sm:space-x-6">
		  {/* Green Circle */}
		  <div className="absolute top-20 -left-2 mt-[-10px] ml-[-10px] w-10 h-10 bg-green-500 rounded-full"></div>

		  {/* Module Info */}
		  <div className={`flex flex-col space-y-4 ${moduleData.image_url ? 'sm:w-[60%]' : 'w-full'}`}>
			<div className="flex flex-row">
			  <Blocks color="gray" />
			  <p className="pl-2 text-gray-600 font-normal">MODULE</p>
			</div>
			<div>
			  <h2 className="text-xl font-medium mb-3">{moduleData.title}</h2>
			  <p>{moduleData.description}</p>
			</div>

			{/* Steps Section */}
			<div className="flex flex-row">
			  <SquareSquare color="gray" className="mt-3" />
			  <p className="pl-2 text-gray-600 mt-3 font-normal">STEPS</p>
			</div>
			<ul className="space-y-3 pt-3">
			  {steps.length > 0 ? (
				steps.map((step) => (
				  <li key={step.step_id} className="flex items-center space-x-2">
					<div className={`flex items-center justify-center w-6 h-6 rounded-full ${step.check === 'true' ? 'bg-green-500' : 'bg-gray-500'}`}>
					  {step.check === 'true' ? <Check className="text-white w-4 h-4" /> : <Loader className="text-white w-4 h-4" />}
					</div>
					<p className="pl-3 text-lg font-normal">{step.title}</p>
				  </li>
				))
			  ) : (
				<p className="text-gray-500">No step for this module.</p>
			  )}
			</ul>
		  </div>

		  {/* Image content */}
		  {moduleData.image_url && (
			<div className="mt-4 sm:mt-0">
			  <img src={moduleData.image_url} alt="module image" className="w-[350px] h-[210px] rounded-sm object-cover" />
			</div>
		  )}
		</div>
	);
};

export default Module;
