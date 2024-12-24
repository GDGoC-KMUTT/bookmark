import courseBg from '../../assets/course_bg.png';
import { Bookmark } from 'lucide-react';

const Hero = ({ courseField, courseName }: { courseField: string; courseName: String }) => (
	<div className="relative w-screen h-[480px] bg-cover bg-center" style={{ backgroundImage: `url(${courseBg})` }}>
		<div className="absolute inset-0 bg-black bg-opacity-30 flex flex-col justify-center pl-10">
			<div className="relative w-[30%] space-y-8 pl-10">
				<div className="relative flex flex-row">
					<Bookmark size={25} color="white" />
					<h1 className="pl-4 text-xl text-white font-light">COURSE / </h1>
					<h1 className="pl-2 text-xl text-white font-light">{courseField}</h1>
				</div>
				<h1 className="text-4xl text-white font-bold">{courseName}</h1>
				<button className="w-36 h-15 bg-primary text-fieldType-foreground rounded-lg text-xl">Enroll</button>
			</div>
		</div>
	</div>
);

export default Hero;
