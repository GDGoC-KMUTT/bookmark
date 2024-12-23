import { BrowserRouter, Route, Routes } from "react-router-dom"
import { router } from "./configs/routes"
import ToastProvider from "./configs/toast"
import Navbar from "@/components/navbar"
import Footer from "./components/footer"
import AppLayout from "./components/layouts/app-layout"

function App() {
    return (
        <ToastProvider>
            <BrowserRouter>
                <Navbar />
                <Routes>
                    <Route path="/" element={<AppLayout />} />
                    {router.map((item) => (
                        <Route key={item.path} element={item.element} path={item.path} />
                    ))}
                </Routes>
            </BrowserRouter>
            <Footer />
        </ToastProvider>
    )
}

export default App

