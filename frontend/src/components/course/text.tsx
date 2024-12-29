import droplet1 from '../../assets/droplet1.png';
import droplet2 from '../../assets/droplet2.png';
import droplet3 from '../../assets/droplet3.png';
import droplet4 from '../../assets/droplet4.png';

const Text = ({ content, backgroundIndex }: { content: string; backgroundIndex: number }) => {
	const backgroundImages = [droplet1, droplet2, droplet3, droplet4];
	const backgroundImage = backgroundImages[backgroundIndex % 4];

	return (
		<div
			className="relative w-screen h-[550px] bg-cover bg-center flex flex-col justify-center items-center text-center"
			style={{
				backgroundImage: `url(${backgroundImage})`,
				backgroundSize: "contain",
				backgroundRepeat: "no-repeat",
			}}
		>
			<p className="w-[70%] text-lg text-black font-medium">{content}</p>
		</div>
	);
};

export default Text;
