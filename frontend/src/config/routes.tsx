/* eslint-disable react-refresh/only-export-components */
import { RouteObject } from "react-router-dom";
import NotFound from "./pages/NotFound";

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const ROUTES: Record<string, any> = import.meta.glob(
  "/src/pages/**/[a-z[]*.tsx",
  { eager: true }
);

const routes: RouteObject[] = Object.keys(ROUTES)
  .filter((item) => item.endsWith(".tsx"))
  .map((route) => {
    const path = route
      .replace(/\/src\/pages|index|\.tsx$/g, "")
      .replace(/\[\.{3}.+\]/, "*")
      .replace(/\[(.+)\]/, ":$1")
      .replace(/\/$/, "");
    const Element = ROUTES[route].default;
    return { path, element: <Element /> };
  });

export const router = [...routes, { path: "*", element: <NotFound /> }];
