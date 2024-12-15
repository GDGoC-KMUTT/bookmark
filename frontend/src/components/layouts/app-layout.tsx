import { Outlet } from "react-router-dom"
import AppLoading from "../ui/app-loading"

export default function AppLayout() {
    return (
        <AppLoading>
            <Outlet />
        </AppLoading>
    )
}

