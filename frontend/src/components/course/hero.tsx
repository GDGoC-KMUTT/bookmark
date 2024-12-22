import courseBg from '../../assets/course_bg.png';
import { Bookmark } from 'lucide-react';


const Hero = () => {
  return (
    <div className="relative w-screen h-[480px] bg-cover bg-center" style={{ backgroundImage: `url(${courseBg})` }}>
      <div className="absolute inset-0 bg-black bg-opacity-30 flex flex-col justify-center pl-10">
		<div className="relative w-[500px] space-y-8 pl-10">
			<div className="relative flex flwx-row">
				<Bookmark size={25} color="white" />
				<h1 className="pl-4 text-xl text-white font-light">COURSE / </h1>
				<h1 className="pl-2 text-xl text-white font-light">INFRASTRUCTURE</h1>
			</div>
			<h1 className="text-4xl text-white font-bold">Setup LED blink with ESP32 and Arduino IDE</h1>
			<button className="w-32 h-12 bg-primary text-fieldType-foreground rounded-lg text-xl">Enroll</button>
		</div>
      </div>
    </div>
  );
};

export default Hero;
