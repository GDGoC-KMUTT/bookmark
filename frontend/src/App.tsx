import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom"
import { router } from "./configs/routes"
import ToastProvider from "./configs/toast"
import Navbar from "@/components/navbar"
import Footer from "./components/footer"

function App() {
    return (
        <ToastProvider>
            <Navbar />
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<Navigate to="/course" />} />
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

