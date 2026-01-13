import Navbar from "./components/Navbar";
import { Outlet } from "react-router-dom";

export default function Layout() {
  return (
    <div className="h-screen mx-auto bg-gray-100 flex flex-col">
      <Navbar />
      <main className="max-w-7xl mx-auto">
        <Outlet />
      </main>
    </div>
  );
}
