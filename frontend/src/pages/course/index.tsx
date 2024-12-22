import Hero from '../../components/course/hero';
import Text from '../../components/course/text';
import Module from '@/components/course/module';


const Course = () => {
	return (
	  <div className="relative w-full mt-[50px] mb-[150px] flex flex-col overflow-y-auto">
		<Hero />
		<Text />
		<div className="flex flex-col items-center justify-center space-y-10 mt-10">
		  <Module />
		  <Module />
		  <Module />
		</div>
	  </div>
	);
  };


export default Course
