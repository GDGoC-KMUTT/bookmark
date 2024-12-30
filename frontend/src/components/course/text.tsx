import { useEffect, useRef } from 'react';
import { gsap } from 'gsap';
import { ScrollTrigger } from 'gsap/ScrollTrigger';
import droplet1 from '../../assets/droplet1.png';
import droplet2 from '../../assets/droplet2.png';
import droplet3 from '../../assets/droplet3.png';
import droplet4 from '../../assets/droplet4.png';

gsap.registerPlugin(ScrollTrigger);

const Text = ({ content, backgroundIndex }: { content: string; backgroundIndex: number }) => {
    const backgroundImages = [droplet1, droplet2, droplet3, droplet4];
    const backgroundImage = backgroundImages[backgroundIndex % 4];
    const textRef = useRef<HTMLDivElement>(null);

    useEffect(() => {
        const element = textRef.current;

        if (element) {
            gsap.fromTo(
                element,
                { opacity: 0 },
                {
                    opacity: 1,
                    duration: 1.5, // Smooth animation
                    scrollTrigger: {
                        trigger: element,
                        start: 'top 60%', // Start fade-in before the component's center reaches viewport center
                        end: 'top 30%', // Fade out as the component's top moves past 40% of the viewport height
                        toggleActions: 'play reverse play reverse',
                        scrub: 1, // Synchronizes with scroll
                    },
                }
            );
        }
    }, []);

    return (
        <div
            ref={textRef}
            className="relative w-screen h-[500px] sm:h-[500px] bg-cover bg-center flex flex-col justify-center items-center text-center mt-10 mb-10"
            style={{
                backgroundImage: `url(${backgroundImage})`,
                backgroundSize: 'contain',
                backgroundRepeat: 'no-repeat',
            }}
        >
            <p className="w-[70%] text-lg text-black font-medium">{content}</p>
        </div>
    );
};

export default Text;
