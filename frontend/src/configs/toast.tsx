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
                toastOptions={{
                    duration: 3000,
                }}
            />
        </>
    )
}

