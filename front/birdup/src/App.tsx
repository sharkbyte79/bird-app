import "./App.css";
import { HashRouter as Router, Routes, Route } from "react-router-dom";
import Layout from "./Layout";
import Home from "./pages/home";

export default function App() {
  return (
    <Router>
      <Routes>
        {/* wrap all other routes in parent layout with navbar */}
        <Route element={<Layout />}>
          <Route path="/" element={<Home />} />
        </Route>
      </Routes>
    </Router>
  );
}
