import droplet from '../../assets/droplet3.png';

const Text = ({ content }: { content: string }) => (
	<div className="relative w-screen h-[650px] bg-cover bg-center flex flex-col justify-center items-center text-center" style={{ backgroundImage: `url(${droplet})`, backgroundSize: "contain", backgroundRepeat: "no-repeat", }}>
		<p className="w-[70%] text-lg text-black font-medium">{content}</p>
	</div>
);


export default Text;
