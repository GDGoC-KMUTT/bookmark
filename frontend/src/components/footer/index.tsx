import SITLogo from "@/assets/sitkmutt.png"

const Footer = () => {
    return (
        <div className="w-full bg-footer h-[3rem] fixed bottom-0 shadow-md flex px-3 py-3 justify-end items-center">
            <img src={SITLogo} alt="SITLogo" className="w-15 h-8" />
        </div>
    )
}

export default Footer

