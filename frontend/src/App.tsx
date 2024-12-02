import { BrowserRouter, Route, Routes } from "react-router-dom";
import { router } from "./configs/routes";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        {router.map((item) => (
          <Route key={item.path} element={item.element} path={item.path} />
        ))}
      </Routes>
    </BrowserRouter>
  );
}

export default App;
