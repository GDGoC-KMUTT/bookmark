import { Blocks, Check, Loader, SquareSquare } from 'lucide-react';

const Module = ({ title, description, imageUrl }: { title: string; description: string; imageUrl: string }) => (
	<div className="sm:w-[70%] xl:w-[800px] bg-white border-2 border-gray shadow-md rounded-lg p-10 flex flex-col sm:flex-row sm:space-x-6">

	  <div className="flex flex-col space-y-4 sm:w-[60%]">
		<div className="flex flex-row">
		  <Blocks color="gray" />
		  <p className="pl-2 text-gray-600">MODULE</p>
		</div>
		<div>
		  <h2 className="text-xl font-bold mb-3">{title}</h2>
		  <p>{description}</p>
		</div>

		<div className="flex flex-row">
		  <SquareSquare color="gray" className='mt-3'/>
		  <p className="pl-2 text-gray-600 mt-3 ">STEPS</p>
		</div>
		<ul className="space-y-3 pt-3">
		  <li className="flex items-center space-x-2">
			<div className="flex items-center justify-center w-6 h-6 bg-green-500 rounded-full">
			  <Check className="text-white w-4 h-4" />
			</div>
			<p className="pl-3 text-lg">ต่อ ESP32 board เข้ากับคอมพิวเตอร์</p>
		  </li>
		  <li className="flex items-center space-x-2">
			<div className="flex items-center justify-center w-6 h-6 bg-gray-500 rounded-full">
			  <Loader className="text-white w-4 h-4" />
			</div>
			<p className="pl-3 text-lg">ต่อ ESP32 board เข้ากับคอมพิวเตอร์</p>
		  </li>
		  <li className="flex items-center space-x-2">
			<div className="flex items-center justify-center w-6 h-6 bg-gray-500 rounded-full">
			  <Loader className="text-white w-4 h-4" />
			</div>
			<p className="pl-3 text-lg">ต่อ ESP32 board เข้ากับคอมพิวเตอร์</p>
		  </li>
		</ul>
	  </div>

	  {/* Image content */}
	  <div className="mt-4 sm:mt-0">
		<img src={imageUrl} alt="module" className="w-[300px] h-[170px] rounded-sm object-cover" />
	  </div>
	</div>
  );

export default Module;
