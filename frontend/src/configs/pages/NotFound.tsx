// import { Box, Heading, VStack, Text } from "@chakra-ui/react";
// import AppBody from "../../components/share/app/AppBody";
import noResultIcon from "@/assets/no-results.png"
const NotFound = () => {
    return (
        <div className="flex flex-col items-center justify-center w-screen">
            <img src={noResultIcon} alt="notFoundImg" className="w-52 h-52" />
            <h1 className="text-2xl">Page Not Found</h1>
        </div>
    )
}

export default NotFound

