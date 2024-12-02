import { useState } from "react";
import bookmarkLogo1 from "./assets/logo2.png";
import bookmarkLogo2 from "./assets/logo1.png";
import "./App.css";

function App() {
  const [count, setCount] = useState(0);

  return (
    <>
      <div style={{ display: "flex", justifyContent: "center" }}>
        <a href="https://react.dev" target="_blank">
          <img src={bookmarkLogo1} className="logo react" alt="React logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={bookmarkLogo2} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Bookmark เทสเทส</h1>
      <h2>test เทสเทส</h2>
      <p>test test เทสเทส</p>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  );
}

export default App;
