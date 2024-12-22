import droplet from '../../assets/droplet3.png';

const Text = () => {
	return (
		<div className="relative w-screen h-[650px] bg-cover bg-center flex flex-col justify-center items-center text-center" style={{ backgroundImage: `url(${droplet})`, backgroundSize: "contain", backgroundRepeat: "no-repeat", }}>
			<p className="w-[70%] text-lg text-black font-medium">“รู้จักกับการใช้งาน Arduino IDE เบื้องต้นเพื่อเขียนโปรแกรมควบคุม ESP32 ซึ่งจะทำให้เราได้ฝึกตั้งค่า Arduino IDE
				ลงไดร์ฟเวอร์ของบอร์ด และทำความเข้าใจกับ ecosystem ของ Arduino และ ESP เพื่อให้พัฒนาโปรแกรมในส่วนอื่นๆต่อไปได้”</p>
		</div>
	);
};

export default Text;
