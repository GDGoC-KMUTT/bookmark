import { PropsWithChildren } from "react"
import { Toaster } from "sonner"

export default function ToastProvider({ children }: PropsWithChildren) {
    return (
        <>
            {children}
            <Toaster
                className="capitalize"
                position="bottom-right"
                closeButton
                // reverseOrder={false}
                // containerStyle={{}}
                toastOptions={{
                    // Define default options
                    // className: "",
                    duration: 3000,
                    // style: {
                    //   background: "#363636",
                    //   color: "#fff",
                    // },
                    // Default options for specific types
                    // success: {
                    //   duration: 3000,
                    //   theme: {
                    //     primary: "green",
                    //     secondary: "black",
                    //   },
                    // },
                }}
            />
        </>
    )
}

