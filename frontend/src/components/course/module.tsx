import { Blocks, Check, Loader, SquareSquare } from 'lucide-react';
import courseBg from '../../assets/course_bg.png';



const Module = () => {
	// const Module = ({ title, steps }) => {
	return (
		<div className="sm:w-[70%] xl:w-[1000px] bg-white border-2 border-gray shadow-md rounded-lg p-10 space-y-2">
			<div className="flex flex-row">
				<Blocks color='gray' />
				<p className='pl-2 text-gray-600'>MODULE</p>
			</div>
			<div className="flex flex-row justify-between">
				<div className='flex flex-col w-[60%] pr-2'>
					<h2 className="text-xl font-bold mb-3">การเชื่อมต่อ ESP32 Microcontroller Board เข้ากับ Arduino IDE</h2>
					<p>คือการจะคุม Microcontroller ได้เนี่ย เราต้องเขียนโปรแกรม ไส่มันผ่าน Arduino IDE ในโมดูลนี้เราจะมาลองตั้งค่าให้ Arduino สามารถแฟลชโปรแกรมลงไปในบอร์ดให้ได้กัน</p>
				</div>
				<img src={courseBg} alt="module" className="w-[290px] h-[150px] rounded-sm" />
			</div>

			<div className="flex flex-row">
				<SquareSquare color='gray' />
				<p className='pl-2 text-gray-600'>STEPS</p>
			</div>
			<ul className="space-y-2 pt-3">
				<li className="flex items-center space-x-2">
					<div className="flex items-center justify-center w-6 h-6 bg-green-500 rounded-full">
						<Check className="text-white w-4 h-4" />
					</div>
					<p className='pl-3 text-lg'>ต่อ ESP32 board เข้ากับคอมพิวเตอร์</p>
				</li>
				<li className="flex items-center space-x-2">
					<div className="flex items-center justify-center w-6 h-6 bg-gray-500 rounded-full">
						<Loader className="text-white w-4 h-4" />
					</div>
					<p className='pl-3 text-lg'>ต่อ ESP32 board เข้ากับคอมพิวเตอร์</p>
				</li>
				<li className="flex items-center space-x-2">
					<div className="flex items-center justify-center w-6 h-6 bg-gray-500 rounded-full">
						<Loader className="text-white w-4 h-4" />
					</div>
					<p className='pl-3 text-lg'>ต่อ ESP32 board เข้ากับคอมพิวเตอร์</p>
				</li>

				{/* {steps.map((step, index) => (
					<li key={index} className="flex items-center space-x-2">
						<span className="text-green-500">✔</span>
						<span>{step}</span>
					</li>
				))} */}
			</ul>
		</div>
	);
};

export default Module;
