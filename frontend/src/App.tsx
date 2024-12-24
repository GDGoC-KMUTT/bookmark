import Navbar from "@/components/navbar"
import { BrowserRouter, Route, Routes } from "react-router-dom"
import Footer from "./components/footer"
import AppLayout from "./components/layouts/app-layout"
import { router } from "./configs/routes"
import ToastProvider from "./configs/toast"

function App() {
    return (
        <ToastProvider>
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<AppLayout />} />
                    {router.map((item) =>
                        item.path == "/welcome" ? (
                            <Route key={item.path} element={item.element} path={item.path} />
                        ) : (
                            <Route
                                key={item.path}
                                element={
                                    <>
                                        <Navbar />
                                        {item.element}
                                    </>
                                }
                                path={item.path}
                            />
                        )
                    )}
                </Routes>
            </BrowserRouter>
            <Footer />
        </ToastProvider>
    )
}

export default App

